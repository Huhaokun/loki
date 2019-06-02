package node

type NodeResource struct {
	cpu    int64 // core number
	memory int64 // memory in bytes
}
type Node interface {
	Start() error
	Stop() error
	Pause() error
	Restart() error
	Connect() error
	Disconnect() error
	Update(r NodeResource) error
}
