package core

import (
	"hash/fnv"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/core/ledger/util"
)

func (config *config) computeBucketHash(data []byte) uint32 {
	return config.hashFunc(data)
}
const MaxGroupingAtEachLevel = 5

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

func (bucketKey *bucketKey) getEncodedBytes() []byte {
	encodedBytes := []byte{}
	encodedBytes = append(encodedBytes, byte(0))
	encodedBytes = append(encodedBytes, proto.EncodeVarint(uint64(bucketKey.level))...)
	encodedBytes = append(encodedBytes, proto.EncodeVarint(uint64(bucketKey.bucketNumber))...)
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

func computeBucketNumber (bucketNumber int) (int){
	BucketNumber := bucketNumber/MaxGroupingAtEachLevel
	if bucketNumber%MaxGroupingAtEachLevel != 0 {
		BucketNumber++
	}
	return BucketNumber
}

func unmarshalCryptoHash(serializedBytes []byte) []byte {
	var unmarshalCryptoHash []byte
	buffer := proto.NewBuffer(serializedBytes)
	for i := 0; i < MaxGroupingAtEachLevel; i++ {
		cryptoHash, _ := buffer.DecodeRawBytes(false)
		if len(cryptoHash)> 10 {
			unmarshalCryptoHash = cryptoHash
			break
		}
	}
	return unmarshalCryptoHash
}

