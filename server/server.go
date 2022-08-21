package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

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
	Address string  `json:"address"`
	Balance float32 `json:"balance"`
}

func (*server) CreateWallet(ctx context.Context, req *pb.CreateWalletRequest) (*pb.CreateWalletReply, error) {

	walletData := req.GetWallet()
	data := Wallet{
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
			"address": map[string]interface{}{
				"$eq": addressFrom,
			},
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
			"address": map[string]interface{}{
				"$eq": addressTo,
			},
		},
	}
	results = database.Find(context.TODO(), query)

	for results.Next() {
		var walletTo Wallet
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
	return &pb.SendReply{Balance: walletFrom.Balance}, nil

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
