package main

type Hash func() (val int)

type HashSet []Hash

// Returns a HashSet with size Hash's.
func NewHashSet(size int) HashSet {
	h := make([]Hash, size)

	for i := 0; i < size; i++ {
		f := func() (val int) {
			return i
		}
		h[i] = Hash(f)
	}
	return h
}
