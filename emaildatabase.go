package main

import (
	"github.com/syndtr/goleveldb/leveldb"
	"os"
	"log"
	"github.com/ethereum/go-ethereum/common"
	"encoding/json"
//	"github.com/SomniaStellarum/StellarUtilities/slog"
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
		return "", err
	}
	var e emailDBEntry
	err = json.Unmarshal(v, &e)
	if err != nil {
		return "", err
	}
	return e.Email, nil
}

func PutEmail(addr common.Address, email string) (err error) {
	var e emailDBEntry
	e.Email = email
	v, err := json.Marshal(e)
	if err != nil {
		return err
	}
	err = db.Put(addr.Bytes(), v, nil)
	return err
}
