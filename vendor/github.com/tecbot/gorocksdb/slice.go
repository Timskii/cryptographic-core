package gorocksdb


// Slice is used as a wrapper for non-copy values
type Slice struct {
	data  string
	size  int
	freed bool
}

// NewSlice returns a slice with the given data.
func NewSlice(data string, size int) *Slice {
	return &Slice{data, size, false}
}

// Data returns the data of the slice.
func (s *Slice) Data() []byte {
	return []byte(s.data /*+ string(s.size)*/)
}

// Size returns the size of the data.
func (s *Slice) Size() int {
	return int(s.size)
}

// Free frees the slice data.
func (s *Slice) Free() {
	s.freed = true
}
