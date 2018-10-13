package main

import (
	"github.com/boltdb/bolt"
	"os"
)

const dbfile  = "blockchain.db"
const blockBucket  = "block_demo"
const lastHashKey  = "genesis"

type BlockChian struct {
	//blocks []*Block
	db *bolt.DB
	lastHash []byte
}

func NewBlockChain(addr string) *BlockChian {
	//return &Blo ckChian{[]*Block{NewGenesisBlock()}}
	//func Open(path string, mode os.FileMode, options *Options) (*DB, error) {
	//	var db = &DB{opened: true}

	db, err := bolt.Open(dbfile, 7777, nil)
	CheckErr(err)
	var lasthash []byte
	//db.View()
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))

		if bucket != nil {
			//读取lasthash
			lasthash = bucket.Get([]byte(lastHashKey))
		} else {
			// 1.创建bucket
			// 2. 创世区块
			genesis := NewGenesisBlock()
			bucket, err := tx.CreateBucket([]byte(blockBucket))
			CheckErr(err)

			err = bucket.Put(genesis.Hash, genesis.Serialize())
			CheckErr(err)
			lasthash = genesis.Hash
			err = bucket.Put([]byte(lastHashKey), genesis.Hash)
			CheckErr(err)
		}
		return nil
	})

	return &BlockChian{db:db, lastHash:lasthash}

}


func (bc *BlockChian)AddBlock(data string) {
	//var prevBlockHash  []byte
	/*err := bc.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		lasthash := bucket.Get([]byte(lastHashKey))
		prevBlockHash = lasthash
		return nil
	})*/


	block := NewBlock(data, bc.lastHash)

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
	db *bolt.DB
	currentHash []byte
}


func (bc *BlockChian)Iterator() *BlockChainIterator {
	return &BlockChainIterator{bc.db, bc.lastHash}
}

func (it *BlockChainIterator)Next() *Block  {
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