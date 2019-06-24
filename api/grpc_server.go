package api

import (
	"context"
	"errors"
	pb "github.com/Huhaokun/loki/api/loki"
	"github.com/Huhaokun/loki/core/node"
	"github.com/Huhaokun/loki/core/trick"
	log "github.com/sirupsen/logrus"
)

type NodeControllerServer struct {
	nodes map[string]node.Node
}

func NewNodeControllerServer() *NodeControllerServer {
	return &NodeControllerServer{
		nodes: make(map[string]node.Node),
	}
}

func (s *NodeControllerServer) ListNode(empty *pb.Empty, stream pb.NodeController_ListNodeServer) error {
	log.Infof("List node")
	for id, _ := range s.nodes {
		err := stream.Send(&pb.Node{
			Id: id,
		})

		if err != nil {
			return err
		}
	}
	return nil
}

func (s *NodeControllerServer) AddNode(ctx context.Context, n *pb.Node) (*pb.BaseResponse, error) {
	log.Infof("Add Node %v", n)

	// TODO support k8s type node
	switch n.GetType() {
	case pb.Node_K8S:
		return nil, errors.New("not impl")
	case pb.Node_DOCKER:
		s.nodes[n.GetId()] = node.NewDockerNode(n.GetId(), n.GetNetworkId())
		return &pb.BaseResponse{}, nil
	default:
		return nil, errors.New("invalid node type")
	}
}

func (s *NodeControllerServer) RemoveNode(ctx context.Context, n *pb.Node) (*pb.BaseResponse, error) {
	log.Infof("Remove Node %v", n)
	delete(s.nodes, n.GetId())
	return &pb.BaseResponse{}, nil
}

type ResourceTrickerServer struct {
	tricker trick.ResourceTricker
}

func (s *ResourceTrickerServer) Apply(ctx context.Context, trick *pb.ResourceTrick) (*pb.BaseResponse, error) {
	log.Infof("Apply resource trick %v", trick)
	return nil, errors.New("not impl")
}

type StateTrickerServer struct {
	nodeRegister *NodeControllerServer
	tricker      *trick.StateTricker
}

func NewStatetrickerServer(n *NodeControllerServer, t *trick.StateTricker) *StateTrickerServer {
	return &StateTrickerServer{
		nodeRegister: n,
		tricker:      t,
	}
}

func (s *StateTrickerServer) Apply(ctx context.Context, trick *pb.StateTrick) (*pb.BaseResponse, error) {
	log.Infof("Apply state trick %v", trick)
	policy := pb2Policy(trick.Policy)

	for _, nPb := range trick.GetNodes() {
		if n, ok := s.nodeRegister.nodes[nPb.GetId()]; ok {
			switch trick.GetType() {
			case pb.StateTrick_NODE_DOWN:
				s.tricker.StopNode(n, policy)
				break
			default:
				return nil, errors.New("not impl")
			}
		}
	}
	return &pb.BaseResponse{}, nil
}

func pb2Policy(policy *pb.TrickPolicy) trick.TrickPolicy {
	return trick.TrickPolicy{
		Keep:     policy.GetKeep(),
		Delay:    policy.GetDelay(),
		Interval: policy.GetInterval(),
	}
}
