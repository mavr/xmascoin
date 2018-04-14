package main

import (
	"fmt"
	"time"

	"github.com/mavr/xmaschain/blockchain"
)

func main() {
	fmt.Println("Welcome to the XMas coin ! \n")

	rootBlock := blockchain.Block{
		Index:     0,
		Timestamp: time.Now().String(),
		BPM:       0,
		Hash:      "thisisrootblock",
		PrevHash:  "",
	}

	var blc, tmpBlc *blockchain.Block
	var err error
	blc = &rootBlock

	for i := 0; i < 10; i++ {
		for bmp := 0; bmp < blockchain.MaxInt; bmp++ {
			tmpBlc, err = blockchain.GenerateBlock(blc, bmp)
			if err != nil {
				panic(err)
			}

			if tmpBlc.ZeroValidate() {
				blc = tmpBlc
				break
			}
		}

		fmt.Printf("%s........................................................................\n\n",
			blc.String())

	}

}
