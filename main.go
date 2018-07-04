package main

import "fmt"

func main() {
	//https://qiita.com/ruiu/items/28c77ed483cec365fe84
	intBitSize := 32 << (^uint(0) >> 63)
	fmt.Printf("int bit size: %d\n", intBitSize)

}
