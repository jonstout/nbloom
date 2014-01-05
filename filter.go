package main

import (
	"log"
	"net"
	"math"
)

type BloomFilter struct {
	index int
	filter []bool
	hashes HashSet
	prefixes map[string]*net.IPNet
	
	// number of hash functions in this bloom filter
	k int
	// The size of this bloom filter.
	m int
	// Number of prefixes stored in this bloom filter
	n int
}

// total amount of embedded memory available for Bloom filters
var M int = 2000
// total number of prefixes supported by this system
var N int = 32

func NewBloomFilter(w int) *BloomFilter {
	f := new(BloomFilter)
	f.index = w
	// (M / N) is the size of this bloom filter. This could be improved to
	// be the size of ((M / N) * p) where p is the probibility that a
	// random prefix will be found in this filter.
	f.m = M / N
	// In the optimal case, when false positive probability is minimized
	// with respect to k, we get the following relationship:
	// k = (m/n)ln(2)
	f.k = int(float64(f.m) * math.Log(2))
	f.n = 0
	f.filter = make([]bool, f.m)
	f.hashes = NewHashSet(f.k)
	f.prefixes = make(map[string]*net.IPNet)
	log.Println("Prefix", f.index, "Bloom Filter size", f.m, "with", f.k, "hash funcs created.")
	return f
}

func (b *BloomFilter) ProgramPrefix(p *net.IPNet, nextHop *net.IPNet) {
	b.prefixes[p.IP.String()] = nextHop
	
	for _, f := range b.hashes {
		n := f(p.IP)
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

func (b *BloomFilter) String() string {
	s := "["
	for _, b := range b.filter {
		if b == true {
			s += "1, "
		} else {
			s += "0, "
		}
	}
	return s[:len(s) - 2] + "]"
}
