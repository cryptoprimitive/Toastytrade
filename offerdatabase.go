package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/syndtr/goleveldb/leveldb"
	"os"
	"log"
)

// to store
// offers -> address
// values:
//  seller -> address
//  buyer -> address

var offerdb *leveldb.DB

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Panic("Error finding working directory", err)
	}

	offerdb, err = leveldb.OpenFile(pwd + "/offerdb", nil)
	if err != nil {
		log.Panic("Error opening offer database", err)
	}
}

type offerDBEntry struct {
	Seller common.Address `json:"seller"`
	Buyer common.Address `json:"buyer"`
}
