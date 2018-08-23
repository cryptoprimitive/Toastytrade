package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/syndtr/goleveldb/leveldb"
	"os"
	"log"
	"github.com/SomniaStellarum/StellarUtilities/slog"
	"encoding/json"
)

// to store
// offers -> address
// values:
//  seller -> address
//  buyer -> address

var ttdb *leveldb.DB

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Panic("Error finding working directory", err)
	}

	ttdb, err = leveldb.OpenFile(pwd + "/ttdb", nil)
	if err != nil {
		log.Panic("Error opening offer database", err)
	}
}

type toastytradeEntry struct {
	Seller common.Address `json:"seller"`
	Buyer common.Address `json:"buyer"`
}

func RegisterAndPopulateToastytrade(addr common.Address) (err error) {
	slog.DebugPrint("Okay, gonna totally register and populate toastytrade for address ", addr.Hex())

	return nil
}

func GetToastytradeAddresses() (addresses []common.Address, err error) {
	iter := ttdb.NewIterator(nil, nil)

	for iter.Next() {
		addresses = append(addresses, common.BytesToAddress(iter.Value()))
	}

	return addresses, nil
}

func GetToastytrade(ToastytradeAddress common.Address) (entry *toastytradeEntry, err error) {
	v, err := ttdb.Get(ToastytradeAddress.Bytes(), nil)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(v, entry)
	if err != nil {
		return nil, err
	}
	return entry, err
}

func PutToastytrade(ToastytradeAddress common.Address, entry *toastytradeEntry) (err error) {
	v, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	err = ttdb.Put(ToastytradeAddress.Bytes(), v, nil)
	return err
}
