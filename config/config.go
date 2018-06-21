package config // import "chainspace.io/prototype/config"

// NOTE(tav): The order of some of the struct fields are purposefully not in
// alphabetical order so as to generate a more pleasing ordering when serialised
// to YAML.
import (
	"crypto/sha512"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Key represents a cryptographic key of some kind.
type Key struct {
	Type    string
	Public  string
	Private string `yaml:",omitempty"`
}

// Keys represents the configuration that individual nodes use for their various
// cryptographic keys.
type Keys struct {
	SigningKey    *Key `yaml:"signing.key"`
	TransportCert *Key `yaml:"transport.cert"`
}

// Network represents the configuration of a Chainspace network as a whole.
type Network struct {
	ID        string           `yaml:",omitempty"`
	SeedNodes map[uint64]*Peer `yaml:"seed.nodes"`
	Shard     *Shard
}

// Hash returns the SHA-512/256 hash of the network's seed configuration by
// first encoding it into YAML. The returned hash should then be used as the ID
// for the network.
func (n *Network) Hash() ([]byte, error) {
	data, err := yaml.Marshal(&Network{
		SeedNodes: n.SeedNodes,
		Shard:     n.Shard,
	})
	if err != nil {
		return nil, err
	}
	h := sha512.New512_256()
	if _, err = h.Write(data); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

// Node represents the configuration of an individual node in a Chainspace
// network.
type Node struct {
	Announce         []string          `yaml:"announce,omitempty"`
	BootstrapFile    string            `yaml:"bootstrap.file,omitempty"`
	BootstrapMDNS    bool              `yaml:"bootstrap.mdns,omitempty"`
	BootstrapStatic  map[uint64]string `yaml:"bootstrap.static,omitempty"`
	BootstrapURL     string            `yaml:"bootstrap.url,omitempty"`
	HostIP           string            `yaml:"host.ip,omitempty"`
	RuntimeDirectory string            `yaml:"runtime.directory"`
}

// Peer represents the cryptographic public keys of a node in a Chainspace
// network.
type Peer struct {
	SigningKey    *PeerKey `yaml:"signing.key"`
	TransportCert *PeerKey `yaml:"transport.cert"`
}

// PeerKey represents a cryptographic public key of some kind.
type PeerKey struct {
	Type  string
	Value string
}

// Shard represents the configuration of the sharding system within a given
// Chainspace network.
type Shard struct {
	Count int
	Size  int
}

// LoadKeys will read the YAML file at the given path and return the
// corresponding Keys config.
func LoadKeys(path string) (*Keys, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	cfg := &Keys{}
	err = yaml.Unmarshal(data, cfg)
	return cfg, err
}

// LoadNetwork will read the YAML file at the given path and return the
// corresponding Network config.
func LoadNetwork(path string) (*Network, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	network := &Network{}
	err = yaml.Unmarshal(data, network)
	return network, err
}

// LoadNode will read the YAML file at the given path and return the
// corresponding Node config.
func LoadNode(path string) (*Node, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	cfg := &Node{}
	err = yaml.Unmarshal(data, cfg)
	return cfg, err
}
