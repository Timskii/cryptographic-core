package core

import (
	"fmt"
	"hash/fnv"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/core/ledger/util"
	ut"github.com/hyperledger/fabric/core/util"
)

func (config *config) computeBucketHash(data []byte) uint32 {
	fmt.Printf("config data = %s\n", data)
	return config.hashFunc(data)
}

type config struct {
	maxGroupingAtEachLevel int
	lowestLevel            int
	levelToNumBucketsMap   map[int]int
	hashFunc               hashFunc
}
type hashFunc func(data []byte) uint32

type hash struct {
	hashingData        []byte
}

type bucketKey struct {
	level        int
	bucketNumber int
}
type dataKey struct {
	bucketKey    *bucketKey
	compositeKey []byte
}

func (key *dataKey) getEncodedBytes() []byte {
	encodedBytes := encodeBucketNumber(key.bucketKey.bucketNumber)
	encodedBytes = append(encodedBytes, key.compositeKey...)
	return encodedBytes
}
func encodeBucketNumber(bucketNumber int) []byte {
	return util.EncodeOrderPreservingVarUint64(uint64(bucketNumber))
}
func fnvHash(data []byte) uint32 {
	fnvHash := fnv.New32a()
	fnvHash.Write(data)
	return fnvHash.Sum32()
}
func (c *hash) appendSizeAndData(b []byte) {
	c.appendSize(len(b))
	c.hashingData = append(c.hashingData, b...)
}

func (c *hash) appendSize(size int) {
	c.hashingData = append(c.hashingData, proto.EncodeVarint(uint64(size))...)
}

func computeCryptoHash(chaincodeID,key string, value []byte)([]byte){
	hash := &hash{nil}
	hash.appendSizeAndData([]byte(chaincodeID))
	hash.appendSize(1)
	hash.appendSizeAndData([]byte(key))
	hash.appendSizeAndData(value)
	return ut.ComputeCryptoHash(hash.hashingData)
}