package main

import (
	"flag"
	"fmt"
	"os"
)

const Usage  = `
	createChain	--address ADDRESS		  "create block Chain"
	addBlock --data DATA  "add a block to block chain"
	printChain            "print all blocks"
`

type CLI struct {
	//bc *BlockChian
}

func (cli *CLI)Run()  {
	if len(os.Args) < 2 {
		fmt.Println("too few parameters!\n", Usage)
		os.Exit(1)
	}
	
	createChainCmd := flag.NewFlagSet("createChain", flag.ExitOnError)
	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)
	
	addBlockCmdPara := addBlockCmd.String("data", "", "block info")
	createChainPara := createChainCmd.String("address", "", "input address")

	switch os.Args[1] {
	case "createChain":
		err := createChainCmd.Parse(os.Args[2:])
		CheckErr(err)
	case "addBlock":
		err := addBlockCmd.Parse(os.Args[2:])
		CheckErr(err)
	case "printChain":
		err := printChainCmd.Parse(os.Args[2:])
		CheckErr(err)
	default:
		fmt.Println("invalid cmd\n", Usage)
		os.Exit(1)

	}

	if createChainCmd.Parsed() {
		if *createChainPara == "" {
			fmt.Println("data is empty")
			fmt.Println(Usage)
			os.Exit(1)
		}
		cli.CreateChain(*addBlockCmdPara)
	}

	if addBlockCmd.Parsed() {
		if *addBlockCmdPara == "" {
			fmt.Println("data is empty")
			os.Exit(1)
		}
		cli.AddBlock(*addBlockCmdPara)
	}

	if printChainCmd.Parsed() {
		cli.PrintChain()
	}



}