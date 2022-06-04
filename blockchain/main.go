package main

// https://dev.to/nheindev/build-the-hello-world-of-blockchain-in-go-bli

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"
)

var hashFilterValue = big.NewInt(1)

type Block struct {
	PrecedingHash []byte
	Data          []byte
	Nonce         int
	Hash          []byte
}

func ToByteArray(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

func GetHashBytes(prevHash []byte, data []byte, nonce int) [32]byte {
	hashableData := bytes.Join(
		[][]byte{
			prevHash,
			data,
			ToByteArray(int64(nonce)),
		},
		[]byte{},
	)

	return sha256.Sum256(hashableData)
}

func NewBlock(data string, prevHash []byte) *Block {
	var intHash big.Int
	var hash [32]byte
	nonce := 0

	for nonce < math.MaxInt64 {
		nonce++
		hash = GetHashBytes(prevHash, []byte(data), nonce)

		fmt.Printf("\rNonce: 0x%x, computed hash: 0x%x", nonce, hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(hashFilterValue) == -1 {
			break
		}
	}
	fmt.Println()

	return &Block{
		Hash:          hash[:],
		Data:          []byte(data),
		PrecedingHash: prevHash,
		Nonce:         nonce,
	}
}

func (block *Block) IsValid() bool {
	var intHash big.Int
	hash := GetHashBytes(block.PrecedingHash, block.Data, block.Nonce)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(hashFilterValue) == -1
}

func logger(data interface{}) {
	jsonBytes, marshalErr := json.Marshal(data)
	if marshalErr != nil {
		fmt.Println(marshalErr)
		os.Exit(1)
	}

	fmt.Println(string(jsonBytes))
}

func main() {
	difficulty, _ := strconv.Atoi(os.Args[1])
	hashFilterValue.Lsh(hashFilterValue, uint(256-difficulty))
	fmt.Printf(fmt.Sprintf("Computed hashes must be less than 0x%x\n", hashFilterValue.Bytes()))

	blockChain := []*Block{NewBlock("genesis data", []byte{})}
	blockChain = append(blockChain, NewBlock("first block data", blockChain[len(blockChain)-1].Hash))
	blockChain = append(blockChain, NewBlock("second block data", blockChain[len(blockChain)-1].Hash))
	blockChain = append(blockChain, NewBlock("third block data", blockChain[len(blockChain)-1].Hash))

	tamperedBlock := NewBlock("fourth block data", blockChain[len(blockChain)-1].Hash)
	tamperedBlock.Data = []byte("tampered data")
	blockChain = append(blockChain, tamperedBlock)

	fmt.Println("Validating blocks . . .")
	for _, block := range blockChain {
		if block.IsValid() {
			logger(block)
		} else {
			fmt.Println(errors.New(fmt.Sprintf("Block hash 0x%x", block.Hash) + " is invalid!"))
		}
	}
}
