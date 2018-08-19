package main

import (
	"math/big"
	"context"
	"html/template"
	"log"
	"net/http"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/SomniaStellarum/StellarUtilities/slog"
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

	err := PutEmail(ethAddress, email)
	if err != nil {
		log.Panic("Error calling PutEmail: ", err)
	}
}

func registerTrade(w http.ResponseWriter, r *http.Request) {
	slog.DebugPrint("registerTrade called")

	err := tpl.ExecuteTemplate(w, "registerTrade.html", nil)
	if err != nil {
		log.Panic("Template error on registerTrade")
	}
}

func doRegisterTrade(w http.ResponseWriter, r *http.Request) {
	slog.DebugPrint("doRegisterTrade called")

	toastytradeAddress := common.HexToAddress(r.PostFormValue("toastytradeAddress"))

	err := RegisterAndPopulateToastytrade(toastytradeAddress)
	if err != nil {
		log.Panic("Error calling RegisterAndPopulateToastytrade", err)
	}

	slog.DebugPrint("Toastytrade registered")
}

func getUserEmail(w http.ResponseWriter, r *http.Request) {
	slog.DebugPrint("getUserEmail called")

	ethAddress := common.HexToAddress(r.PostFormValue("ethAddress"))

	email, err := GetEmail(ethAddress)
	if err != nil {
		log.Panic("Error calling GetEmail: ", err)
	}

	w.Write([]byte(email))
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Running test func"))

	cli, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	defer cli.Close()

	if err != nil {
		log.Panic("Didn't work: ", err)
	}

	ctx := context.Background()

	fq := ethereum.FilterQuery {}
	fq.FromBlock = big.NewInt(3810620)
	fq.Addresses = []common.Address{common.HexToAddress("0x69FC16ee0b64c53FAfE516b458CC42D2e19Af566")}

	logs, err := cli.FilterLogs(ctx, fq)

	if err != nil {
		log.Panic("Error fetching logs: ", err)
	}

	for i := 0; i < len(logs); i++ {
		if logs[i].Topics[0] == crypto.Keccak256Hash([]byte("Committed(address)")) {

			//create new contract
			bp, err := NewBurnablePayment(logs[i].Address, cli)
			if err != nil {
				log.Panic("NewBurnablePayment error: ", err)
			}

			//decode event
			event := new(BurnablePaymentCommitted)
			bp.BurnablePaymentFilterer.contract.UnpackLog(event, "Committed", logs[i])

			slog.DebugPrint(event.Committer.Hex())

			//pass contract addr and decoded event off to be massaged into an email notification and sent out
		}
	}

	//slog.DebugPrint(logs[0].Topics[0].Hex())

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
	http.HandleFunc("/", index)
	http.HandleFunc("/register", register)
	http.HandleFunc("/doRegister", doRegister)
	http.HandleFunc("/getUserEmail", getUserEmail)
	http.HandleFunc("/registerTrade", registerTrade)
	http.HandleFunc("/doRegisterTrade", doRegisterTrade)
	http.HandleFunc("/test", test)
	http.Handle("/css/", slog.DebugPrintURL(http.StripPrefix("/css" ,http.FileServer(http.Dir("./css")))))
	http.Handle("/js/", slog.DebugPrintURL(http.StripPrefix("/js" ,http.FileServer(http.Dir("./js")))))
	http.ListenAndServe(":8080", nil)
}
