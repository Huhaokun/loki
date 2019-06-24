package cmd

import (
	"context"
	"github.com/Huhaokun/loki/api/loki"
	"github.com/Huhaokun/loki/core/trick"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var stateTrickCommand = &cobra.Command{
	Use:   "state-trick",
	Short: "apply state trick to nodes with a yaml file",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		applyYaml(args[0])
	},
}

func init() {
	rootCmd.AddCommand(stateTrickCommand)
}

type StateTrickCommand struct {
	Policy trick.TrickPolicy `yaml:"policy"`
	Nodes  []string          `yaml:"nodes,flow"`
	Type   string            `yaml:"type"`
}

func (cmd *StateTrickCommand) toProto() *loki.StateTrick {
	// transform node

	var nodes []*loki.Node
	for _, nodeId := range cmd.Nodes {
		nodes = append(nodes, &loki.Node{
			Id: nodeId,
		})
	}

	var t loki.StateTrick_Type
	if cmd.Type == "shutdown" {
		t = loki.StateTrick_NODE_DOWN
	} else if cmd.Type == "restart" {
		t = loki.StateTrick_NODE_RESTRT
	} else {
		log.Fatalf("not support type %s", cmd.Type)
	}

	return &loki.StateTrick{
		Nodes: nodes,
		Type:  t,
		Policy: &loki.TrickPolicy{
			Delay:    cmd.Policy.Delay,
			Interval: cmd.Policy.Interval,
			Keep:     cmd.Policy.Keep,
		},
	}

}

func applyYaml(filePath string) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("read file from %s error, err: %s", filePath, err.Error())
	}

	trickCommand := StateTrickCommand{}
	if err := yaml.Unmarshal(bytes, &trickCommand); err != nil {
		log.Fatalf("parse file as yaml failed, err: %s", err.Error())
	}

	_, err = newClient().Apply(context.Background(), trickCommand.toProto())

	if err != nil {
		log.Error("apply policy error", err)
	}
}

func newClient() loki.StateTrickerClient {
	return loki.NewStateTrickerClient(NewGrpcConn())
}
