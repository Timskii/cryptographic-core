package gorocksdb



// Snapshot provides a consistent view of read operations in a DB.
type Snapshot struct {

}

// NewNativeSnapshot creates a Snapshot object.
/*func NewNativeSnapshot(c *C.rocksdb_snapshot_t, cDb *C.rocksdb_t) *Snapshot {
	return &Snapshot{c, cDb}
}*/

// Release removes the snapshot from the database's list of snapshots.
func (s *Snapshot) Release() {
}
