package main

import (
	"dstruct/bitpack"
	"dstruct/bloommap"
	"dstruct/intmap"
	"dstruct/util"
	"fmt"
)

var size uint32 = 100000

func main() {
	fmt.Println("Bloom map:")
	util.RunBase(&bloommap.BloomMap{}, size, 0) // Too slow to test with size*10 and size*100

	// With optimizations off this is significantly slower than int map.
	fmt.Println("\nBit pack:")
	util.RunBase(&bitpack.HashTable{}, size, 0)
	util.RunBase(&bitpack.HashTable{}, size*10, 0)
	util.RunBase(&bitpack.HashTable{}, size*100, 0)

	fmt.Println("\nInt map:")
	util.RunBase(&intmap.HashTable{}, size, 0)
	util.RunBase(&intmap.HashTable{}, size*10, 0)
	util.RunBase(&intmap.HashTable{}, size*100, 0)
}
