package main

import (
	"log"
	"math/rand"
	"time"
)

// A function that returns an integer val.
type Hash func([]byte) (val int)

type HashSet []Hash

// Returns a HashSet with size Hash's.
func NewHashSet(size int) HashSet {
	set := make([]Hash, size)

	for i := 0; i < size; i++ {
		set[i] = NewHashFunction(size)
	}
	return set
}

// Returns a Hash that returns the md5hash value of addr plus a set
// random number modulo size.
func NewHashFunction(size int) Hash {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	
	f := func(addr []byte) (i int) {
		h := hashValue(addr)
		log.Println(h, r.Int(), size)
		return (h + r.Int()) % size
	}
	return f
}

// Returns the md5 checksum
func hashValue(buf []byte) int {
	//h := md5.New()
	//return h.Sum(data)
	return 10
}
