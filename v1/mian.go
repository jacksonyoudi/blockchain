package v1

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"os"
	"time"
)

// 实现int转换成byte数组
func Int2Byte(num int64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	CheckErr(err)
	return buffer.Bytes()
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}
}

//区块的结构
type Block struct {
	Version       int64
	PrevBlockHash []byte
	Hash          []byte
	TimeStamp     int64
	TargetBits    int64
	Nonce         int64
	MerkeRoot     []byte
	Data          []byte
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Version:       1,
		PrevBlockHash: prevBlockHash,
		TimeStamp:     time.Now().Unix(),
		TargetBits:    10,
		Nonce:         5,
		MerkeRoot:     []byte{},
		Data:          []byte(data)}
	block.SetHash()

	return block

}

// int -> byte

func (block *Block) SetHash() {
	tmp := [][]byte{
		//实现将int转成byte数组的函数
		Int2Byte(block.Version),
		block.PrevBlockHash,
		Int2Byte(block.TimeStamp),
		block.MerkeRoot,
		Int2Byte(block.Nonce),
		block.Data}
	//将区块个字段连接成一个切片，使用[]byte{}进行连接

	data := bytes.Join(tmp, []byte{})

	//算出hash的值
	hash := sha256.Sum256(data)
	block.Hash = hash[:]
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block!", []byte{})
}

type BlockChian struct {
	//使用切片保存区块，用于模拟区块链
	blocks []*Block
}

func NewBlockChain() *BlockChian {
	//创建一个区块链
	return &BlockChian{[]*Block{NewGenesisBlock()}}
}

func (bc *BlockChian) AddBlock(data string) {
	// 防止区块越界
	if len(bc.blocks) <= 0 {
		os.Exit(1)
	}
	lastBlock := bc.blocks[len(bc.blocks)-1]
	block := NewBlock(data, lastBlock.Hash)
	bc.blocks = append(bc.blocks, block)
}

func main() {
	// 实例化一个区块链
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
		fmt.Printf("None:%d\n", block.Nonce)
	}
}
