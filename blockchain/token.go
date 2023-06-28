package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"

	"github.com/mrui0905/Blockchain/wallet"
)

type Tokens struct {
	AllTokens		[]*Token
}

type Token struct {
	Owner		wallet.Wallet
	Id			[32]byte
}

func (ts *Tokens) CreateToken(owner wallet.Wallet) bool {
	var id [32]byte
	newToken := &Token{owner, id}
	newToken.Id = CreateId(newToken.Serialize(len(ts.AllTokens)))

	ts.AllTokens = append(ts.AllTokens, newToken)

	return true
}

func (t *Token) Serialize(number int) []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res) 

	err := encoder.Encode(t)

	Handle(err)

	return res.Bytes()
}

func CreateId(b []byte) [32]byte {
	id := sha256.Sum256(b)

	return id
}

