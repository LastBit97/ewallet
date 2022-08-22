/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/LastBit97/ewallet/ewallet"
	"github.com/spf13/cobra"
)

// getlastCmd represents the getlast command
var getlastCmd = &cobra.Command{
	Use:   "getlast",
	Short: "Returns an array of transactions",
	Long:  `Returns an array of transactions`,
	Run: func(cmd *cobra.Command, args []string) {

		req := &pb.GetLastRequest{}
		stream, err := client.GetLast(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		// Iterate
		for {
			res, err := stream.Recv()
			// If it is end of the stream, then break the loop
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(res.GetTransaction())
		}
	},
}

func init() {
	rootCmd.AddCommand(getlastCmd)
}
