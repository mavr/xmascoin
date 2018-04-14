package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

const (
	MaxUint   = ^uint(0)
	MaxInt    = int(MaxUint >> 1)
	zerocount = 3
)

var (
	Blockchain []Block
)

func calculateHash(block *Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)

	return hex.EncodeToString(hashed)
}

func GenerateBlock(oldBlock *Block, BPM int) (*Block, error) {

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
