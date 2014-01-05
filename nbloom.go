package main

import (
	//"log"
	"net"
)

// A slice of BloomFilters sorted by prefix length.
type NBloom struct {
	size int
	filters []*BloomFilter
}

// Creates a new NBloom where w is the length of input addresses.
// i.e. 32 for IPv4 and 128 for IPv6.
func NewNBloom(w int) *NBloom {
	b := new(NBloom)
	b.size = w
	b.filters = make([]*BloomFilter, w)
	for i := 0; i < len(b.filters); i++ {
		b.filters[i] = NewBloomFilter(i)
	}
	return b
}

// Returns the longest prefix, and true if there exists a match
// for ip. Else returns nil and false.
func (n *NBloom) Search(ip *net.IP) (*net.IPNet, bool) {
	// []bool indicates the index of each BloomFilter that
	// reported a positive search result.
	res := make([]bool, n.size)
	for i, f := range n.filters {
		go func() {
			if _, ok := f.Search(ip); ok {
				res[i] = true
			}
		}()
	}
	// TODO - Wait for all funcs to finish.

	// For each prefix that reported a positive search
	// result, check that prefix's hashmap starting with
	// the longest prefix first.
	for i := len(res) - 1; i > -1; i-- {
		if res[i] {
			if nextHop, ok := n.filters[i].Lookup(ip); ok {
				return nextHop, true
			}
		}
	}
	return nil, false
}

func (n *NBloom) ProgramPrefix(ip *net.IPNet, nHop *net.IPNet) {
	for _, f := range n.filters {
		f.ProgramPrefix(ip, nHop)
	}
}

func main() {
	h := NewNBloom(32)
	_, p, _ := net.ParseCIDR("10.10.0.0/24")
	_, nHop, _ := net.ParseCIDR("10.10.0.1/24")
	h.ProgramPrefix(p, nHop)
}
