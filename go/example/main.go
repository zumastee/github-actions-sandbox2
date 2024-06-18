package main

import "fmt"


var version string # ビルド時にidflas フラグ経由でバージョンを埋め込む

func main() {
	fmt.Println("Example: %s¥n", version)
}