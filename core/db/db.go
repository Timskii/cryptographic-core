/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package db

import (
	"io"
	"os"

	"github.com/op/go-logging"
	"github.com/tecbot/gorocksdb"
	"bytes"
)

var dbLogger = logging.MustGetLogger("db")

const blockchainCF = "blockchainCF"
const stateCF = "stateCF"
const stateDeltaCF = "stateDeltaCF"
const indexesCF = "indexesCF"
const persistCF = "persistCF"

var columnfamilies = []string{
	blockchainCF, // blocks of the block chain
	stateCF,      // world state
	stateDeltaCF, // open transaction state
	indexesCF,    // tx uuid -> blockno
	persistCF,    // persistent per-peer state (consensus)
}

// OpenchainDB encapsulates rocksdb's structures
type OpenchainDB struct {
	DB           *gorocksdb.DB
	BlockchainCF *gorocksdb.ColumnFamilyHandle
	StateCF      *gorocksdb.ColumnFamilyHandle
	StateDeltaCF *gorocksdb.ColumnFamilyHandle
	IndexesCF    *gorocksdb.ColumnFamilyHandle
	PersistCF    *gorocksdb.ColumnFamilyHandle
}

var openchainDB = create()

// Create create an openchainDB instance
func create() *OpenchainDB {
	return &OpenchainDB{}
}

// GetDBHandle gets an opened openchainDB singleton. Note that method Start must always be invoked before this method.
func GetDBHandle() *OpenchainDB {
	return openchainDB
}


// GetFromBlockchainCF get value for given key from column family - blockchainCF
func (openchainDB *OpenchainDB) GetFromBlockchainCF(key []byte) ([]byte, error) {
	openchainDB.BlockchainCF = &gorocksdb.ColumnFamilyHandle{}
	if bytes.Equal(key,[]byte("blockCount")){openchainDB.BlockchainCF.Type = 1}else{openchainDB.BlockchainCF.Type = gorocksdb.BLOCKCHAIN}
	return openchainDB.Get(openchainDB.BlockchainCF, key,0)
}

// GetFromBlockchainCFSnapshot get value for given key from column family in a DB snapshot - blockchainCF
func (openchainDB *OpenchainDB) GetFromBlockchainCFSnapshot(snapshot *gorocksdb.Snapshot, key []byte) ([]byte, error) {
	return openchainDB.getFromSnapshot(snapshot, openchainDB.BlockchainCF, key)
}

// GetFromStateCF get value for given key from column family - stateCF
func (openchainDB *OpenchainDB) GetFromStateCF(key []byte) ([]byte, error) {
	openchainDB.StateCF = &gorocksdb.ColumnFamilyHandle{}
	openchainDB.StateCF.Type = gorocksdb.STATE
	return openchainDB.Get(openchainDB.StateCF, key,0)
}

func (openchainDB *OpenchainDB) GetFromStateCFForBlockNumber(key []byte,blockNumber int) ([]byte, error) {
	openchainDB.StateCF = &gorocksdb.ColumnFamilyHandle{}
	openchainDB.StateCF.Type = gorocksdb.STATE
	return openchainDB.Get(openchainDB.StateCF, key,blockNumber)
}

// GetFromStateDeltaCF get value for given key from column family - stateDeltaCF
func (openchainDB *OpenchainDB) GetFromStateDeltaCF(key []byte) ([]byte, error) {
	openchainDB.StateDeltaCF = &gorocksdb.ColumnFamilyHandle{}
	openchainDB.StateDeltaCF.Type = gorocksdb.STATEDELTA
	return openchainDB.Get(openchainDB.StateDeltaCF, key,0)
}

// GetFromIndexesCF get value for given key from column family - indexCF
func (openchainDB *OpenchainDB) GetFromIndexesCF(key []byte) ([]byte, error) {
	openchainDB.IndexesCF = &gorocksdb.ColumnFamilyHandle{}
	openchainDB.IndexesCF.Type = gorocksdb.INDEXES
	return openchainDB.Get(openchainDB.IndexesCF, key,0)
}

