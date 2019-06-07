package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "loki-ctl",
	Short: "command tool line for loki",
	Long:  `A command tool line for loki`,
}

var endpoint *string

func init() {
	// add root flag
	endpoint = rootCmd.PersistentFlags().String("endpoint", "127.0.0.1:33033", "loki server endpoint")

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func NewGrpcConn() *grpc.ClientConn {

	conn, err := grpc.Dial(*endpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial endpoint %s failed", *endpoint)
	}

	return conn
}
