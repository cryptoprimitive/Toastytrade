package main

import (
	"encoding/json"
	_ "github.com/SomniaStellarum/StellarUtilities/slog"
	"github.com/ethereum/go-ethereum/common"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"math/big"
	"os"
)

var (
	ttdb    *leveldb.DB
	accountsdb *leveldb.DB
)

var (
	createdUpdateChan        chan *createdUpdate
	committedUpdateChan      chan *committedUpdate
	fundsBurnedUpdateChan    chan *fundsBurnedUpdate
	fundsReleasedUpdateChan  chan *fundsReleasedUpdate
	fundsRecoveredUpdateChan chan *noArgUpdate
	claimStartedUpdateChan   chan *noArgUpdate
	claimCanceledUpdateChan  chan *noArgUpdate
	claimTriggeredUpdateChan chan *noArgUpdate
	closedUpdateChan         chan *noArgUpdate

	putAccountReqChan chan *putAccountRequest

	getAllToastytradeAddressesReqChan chan bool
	getAllToastytradeAddressesResChan chan getAllToastytradeAddressesResult
	getToastytradeReqChan             chan common.Address
	getToastytradeResChan             chan getToastytradeResult
	getAccountReqChan                   chan common.Address
	getAccountResChan                   chan getAccountResult
)

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Panic("Error finding working directory", err)
	}

	ttdb, err = leveldb.OpenFile(pwd+"/ttdb", nil)
	if err != nil {
		log.Panic("Error opening toastytrade database: ", err)
	}

	accountsdb, err = leveldb.OpenFile(pwd+"/accountsdb", nil)
	if err != nil {
		log.Panic("Error opening email database: ", err)
	}

	createdUpdateChan = make(chan *createdUpdate)
	committedUpdateChan = make(chan *committedUpdate)
	fundsBurnedUpdateChan = make(chan *fundsBurnedUpdate)
	fundsReleasedUpdateChan = make(chan *fundsReleasedUpdate)
	fundsRecoveredUpdateChan = make(chan *noArgUpdate)
	claimStartedUpdateChan = make(chan *noArgUpdate)
	claimCanceledUpdateChan = make(chan *noArgUpdate)
	claimTriggeredUpdateChan = make(chan *noArgUpdate)
	closedUpdateChan = make(chan *noArgUpdate)

	putAccountReqChan = make(chan *putAccountRequest)

	getAllToastytradeAddressesReqChan = make(chan bool)
	getAllToastytradeAddressesResChan = make(chan getAllToastytradeAddressesResult)
	getToastytradeReqChan = make(chan common.Address)
	getToastytradeResChan = make(chan getToastytradeResult)
	getAccountReqChan = make(chan common.Address)
	getAccountResChan = make(chan getAccountResult)
}

type createdUpdate struct {
	ttAddr common.Address
	entry  *toastytradeEntry
}

type committedUpdate struct {
	ttAddr    common.Address
	buyerAddr common.Address
}

type fundsBurnedUpdate struct {
	ttAddr common.Address
	amount big.Int
}

type fundsReleasedUpdate struct {
	ttAddr common.Address
	amount big.Int
}

type noArgUpdate struct {
	ttAddr common.Address
}

type putAccountRequest struct {
	ethAddr common.Address
	entry *AccountEntry
}

type getAccountResult struct {
	entry *AccountEntry
	err error
}

type getAllToastytradeAddressesResult struct {
	addresses []common.Address
	err       error
}

type getToastytradeResult struct {
	entry *toastytradeEntry
	err   error
}

type toastytradeEntry struct {
	Seller              common.Address `json:"seller"`
	Buyer               common.Address `json:"buyer"`
	TradeAmount         big.Int        `json:"amount"`
	SummonFee           big.Int        `json:"summon"`
	CommitThreshold     big.Int        `json:"commit"`
	AutoreleaseInterval big.Int        `json:"autoreleaseInterval"`
	AutoreleaseTime     big.Int        `json:"autoreleaseTime"`
	State               uint8          `json:"state"`

	Balance   big.Int `json:"balance"`
	Deposited big.Int `json:"deposited"`
	Burned    big.Int `json:"burned"`
	Released  big.Int `json:"released"`

	Banks           []string `json:"banks"`
	OtherDeposits   []string `json:"otherDeposits"`
	DepositInterval big.Int  `json:"depositInterval"`
	DepositDeadline big.Int  `json:"depositDeadline"`
}

