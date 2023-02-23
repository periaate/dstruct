package intmap

import (
	"dstruct/util"
	"math"
	"sync"
)

const (
	resizeThreshold = 0.65
	resizeMax       = 20
	resizeMin       = 2
)

var MAX = 0

type HashTable struct {
	entries [][2]uint32
	mutex   sync.Mutex
	count   int
	size    uint32
	hashKey func(uint32) uint32
}

func New(size uint32) *HashTable {
	return &HashTable{
		entries: make([][2]uint32, size),
		size:    size,
		hashKey: util.Hash,
	}
}

func (bm *HashTable) Init(size uint32) {
	bm.entries = make([][2]uint32, size)
	bm.size = size
	bm.hashKey = util.Hash
}

func (ht *HashTable) resize() {
	MaxJumps = 0
	old := ht.entries
	res := uint32(interpolate(ht.size))
	newSize := ht.size * res
	ht.entries = make([][2]uint32, newSize)

	ht.size = newSize
	ht.count = 0

	// rehash existing entries into the new table
	for _, v := range old {
		if v[0] != 0 {
			index := ht.hashKey(v[0]) % ht.size
			for {
				if ht.entries[index][0] == 0 {
					ht.entries[index] = [2]uint32{v[0], v[1]}
					ht.count++
					break
				}
				index = (index + 1) % ht.size
			}
		}
	}
	defer ht.mutex.Unlock()
}

func interpolate(value uint32) float64 {
	var max uint32 = 1000000
	var maxv float64 = resizeMin
	var minv float64 = resizeMax
	if value >= max {
		return maxv
	}

	minLog := math.Log10(float64(max))
	logFactor := 0 - minLog

	logValue := (math.Log10(float64(value)) - minLog) / logFactor

	linearFactor := (maxv - minv) / (1.0 - 0.0)
	res := minv + (linearFactor * (1.0 - logValue))
	return res
}

var index uint32
var MaxJumps = 0
var JumpAvg float64 = 0.0

func (ht *HashTable) Set(key uint32, value uint32) {
	ht.mutex.Lock()
	jumps := 0
	index = ht.hashKey(key) % ht.size
	for {
		if ht.entries[index][0] == 0 {
			ht.entries[index] = [2]uint32{key, value}
			ht.count++

			// update stats
			if jumps > MaxJumps {
				MaxJumps = jumps
			}
			JumpAvg += float64(jumps) / float64(MAX)

			break
		}

		if ht.entries[index][0] == key {
			ht.entries[index] = [2]uint32{key, value}
			ht.mutex.Unlock()

			// update stats
			if jumps > MaxJumps {
				MaxJumps = jumps
			}
			JumpAvg += float64(jumps) / float64(MAX)
			return
		}
		index = (index + 1) % ht.size
		jumps++
	}

	// check if the table needs to be resized
	if float64(ht.count)/float64(ht.size) > resizeThreshold {
		ht.resize()
		return
	}
	ht.mutex.Unlock()
}

var getIndex uint32

func (ht *HashTable) Get(key uint32) (uint32, bool) {
	getIndex = ht.hashKey(key) % ht.size
	for {
		entry := ht.entries[getIndex]
		// If an empty slot is found, the key is not in the table
		if entry[0] == 0 {
			break
		}

		if entry[0] == key {
			return entry[1], true
		}
		getIndex = (getIndex + 1) % ht.size
	}
	return 0, false
}
