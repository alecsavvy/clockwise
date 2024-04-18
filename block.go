package main

type Block struct {
	BlockNumber  uint64
	Transactions []Transaction
}

type Transaction struct {
	ID   string
	Data interface{}
}
