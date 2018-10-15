package main

//import "fmt"

func main() {
	bc := NewBlockChain()
	bc.AddBlock("111111111111111")
	bc.AddBlock("222222222222222")

	/*
	for i, block := range bc.blocks {

		fmt.Printf("======== 当前区块高度： %d ========\n", i)
		fmt.Printf("前区块哈希值： %x\n", block.PrevHash)
		fmt.Printf("当前区块哈希值： %x\n", block.Hash)
		fmt.Printf("区块数据 :%s\n", block.Data)
	}
	*/
}
