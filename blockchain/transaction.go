package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

// Stores tranaction ID, inputs and outputs
type Transaction struct {
	ID 		[]byte
	Inputs	[]TxInput
	Outputs []TxOutput
}

// Transaction ouput- value contains the information, PubKey is needed to unlock transaction
type TxOutput struct {
	Value	int
	PubKey	string
}

// Transaction input- Refrences transaction by ID and Out (index), uses Sig to unlock transaction
type TxInput struct {
	ID		[]byte
	Out		int
	Sig		string
}


// Creates Coinbase
func CoinbaseTx(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Coins to %s", to)
	}

	txin := TxInput{[]byte{}, -1, data}
	txout := TxOutput{100, to}

	tx := Transaction{nil, []TxInput{txin}, []TxOutput{txout}}
	tx.SetID()

	return &tx
}

// Transaction method whcih sets a transactions ID based on a transaction
func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	encode := gob.NewEncoder(&encoded)
	err := encode.Encode(tx)
	Handle(err)

	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

// Returns true if tx is the coinbase
func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Inputs) == 1 && len(tx.Inputs[0].ID) == 0 && tx.Inputs[0].Out == -1
}

// Returns treue if TxInput can be unlocked
func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

// Return True if TxOutput can be unlocked
func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}

// Creates transaction baseed on the sender, receiver, and amount
func NewTransaction(from, to string, amount int, chain *BlockChain) *Transaction {
	var inputs []TxInput
	var outputs []TxOutput

	acc, validOutputs := chain.FindSpendableOutputs(from, amount)

	if acc < amount {
		log.Panic("Error: not enough funds")
	}

	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		Handle(err)

		for _, out := range outs {
			input := TxInput{txID, out, from}
			inputs = append(inputs, input)
		}
	}

	outputs = append(outputs, TxOutput{amount, to})

	if acc > amount {
		outputs = append(outputs, TxOutput{acc-amount, from})
	}

	tx := Transaction{nil, inputs, outputs}
	tx.SetID()

	return &tx
}