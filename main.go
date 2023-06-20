package main

import (
	"fmt"
	
)



func main() {
	chain := InitBlockChain()

	chain.AddBlock("First Block")

	for _, block := range chain.blocks {
		fmt.Printf("Data: %s\n", block.Data)
		//fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Hash: %x\n", block.Hash)
	}


}