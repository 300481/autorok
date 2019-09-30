package autorok

import (
	"log"
)

// Node data
type node struct {
	Hostname    string   // the hostname of the node
	AddressCIDR string   // the IP address of the node with netmask as CIDR
	Address     string   // the IP address of the node
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
	StartCIDR *CIDR            // The CIDR struct for the starting IP address
}

func (c *config) newCluster() *cluster {
	cidr, err := ParseCIDR(c.StartCIDR)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &cluster{
		Name:      c.ClusterName,
		NodeCount: c.NodeCount,
		Nodes:     make(map[string]*node),
		StartCIDR: cidr,
	}
}

func (c *cluster) size() int {
	return len(c.Nodes)
}

func (c *cluster) maxNodes() bool {
	return c.NodeCount == c.size()
}

func (a *Autorok) getNode(uuid string) *node {
	n, ok := a.Cluster.Nodes[uuid]
	if ok {
		return n
	}

	if a.Cluster.maxNodes() {
		return nil
	}

	cidr, err := a.Cluster.StartCIDR.Relative(a.Cluster.size())
	if err != nil {
		log.Println(err)
		return nil
	}

	n = &node{
		Hostname:    uuid,
		AddressCIDR: cidr.String(),
		Address:     cidr.IP.String(),
		Gateway:     a.Config.Gateway,
		MTU:         a.Config.MTU,
		Nameservers: a.Config.Nameservers,
		PublicKey:   a.Config.PublicKey,
		DHCP:        a.Config.DHCP,
	}

	a.Cluster.Nodes[uuid] = n

	return n
}
