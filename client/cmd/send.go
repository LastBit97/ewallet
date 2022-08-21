/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"log"
	"time"

	pb "github.com/LastBit97/ewallet/ewallet"
	"github.com/spf13/cobra"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "send money from one wallet to another",
	Long:  `send money from one wallet to another`,
	Run: func(cmd *cobra.Command, args []string) {

		addressFrom, err := cmd.Flags().GetString("addressFrom")
		if err != nil {
			log.Fatal(err)
		}
		addressTo, err := cmd.Flags().GetString("addressTo")
		if err != nil {
			log.Fatal(err)
		}
		amount, err := cmd.Flags().GetFloat32("amount")
		if err != nil {
			log.Fatal(err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		reply, err := client.Send(ctx, &pb.SendRequest{
			From:   addressFrom,
			To:     addressTo,
			Amount: float32(amount)})

		if err != nil {
			log.Fatalf("Unexpected error: %v", err)
		}
		log.Printf("Fund balance: %v", reply.GetBalance())
	},
}

func init() {
	sendCmd.Flags().StringP("addressFrom", "f", "", "the address of the wallet from which you need to transfer money")
	sendCmd.Flags().StringP("addressTo", "t", "", "the address of the wallet to which you need to transfer money")
	sendCmd.Flags().Float32P("amount", "a", 100, "transfer amount")
	sendCmd.MarkFlagRequired("addressFrom")
	sendCmd.MarkFlagRequired("addressTo")
	sendCmd.MarkFlagRequired("amount")
	rootCmd.AddCommand(sendCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
