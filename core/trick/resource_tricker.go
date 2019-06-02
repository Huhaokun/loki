package trick

import "github.com/Huhaokun/loki/core/node"

func LimitCPU(n node.Node, policy TrickPolicy) {
	policy.Apply()
}

func LimitMemory(n node.Node, policy TrickPolicy) {
	policy.Apply()
}
