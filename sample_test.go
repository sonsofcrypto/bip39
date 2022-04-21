// SPDX-License-Identifier: MIT
package bip39

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"testing"
)

func TestGetWordList(t *testing.T) {
	// Load word list, need to provide path to word lists folder or copy folder
	// to same folder as binary
	wordList, _ := LoadWordList(ListLangEn, "wordLists")
	// Generate bip39 compliant entropy
	entropy, _ := NewEntropy(EntropySize128)
	// Generate mnemonic from entropy
	mnemonic, _ := NewMnemonic(entropy, wordList)
	// Generate seed from mnemonic
	seed := NewSeed(mnemonic, "Secret Passphrase")

	if !IsMnemonicValid(mnemonic, WordMap(wordList)) {
		t.Error("Invalid mnemonic")
	}

	// Display mnemonic and keys
	fmt.Println("Mnemonic:", mnemonic)
	fmt.Println("Seed:", seed)
	fmt.Println("Hex:", hexutil.Encode(seed))
	fmt.Println("Int:", big.NewInt(0).SetBytes(seed[:]))
}
