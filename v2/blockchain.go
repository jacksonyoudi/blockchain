package main

import "os"

type BlockChian struct {
	blocks []*Block
}

func NewBlockChain() *BlockChian {
	return &BlockChian{[]*Block{NewGenesisBlock()}}
}


func (bc *BlockChian)AddBlock(data string) {
	if len(bc.blocks) <= 0 {
		os.Exit(1)
	}
	lastBlock := bc.blocks[len(bc.blocks) -1]
	block := NewBlock(data, lastBlock.Hash)
	bc.blocks = append(bc.blocks, block)
}