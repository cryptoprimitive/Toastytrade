package main

import (
	"github.com/SomniaStellarum/StellarUtilities/slog"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("html/*.html"))
}

func index(w http.ResponseWriter, r *http.Request) {
	slog.DebugPrint("Path requested: ", r.URL.String())
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Panic("Template error on index: ", err)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	slog.DebugPrint("Path requested: ", r.URL.String())
	err := tpl.ExecuteTemplate(w, "register.html", nil)
	if err != nil {
		log.Panic("Template error on register: ", err)
	}
}

func doRegister(w http.ResponseWriter, r *http.Request) {
	slog.DebugPrint("doRegister called")

	ethAddress := common.HexToAddress(r.PostFormValue("ethAddress"))
	email := r.PostFormValue("email")

	putAccountReqChan <- &putAccountRequest{ethAddress, &AccountEntry{email, 1}}
}

func getUserEmail(w http.ResponseWriter, r *http.Request) {
	slog.DebugPrint("getUserEmail called")

	ethAddress := common.HexToAddress(r.PostFormValue("ethAddress"))

	getAccountReqChan <- ethAddress
	result := <-getAccountResChan
	if result.err != nil {
		// Crude, but return empty string to represent no email found
		w.Write([]byte(""))
		return
	}
	email := result.entry.Email

	w.Write([]byte(email))
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test func finished"))
}

type logRequester struct {
	h http.Handler
}

func (l logRequester) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	slog.DebugPrint("Static Path requested: ", r.URL.String())
	l.h.ServeHTTP(w, r)
}

func logRequest(f http.Handler) http.Handler {
	return logRequester{h: f}
}

func main() {
	slog.SetDebug()

	err := initEthStuff()
	if err != nil {
		log.Panic(err)
	}

	go dbRequestsHandler()
	go mainLoop(4200000)

	http.HandleFunc("/", index)
	http.HandleFunc("/register", register)
	http.HandleFunc("/doRegister", doRegister)
	http.HandleFunc("/getUserEmail", getUserEmail)
	http.HandleFunc("/test", test)
	http.Handle("/css/", slog.DebugPrintURL(http.StripPrefix("/css", http.FileServer(http.Dir("./css")))))
	http.Handle("/js/", slog.DebugPrintURL(http.StripPrefix("/js", http.FileServer(http.Dir("./js")))))
	http.ListenAndServe(":8080", nil)
}

func initEthStuff() error {
	CreatedTopic = crypto.Keccak256Hash([]byte("Created(address,bool,address,uint256,uint256,string)"))
	FundsAddedTopic = crypto.Keccak256Hash([]byte("FundsAdded(address,uint256)"))
	PayerStatementTopic = crypto.Keccak256Hash([]byte("PayerStatement(string)"))
	WorkerStatementTopic = crypto.Keccak256Hash([]byte("WorkerStatement(string)"))
	FundsRecoveredTopic = crypto.Keccak256Hash([]byte("FundsRecovered()"))
	CommittedTopic = crypto.Keccak256Hash([]byte("Committed(address)"))
	FundsBurnedTopic = crypto.Keccak256Hash([]byte("FundsBurned(uint256)"))
	FundsReleasedTopic = crypto.Keccak256Hash([]byte("FundsReleased(uint256)"))
	ClaimStartedTopic = crypto.Keccak256Hash([]byte("ClaimStarted()"))
	ClaimCanceledTopic = crypto.Keccak256Hash([]byte("ClaimCanceled()"))
	ClaimTriggeredTopic = crypto.Keccak256Hash([]byte("ClaimTriggered()"))
	ClosedTopic = crypto.Keccak256Hash([]byte("Closed()"))
	UnclosedTopic = crypto.Keccak256Hash([]byte("Unclosed()"))

	factoryAddress = common.HexToAddress(ROPSTEN_FACTORY_ADDRESS)

	var err error

	cli, err = ethclient.Dial("wss://ropsten.infura.io/ws")
	if err != nil {
		return err
	}

	toastytradeFactory, err = NewToastytradeFactory(common.HexToAddress(ROPSTEN_FACTORY_ADDRESS), cli)
	if err != nil {
		return err
	}

	return nil
}
