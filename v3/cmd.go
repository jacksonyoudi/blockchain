package main

import "fmt"

func (cli *CLI)AddBlock(data string)  {
	cli.bc.AddBlock(data)
}


func (cli *CLI)PrintChain()  {
	bc := cli.bc
	it := bc.Iterator()
	for {
		// 取回当前hash指定的block，并且将当前的hash指向上一个区块的hash
		block := it.Next()

		fmt.Println("data:", string(block.Data))
		fmt.Println("Version:", block.Version)
		fmt.Printf("Hash:%x\n", block.Hash)
		fmt.Printf("TimeStamp:%d\n", block.TimeStamp)
		fmt.Printf("MerkeRoot:%x\n", block.MerkeRoot)
		fmt.Printf("Nonce:%d\n", block.Nonce)
		fmt.Printf("preblock Hash:%x\n", block.PrevBlockHash)
		println("  ")

		pow := NewProofOfWork(block)
		fmt.Println("Vaild:", pow.IsVaild())

		if len(block.PrevBlockHash) == 0 {
			break
		}


	}
}