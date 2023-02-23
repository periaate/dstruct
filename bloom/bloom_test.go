package bloom

import (
	"testing"
)

func TestBloom(t *testing.T) {
	// i := findSize(14, 0.03)
	// 	fmt.Println(i / 8)
	// 	return
	// 	p := message.NewPrinter(language.English)
	// 	fmt.Printf("Testing hash table — Size: %s\n", p.Sprintf("%v", size))

	// 	keys := make([]uint32, size)

	// 	keyGen := timer{}
	// 	keyGen.lastTime = time.Now()
	// 	fmt.Print("Generating keys")
	// 	for i := uint32(0); i < size; i++ {
	// 		keys[i] = util.HashXx(i)

	// 		if i%(size/10) == 0 {
	// 			fmt.Printf(".")
	// 		}

	// 	}
	// 	fmt.Printf("OK!")
	// 	keyGen.count("", size*2)

	// 	// // time how long it takes to set the keys
	// 	// iterationTimer := timer{}
	// 	// iterationTimer.lastTime = time.Now()

	// 	// //iteratation test. get all after every insertion
	// 	// fmt.Print(" Iterating keys")
	// 	// for i := 0; i < len(keys); i++ {
	// 	// 	ht.Set(keys[i], keys[i])
	// 	// 	for v := range keys {
	// 	// 		_ = v
	// 	// 	}
	// 	// }
	// 	// fmt.Printf("OK!")
	// 	// iterationTimer.count("", 1250025000)
	// 	// return

	// 	bitpack.MAX = size
	// 	ht := bitpack.New(start)
	// 	total := timer{}
	// 	total.lastTime = time.Now()

	// 	// time how long it takes to set the keys
	// 	setTimer := timer{}
	// 	setTimer.lastTime = time.Now()

	// 	fmt.Print("   Setting keys")
	// 	for i := range keys {
	// 		ht.Set(keys[i], keys[i])

	// 		if i%int(size/10) == 0 {
	// 			fmt.Printf(".")
	// 		}
	// 	}
	// 	fmt.Printf("OK!")
	// 	setTimer.count("", size)

	// 	getTimer := timer{}
	// 	getTimer.lastTime = time.Now()
	// 	fmt.Print("   Getting keys")
	// 	for i := range keys {
	// 		v, ok := ht.Get(keys[i])
	// 		if !ok || v != keys[i] {
	// 			fmt.Printf("Key %v not found", keys[i])
	// 			return
	// 		}

	// 		if i%int(size/10) == 0 {
	// 			fmt.Printf(".")
	// 		}
	// 	}
	// 	fmt.Printf("OK!")
	// 	getTimer.count("", size)

	// 	total.count("Total", size*2)
	// 	fmt.Println(bitpack.MaxJumps, bitpack.JumpAvg)

	// 	// bf := bloom.New(100, 4)
	// 	// buf := new(bytes.Buffer)
	// 	// // int to bytes
	// 	// binary.Write(buf, binary.LittleEndian, uint64(1))
	// 	// bf.Add(buf.Bytes())
	// 	// buf.Reset()
	// 	// binary.Write(buf, binary.LittleEndian, uint64(2))
	// 	// bf.Add(buf.Bytes())
	// 	// buf.Reset()
	// 	// binary.Write(buf, binary.LittleEndian, uint64(3))
	// 	// bf.Add(buf.Bytes())
	// 	// buf.Reset()
	// 	// binary.Write(buf, binary.LittleEndian, uint64(4))
	// 	// bf.Add(buf.Bytes())
	// 	// buf.Reset()

	// 	// bar := [][]byte{}
	// 	// for i := 1; i < 11; i++ {
	// 	// 	buf := new(bytes.Buffer)
	// 	// 	binary.Write(buf, binary.LittleEndian, uint64(i))
	// 	// 	byt := buf.Bytes()
	// 	// 	bar = append(bar, byt)
	// 	// }

	// 	// a := bf.Intersect(bar)
	// 	// fmt.Println("Found:", len(a))
	// 	// res := []uint64{}
	// 	// for _, v := range a {
	// 	// 	res = append(res, binary.LittleEndian.Uint64(v))
	// 	// }

	// 	// fmt.Println("Values:", res)
}

// func findSize(size int, fpRate float64) int {
// 	m := math.Ceil(-float64(size) * math.Log(fpRate) / math.Pow(math.Log(2), 2))
// 	return int(m)
// }
