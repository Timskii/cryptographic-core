package gorocksdb

//import (	"unsafe")

// Iterator provides a way to seek to specific keys and iterate through
// the keyspace from that point, as well as access the values of those keys.
//
// For example:
//
//      it := db.NewIterator(readOpts)
//      defer it.Close()
//
//      it.Seek([]byte("foo"))
//		for ; it.Valid(); it.Next() {
//          fmt.Printf("Key: %v Value: %v\n", it.Key().Data(), it.Value().Data())
// 		}
//
//      if err := it.Err(); err != nil {
//          return err
//      }
//
type Iterator struct {
	
}

// NewNativeIterator creates a Iterator object.
func NewNativeIterator(/*c unsafe.Pointer*/) *Iterator {
	return &Iterator{}
}

// Valid returns false only when an Iterator has iterated past either the
// first or the last key in the database.
func (iter *Iterator) Valid() bool {
	return true
}

// ValidForPrefix returns false only when an Iterator has iterated past the
// first or the last key in the database or the specified prefix.
func (iter *Iterator) ValidForPrefix(prefix []byte) bool {
	return true
}

// Key returns the key the iterator currently holds.
func (iter *Iterator) Key() *Slice {

	return &Slice{"zb76c3e8", 4, false}

}

// Value returns the value in the database the iterator currently holds.
func (iter *Iterator) Value() *Slice {
	return &Slice{"datadat", 4, false}
}

// Next moves the iterator to the next sequential key in the database.
func (iter *Iterator) Next() {

}

// Prev moves the iterator to the previous sequential key in the database.
func (iter *Iterator) Prev() {

}

// SeekToFirst moves the iterator to the first key in the database.
func (iter *Iterator) SeekToFirst() {

}

// Seek moves the iterator to the position greater than or equal to the key.
func (iter *Iterator) Seek(key []byte) {

}

// Close closes the iterator.
func (iter *Iterator) Close() {
	iter = nil
}
