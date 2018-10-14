package main

import "fmt"

func (cli *CLI)AddBlock(data string)  {
	//cli.AddBlock(data)
}

func (cli *CLI)CreateChain(addr string) {
	bc := NewBlockChain(addr)
	defer bc.db.Close()
	fmt.Println("create blockchain successfully!")

}


func (cli *CLI)PrintChain()  {
	bc := GetBlockChain()
	it := bc.Iterator()
	for {
		// 取回当前hash指定的block，并且将当前的hash指向上一个区块的hash
		block := it.Next()

		fmt.Printf("transaction %v\n:", block.Transactions)
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

func (cli *CLI)GetBalanece(address string)  {
	bc := GetBlockChain()
	defer bc.db.Close()

	var total float64
	utxos := bc.FindUTXOs(address)
	for _, utxo := range utxos {
		total += utxo.Value
	}
	fmt.Printf("The balance of '%s' is %f", total)
}