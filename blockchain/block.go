package blockchain

import (
	"bytes"
	"crypto/sha256"
)

// Create BlockChain structure
type BlockChain struct{
	Blocks 		[]*Block
}

// BlockChain method to add a new block with info 'data'
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

// Create Block structure
type Block struct {
	Hash		[]byte
	Data		[]byte
	PrevHash	[]byte
}

// Block method to create the hash of the block based on 'Block.Data' and 'Block.PrevHash'
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// Method to create a block with info 'data' and previous hash 'prevHash'. Returns refrence to new block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

// Method to create the genesis block of BlockChain. Returns refrence to new block with Block.Data == "Genesis"
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// Method to initialize the BlockChain structure. Returns refrence to BlockChain with Genesis block already created
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}