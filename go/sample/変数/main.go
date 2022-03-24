package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 明示的な定義
	// var 変数名 型 = 値
	var i int = 100
	var i2 int64 = 200
	fmt.Println(i)
	// 型変換
	fmt.Println(i + int(i2))
	// 型表示
	fmt.Printf("%T\n", i)

	var zero float64 = 0.0
	pinf := 1.0 / zero
	ninf := -1.0 / zero
	// 正の無限大
	fmt.Println(pinf)
	// 負の無限大
	fmt.Println(ninf)
	// 正の整数
	var ui uint = 10
	fmt.Println(ui)
	// これはエラーになる
	// ui = -10
	// 複素数
	var comp complex64 = 2 + 5i
	fmt.Println(comp)
	// 文字列
	var s string = "Hello World"
	fmt.Println(s)
	convi, _ := strconv.Atoi("10")
	fmt.Printf("%T\n", convi)
	// byte型
	// 文字列をバイト配列
	var s2 []byte = []byte("Hellow World")
	fmt.Println(s2)
	// バイト配列を文字列
	fmt.Println(string(s2))

	// 配列
	// 要素数は変更できない
	var arr [2]int = [2]int{10, 100}
	fmt.Println(arr)
	// for分で配列処理
	for index := 0; index < len(arr); index++ {
		fmt.Println(arr[index])
	}
	// 同じ型をまとめて宣言
	var b1, b2 bool = true, false
	fmt.Println(b1, b2)

	// 違う型をまとめて宣言
	var (
		number int    = 0
		char   string = "hoge"
	)
	fmt.Println(number, char)
	// 暗黙的な定義
	// 変数名 := 値
	foo := 100
	fmt.Println(foo)

	// interface型
	// どんな型も受け付ける
	var x interface{} = 0
	fmt.Println(x)
	x = "string"
	fmt.Println(x)
}
