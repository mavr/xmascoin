package blockchain

import "fmt"

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
