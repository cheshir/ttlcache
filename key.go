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

// IntKey creates key from int value.
func IntKey(k int) uint64 {
	return uint64(k)
}

// ByteKey creates key from byte value.
func ByteKey(k byte) uint64 {
	return uint64(k)
}

// Int8Key creates key from int8 value.
func Int8Key(k int8) uint64 {
	return uint64(k)
}

// Uint8Key creates key from uint8 value.
func Uint8Key(k uint8) uint64 {
	return uint64(k)
}

// Int16Key creates key from int16 value.
func Int16Key(k int16) uint64 {
	return uint64(k)
}

// Uint16Key creates key from uint16 value.
func Uint16Key(k uint16) uint64 {
	return uint64(k)
}

// Int32Key creates key from int32 value.
func Int32Key(k int32) uint64 {
	return uint64(k)
}

// Uint32Key creates key from uint32 value.
func Uint32Key(k uint32) uint64 {
	return uint64(k)
}

// Int64Key creates key from int64 value.
func Int64Key(k int64) uint64 {
	return uint64(k)
}

// Uint64Key creates key from uint64 value.
func Uint64Key(k uint64) uint64 {
	return uint64(k)
}

// BytesKey creates key from slice of bytes value.
func BytesKey(k []byte) uint64 {
	return newKeyFromBytes(k)
}

// StringKey creates key from string value.
func StringKey(k string) uint64 {
	return newKeyFromBytes([]byte(k))
}

// AnyKey creates key from anything.
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
