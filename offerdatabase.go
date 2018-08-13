package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/syndtr/goleveldb/leveldb"
	"os"
	"log"
	"github.com/SomniaStellarum/StellarUtilities/slog"
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

func RegisterAndPopulateToastytrade(addr common.Address) (err error) {
	slog.DebugPrint("Okay, gonna totally register and populate toastytrade for address ", addr.Hex())

	return nil
}
