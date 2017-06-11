package gorocksdb


// WriteOptions represent all of the available options when writing to a
// database.
type WriteOptions struct {
	sync 			bool
	disableWAL 		bool
	ignore_missing_column_families bool
	no_slowdown		bool
	low_pri			bool

}

// NewDefaultWriteOptions creates a default WriteOptions object.
func NewDefaultWriteOptions() *WriteOptions {
	return &WriteOptions{}
}

// NewNativeWriteOptions creates a WriteOptions object.
/*func NewNativeWriteOptions(c *C.rocksdb_writeoptions_t) *WriteOptions {
	return &WriteOptions{c}
}
*/
// SetSync sets the sync mode. If true, the write will be flushed
// from the operating system buffer cache before the write is considered complete.
// If this flag is true, writes will be slower.
// Default: false
/*func (opts *WriteOptions) SetSync(value bool) {
	C.rocksdb_writeoptions_set_sync(opts.c, boolToChar(value))
}*/

// DisableWAL sets whether WAL should be active or not.
// If true, writes will not first go to the write ahead log,
// and the write may got lost after a crash.
// Default: false
/*func (opts *WriteOptions) DisableWAL(value bool) {
	C.rocksdb_writeoptions_disable_WAL(opts.c, C.int(btoi(value)))
}*/

// Destroy deallocates the WriteOptions object.
func (opts *WriteOptions) Destroy() {
	
	opts = nil
}
