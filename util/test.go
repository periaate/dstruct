package util

import (
	"fmt"
	"testing"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	size  uint32 = 100000
	start uint32 = 12
)

var keys []uint32

type hashmap interface {
	Set(key uint32, value uint32)
	Get(key uint32) (uint32, bool)
	Init(size uint32)
}

func RunBaseTest(hm hashmap, t *testing.T, defSize, defStart uint32) bool {
	if defSize != 0 {
		size = defSize
	}
	if defStart != 0 {
		start = defStart
	}
	GenerateKeys(int(defSize))
	hm.Init(size)
	return SetGet(hm, int(size), t)
}

func RunBase(hm hashmap, defSize, defStart uint32) bool {
	if defSize != 0 {
		size = defSize
	}
	if defStart != 0 {
		start = defStart
	}
	GenerateKeys(int(defSize))
	hm.Init(size)
	return SetGet(hm, int(size), &testing.T{})
}

func GenerateKeys(size int) {
	keys = make([]uint32, size)

	fmt.Print("Generating keys")
	for i := 0; i < size; i++ {
		keys[i] = HashXx(uint32(i))

		if i%(size/10) == 0 {
			fmt.Printf(".")
		}
	}
	fmt.Printf("OK!\n")
}

func SetGet(hm hashmap, size int, t *testing.T) bool {
	p := message.NewPrinter(language.English)
	fmt.Printf("Testing hash table — Size: %s\n", p.Sprintf("%v", size))

	total := timer{}
	total.lastTime = time.Now()

	// time how long it takes to set the keys
	setTimer := timer{}
	setTimer.lastTime = time.Now()

	fmt.Print("   Setting keys")
	for i := range keys {
		hm.Set(keys[i], keys[i])

		if i%int(size/10) == 0 {
			fmt.Printf(".")
		}
	}
	fmt.Printf("OK!")
	setTimer.count("", size)

	getTimer := timer{}
	getTimer.lastTime = time.Now()
	fmt.Print("   Getting keys")
	for i := range keys {
		v, ok := hm.Get(keys[i])
		if !ok || v != keys[i] {
			fmt.Printf("Key %v not found\n", keys[i])
			t.Fail()
		}

		if i%int(size/10) == 0 {
			fmt.Printf(".")
		}
	}
	fmt.Printf("OK!")
	getTimer.count("", size)

	total.count("Total", size*2)

	return true
}

func Iterate(hm hashmap, size int, t *testing.T) {
	iterationTimer := timer{}
	iterationTimer.lastTime = time.Now()
	//iteratation test. get all after every insertion
	j := 0
	fmt.Print(" Iterating keys")
	for i := 0; i < len(keys); i++ {
		hm.Set(keys[i], keys[i])
		for range keys {
			j++
		}
	}
	fmt.Printf("OK!")
	iterationTimer.count("", j)
}

func TestBuiltin(size int, t *testing.T) {
	if len(keys) == 0 {

		hmTimer := timer{}
		hmTimer.lastTime = time.Now()
		hm := make(map[uint32]uint32, start)
		fmt.Print(" Comparing maps\n")
		fmt.Print("   Setting keys")
		for i := 0; i < len(keys); i++ {
			hm[keys[i]] = keys[i]

			if i%int(size/10) == 0 {
				fmt.Printf(".")
			}
		}
		fmt.Printf("OK!")
		hmTimer.count("", size)

		hmTimer = timer{}
		hmTimer.lastTime = time.Now()
		fmt.Print("   Getting keys")
		for i := 0; i < len(keys); i++ {
			v, ok := hm[keys[i]]
			if !ok || v != keys[i] {
				fmt.Printf("Key %v not found", keys[i])
				return
			}

			if i%int(size/10) == 0 {
				fmt.Printf(".")
			}
		}
		fmt.Printf("OK!")
		hmTimer.count("", size)
	}
}
