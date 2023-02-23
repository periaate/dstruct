package bloom

import (
	"encoding/binary"
	"fmt"
	"hash"
	"hash/fnv"
	"math"
)

const (
	// False positive rate
	fpRate = 0.01
)

type Filter struct {
	bitArray []byte
	hashFunc hash.Hash64
	k        int
}

func New(size int, k int) *Filter {
	byteSize := (findSize(size, fpRate) + 7) / 8
	return &Filter{
		bitArray: make([]byte, byteSize),
		hashFunc: fnv.New64(),
		k:        k,
	}
}

func (bf *Filter) Add(key []byte) {
	for i := 0; i < bf.k; i++ {
		bf.hashFunc.Reset()
		bf.hashFunc.Write(key)
		hashValue := bf.hashFunc.Sum64()
		index := hashValue % uint64(len(bf.bitArray)*8)
		byteIndex := index / 8
		bitIndex := index % 8
		bf.bitArray[byteIndex] |= (1 << bitIndex)
	}
}

func (bf *Filter) Test(key []byte) bool {
	for i := 0; i < bf.k; i++ {
		bf.hashFunc.Reset()
		bf.hashFunc.Write(key)
		hashValue := bf.hashFunc.Sum64()
		index := hashValue % uint64(len(bf.bitArray)*8)
		byteIndex := index / 8
		bitIndex := index % 8
		if bf.bitArray[byteIndex]&(1<<bitIndex) == 0 {
			return false
		}
	}
	return true
}

func (bf *Filter) Intersect(toCheck [][]byte) [][]byte {
	var res [][]byte
	for _, key := range toCheck {

		keyNum := binary.LittleEndian.Uint64(key)
		val := key
		if bf.Test(val) {
			res = append(res, key)
			fmt.Println("Found:", keyNum)
			continue
		}
		fmt.Println("Not found:", keyNum)
	}
	return res
}

func findSize(size int, fpRate float64) int {
	m := math.Ceil(-float64(size) * math.Log(fpRate) / math.Pow(math.Log(2), 2))
	return int(m)
}
