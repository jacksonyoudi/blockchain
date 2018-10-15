package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

const targetBits  = 8

type ProofOfWork struct {
	block *Block
	targetBit *big.Int

}

func NewProofOfWork(block *Block) *ProofOfWork  {
	// 设置64位全1
	var IntTarget = big.NewInt(1)
	//00000000000000000000000000001
	//10000000000000000000000000000
	//00000000000100000000000000000
	//0000001
	// 右移 targetBits位
	IntTarget.Lsh(IntTarget, uint(256 - targetBits))

	return &ProofOfWork{block:block, targetBit:IntTarget}
}

func (pow *ProofOfWork)PrepareRawData(nonce int64)[]byte  {

	block := pow.block
	tmp := [][]byte{
		Int2Byte(block.Version),
		block.PrevBlockHash,
		Int2Byte(block.TimeStamp),
		block.MerkeRoot,
		Int2Byte(nonce),
		Int2Byte(targetBits),
		block.Data}

	data := bytes.Join(tmp, []byte{})
	return data
}

func (pow *ProofOfWork)Run() (int64, []byte) {

	var nonce int64
	var hash [32]byte
	var HashInt big.Int
	for nonce < math.MaxInt64 {
		data := pow.PrepareRawData(nonce)
		hash = sha256.Sum256(data)

		HashInt.SetBytes(hash[:])
		//fmt.Println(nonce)
		// 这里用于 判断算出的hash值（int）只要比最大的IntTarget小就是正确的。
		if HashInt.Cmp(pow.targetBit) == -1 {
			fmt.Printf("Found Hash: %x\n", hash)
			break
		} else {
			nonce++
		}

	}
	return nonce, hash[:]
}

// 对block的数据校验
func (pow *ProofOfWork)IsVaild() bool {
	data := pow.PrepareRawData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	var IntHash big.Int
	IntHash.SetBytes(hash[:])
	return IntHash.Cmp(pow.targetBit) == -1

}