package main

import (
	"fmt"
	"net"
)

// A slice of BloomFilters sorted by prefix length.
type NBloom struct {
	size int
	filters []*BloomFilter
}

// Creates a new NBloom where w is the length of input addresses.
// i.e. 24 for IPv4 and 128 for IPv6.
func NewNBloom(w int) *NBloom {
	b := new(NBloom)
	b.size = w
	b.filters = make([]*BloomFilter, w)
	for i, _ := range b.filters {
		b.filters[i] = NewBloomFilter(i)
	}
	return b
}

// Returns a []int indicating the index of each BloomFilter that
// reported a positive search result.
func (n *NBloom) Search(ip *net.IP) (*net.IPNet, bool) {
	res := make([]bool, n.size)
	for i, f := range n.filters {
		go func() {
			if _, ok := f.Search(ip); ok {
				res[i] = true
			}
		}()
	}
	// TODO - Wait for all funcs to finish.

	for i := len(res) - 1; i > -1; i-- {
		if res[i] {
			if nextHop, ok := n.filters[i].Lookup(ip); ok {
				return nextHop, true
			}
		}
	}
	return nil, false
}

type BloomFilter struct {
	index int
	hashes HashSet
	prefixes map[string]*net.IPNet
}

func NewBloomFilter(p int) *BloomFilter {
	f := new(BloomFilter)
	f.index = p
	f.hashes = NewHashSet(p)
	f.prefixes = make(map[string]*net.IPNet)
	return f
}

func (b *BloomFilter) ProgramPrefix(p *net.IPNet, nextHop *net.IPNet) {
	b.prefixes[p.IP.String()] = nextHop
}

func (b *BloomFilter) Search(ip *net.IP) (int, bool) {
	if true {
		return b.index, true
	}
	return -1, false
}

func (b *BloomFilter) Lookup(ip *net.IP) (n *net.IPNet, ok bool) {
	n, ok = b.prefixes[ip.String()]
	return
}

func main() {
	h := NewNBloom(10)
	fmt.Println("Hello nbloom", h)
}
