package trick

import (
	"github.com/Huhaokun/loki/core/node"
)

type StateTricker struct {
}

func (t *StateTricker) StopNode(n node.Node, policy TrickPolicy) {
	policy.Apply(n.Stop, n.Start)
}

