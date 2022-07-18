package main

import "fmt"

func main() {
	pow := make([]int, 10)
	// "_"へ代入すると捨てることができる
	// for i, _ := range pow
	// for _, value := range pow
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	fmt.Println(pow)
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

// [1 2 4 8 16 32 64 128 256 512]
// 1
// 2
// 4
// 8
// 16
// 32
// 64
// 128
// 256
// 512
