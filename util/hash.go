package util

import (
	"encoding/binary"
	"hash/crc32"
	"hash/fnv"
)

// var buf bytes.Buffer
var fnv32 = fnv.New32a()
var fnvBytes = make([]byte, 4)

func HashFnv(key uint32) uint32 {
	fnv32.Reset()
	binary.LittleEndian.PutUint32(fnvBytes, key)
	fnv32.Write(fnvBytes)
	return fnv32.Sum32()
}

// var buf bytes.Buffer
var crcBytes = make([]byte, 4)

func HashCrc(key uint32) uint32 {
	binary.BigEndian.PutUint32(crcBytes, key)
	return uint32(crc32.ChecksumIEEE(crcBytes))
}

const prime = 2654435761

var hash uint32

// Upon closer inspection this is a variant of MurmurHash3. It's fast so 🤷
func Hash(key uint32) uint32 {

	hash = prime + (key>>0)*prime
	hash = (hash << 13) | (hash >> 19)
	hash *= prime

	hash += (key >> 8) * prime
	hash = (hash << 13) | (hash >> 19)
	hash *= prime

	hash += (key >> 16) * prime
	hash = (hash << 13) | (hash >> 19)
	hash *= prime

	hash += (key >> 24) * prime
	hash = (hash << 13) | (hash >> 19)
	hash *= prime

	hash ^= 4
	hash ^= hash >> 16
	hash *= prime
	hash ^= hash >> 13
	hash *= prime
	hash ^= hash >> 16

	return hash
}
