package main

import "fmt"

// slice
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

// range
// i: インデックス、v: 要素を返す
func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
