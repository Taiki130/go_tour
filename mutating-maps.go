package main

import "fmt"

func main() {
	m := make(map[string]int)

	// 代入、更新
	// m[key] = elem
	// 要素の取得
	// elem = m[key]
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// 要素の削除
	// delete(m, key)
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// キーに対する要素が存在するかどうか
	// elem, ok = m[key]
	// ok: trueのとき要素は存在する
	// ok: falseのとき要素は存在しない
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
