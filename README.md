# loki -- A fault injection framework for building cloud native applications.

## core concept
### Node
Node: in k8s it's a pod, in docker it's a container, in real world, it's a phycical machine. 

### Trick
you can set different tricks to mess around with your system.
#### resource limit trick 
- CPU generates high load for one or more CPU cores.
- Memory Allocates a specific amount of RAM.
- IO puts read/write pressure on I/O devices
- Disk writes files to disk to fill it.
#### state trick 
- Shutdown reboot or halts the host os.
- Time travel changes the host's system time
- Process killer kill the specified process

#### network trick 
- Blackhole Drops all matching network traffic.
- Latency Injects latency into all matching egress network traffic.
- Packet loss Induces packet loss into all matching egress network traffic.
- DNS Blocks access to DNS servers