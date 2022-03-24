package main

import "fmt"

// 定数
// パッケージ外からアクセスしたい場合は大文字から始める
const Hoge = "hoge"

// まとめて宣言可能
const (
	foo = "foo"
	poo = "poo"
)

func main() {
	fmt.Println(Hoge)
	fmt.Println(foo, poo)
}
