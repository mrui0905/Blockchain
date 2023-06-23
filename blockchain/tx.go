package blockchain

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

// Returns treue if TxInput can be unlocked
func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

// Return True if TxOutput can be unlocked
func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}