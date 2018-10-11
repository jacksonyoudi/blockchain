package main

import (
	"flag"
	"fmt"
	"os"
)

const Usage  = `
	addBlock --data DATA  "add a block to block chain"
	printChain            "print all blocks"
`

type CLI struct {
	bc *BlockChian
}

func (cli *CLI)Run()  {
	if len(os.Args) < 2 {
		fmt.Println("too few parameters!\n", Usage)
		os.Exit(1)
	}
	
	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)
	
	addBlockCmdPara := addBlockCmd.String("data", "", "block info")

	switch os.Args[1] {
	case "addBlock":
		err := addBlockCmd.Parse(os.Args[2:])
		CheckErr(err)
		if addBlockCmd.Parsed() {
			if *addBlockCmdPara == "" {
				fmt.Println("data is empty")
				os.Exit(1)
			}
			cli.AddBlock(*addBlockCmdPara)
		}
	case "printChain":
		err := printChainCmd.Parse(os.Args[2:])
		CheckErr(err)
		if printChainCmd.Parsed() {
			cli.PrintChain()
		}
	default:
		fmt.Println("invalid cmd\n", Usage)
		os.Exit(1)

	
	}
	
}