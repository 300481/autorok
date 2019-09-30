package autorok

import (
	"errors"
	"net"
	"strconv"
)

const (
	ERROR_ADDR_NOT_IN_NETWORK = "Address not in network."
)

type CIDR struct {
	IP   net.IP
	Net  *net.IPNet
	Size string
}

// ParseCIDR parses a string and returns a *CIDR
func ParseCIDR(s string) (*CIDR, error) {
	addr, network, err := net.ParseCIDR(s)
	if err != nil {
		return nil, err
	}

	size, _ := network.Mask.Size()
	cidr := &CIDR{
		IP:   addr,
		Net:  network,
		Size: strconv.Itoa(size),
	}

	return cidr, nil
}

// Relative returns a new *CIDR with an IP address increased/decreased by r
func (c *CIDR) Relative(r int) (*CIDR, error) {
	cidr, err := ParseCIDR(c.String())
	if err != nil {
		return nil, err
	}

	if r == 0 {
		return cidr, nil
	}

	increase := r > 0
	if !increase {
		r = -r
	}

	for i := 0; i < r; i++ {
		for j := len(cidr.IP) - 1; j >= 0; j-- {
			if increase {
				cidr.IP[j]++
				if cidr.IP[j] > 0 {
					break
				}
			} else {
				cidr.IP[j]--
				if cidr.IP[j] < 255 {
					break
				}
			}
		}
	}

	if cidr.Net.Contains(cidr.IP) {
		return cidr, nil
	}

	return nil, errors.New(ERROR_ADDR_NOT_IN_NETWORK)
}

// String returns the string form of the CIDR
func (c *CIDR) String() string {
	return c.IP.String() + "/" + c.Size
}
