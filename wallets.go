package main

import (
	"io/ioutil"
	"bytes"
	"encoding/gob"
)

//定一个 Wallets结构，它保存所有的wallet以及它的地址
type Wallets struct {
	//map[地址]钱包
	WalletsMap map[string]*Wallet
}

//创建方法，
func NewWallets() *Wallets {
	var ws Wallets
	ws.WalletsMap = make(map[string]*Wallet)
	//ws := loadFile()
	return &ws
}

func (ws *Wallets) CreateWallet() string {
	wallet := NewWallet()
	address := wallet.NewAddress()

	ws.WalletsMap[address] = wallet

	ws.saveToFile()
	return address
}

//保存方法，把新建的wallet添加进去
func (ws *Wallets) saveToFile() {

	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)
	encoder.Encode(ws)

	ioutil.WriteFile("wallet.dat", buffer.Bytes(), 0600)
}

//读取文件方法，把所有的wallet读出来
