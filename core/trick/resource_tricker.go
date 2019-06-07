package trick

import "github.com/Huhaokun/loki/core/node"

type ResourceTricker struct {
}

func (r *ResourceTricker) LimitCPU(n node.Node, policy TrickPolicy, cpuNum int64) {
	policy.Apply(func() error {
		return n.Update(node.NodeResource{
			Cpu: cpuNum,
		})
	}, func() error {
		return n.Update(node.NodeResource{
			Cpu: 0,
		})
	})
}

func (r *ResourceTricker) LimitMemory(n node.Node, policy TrickPolicy, memoryInByte int64) {
	policy.Apply(func() error {
		return n.Update(node.NodeResource{
			Memory: memoryInByte,
		})
	}, func() error {
		return n.Update(node.NodeResource{
			Memory: 0,
		})
	})
}
