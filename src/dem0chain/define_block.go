package dem0chain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
	TargetBits    int
}

func (block *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))
	headers := bytes.Join([][]byte{block.PrevBlockHash, block.Data, timestamp}, []byte{})
	next_hash := sha256.Sum256(headers)
	block.Hash = next_hash[:]
	// 当前块hash包括prevhash, data, timestamp
}

func (b *Block) PrintBlockInfo() {
	fmt.Println("prevhash:", b.PrevBlockHash)
	fmt.Println("timestamp:", b.Timestamp)
	fmt.Println("data:", string(b.Data))
	fmt.Println("hash:", b.Hash)
	fmt.Println("targetbits:", b.TargetBits)
	fmt.Println("nonce:", b.Nonce)
}
func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevHash, []byte{}, 0, 0}
	return block
	// NewBlock在数据准备完毕之后被调用，打包后直接写入链，并且不再改动
	// 因此记录打包时间戳，打包完成后计算Hash
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
	// return byte array consists of prevhash, data, timestamp
}
