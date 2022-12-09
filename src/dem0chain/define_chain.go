package dem0chain

import "fmt"

type Blockchain struct {
	blocks []*Block // 区块的指针数组
}

func (chain *Blockchain) AddBlock(data_prepared string, targetBits int) {
	new_block := NewBlock(data_prepared, chain.blocks[len(chain.blocks)-1].Hash)
	nonce, _ := ProofOfWork(new_block, targetBits)
	if nonce == -1 {
		panic("nonce not found.")
	}
	// proof of work should be done before set hash
	new_block.SetHash()
	new_block.Nonce = nonce
	new_block.TargetBits = targetBits
	// the latest block finally was added to the end of chain
	chain.blocks = append(chain.blocks, new_block)
}

func (chain *Blockchain) PrintChain() {
	fmt.Println("######## block chain info ########")
	for i, b := range chain.blocks {
		fmt.Println("===> block", i)
		b.PrintBlockInfo()
	}
	fmt.Println("######## block chain ended ########")
}

func NewBlockchain() *Blockchain {
	genesis_block := NewBlock("Genesis block created by redh3t.", []byte{})
	genesis_block.SetHash()
	genesis_block.TargetBits = 0 // genesis block doesn't neet to calc the nonce
	return &Blockchain{[]*Block{genesis_block}}
}
