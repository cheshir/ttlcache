package ttlcache

import (
	"fmt"
	"hash/crc64"
	"sync"
)

var hashPool = sync.Pool{
	New: func() interface{} {
		return crc64.MakeTable(crc64.ISO)
	},
}

func IntKey(k int) uint64 {
	return uint64(k)
}

func ByteKey(k byte) uint64 {
	return uint64(k)
}

func Int8Key(k int8) uint64 {
	return uint64(k)
}

func Uint8Key(k uint8) uint64 {
	return uint64(k)
}

func Int16Key(k int16) uint64 {
	return uint64(k)
}

func Uint16Key(k uint16) uint64 {
	return uint64(k)
}

func Int32Key(k int32) uint64 {
	return uint64(k)
}

func Uint32Key(k uint32) uint64 {
	return uint64(k)
}

func Int64Key(k int64) uint64 {
	return uint64(k)
}

func Uint64Key(k uint64) uint64 {
	return uint64(k)
}

func BytesKey(k []byte) uint64 {
	return newKeyFromBytes(k)
}

func StringKey(k string) uint64 {
	return newKeyFromBytes([]byte(k))
}

// Should not be used for large datasets.
// For complex keys you can write your own hashing implementation.
func AnyKey(k interface{}) uint64 {
	s := fmt.Sprintf("%#v", k)

	return newKeyFromBytes([]byte(s))
}

func newKeyFromBytes(k []byte) uint64 {
	hashTable := hashPool.Get().(*crc64.Table)
	hash := crc64.Checksum(k, hashTable)
	hashPool.Put(hashTable)

	return hash
}
