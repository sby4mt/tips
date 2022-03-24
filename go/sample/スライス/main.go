package main

import "fmt"

// スライス
func main() {
	// 配列の要素数をしていない
	var sl []int = []int{10, 20, 30}
	fmt.Println(sl)
	// スライスは要素数可変
	sl = append(sl, 40)
	fmt.Println(sl)
	sl = append(sl, 50, 51, 52)
	fmt.Println(sl)

	// make関数
	// スライスを生成してくれる
	sl2 := make([]int, 5)
	fmt.Println(sl2)
	fmt.Println(len(sl2))
	fmt.Println(cap(sl2))
	sl2 = append(sl2, 1)
	fmt.Println(cap(sl2))
	// makeの第3引数でキャパシティを指定できる
	sl2 = make([]int, 5, 10)
	fmt.Println(cap(sl2))
	// 要素数を超えるとキャパシティが倍増する
	// パフォーマンス
	sl2 = append(sl2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(cap(sl2))

	sl3 := []int{1, 2, 3, 4, 5}
	// 範囲を指定して取り出せる
	fmt.Println(sl3[2:4])
	fmt.Println(sl3[:4])
	fmt.Println(sl[2:])
}
