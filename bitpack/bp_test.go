package bitpack_test

import (
	"dstruct/bitpack"
	"dstruct/util"
	"fmt"
	"testing"
	"time"
)

const (
	size  uint32 = 100000
	start uint32 = 12
)

type timer struct {
	totalDuration time.Duration
	lastTime      time.Time
}

func (t *timer) count(s string, i int) {
	currentTime := time.Now()

	if t.lastTime.IsZero() {
		t.lastTime = currentTime
	}

	elapsedTime := currentTime.Sub(t.lastTime)
	t.totalDuration += elapsedTime
	avgDuration := t.totalDuration / time.Duration(i+1)

	fmt.Printf("%s — Total: %v — avg/item: %v\n", s, t.totalDuration, avgDuration)

	t.lastTime = currentTime
}

func TestHashTable(t *testing.T) {
	util.RunBaseTest(&bitpack.HashTable{}, t, size, start)
	util.RunBaseTest(&bitpack.HashTable{}, t, size*10, start)
	util.RunBaseTest(&bitpack.HashTable{}, t, size*100, start)

}
