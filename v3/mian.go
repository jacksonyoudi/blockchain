package main

import (
	"fmt"
)

func main() {
	// 实例化一个区块链
	fmt.Println("begin")
	bc := NewBlockChain()

	//添加block
	bc.AddBlock("测试第一个BTC")
	bc.AddBlock("测试第二个EOS")

	//打印出信息
	for i, block := range bc.blocks {
		fmt.Println("=========block num:", i)
		fmt.Println("data", string(block.Data))
		fmt.Println("Version:", block.Version)
		fmt.Printf("Hash:%x\n", block.Hash)
		fmt.Printf("TimeStamp:%d\n", block.TimeStamp)
		fmt.Printf("MerkeRoot:%x\n", block.MerkeRoot)
		fmt.Printf("Nonce:%d\n", block.Nonce)

		pow := NewProofOfWork(block)
		fmt.Println("Vaild:", pow.IsVaild())
	}
}
