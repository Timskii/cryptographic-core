package gorocksdb



// Snapshot provides a consistent view of read operations in a DB.
type Snapshot struct {

}

// Release removes the snapshot from the database's list of snapshots.
func (s *Snapshot) Release() {
}
