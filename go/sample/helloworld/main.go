package main

// パッケージ宣言は必須

// go標準パッケージ https://pkg.go.dev/std
import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(time.Now())
}

// GOPATH/src配下にプロジェクトを作成することで
// 独自のパッケージをインポート出来る
// go env を実行することでGOPAHTを確認できる
// set GOPATH=C:\Users\hoge\go
