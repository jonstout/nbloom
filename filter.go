package main

import (
	"log"
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
	k := p
	
	f.index = p
	f.filter = make([]bool, p)
	f.hashes = NewHashSet(k)
	f.prefixes = make(map[string]*net.IPNet)
	return f
}

func (b *BloomFilter) ProgramPrefix(p *net.IPNet, nextHop *net.IPNet) {
	b.prefixes[p.IP.String()] = nextHop
	
	for _, f := range b.hashes {
		n := f(p.IP)
		log.Println(n)
		b.filter[n] = true
	}
}

func (b *BloomFilter) Search(ip *net.IP) (int, bool) {
	for f := range b.hashes {
		n := f
		if !b.filter[n] {
			return -1, false
		}
	}
	return b.index, true
}

func (b *BloomFilter) Lookup(ip *net.IP) (n *net.IPNet, ok bool) {
	n, ok = b.prefixes[ip.String()]
	return
}
