package main

import (
	"fmt"
	"time"
)

type Transaction struct {
	Sender    string
	Recipient string
	Amount    int
}
type Block struct {
	ID                int
	Timestamp         int64
	Transactions      []Transaction
	Proof             int64
	PreviousBlockHash string
	LastBlock         int64
}

type Blockchain struct {
	Blocks []Block
}

// add new transaction
// create new block
// hash block
// get last block
var blockchain []Block

func init() {
	// Genesis block
	blockchain = []Block{
		Block{
			PreviousBlockHash: "1",
			Proof:             100,
		},
	}
}

func AddTransaction(sender string, recipient string, amount int) Transaction {
	transaction := Transaction{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
	}
	return transaction // probably also need to return the index of the block in relation to the last (+1)
}

func NewBlock(previousHash string, proof int64) Block {
	// when do I create the genesis block? should that be it's own function? should I just pass hardcoded hash and proof values?
	// needs to have no previous has

	// Where would I initialize everything? I would need to start the blockchain and keep it going but I'm not sure where!
	block := Block{}
	// ID                int64
	// Timestamp         time.Time
	// Transactions      []Transaction
	// Proof             int64
	// PreviousBlockHash string
	// LastBlock         int64
	block.ID = len(blockchain) + 1
	block.Timestamp = time.Now().Unix()
	// block.Transactions =
	block.PreviousBlockHash = previousHash
	block.Proof = proof

	return block
}

func HashBlock() string {
	return ""
}

func main() {
	newBlock := NewBlock("hashyhashyhasy", 12345)
	blockchain = append(blockchain, newBlock)
	fmt.Printf("blockchain %+v", blockchain)
}
