package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

// Create Block structure
type Block struct {
	Hash		[]byte
	Transactions		[]*Transaction
	PrevHash	[]byte
	Nonce 		int
}

// Combines the IDs of every transaction in the block.
func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}

	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]


}

// Method to create a block with info 'data' and previous hash 'prevHash'. Returns refrence to new block
func CreateBlock(txs []*Transaction, prevHash []byte) *Block {
	block := &Block{[]byte{}, txs, prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// Method to create the genesis block of BlockChain. Returns refrence to new block with Block.Data == "Genesis"
func Genesis(coinbase *Transaction) *Block {
	return CreateBlock([]*Transaction{coinbase}, []byte{})
}

// Block method which takes a Block object and returns a splice of its bytes
func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res) 

	err := encoder.Encode(b)

	Handle(err)

	return res.Bytes()
}

// Method that takes a splice of bytes and returns a block object
func Deserialize(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)

	Handle(err)

	return &block
}

// Handles encode/decode errors
func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
