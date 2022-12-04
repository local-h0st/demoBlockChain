package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"time"
)

const targetBits = 12

// proof of work
// 为4的整数倍的话target的十六进制是一堆0跟1个1再跟一堆0
// 如果不是4的整数被可能不是1而是其他的数字，但感觉问题不大
// 二进制下是targetbits-1个0跟1个1再跟256-targetbits个0

func main() {
	fmt.Println("This is the start of project 'demoBlockChain'")
	blockchain := NewBlockchain()
	blockchain.AddBlock("this is the second block's data")
	blockchain.AddBlock("third one here!")
	for _, b := range blockchain.blocks {
		fmt.Println("\nprev-hash:", b.PrevBlockHash, "\ndata:", b.Data, "\ntimestamp:", b.Timestamp, "\ncurrent block hash:", b.Hash)
	}
	fmt.Println()
	fmt.Println()
	fmt.Println(ProofOfWork(blockchain.blocks[0], targetBits))
}

// 基本结构体和方法

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func (block *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))
	headers := bytes.Join([][]byte{block.PrevBlockHash, block.Data, timestamp}, []byte{})
	next_hash := sha256.Sum256(headers)
	block.Hash = next_hash[:]
}

type Blockchain struct {
	blocks []*Block // 区块的指针数组
}

func (chain *Blockchain) AddBlock(data_prepared string) {
	new_block := NewBlock(data_prepared, chain.blocks[len(chain.blocks)-1].Hash)
	chain.blocks = append(chain.blocks, new_block)
}

// 以下是NewBlock函数和NewBlockchain函数

func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevHash, []byte{}}
	block.SetHash()
	return block
	// NewBlock在数据准备完毕之后被调用，打包后直接写入链，并且不再改动
	// 因此记录打包时间戳，打包完成后计算Hash
}
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewBlock("Genesis block created by redh3t.", []byte{})}}
}

// proof of work
func (block *Block) PrepareData() []byte {
	var data []byte = bytes.Join(
		[][]byte{
			block.PrevBlockHash,
			block.Data,
			// 原教程用的函数的IntToHex但是我没有
			[]byte(strconv.FormatInt(block.Timestamp, 16)),
		},
		[]byte{},
	)
	return data
	// 前区块hash+当前区块内容+当前区块创建时间戳
}

func ProofOfWork(block *Block, targetbits int) (int, []byte) {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetbits))
	// target是proof of work的目标，左移的目的是造出前面有targetbits的0
	// target准备完毕, 开始计算nonce
	const maxNonce = math.MaxInt64
	var nonce int // 从0开始还是1开始呢
	var origin_data []byte = bytes.Join(
		[][]byte{
			block.PrepareData(),
			[]byte(strconv.FormatInt(int64(targetbits), 16)),
			// 添加难度信息
		},
		[]byte{},
	)
	var data []byte
	// 正式开始proof of work
	for nonce < maxNonce {
		// step 1. merge nonce with origin_data
		data = bytes.Join(
			[][]byte{
				origin_data,
				[]byte(strconv.FormatInt(int64(nonce), 16)),
			},
			[]byte{},
		)
		// step 2. calc the hash
		hash := sha256.Sum256(data)
		var hashInt big.Int
		hashInt.SetBytes(hash[:])
		// step 3/ compare current and target
		if hashInt.Cmp(target) == -1 {
			// current < target, which means the number of 0s from the start is more than target
			return nonce, data
		} else {
			nonce++
		}
	}
	return -1, []byte{}
	// TODO 写完还没测试
}
