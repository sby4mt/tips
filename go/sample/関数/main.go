package main

import "fmt"

func Plus(x int, y int) int {
	return x + y
}

// 返り値が複数ある場合
func SumDif(x int, y int) (int, int) {
	sum := x + y
	dif := x - y
	return sum, dif
}

// 返却する変数を予め宣言する
func Sum(x int, y int) (result int) {
	result = x + y
	return
}

// ジェネレーター
func Increment() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	res := Plus(10, 20)
	fmt.Println(res)
	s, d := SumDif(10, 20)
	fmt.Println(s, d)
	// 返り値を破棄したい場合_
	s2, _ := SumDif(10, 20)
	fmt.Println(s2)
	s3 := Sum(10, 20)
	fmt.Println(s3)

	// 無名関数
	f := func(x, y int) int {
		return x + y
	}
	fmt.Println(f(20, 30))
	// 関数宣言しながら引数を渡せる
	f1 := func(x, y int) int {
		return x + y
	}(2, 5)

	fmt.Println(f1)

	generater := Increment()
	fmt.Println(generater())
	fmt.Println(generater())
	fmt.Println(generater())
	fmt.Println(generater())
}
