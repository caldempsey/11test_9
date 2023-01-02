package types

const (
	RuntimeProviderKey = "RuntimeProvider"
)

type Status string

const (
	StatusInitializingNode Status = "initializingNode"
	StatusRunning          Status = "running"
	StatusStoppingNode     Status = "stoppingNode"
)

// RuntimeProvider is the platform that the node is spawned on. This is where the user
// code runs
type RuntimeProvider string

func (r RuntimeProvider) String() string {
	return string(r)
}

const (
	LambdaLabsProvider RuntimeProvider = "LambdaLabs"
	UnweaveProvider    RuntimeProvider = "Unweave"
)

type NodeSpecs struct {
	VCPUs int `json:"vCPUs"`
	// Memory is the RAM in GB
	Memory int `json:"memory"`
	// GPUMemory is the GPU RAM in GB
	GPUMemory *int `json:"gpuMemory"`
}

type NodeType struct {
	ID          string          `json:"id"`
	Available   bool            `json:"available"`
	Name        *string         `json:"name"`
	Price       *int            `json:"price"`
	Region      *string         `json:"region"`
	Description *string         `json:"description"`
	Provider    RuntimeProvider `json:"provider"`
	Specs       NodeSpecs       `json:"specs"`
}

type Node struct {
	ID       string  `json:"id"`
	Region   *string `json:"region"`
	KeyPair  SSHKey  `json:"sshKeyPair"`
	Status   Status  `json:"status"`
	NodeType `json:"nodeType"`
}

type SSHKey struct {
	Name       *string `json:"name,omitempty"`
	PrivateKey *string `json:"privateKey,omitempty"`
	PublicKey  *string `json:"publicKey,omitempty"`
}

type Session struct {
	ID     string `json:"id"`
	SSHKey SSHKey `json:"sshKey"`
	Status Status `json:"runtimeStatus"`
}

type ExecParams struct {
	Cmd   []string `json:"cmd"`
	Image []string `json:"containerImage"`
}
