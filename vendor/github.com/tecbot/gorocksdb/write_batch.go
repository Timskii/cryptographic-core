package gorocksdb

import "io"
import "encoding/json"
import (
	"github.com/hyperledger/fabric/util"
	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("writeBatch")

// WriteBatch is a batching of Puts, Merges and Deletes.
type WriteBatch struct {
	}

type ColumnFamilyHandle struct {
	Type int
}

const (
	BLOCKCHAIN 	int = 0
	STATE		int = 1
	STATEDELTA	int = 2
	INDEXES		int = 3
)

// NewWriteBatch create a WriteBatch object.
func NewWriteBatch() *WriteBatch {
	return  &WriteBatch{}
}

// WriteBatchFrom creates a write batch from a serialized WriteBatch.
func WriteBatchFrom(data []byte) *WriteBatch {
	return &WriteBatch{}
}

// Put queues a key-value pair.
func (wb *WriteBatch) Put(key, value []byte) {
	logger.Debugf("Put \n")
}

// PutCF queues a key-value pair in a column family.
func (wb *WriteBatch) PutCF(cf *ColumnFamilyHandle, key, value []byte) {
	logger.Debugf("\nWriteBatch PutCF value:%x \nWriteBatch PutCF key:%x\n",value,key)
	logger.Debugf("WriteBatch PutCF value:%s \nWriteBatch PutCF key:%v\n",value,key)
	dataJson := new(DataJson)
	dataJson.Key = key
	dataJson.Value = value
	data ,_ := json.Marshal(dataJson)
	//data = append(data,[]byte(",\n")...)
	util.PrintData(data,"db")
}

// Merge queues a merge of "value" with the existing value of "key".
func (wb *WriteBatch) Merge(key, value []byte) {
	logger.Debugf("gorocksdb/writeBatch Merge \n")
}


// Delete queues a deletion of the data at key.
func (wb *WriteBatch) Delete(key []byte) {

	logger.Debugf("gorocksdb/writeBatch Delete \n")
}

// DeleteCF queues a deletion of the data at key in a column family.
func (wb *WriteBatch) DeleteCF(cf *ColumnFamilyHandle, key []byte) {

	logger.Debugf("gorocksdb/writeBatch DeleteCF \n")
}

// Data returns the serialized version of this batch.
func (wb *WriteBatch) Data() []byte {

	return []byte("data")
}

// Count returns the number of updates in the batch.
func (wb *WriteBatch) Count() int {
	return 10
}

// NewIterator returns a iterator to iterate over the records in the batch.
func (wb *WriteBatch) NewIterator() *WriteBatchIterator {
	data := wb.Data()
	if len(data) < 8+4 {
		return &WriteBatchIterator{}
	}
	return &WriteBatchIterator{data: data[12:]}
}


// Destroy deallocates the WriteBatch object.
func (wb *WriteBatch) Destroy() {

	
}

// WriteBatchRecordType describes the type of a batch record.
type WriteBatchRecordType byte

// Types of batch records.
const (
	WriteBatchRecordTypeDeletion WriteBatchRecordType = 0x0
	WriteBatchRecordTypeValue    WriteBatchRecordType = 0x1
	WriteBatchRecordTypeMerge    WriteBatchRecordType = 0x2
	WriteBatchRecordTypeLogData  WriteBatchRecordType = 0x3
)

// WriteBatchRecord represents a record inside a WriteBatch.
type WriteBatchRecord struct {
	Key   []byte
	Value []byte
	Type  WriteBatchRecordType
}

// WriteBatchIterator represents a iterator to iterator over records.
type WriteBatchIterator struct {
	data   []byte
	record WriteBatchRecord
	err    error
}

// Next returns the next record.
// Returns false if no further record exists.
func (iter *WriteBatchIterator) Next() bool {
	if iter.err != nil || len(iter.data) == 0 {
		return false
	}
	// reset the current record
	iter.record.Key = nil
	iter.record.Value = nil

	// parse the record type
	recordType := WriteBatchRecordType(iter.data[0])
	iter.record.Type = recordType
	iter.data = iter.data[1:]

	// parse the key
	x, n := iter.decodeVarint(iter.data)
	if n == 0 {
		iter.err = io.ErrShortBuffer
		return false
	}
	k := n + int(x)
	iter.record.Key = iter.data[n:k]
	iter.data = iter.data[k:]

	// parse the data
	if recordType == WriteBatchRecordTypeValue || recordType == WriteBatchRecordTypeMerge {
		x, n := iter.decodeVarint(iter.data)
		if n == 0 {
			iter.err = io.ErrShortBuffer
			return false
		}
		k := n + int(x)
		iter.record.Value = iter.data[n:k]
		iter.data = iter.data[k:]
	}
	return true
}

// Record returns the current record.
func (iter *WriteBatchIterator) Record() *WriteBatchRecord {
	return &iter.record
}

// Error returns the error if the iteration is failed.
func (iter *WriteBatchIterator) Error() error {
	return iter.err
}

func (iter *WriteBatchIterator) decodeVarint(buf []byte) (x uint64, n int) {
	// x, n already 0
	for shift := uint(0); shift < 64; shift += 7 {
		if n >= len(buf) {
			return 0, 0
		}
		b := uint64(buf[n])
		n++
		x |= (b & 0x7F) << shift
		if (b & 0x80) == 0 {
			return x, n
		}
	}
	// The number is too large to represent in a 64-bit value.
	return 0, 0
}
