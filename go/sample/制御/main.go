package main

import (
	"fmt"
	"strconv"
	"time"
)

func testDefer() {
	defer fmt.Println("func End")
	fmt.Println("func Start")
}
func SubRoutin() {
	for {
		fmt.Println("Sub loop")
		time.Sleep(1000 * time.Millisecond)
	}
}

// init 関数
// main よりも先に実行される特別な関数
func init() {
	fmt.Println("init")
}
func main() {
	// if基本
	hoge := 0
	if hoge == 0 {
		fmt.Println("zero")
	} else if hoge == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("other")
	}
	// 簡易文
	if x := 10; x == 10 {
		fmt.Println("x is 10")
	}

	// エラーハンドリング
	// 文字列からintに変換
	var s = "hoge"
	if i, err := strconv.Atoi(s); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	// for
	for in := 0; in < 10; in++ {
		// continue
		if in == 4 {
			continue
		}
		if in == 6 {
			break
		}
		fmt.Println(in)
	}
	// 無限ループもfor
	// for {
	// 	fmt.Println("loop")
	// }
	// 条件を満たした時break
	o := 0
	for o < 10 {
		fmt.Println(o)
		o++
	}

	var arr [3]int = [3]int{10, 100, 1000}
	fmt.Println(arr)
	// for分で配列処理
	for index := 0; index < len(arr); index++ {
		fmt.Println(arr[index])
	}
	// indexとvalueのペアを取り出す
	for index, value := range arr {
		fmt.Println(index, value)
	}
	// ラベル付きfor
	// break時にループをどこまで抜けるか指定
Loop:
	for {
		for {
			for {
				fmt.Println("Start")
				break Loop
			}
			fmt.Println("何らかの処理")
		}
		fmt.Println("何らかの処理")
	}
	fmt.Println("END")

	// switch　式スイッチ
	n := 0
	switch n {
	case 0, 1:
		fmt.Println("0or1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("unknown")
	}

	switch {
	case n >= 0 && n < 2:
		fmt.Println("n is 0 or 1")
	case n >= 2 && n < 4:
		fmt.Println("n is 2 or 3")
	default:
		fmt.Println("unknown")
	}
	// switch 型スイッチ
	var x interface{} = 10
	x = "hoge"
	// if版
	if x == nil {
		fmt.Println("x is nil")
	} else if intx, isInt := x.(int); isInt {
		fmt.Println(intx)
	} else if stringx, isString := x.(string); isString {
		fmt.Println(stringx)
	} else {
		fmt.Printf("unknown")
	}
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
	// 値を受け取る
	switch v := x.(type) {
	case bool:
		fmt.Println(v, "bool")
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	}

	// defer
	// 関数の実行後に処理する関数を登録出来る
	testDefer()

	// goroutin
	go SubRoutin()
	for {
		fmt.Println("Main loop")
		time.Sleep(2000 * time.Millisecond)
	}
}
