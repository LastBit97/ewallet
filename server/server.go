package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/LastBit97/ewallet/ewallet"
	couchdb "github.com/go-kivik/couchdb/v4" // The CouchDB driver
	kivik "github.com/go-kivik/kivik/v4"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

var database *kivik.DB

const (
	driverName     string = "couch"
	dataSourceName string = "http://localhost:5984"
	user           string = "admin"
	password       string = "admin"
	databaseName   string = "ewallet_database"
)

type server struct {
	pb.UnimplementedEWalletServer
}

type Wallet struct {
	Id      string  `json:"_id,omitempty"`
	Rev     string  `json:"_rev,omitempty"`
	Type    string  `json:"type"`
	Address string  `json:"address"`
	Balance float32 `json:"balance"`
}

type Transaction struct {
	Id       string  `json:"_id,omitempty"`
	Rev      string  `json:"_rev,omitempty"`
	Type     string  `json:"type"`
	Datetime string  `json:"datetime"`
	Address  string  `json:"address"`
	Amount   float32 `json:"amount"`
	Flag     bool    `json:"flag"`
}

func (*server) CreateWallet(ctx context.Context, req *pb.CreateWalletRequest) (*pb.CreateWalletReply, error) {

	walletData := req.GetWallet()
	data := Wallet{
		Type:    "wallet",
		Address: walletData.GetAddress(),
		Balance: walletData.GetBalance(),
	}
	//create doc in db
	theId := uuid.New().String()
	_, err := database.Put(context.TODO(), theId, data)
	if err != nil {
		panic(err)
	}

	return &pb.CreateWalletReply{
		Wallet: &pb.Wallet{
			Address: walletData.GetAddress(),
			Balance: walletData.GetBalance(),
		},
	}, nil
}

func (s *server) Send(ctx context.Context, req *pb.SendRequest) (*pb.SendReply, error) {

	addressFrom := req.GetFrom()
	addressTo := req.GetTo()
	amount := req.GetAmount()

	query := map[string]interface{}{
		"selector": map[string]interface{}{
			"type":    "wallet",
			"address": addressFrom,
		},
	}

	results := database.Find(context.TODO(), query)

	var walletFrom Wallet
	for results.Next() {
		if err := results.ScanDoc(&walletFrom); err != nil {
			panic(err)
		}

		if walletFrom.Balance-amount <= 0 {
			return &pb.SendReply{Message: "insufficient funds in the wallet"}, nil
		}

		walletFrom.Balance -= amount
		newRev, err := database.Put(context.TODO(), walletFrom.Id, walletFrom)
		if err != nil {
			panic(err)
		}
		walletFrom.Rev = newRev
	}

	//create itransaction
	timeOutput := time.Now().Format("2006-01-02T15:04:05-0700")
	outputTransaction := Transaction{
		Type:     "output transaction",
		Datetime: timeOutput,
		Address:  walletFrom.Address,
		Amount:   amount,
	}

	theId := uuid.New().String()
	_, err := database.Put(context.TODO(), theId, outputTransaction)
	if err != nil {
		panic(err)
	}

	query = map[string]interface{}{
		"selector": map[string]interface{}{
			"type":    "wallet",
			"address": addressTo,
		},
	}
	results = database.Find(context.TODO(), query)

	var walletTo Wallet
	for results.Next() {
		if err := results.ScanDoc(&walletTo); err != nil {
			panic(err)
		}

		walletTo.Balance += amount
		newRev, err := database.Put(context.TODO(), walletTo.Id, walletTo)
		if err != nil {
			panic(err)
		}
		walletTo.Rev = newRev
	}

	//create transaction
	timeInput := time.Now().Format("2006-01-02T15:04:05-0700")
	inputTransaction := Transaction{
		Type:     "input transaction",
		Datetime: timeInput,
		Address:  walletTo.Address,
		Amount:   amount,
		Flag:     false,
	}

	theId = uuid.New().String()
	_, err = database.Put(context.TODO(), theId, inputTransaction)
	if err != nil {
		panic(err)
	}

	return &pb.SendReply{Message: "success"}, nil

}

func (s *server) GetLast(req *pb.GetLastRequest, stream pb.EWallet_GetLastServer) error {

	query := map[string]interface{}{
		"selector": map[string]interface{}{
			"type": "input transaction",
			"flag": false,
		},
	}
	results := database.Find(context.TODO(), query)
	data := &Transaction{}

	for results.Next() {
		if err := results.ScanDoc(&data); err != nil {
			panic(err)
		}

		stream.Send(&pb.GetLastReply{
			Transaction: &pb.Transaction{
				Datetime: data.Datetime,
				Address:  data.Address,
				Amount:   data.Amount,
			},
		})

		data.Flag = true
		_, err := database.Put(context.TODO(), data.Id, data)
		if err != nil {
			panic(err)
		}
	}
	if results.Err() != nil {
		panic(results.Err())
	}
	return nil
}

func main() {

	//connect to db
	client, err := kivik.New(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	err = client.Authenticate(context.TODO(), couchdb.BasicAuth(user, password))
	if err != nil {
		panic(err)
	}
	database = client.DB(databaseName)

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEWalletServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
