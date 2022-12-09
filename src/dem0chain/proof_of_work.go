package dem0chain

import (
	"bytes"
	"crypto/sha256"
	"math"
	"math/big"
	"strconv"
)

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
		// hash consists of prevhash, data, timestamp, targetbits, nonce altogether
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
	// 测试成功
}
