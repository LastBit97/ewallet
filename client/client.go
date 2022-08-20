package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/LastBit97/ewallet/ewallet"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	r, err = c.SayHelloAgain(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	wallet1 := &pb.Wallet{
		Address: "e240d825d255af751f5f55af8d9671beabdf2236c0a3b4e2639b3e182d994c88",
		Balance: 7525,
	}

	wallet2 := &pb.Wallet{
		Address: "9ee4b55aa524c869fda5d86dde14c512cec65fb6ff315ce2b45dd76631b2cfcb",
		Balance: 9255,
	}

	wallet3 := &pb.Wallet{
		Address: "8vpn0a5zi0qhnej43ofqtarl1r9zml7b9q4nouyq00n9z2pu9rjdxpj0w6u49qux",
		Balance: 5408,
	}

	reply, err := c.CreateWallet(context.Background(), &pb.CreateWalletRequest{Wallet: wallet1})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	log.Printf("Create wallet with address: %v", reply.Wallet.GetAddress())

	reply, err = c.CreateWallet(context.Background(), &pb.CreateWalletRequest{Wallet: wallet2})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	log.Printf("Create wallet with address: %v", reply.Wallet.GetAddress())
	reply, err = c.CreateWallet(context.Background(), &pb.CreateWalletRequest{Wallet: wallet3})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	log.Printf("Create wallet with address: %v", reply.Wallet.GetAddress())
}
