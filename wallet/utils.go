package wallet

import (
	"log"

	"github.com/mr-tron/base58"
)

// Converts hash to Base 58
func Base58Encode(input []byte) []byte {
	encode := base58.Encode(input)

	return []byte(encode)
}

// Converst Base 58 to hash
func Base58Decode(input []byte) []byte {
	decode, err := base58.Decode(string(input[:]))
	if err != nil {
		log.Panic(err)
	}

	return decode
}