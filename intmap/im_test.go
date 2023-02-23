package intmap_test

import (
	"dstruct/intmap"
	"dstruct/util"
	"testing"
)

const (
	size  uint32 = 100000
	start uint32 = 12
)

func TestHashTable(t *testing.T) {
	util.RunBaseTest(&intmap.HashTable{}, t, size, start)
	util.RunBaseTest(&intmap.HashTable{}, t, size*10, start)
	util.RunBaseTest(&intmap.HashTable{}, t, size*100, start)
}
