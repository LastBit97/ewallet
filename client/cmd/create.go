/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"log"

	pb "github.com/LastBit97/ewallet/ewallet"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create wallet",
	Long:  `Create wallet`,
	Run: func(cmd *cobra.Command, args []string) {

		address, err := cmd.Flags().GetString("address")
		if err != nil {
			log.Fatal(err)
		}
		balance, err := cmd.Flags().GetFloat32("balance")
		if err != nil {
			log.Fatal(err)
		}

		wallet := &pb.Wallet{
			Address: address,
			Balance: balance,
		}

		reply, err := client.CreateWallet(context.Background(), &pb.CreateWalletRequest{Wallet: wallet})
		if err != nil {
			log.Fatalf("Unexpected error: %v", err)
		}
		log.Printf("Create wallet with address: %v", reply.Wallet.GetAddress())
	},
}

func init() {
	createCmd.Flags().StringP("address", "a", "", "wallet address")
	createCmd.Flags().Float32P("balance", "b", 5000, "wallet balance")
	createCmd.MarkFlagRequired("address")
	createCmd.MarkFlagRequired("balance")
	rootCmd.AddCommand(createCmd)
}
