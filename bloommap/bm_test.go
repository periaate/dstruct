package bloommap_test

import (
	"dstruct/bloommap"
	"dstruct/util"
	"testing"
)

func TestHashTable(t *testing.T) {
	util.RunBaseTest(&bloommap.BloomMap{}, t, 0, 0)
}
