package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

// Difficulty Level
const Difficulty = 18

// Creates ProofOfWork structure
type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

// ProofOfWork method that combines and returns a block's Data, PrevHash, the current nonce and Difficulty into a splice of bytes
func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.HashTransactions(),
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)
	return data
}

// ProofOfWork method that increments the nonce by one until the hash of data is less than target. Returns the correct nonce and hash
func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce<math.MaxInt64 { // Essentially infinite loop
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		}else {
			nonce++
		}
	}

	fmt.Println()

	return nonce, hash[:]
}

// Validates that ProofOfWork's Nonce is valid. Return true if valid
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int
	
	data := pow.InitData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}

// Method that creates and returns a refrence to ProofOfWork Structure
func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))

	pow := &ProofOfWork{b, target}

	return pow
}

// Method that coonverts and returns an int64 into it's hexadecimal equivalent
func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}