// GetBlockchainCFIterator get iterator for column family - blockchainCF
func (openchainDB *OpenchainDB) GetBlockchainCFIterator() *gorocksdb.Iterator {
	return openchainDB.GetIterator(openchainDB.BlockchainCF)
}

// GetStateCFIterator get iterator for column family - stateCF
func (openchainDB *OpenchainDB) GetStateCFIterator() *gorocksdb.Iterator {
	return openchainDB.GetIterator(openchainDB.StateCF)
}

// GetStateCFSnapshotIterator get iterator for column family - stateCF. This iterator
// is based on a snapshot and should be used for long running scans, such as
// reading the entire state. Remember to call iterator.Close() when you are done.
func (openchainDB *OpenchainDB) GetStateCFSnapshotIterator(snapshot *gorocksdb.Snapshot) *gorocksdb.Iterator {
	return openchainDB.getSnapshotIterator(snapshot, openchainDB.StateCF)
}

// GetStateDeltaCFIterator get iterator for column family - stateDeltaCF
func (openchainDB *OpenchainDB) GetStateDeltaCFIterator() *gorocksdb.Iterator {
	return openchainDB.GetIterator(openchainDB.StateDeltaCF)
}

// GetSnapshot returns a point-in-time view of the DB. You MUST call snapshot.Release()
// when you are done with the snapshot.
func (openchainDB *OpenchainDB) GetSnapshot() *gorocksdb.Snapshot {
	return nil
}

// DeleteState delets ALL state keys/values from the DB. This is generally
// only used during state synchronization when creating a new state from
// a snapshot.
func (openchainDB *OpenchainDB) DeleteState() error {
	return nil
}

// Get returns the valud for the given column family and key
func (openchainDB *OpenchainDB) Get(cfHandler *gorocksdb.ColumnFamilyHandle, key []byte, blockNumber int) ([]byte, error) {

	slice, _ := openchainDB.DB.GetCF( cfHandler, key, blockNumber)
	if slice.Data() == nil {
		return nil, nil
	}
	data := makeCopy(slice.Data())
	if len(data) < 1 {
		return nil,nil
	}
	return data, nil	// TIM get from file
}

// Put saves the key/value in the given column family
func (openchainDB *OpenchainDB) Put(cfHandler *gorocksdb.ColumnFamilyHandle, key []byte, value []byte) error {
	err := openchainDB.DB.PutCF(cfHandler, key, value)
	if err != nil {
		dbLogger.Errorf("Error while trying to write key: %s", key)
		return err
	}
	return nil
}

// Delete delets the given key in the specified column family
func (openchainDB *OpenchainDB) Delete(cfHandler *gorocksdb.ColumnFamilyHandle, key []byte) error {

	return nil
}

func (openchainDB *OpenchainDB) getFromSnapshot(snapshot *gorocksdb.Snapshot, cfHandler *gorocksdb.ColumnFamilyHandle, key []byte) ([]byte, error) {

	return nil, nil
}

// GetIterator returns an iterator for the given column family
func (openchainDB *OpenchainDB) GetIterator(cfHandler *gorocksdb.ColumnFamilyHandle) *gorocksdb.Iterator {
	//opt := gorocksdb.NewDefaultReadOptions()
	//defer opt.Destroy()
	iter := openchainDB.DB.NewIteratorCF(/*opt, cfHandler*/)
	return iter
}

func (openchainDB *OpenchainDB) getSnapshotIterator(snapshot *gorocksdb.Snapshot, cfHandler *gorocksdb.ColumnFamilyHandle) *gorocksdb.Iterator {

	return nil
}

func dirMissingOrEmpty(path string) (bool, error) {
	dirExists, err := dirExists(path)
	if err != nil {
		return false, err
	}
	if !dirExists {
		return true, nil
	}

	dirEmpty, err := dirEmpty(path)
	if err != nil {
		return false, err
	}
	if dirEmpty {
		return true, nil
	}
	return false, nil
}

func dirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func dirEmpty(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err
}

func makeCopy(src []byte) []byte {
	dest := make([]byte, len(src))
	copy(dest, src)
	return dest
}
