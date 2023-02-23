package bloommap

import (
	"dstruct/util"
	"encoding/binary"
	"hash/fnv"
	"math"
	"sync"
)

const (
	bfLen = 8 * 8
)

const (
	resizeThreshold = 0.8
	resizeMax       = 20
	resizeMin       = 2
)

func interpolate(value uint32) float64 {
	var max uint32 = 10000000
	var maxv float64 = resizeMin
	var minv float64 = resizeMax
	if value >= max {
		return maxv
	}

	minLog := math.Log10(float64(max))
	logFactor := 0 - minLog

	loguint32 := (math.Log10(float64(value)) - minLog) / logFactor

	linearFactor := (maxv - minv) / (1.0 - 0.0)
	res := minv + (linearFactor * (1.0 - loguint32))
	return res
}

var MAX = 0

// 8 byte bloom filter, 4 byte key, 4 byte value
// bloom filter and (7) kv pairs
// first bit of first byte is bool for whether or not bucket is full
type entry [64]byte
type entries []entry

type BloomMap struct {
	entries     entries
	mutex       sync.Mutex
	currentSize uint32
	maxSize     uint32
	hashKey     func(uint32) uint32
}

func New(size uint32) *BloomMap {
	return &BloomMap{
		entries: make(entries, size),
		maxSize: size,
		hashKey: util.HashXx,
	}
}

func (bm *BloomMap) Init(size uint32) {
	bm.entries = make(entries, size)
	bm.maxSize = size
	bm.hashKey = util.HashXx
}

func intToBytes(value uint32) []byte {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, value)
	return bytes
}

func (bm *BloomMap) Set(key uint32, value uint32) {
	bm.mutex.Lock()

	index := bm.hashKey(key) % bm.maxSize
	for {
		if bm.entries[index].checkBloomFilter(intToBytes(key)) {
			// bloom filter contains key, check if key exists in kv pairs
			for i := 0; i < 7; i++ {
				if binary.LittleEndian.Uint32(bm.entries[index][8+i*8:12+i*8]) == key {
					// key exists in kv pairs, update value
					binary.LittleEndian.PutUint32(bm.entries[index][12+i*8:16+i*8], value)
					bm.mutex.Unlock()
					return
				}
			}
		}
		// bloom filter does not contain key, check if current bucket is full
		if bm.entries[index].isFull() {
			// current bucket is full, find next empty bucket or bucket with same key
			index = (index + 1) % bm.maxSize
			continue
		} else {
			// current bucket is not full, find first empty slot and add kv pair to that bucket
			for i := 0; i < 7; i++ {
				if binary.LittleEndian.Uint32(bm.entries[index][8+i*8:12+i*8]) == 0 {
					binary.LittleEndian.PutUint32(bm.entries[index][8+i*8:12+i*8], key)
					binary.LittleEndian.PutUint32(bm.entries[index][12+i*8:16+i*8], value)
				}
			}

			bm.entries[index].updateBloomFilter(intToBytes(key))
			bm.currentSize++
			bm.mutex.Unlock()
			if float64(bm.currentSize)/float64(bm.maxSize) > resizeThreshold {
				bm.resize()
			}
			return
		}

	}

}

var hashes = []byte{}
var indices = [4]uint32{}
var byteIndices = [4]uint32{}
var bitIndices = [4]uint32{}

func (en *entry) updateBloomFilter(key []byte) {
	hashes = fnv.New128().Sum(key)
	indices = [4]uint32{
		binary.LittleEndian.Uint32(hashes[:4]) % bfLen,
		binary.LittleEndian.Uint32(hashes[4:8]) % bfLen,
		binary.LittleEndian.Uint32(hashes[8:12]) % bfLen,
		binary.LittleEndian.Uint32(hashes[12:16]) % bfLen,
	}
	byteIndices = [4]uint32{
		indices[0] / 8,
		indices[1] / 8,
		indices[2] / 8,
		indices[3] / 8,
	}
	bitIndices = [4]uint32{
		indices[0] % 8,
		indices[1] % 8,
		indices[2] % 8,
		indices[3] % 8,
	}

	en[byteIndices[0]] |= 1 << bitIndices[0]
	en[byteIndices[1]] |= 1 << bitIndices[1]
	en[byteIndices[2]] |= 1 << bitIndices[2]
	en[byteIndices[3]] |= 1 << bitIndices[3]
}

func (en *entry) checkBloomFilter(key []byte) bool {
	hashes = fnv.New128().Sum(key)
	indices = [4]uint32{
		binary.LittleEndian.Uint32(hashes[:4]) % bfLen,
		binary.LittleEndian.Uint32(hashes[4:8]) % bfLen,
		binary.LittleEndian.Uint32(hashes[8:12]) % bfLen,
		binary.LittleEndian.Uint32(hashes[12:16]) % bfLen,
	}
	byteIndices = [4]uint32{
		indices[0] / 8,
		indices[1] / 8,
		indices[2] / 8,
		indices[3] / 8,
	}
	bitIndices = [4]uint32{
		indices[0] % 8,
		indices[1] % 8,
		indices[2] % 8,
		indices[3] % 8,
	}

	return (en[byteIndices[0]]&(1<<bitIndices[0])) != 0 && (en[byteIndices[1]]&(1<<bitIndices[1])) != 0 && (en[byteIndices[2]]&(1<<bitIndices[2])) != 0 && (en[byteIndices[3]]&(1<<bitIndices[3])) != 0
}

func (en *entry) isFull() bool {
	// check if last key value key is not 0
	return binary.LittleEndian.Uint32(en[8+6*8:12+6*8]) != 0
}

func (bm *BloomMap) Get(key uint32) (uint32, bool) {
	index := bm.hashKey(key) % bm.maxSize

	for {
		if !bm.entries[index].checkBloomFilter(intToBytes(key)) && !bm.entries[index].isFull() {
			return 0, false
		}

		for i := 0; i < 7; i++ {
			if binary.LittleEndian.Uint32(bm.entries[index][8+i*8:12+i*8]) == key {
				return binary.LittleEndian.Uint32(bm.entries[index][12+i*8 : 16+i*8]), true
			}
		}

		index = (index + 1) % bm.maxSize
	}

}

func (bm *BloomMap) resize() {
	oldLen := bm.maxSize
	bm.maxSize *= uint32(interpolate(bm.maxSize))
	oldEntries := bm.entries
	bm.currentSize = 0
	bm.entries = make(entries, bm.maxSize)

	for i := 0; i < int(oldLen); i++ {
		for j := 0; j < 7; j++ {
			if binary.LittleEndian.Uint32(oldEntries[i][8+j*8:12+j*8]) != 0 {
				bm.Set(binary.LittleEndian.Uint32(oldEntries[i][8+j*8:12+j*8]), binary.LittleEndian.Uint32(oldEntries[i][12+j*8:16+j*8]))
			} else {
				break
			}
		}
	}
}
