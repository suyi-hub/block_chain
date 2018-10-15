package main

import "math/big"

//定义一个工作量证明的结构ProofOfWork
//
type ProofOfWork struct {
	//a. block
	block *Block
	//b. 目标值
	//一个非常大数，它有很丰富的方法：比较，赋值方法
	target *big.Int
}

//2. 提供创建POW的函数
//
//- NewProofOfWork(参数)
func NewProofOfWork(block *Block) *ProofOfWork {
	pow := ProofOfWork{
		block: block,
	}

	//我们指定的难度值，现在是一个string类型，需要进行转换
	targetStr := "0000100000000000000000000000000000000000000000000000000000000000"
	//
	//引入的辅助变量，目的是将上面的难度值转成big.int
	tmpInt := big.Int{}
	//将难度值赋值给big.int，指定16进制的格式
	tmpInt.SetString(targetStr, 16)

	pow.target = &tmpInt
	return &pow
}

//
//3. 提供计算不断计算hash的哈数
//
//- Run()
//
//4. 提供一个校验函数
//
//- IsValid()
