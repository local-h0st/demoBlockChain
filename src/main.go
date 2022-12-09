package main

import (
	"fmt"

	"github.com/local-h0st/demoBlockChain/src/dem0chain"
)

const targetBits = 24

// proof of work
// 为4的整数倍的话target的十六进制是一堆0跟1个1再跟一堆0
// 如果不是4的整数被可能不是1而是其他的数字，但感觉问题不大
// 二进制下是targetbits-1个0跟1个1再跟256-targetbits个0

func main() {
	fmt.Println("This is the start of project 'demoBlockChain'")
	fmt.Println(dem0chain.PackageInfo())
	blockchain := dem0chain.NewBlockchain()
	blockchain.PrintChain()
	fmt.Println()
	blockchain.AddBlock("transactions: redh3t get 100$ as a miner.", targetBits)
	blockchain.PrintChain()
	blockchain.AddBlock("transactions: redh3t send 50$ to localh0st.", targetBits)
	blockchain.PrintChain()
}
