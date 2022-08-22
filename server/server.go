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
			log.Printf("insufficient funds in the wallet: %v", walletFrom.Address)
			return &pb.SendReply{Balance: walletFrom.Balance}, nil
		}

		walletFrom.Balance -= amount
		newRev, err := database.Put(context.TODO(), walletFrom.Id, walletFrom)
		if err != nil {
			panic(err)
		}
		walletFrom.Rev = newRev
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
	time := time.Now().Format("02.01.2006 15:04")
	transaction := Transaction{
		Type:     "transaction",
		Datetime: time,
		Address:  walletTo.Address,
		Amount:   amount,
	}

	theId := uuid.New().String()
	_, err := database.Put(context.TODO(), theId, transaction)
	if err != nil {
		panic(err)
	}

	return &pb.SendReply{Balance: walletFrom.Balance}, nil

}

func (s *server) GetLast(req *pb.GetLastRequest, stream pb.EWallet_GetLastServer) error {

	query := map[string]interface{}{
		"selector": map[string]interface{}{
			"type": "transaction",
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
	}
	if results.Err() != nil {
		panic(results.Err())
	}
	return nil
}

func main() {

	//connect to db
	client, err := kivik.New("couch", "http://localhost:5984")
	if err != nil {
		panic(err)
	}
	err = client.Authenticate(context.TODO(), couchdb.BasicAuth("admin", "admin"))
	if err != nil {
		panic(err)
	}
	database = client.DB("ewallet_database")

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
