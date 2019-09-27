package autorok

// Node data
type node struct {
	Hostname    string   // the hostname of the node
	AddressCIDR string   // the IP address of the node with netmask as CIDR
	Gateway     string   // the default gateway for the node
	MTU         int      // the MTU for the network interface
	Nameservers []string // a list of nameservers
	PublicKey   string   // the SSH public key
	DHCP        bool     // enable DHCP true/false
}

// Cluster data
type cluster struct {
	Name      string           // a name for the cluster
	NodeCount int              // the configured size of the cluster
	Nodes     map[string]*node // a slice with all cluster nodes
}

func (c *config) newCluster() *cluster {
	return &cluster{
		Name:      c.ClusterName,
		NodeCount: c.NodeCount,
		Nodes:     make(map[string]*node),
	}
}
