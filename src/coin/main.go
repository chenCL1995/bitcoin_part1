package main

import (
	"core"
	"fmt"
)

func main() {
	bc := core.NewBlockchain() //初始化区块链，创建第一个区块（创世纪区块）

	bc.AddBlock("Send 1 BTC to Ivan") //加入一个区块链
	bc.AddBlock("Send 2 BTC to Ivan") //加入第二个区块链

	//i是下标，可以省略，写成：for _, block := range bc.Blocks
	for i, block := range bc.Blocks {
		fmt.Printf("i:%x\n", i)
		fmt.Printf("Prev. Hash:%x\n", block.PrevBlockHash)
		fmt.Printf("Data:%s\n", block.Data)
		fmt.Printf("Hash:%x\n", block.Hash)
		fmt.Println()
	}
}
