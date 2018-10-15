package main

import (
	"./bolt"
)

type BlockChainIterator struct {
	db *bolt.DB
	//游标，用于不断索引
	currentHashPointer []byte
}

//func NewIterator(bc *BlockChain)  {
//
//}

func (bc *BlockChain) NewIterator() *BlockChainIterator {
	return &BlockChainIterator{
		bc.db,
		//最初指向区块链的最后一个区块，随着Next的调用，不断变化
		bc.tail,
	}
}
