package node

type NodeResource struct {
	Cpu    int64 // core number
	Memory int64 // memory in bytes
}
type Node interface {
	Start() error
	Stop() error
	Pause() error
	Restart() error
	Connect() error
	Disconnect() error
	UpdateResource(r NodeResource) error
}
