package main

import (
	"github.com/syndtr/goleveldb/leveldb"
	"os"
	"log"
	"github.com/ethereum/go-ethereum/common"
	"encoding/json"
)

var db *leveldb.DB

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Panic("Error finding working directory", err)
	}

	db, err = leveldb.OpenFile(pwd + "/emaildb", nil)
	if err != nil {
		log.Panic("Error opening email database", err)
	}
}

type emailDBEntry struct {
	Email string `json:"email"`
}

func GetEmail(addr common.Address) (email string, err error) {
	v, err := db.Get(addr.Bytes(), nil)
	if err != nil {
		return nil, err
	}
	var e emailDBEntry
	err = json.Unmarshal(v, e)
	if err != nil {
		return nil, err
	}
	return e.Email, nil
}

