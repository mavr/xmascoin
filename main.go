package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

const (
	MaxUint = ^uint(0)
	MaxInt  = int(MaxUint >> 1)

	zerocount = 5
)

type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

func (b *Block) String() string {
	return fmt.Sprintf("Index : %d\nSalt  : %d\nHash  : %s\nPrev  : %s\n",
		b.Index, b.BPM, b.Hash, b.PrevHash)
}

func (b *Block) ZeroValidate() bool {
	for i := 0; i < zerocount; i++ {
		if b.Hash[i] != '0' {

			return false
		}
	}
	return true
}

var (
	Blockchain []Block
)

func main() {
	fmt.Println("Welcome to the XMas coin ! \n")

	rootBlock := Block{
		Index:     0,
		Timestamp: time.Now().String(),
		BPM:       0,
		Hash:      "thisisrootblock",
		PrevHash:  "",
	}

	var blc, tmpBlc *Block
	var err error
	blc = &rootBlock

	for i := 0; i < 100; i++ {
		for bmp := 0; bmp < MaxInt; bmp++ {
			tmpBlc, err = generateBlock(blc, bmp)
			if err != nil {
				panic(err)
			}

			if tmpBlc.ZeroValidate() {
				blc = tmpBlc
				break
			}
		}

		fmt.Printf("............\n%s\n\n", blc.String())
	}

}

func calculateHash(block *Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)

	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock *Block, BPM int) (*Block, error) {

	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(&newBlock)

	return &newBlock, nil
}

func isBlockValid(newBlock, oldBlock *Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

func replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}
