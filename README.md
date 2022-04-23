# bip39

Implementation of [BIP39](https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki) in go.
Fork of [go-bip39](https://https://github.com/tyler-smith/go-bip39), its own repo due to versioning issues with `go mod`.

## Example

```go
package main

import (
    "fmt"
    "github.com/sonsofcrypto/bip39"
	"encoding/hex"
    "math/big"
)

func main() {
	// Load word list, need to provide path to word lists folder or copy folder
	// to same folder as binary
	wordList, _ := bip39.LoadWordList(bip39.ListLangEn, "wordLists")
	// Generate bip39 compliant entropy
	entropy, _ := bip39.NewEntropy(bip39.EntropySize128)
	// Generate mnemonic from entropy
	mnemonic, _ := bip39.NewMnemonic(entropy, wordList)
	// Generate seed from mnemonic
	seed := bip39.NewSeed(mnemonic, "Secret Passphrase")

	// Display mnemonic and keys
	fmt.Println("Mnemonic:", mnemonic)
	fmt.Println("Seed:", seed)
	fmt.Println("Hex:", hex.EncodeToString(seed))
	fmt.Println("Int:", big.NewInt(0).SetBytes(seed[:]))
}
```
