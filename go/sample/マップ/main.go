package main

import "fmt"

func main() {
	// mapの明示的宣言
	var a = map[string]int{"a": 10, "b": 20}
	fmt.Println(a)
	fmt.Println(a["a"])

	// make を使ってマップを宣言
	m := make(map[int]string)
	fmt.Println(m)
	// 値の追加
	m[1] = "ja"
	m[3] = "en"
	fmt.Println(m)
	fmt.Println(m[3])
	// 2個目の返り値に値が存在するかのフラグが返ってくる
	v, exist := m[2]
	fmt.Println(v, exist)
	// 値の削除
	delete(m, 3)
	fmt.Println(m)

	m2 := map[int]string{0: "a", 1: "b", 2: "c"}
	for v, k := range m2 {
		fmt.Println(v, k)
	}
}
