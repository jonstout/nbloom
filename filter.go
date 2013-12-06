package main

import (
	"net"
)

type BloomFilter struct {
	index int
	filter []bool
	hashes HashSet
	prefixes map[string]*net.IPNet
}

func NewBloomFilter(p int) *BloomFilter {
	f := new(BloomFilter)
	f.index = p
	f.filter = make([]bool, p)
	f.hashes = NewHashSet(p) // Should be size k
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
