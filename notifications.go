package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	//"github.com/SomniaStellarum/StellarUtilities/slog"
)

type eventDescriber interface {
	bodyToSeller(BPAddress common.Address) []byte
	bodyToBuyer(BPAddress common.Address) []byte
	subjectToSeller(BPAddress common.Address) string
	subjectToBuyer(BPAddress common.Address) string
}

func notifyParties(toastytradeAddress common.Address, event eventDescriber) error {
	getToastytradeReqChan <- toastytradeAddress
	entryResult := <-getToastytradeResChan
	if entryResult.err != nil {
		log.Panic(entryResult.err)
	}
	entry := entryResult.entry

	getEmailReqChan <- entry.Seller
	result := <-getEmailResChan
	if result.err == nil {
		sendEmail(result.email, event.subjectToSeller(toastytradeAddress), event.bodyToSeller(toastytradeAddress))
	} else if result.err == leveldb.ErrNotFound {
		//no email was found, so do nothing
	} else {
		log.Panic(result.err)
	}

	if entry.Buyer != common.HexToAddress("0x0") {
		getEmailReqChan <- entry.Buyer
		result := <-getEmailResChan
		if result.err == nil {
			sendEmail(result.email, event.subjectToBuyer(toastytradeAddress), event.bodyToBuyer(toastytradeAddress))
		} else if result.err == leveldb.ErrNotFound {
			//no email was found, so do nothing
		} else {
			log.Panic(result.err)
		}
	}

	return nil
}

func sendEmail(to string, subject string, body []byte) error {
	fmt.Printf("would send email to %s, with:\n", to)
	fmt.Printf("subject: %s\n", subject)
	fmt.Printf("body:\n")
	fmt.Printf("%s\n\n", body)

	return nil
}

//No need to define new structs; we'll use the same ones defined in BurnablePayment.go

func (event BurnablePaymentCreated) subjectToSeller(BPAddress common.Address) string {
	return fmt.Sprintf("You have created a Toastytrade at %s.", BPAddress.Hex())
}

func (event BurnablePaymentCreated) bodyToSeller(BPAddress common.Address) []byte {
	return []byte(fmt.Sprintf("You have created a Toastytrade at %s.", BPAddress.Hex()))
}

func (event BurnablePaymentCreated) subjectToBuyer(BPAddress common.Address) string {
	return ""
}

func (event BurnablePaymentCreated) bodyToBuyer(BPAddress common.Address) []byte {
	return []byte{}
}

func (event BurnablePaymentCommitted) subjectToSeller(BPAddress common.Address) string {
	return fmt.Sprintf("A buyer has committed to your Toastytrade %s", BPAddress.Hex())
}

func (event BurnablePaymentCommitted) bodyToSeller(BPAddress common.Address) []byte {
	return []byte(fmt.Sprintf("The buyer %s has committed to your Toastytrade at address %s.\n\n"+
		"You will recieve another notification when the Buyer marks the deposit complete.",
		event.Committer.Hex(), BPAddress.Hex()))
}

func (event BurnablePaymentCommitted) subjectToBuyer(BPAddress common.Address) string {
	return "Commit confirmation"
}

func (event BurnablePaymentCommitted) bodyToBuyer(BPAddress common.Address) []byte {
	return []byte(fmt.Sprintf("You have committed to the Toastytrade at address %s.\n\n"))
}

func (event BurnablePaymentFundsBurned) subjectToSeller(BPAddress common.Address) string {
	return fmt.Sprintf("You have burned %v ETH for the Toastytrade %s.", weiToEth(event.Amount), BPAddress.Hex())
}

func (event BurnablePaymentFundsBurned) bodyToSeller(BPAddress common.Address) []byte {
	return []byte(fmt.Sprintf("You have burned %v ETH for the Toastytrade %s.", weiToEth(event.Amount), BPAddress.Hex()))
}

func (event BurnablePaymentFundsBurned) subjectToBuyer(BPAddress common.Address) string {
	return fmt.Sprintf("The Seller has burned %v ETH for the Toastytrade %s.", weiToEth(event.Amount), BPAddress.Hex())
}

func (event BurnablePaymentFundsBurned) bodyToBuyer(BPAddress common.Address) []byte {
	return []byte(fmt.Sprintf("The Seller has burned %v ETH for the Toastytrade %s.\n\n"+
		"This likely happened because the Seller did not see the money appear in his bank account by the deadline.\n\n"+
		"For more information, see the post here[LINK NEEDED].",
		weiToEth(event.Amount), BPAddress.Hex()))
}

func (event BurnablePaymentFundsReleased) subjectToSeller(BPAddress common.Address) string {
	return fmt.Sprintf("%v ETH has been released to the Buyer for the Toastytrade %s.", weiToEth(event.Amount), BPAddress.Hex())
}

func (event BurnablePaymentFundsReleased) bodyToSeller(BPAddress common.Address) []byte {
	return []byte(fmt.Sprintf("%v ETH has been released to the Buyer for the Toastytrade %s.", weiToEth(event.Amount), BPAddress.Hex()))
}

func (event BurnablePaymentFundsReleased) subjectToBuyer(BPAddress common.Address) string {
	return fmt.Sprintf("%v ETH has been released to you for the Toastytrade %s!", weiToEth(event.Amount), BPAddress.Hex())
}

func (event BurnablePaymentFundsReleased) bodyToBuyer(BPAddress common.Address) []byte {
	return []byte(fmt.Sprintf("%v ETH has been released to you for the Toastytrade %s!\n\n"+
		"You should now see the ETH in your Ethereum address.",
		weiToEth(event.Amount), BPAddress.Hex()))
}

/*
func (event BurnablePaymentClaimStarted) subjectToSeller(BPAddress common.Address) string {
	return fmt.Sprintf("The Buyer has started a claim for the ETH held in the Toastytrade %s.\n\n"+
		"",
		BPAddress.Hex())
}

func (event BurnablePaymentClaimStarted) bodyToSeller(BPAddress common.Address) []byte {
	return []byte(fmt.Sprintf("%v ETH has been released to the Buyer for the Toastytrade %s.", weiToEth(event.Amount), BPAddress.Hex()))
}

func (event BurnablePaymentClaimStarted) subjectToBuyer(BPAddress common.Address) string {
	return fmt.Sprintf("%v ETH has been released to you for the Toastytrade %s!", weiToEth(event.Amount), BPAddress.Hex())
}

func (event BurnablePaymentClaimStarted) bodyToBuyer(BPAddress common.Address) []byte {
	return []byte(fmt.Sprintf("%v ETH has been released to you for the Toastytrade %s!\n\n"+
		"You should now see the ETH in your Ethereum address.",
		weiToEth(event.Amount), BPAddress.Hex()))
}*/

//BurnablePaymentClaimCanceled
//BurnablePaymentClaimTriggered
//BurnablePaymentClosed
