package main

import (
	"fmt"
)

func main() {
	// 实例化一个区块链
	fmt.Println("begin")
	bc := NewBlockChain()
	cli := CLI{bc:bc}
	cli.Run()
}
