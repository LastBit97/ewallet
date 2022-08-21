/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"flag"
	"log"
	"os"

	pb "github.com/LastBit97/ewallet/ewallet"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client pb.EWalletClient

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "client",
	Short: "a gRPC client to communicate with the Ewallet server",
	Long: `a gRPC client to communicate with the Ewallet server.
	You can use this client to create wallet and send money.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.client.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client = pb.NewEWalletClient(conn)
}

// package main

// import (
// 	"context"
// 	"flag"
// 	"log"
// 	"strconv"
// 	"time"

// 	pb "github.com/LastBit97/ewallet/ewallet"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// )

// var addr = flag.String("addr", "localhost:50051", "the address to connect to")

// func main() {
// 	flag.Parse()
// 	if flag.NArg() < 3 {
// 		log.Fatal("not enough arguments")
// 	}
// 	addressFrom := flag.Arg(0)
// 	addressTo := flag.Arg(1)
// 	amount, err := strconv.ParseFloat(flag.Arg(2), 32)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Fatalf("did not connect: %v", err)
// 	}
// 	defer conn.Close()
// 	c := pb.NewEWalletClient(conn)

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()

// 	reply, err := c.Send(ctx, &pb.SendRequest{
// 		From:   addressFrom,
// 		To:     addressTo,
// 		Amount: float32(amount)})

// 	if err != nil {
// 		log.Fatalf("Unexpected error: %v", err)
// 	}
// 	log.Printf("Fund balance: %v", reply.GetBalance())

// 	// wallet1 := &pb.Wallet{
// 	// 	Address: "e240d825d255af751f5f55af8d9671beabdf2236c0a3b4e2639b3e182d994c88",
// 	// 	Balance: 7525,
// 	// }

// 	// wallet2 := &pb.Wallet{
// 	// 	Address: "9ee4b55aa524c869fda5d86dde14c512cec65fb6ff315ce2b45dd76631b2cfcb",
// 	// 	Balance: 9255,
// 	// }

// 	// wallet3 := &pb.Wallet{
// 	// 	Address: "8vpn0a5zi0qhnej43ofqtarl1r9zml7b9q4nouyq00n9z2pu9rjdxpj0w6u49qux",
// 	// 	Balance: 5408,
// 	// }

// 	// reply, err := c.CreateWallet(context.Background(), &pb.CreateWalletRequest{Wallet: wallet1})
// 	// if err != nil {
// 	// 	log.Fatalf("Unexpected error: %v", err)
// 	// }
// 	// log.Printf("Create wallet with address: %v", reply.Wallet.GetAddress())

// 	// reply, err = c.CreateWallet(context.Background(), &pb.CreateWalletRequest{Wallet: wallet2})
// 	// if err != nil {
// 	// 	log.Fatalf("Unexpected error: %v", err)
// 	// }
// 	// log.Printf("Create wallet with address: %v", reply.Wallet.GetAddress())
// 	// reply, err = c.CreateWallet(context.Background(), &pb.CreateWalletRequest{Wallet: wallet3})
// 	// if err != nil {
// 	// 	log.Fatalf("Unexpected error: %v", err)
// 	// }
// 	// log.Printf("Create wallet with address: %v", reply.Wallet.GetAddress())
// }