func dbRequestsHandler() {
	for {
		select {
		case r := <-putAccountReqChan:
			putAccount(r.ethAddr, r.entry)

		case u := <-createdUpdateChan:
			err := putToastytrade(u.ttAddr, u.entry)
			if err != nil {
				log.Panic(err)
			}

		case u := <-committedUpdateChan:
			entry, err := getToastytrade(u.ttAddr)

			entry.Buyer = u.buyerAddr
			entry.State = 2 // state.Committed = 2

			err = putToastytrade(u.ttAddr, entry)
			if err != nil {
				log.Panic(err)
			}

		case u := <-fundsBurnedUpdateChan:
			entry, err := getToastytrade(u.ttAddr)

			entry.Balance.Sub(&entry.Balance, &u.amount)

			err = putToastytrade(u.ttAddr, entry)
			if err != nil {
				log.Panic(err)
			}

		case u := <-fundsReleasedUpdateChan:
			entry, err := getToastytrade(u.ttAddr)

			entry.Balance.Sub(&entry.Balance, &u.amount)

			err = putToastytrade(u.ttAddr, entry)
			if err != nil {
				log.Panic(err)
			}

		case u := <-fundsRecoveredUpdateChan:
			err := delToastytrade(u.ttAddr)
			if err != nil {
				log.Panic(err)
			}

		case u := <-claimStartedUpdateChan:
			entry, err := getToastytrade(u.ttAddr)

			entry.State = 3 // state.Claimed = 3

			err = putToastytrade(u.ttAddr, entry)
			if err != nil {
				log.Panic(err)
			}

		case u := <-claimCanceledUpdateChan:
			entry, err := getToastytrade(u.ttAddr)

			entry.State = 2 // state.Committed = 2

			err = putToastytrade(u.ttAddr, entry)
			if err != nil {
				log.Panic(err)
			}

		case u := <-claimTriggeredUpdateChan:
			entry, err := getToastytrade(u.ttAddr)

			//We don't need to change entry.State here; it's handled with the Closed event

			err = putToastytrade(u.ttAddr, entry)
			if err != nil {
				log.Panic(err)
			}

		case u := <-closedUpdateChan:
			entry, err := getToastytrade(u.ttAddr)

			entry.State = 4 // state.Closed = 4

			err = putToastytrade(u.ttAddr, entry)
			if err != nil {
				log.Panic(err)
			}

		case <-getAllToastytradeAddressesReqChan:
			r, err := getToastytradeAddresses()
			getAllToastytradeAddressesResChan <- getAllToastytradeAddressesResult{r, err}

		case addr := <-getToastytradeReqChan:
			r, err := getToastytrade(addr)
			getToastytradeResChan <- getToastytradeResult{r, err}

		case addr := <-getAccountReqChan:
			r, err := getAccount(addr)
			getAccountResChan <- getAccountResult{r, err}
		}
	}
}

func getToastytradeAddresses() (addresses []common.Address, err error) {
	iter := ttdb.NewIterator(nil, nil)

	for iter.Next() {
		addresses = append(addresses, common.BytesToAddress(iter.Key()))
	}

	return addresses, nil
}

func getToastytrade(toastytradeAddress common.Address) (entry *toastytradeEntry, err error) {
	v, err := ttdb.Get(toastytradeAddress.Bytes(), nil)
	if err != nil {
		return nil, err
	}
	entry = new(toastytradeEntry)
	err = json.Unmarshal(v, entry)
	if err != nil {
		return nil, err
	}
	return entry, err
}

func delToastytrade(toastytradeAddress common.Address) (err error) {
	err = ttdb.Delete(toastytradeAddress.Bytes(), nil)
	return
}

func putToastytrade(toastytradeAddress common.Address, entry *toastytradeEntry) (err error) {
	v, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	err = ttdb.Put(toastytradeAddress.Bytes(), v, nil)
	return err
}

type AccountEntry struct {
	Email string `json:"email"`
	WarningIntervalMinutes int `json:"warninginterval"`
}

func getAccount(addr common.Address) (account *AccountEntry, err error) {
	v, err := accountsdb.Get(addr.Bytes(), nil)
	if err != nil {
		return &AccountEntry{}, err
	}
	account = new(AccountEntry)
	err = json.Unmarshal(v, account)
	if err != nil {
		return account, err
	}
	return account, nil
}

func putAccount(addr common.Address, entry *AccountEntry) (err error) {
	v, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	err = accountsdb.Put(addr.Bytes(), v, nil)
	return err
}
