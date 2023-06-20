package blockchain

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
	Nonce 		int
}

// Method to create a block with info 'data' and previous hash 'prevHash'. Returns refrence to new block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

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