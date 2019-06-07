package cmd

import (
	"context"
	"github.com/Huhaokun/loki/api/loki"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var nodeCommand = &cobra.Command{
	Use:   "node",
	Short: "manage node in loki",
}

var nodeAddCommand = &cobra.Command{
	Use:   "add",
	Short: "add node to loki server",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("add container id ", args[0])
		_, err := NewClient().AddNode(context.Background(), &loki.Node{
			Id:   args[0],
			Type: loki.Node_DOCKER,
		})

		if err != nil {
			log.Error("add node failed", err)
		}
	},
}

var nodeRemoveCommand = &cobra.Command{
	Use:   "rm",
	Short: "remove node from loki server",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("remove container id ", args[0])
		_, err := NewClient().RemoveNode(context.Background(), &loki.Node{
			Id: args[0],
		})

		if err != nil {
			log.Error("remove node failed", err)
		}
	},
}

var nodeListCommand = &cobra.Command{
	Use:   "ls",
	Short: "ls node from loki server",
	Run: func(cmd *cobra.Command, args []string) {
		NewClient().ListNode(context.Background(), &loki.Empty{})
	},
}

func init() {
	nodeCommand.AddCommand(nodeAddCommand)
	nodeCommand.AddCommand(nodeRemoveCommand)
	nodeCommand.AddCommand(nodeListCommand)

	rootCmd.AddCommand(nodeCommand)

}

func NewClient() loki.NodeControllerClient {
	return loki.NewNodeControllerClient(NewGrpcConn())
}
