package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"time"
)

type Block struct {
	Version       int64
	PrevBlockHash []byte
	Hash          []byte
	TimeStamp     int64
	TargetBits    int64
	Nonce         int64
	MerkeRoot     []byte
	//Data          []byte
	Transactions []*Transaction
}

func (block *Block) Serialize() []byte {
	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)

	err := encoder.Encode(block)

	CheckErr(err)
	return buffer.Bytes()
}

func Deserialize(data []byte) *Block {
	if len(data) == 0 {
		fmt.Println("data is empty!")
		os.Exit(1)
	}

	decoder := gob.NewDecoder(bytes.NewReader(data))

	var block Block
	err := decoder.Decode(&block)
	CheckErr(err)

	return &block

}

func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{
		Version:       1,
		PrevBlockHash: prevBlockHash,
		TimeStamp:     time.Now().Unix(),
		TargetBits:    targetBits,
		Nonce:         0,
		MerkeRoot:     []byte{},
		//Data:          []byte(data)}
		Transactions: transactions}
		//block.SetHash()

		pow := NewProofOfWork(block)
		nonce, hash := pow.Run()
		block.Nonce = nonce
		block.Hash = hash
		return block
	}

	// int -> byte
	/*
	func (block *Block)SetHash()  {
		tmp := [][]byte{
			Int2Byte(block.Version),
			block.PrevBlockHash,
			Int2Byte(block.TimeStamp),
			block.MerkeRoot,
			Int2Byte(block.Nonce),
			block.Data}
		data := bytes.Join(tmp, []byte{})
		hash := sha256.Sum256(data)
		block.Hash = hash[:]
	}
	*/

func NewGenesisBlock(coinbase *Transaction) *Block {
	//return NewBlock("Genesis Block!", []byte{})
	return NewBlock([]*Transaction{coinbase}, []byte{})
}
