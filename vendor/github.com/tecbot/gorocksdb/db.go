package gorocksdb

import(
    "fmt"
    "encoding/json"
    "io/ioutil"
	"strings"
	"bytes"

    "github.com/hyperledger/fabric/util"
	"github.com/hyperledger/fabric/protos"
)


// Range is a range of keys in the database. GetApproximateSizes calls with it
// begin at the key Start and end right before the key Limit.
type Range struct {
	Start []byte
	Limit []byte
}

// DB is a reusable handle to a RocksDB database on disk, created by Open.
type DB struct {
	name string
}

type DataJson struct {
	Key 	[]byte
	Value	[]byte
}
// GetCF returns the data associated with the key from the database and column family.
func (db *DB) GetCF( cf *ColumnFamilyHandle, key []byte, blockNumber int) (*Slice, error) {
	var (
		cValue  []byte
		cValLen int
		countNumber int
	)

	file, e := ioutil.ReadFile("db.txt")
	if e == nil {
		fileS := "[" + strings.TrimRight(string(file),",\n") + "]"
		fileS = strings.Replace(fileS,"}{","},{",-1)

		var jsontype []*DataJson
		block := &protos.Block{}
		json.Unmarshal([]byte(fileS), &jsontype)

		if blockNumber != 0 {
			for i := 0; i< len(jsontype);  i++ {
				if bytes.Equal(key, jsontype[i].Key) {
					countNumber++
					if  countNumber == (blockNumber) {
							cValue = jsontype[i].Value
							cValLen = len(jsontype[i].Value)
							break
					}
				}
			}
		}else {
			for i := len(jsontype) - 1; i > 0; i-- {
				if cf.Type == BLOCKCHAIN {
					block = nil
					json.Unmarshal(jsontype[i].Value, &block)
				}
				if bytes.Equal(key, jsontype[i].Key) {
					if (cf.Type == BLOCKCHAIN && block != nil) || cf.Type != BLOCKCHAIN {

						cValue = jsontype[i].Value
						cValLen = len(jsontype[i].Value)
						break
					}
				}
			}
		}
	}
	return NewSlice(string(cValue), cValLen), nil
}

// PutCF writes data associated with a key to the database and column family.
func (db *DB) PutCF(cf *ColumnFamilyHandle, key, value []byte) error {
	var dataJson DataJson
	dataJson.Key = key
	dataJson.Value = value
	util.PrintData([]byte(fmt.Sprintf("%+v\n", dataJson)),"db")
	return nil
}

// Write writes a WriteBatch to the database
func (db *DB) Write(opts *WriteOptions, batch *WriteBatch) error {
	return nil
}



// NewIteratorCF returns an Iterator over the the database and column family
// that uses the ReadOptions given.
func (db *DB) NewIteratorCF(/*opts *ReadOptions, cf *ColumnFamilyHandle*/) *Iterator {
	//cIter := C.rocksdb_create_iterator_cf(db.c, opts.c, cf.c)
	return NewNativeIterator()//unsafe.Pointer(cIter))
}
