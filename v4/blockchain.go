package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"os"
)

const dbfile = "blockchain.db"
const blockBucket = "block_demo"
const lastHashKey = "genesis"
const reward = 12.5

type BlockChian struct {
	//blocks []*Block
	db       *bolt.DB
	lastHash []byte
}

func NewBlockChain(addr string) *BlockChian {
	//return &Blo ckChian{[]*Block{NewGenesisBlock()}}
	//func Open(path string, mode os.FileMode, options *Options) (*DB, error) {
	//	var db = &DB{opened: true}

	if IsBlockChainExist() {
		fmt.Println("Block Chain already exits!")
		os.Exit(1)
	}


	db, err := bolt.Open(dbfile, 7777, nil)
	CheckErr(err)
	var lasthash []byte
	//db.View()
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))

		// 1.创建bucket
		// 2. 创世区块
		coinbase := NewCoinbaseTx(addr, "gensis")
		genesis := NewGenesisBlock(coinbase)
		bucket, err := tx.CreateBucket([]byte(blockBucket))
		CheckErr(err)

		err = bucket.Put(genesis.Hash, genesis.Serialize())
		CheckErr(err)
		lasthash = genesis.Hash
		err = bucket.Put([]byte(lastHashKey), genesis.Hash)
		CheckErr(err)
		return nil
	})

	return &BlockChian{db: db, lastHash: lasthash}

}

func GetBlockChain() *BlockChian {
	if !IsBlockChainExist() {
		fmt.Println("Block Chain not exist, please create first!")
		os.Exit(1)
	}


	db, err := bolt.Open(dbfile, 0600, nil)
	CheckErr(err)
	var lasthash []byte
	//db.View()
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket != nil {
			lasthash = bucket.Get([]byte(lastHashKey))
		}
		return nil
	})
	CheckErr(err)
	return &BlockChian{db: db, lastHash: lasthash}
}

func (bc *BlockChian) AddBlock(transactions []*Transaction) {
	//var prevBlockHash  []byte
	/*err := bc.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		lasthash := bucket.Get([]byte(lastHashKey))
		prevBlockHash = lasthash
		return nil
	})*/

	block := NewBlock(transactions, bc.lastHash)

	err := bc.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		err := bucket.Put(block.Hash, block.Serialize())
		CheckErr(err)

		err = bucket.Put([]byte(lastHashKey), block.Hash)
		CheckErr(err)
		return nil
	})

	CheckErr(err)
}

type BlockChainIterator struct {
	db          *bolt.DB
	currentHash []byte
}

func (bc *BlockChian) Iterator() *BlockChainIterator {
	return &BlockChainIterator{bc.db, bc.lastHash}
}

func (it *BlockChainIterator) Next() *Block {
	var block *Block

	err := it.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			os.Exit(1)
		}

		blockTmp := bucket.Get(it.currentHash)
		block = Deserialize(blockTmp)

		it.currentHash = block.PrevBlockHash
		return nil

	})

	CheckErr(err)
	return block

}

func NewCoinbaseTx(addr string, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("current Reward is: %f\n", reward)
	}

	inputs := Input{nil, 0, data}
	outputs := Output{reward, addr}

	tx := Transaction{nil, []Input{inputs}, []Output{outputs}}
	tx.SetTXID()
	return &tx
}


func IsBlockChainExist() bool {
	_, err := os.Stat(dbfile)
	if os.IsNotExist(err) {
		return false
	}
	return true

}

func (bc *BlockChian)FindUnspendTransactions(address string) []Transaction  {
	// 遍历每一个区块，找出所有的的，余额
	var transactions []Transaction
	var  spentUTXOs = make(map[string/*交易的txid*/][]int64)

	bci := bc.Iterator()
	for {
		// 返回区块， 将current 前移
		block := bci.Next()

		for _, currentTx := range block.Transactions {
			txid := string(currentTx.TXID)

			//遍历当前交易的outputs，通过output的解锁条件，确定满足条件的交易
			for outputIndex, output := range currentTx.TXOutputs {
				if output.ConBeUnlockedByAddress(address) {
					transactions = append(transactions, *currentTx)
				}
			}

			for iutputIndex, input := range currentTx.TXInputs {
				if input.ConUnlockedByAddress(address) {
					spentUTXOs[txid] = append(spentUTXOs[txid], input.ReferOutputIndex )
				}
			}
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}

	}

	return
}


func (bc *BlockChian)FindUTXOs(address string) []Output {
	var outputs []Output
	txs := bc.FindUnspendTransactions(address)

	for _, tx := range txs {
		for _, output := range tx.TXOutputs {
			outputs = append(outputs, output)
		}
	}
}