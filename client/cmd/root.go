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
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client = pb.NewEWalletClient(conn)
}
