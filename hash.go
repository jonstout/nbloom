package main

import (
	//"log"
	"math/rand"
	"time"
)

// A function that returns an integer val representing the index of an
// array to be set.
type Hash func([]byte) (val int)

// A slice of functions that return an integer which represent all
// indexes of an array to be set.
type HashSet []Hash

// Returns a HashSet of length size.
func NewHashSet(size int) HashSet {
	set := make([]Hash, size)

	for i := 0; i < size; i++ {
		set[i] = NewHashFunction(size)
	}
	return set
}

// Returns a hash function to generate indexes based on a given byte
// array. The returned Hash will return indexes between size - 1 and 0.
func NewHashFunction(size int) Hash {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	
	f := func(addr []byte) (i int) {
		h := hashValue(addr)
		return (h + r.Int()) % size
	}
	return f
}

// Returns the md5 checksum of buf.
func hashValue(buf []byte) int {
	//h := md5.New()
	//return h.Sum(data)
	return 10
}
