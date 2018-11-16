package main

import (
	"fmt"
	"time"
	"github.com/ethereum/go-ethereum/common"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/dustin/go-humanize"
	"log"
	//"github.com/SomniaStellarum/StellarUtilities/slog"
)

type eventDescriber interface {
	describeToSeller(BPAddress common.Address) []byte
	describeToBuyer(BPAddress common.Address) []byte
}

func notifyParties(toastytradeAddress common.Address, event eventDescriber) error {
	getToastytradeReqChan <- toastytradeAddress
	entryResult := <-getToastytradeResChan
	if entryResult.err != nil {
		log.Panic(entryResult.err)
	}
	entry := entryResult.entry

	getAccountReqChan <- entry.Seller
	accountResult := <-getAccountResChan

	subject := fmt.Sprintf("Action for Toastytrade %s", toastytradeAddress.Hex())

	if accountResult.err == nil {
		sendEmail(accountResult.entry.Email, subject, event.describeToSeller(toastytradeAddress))
	} else if accountResult.err == leveldb.ErrNotFound {
		//no email was found, so do nothing
	} else {
		log.Panic(accountResult.err)
	}

	if entry.Buyer != common.HexToAddress("0x0") {
		getAccountReqChan <- entry.Buyer
		accountResult := <-getAccountResChan
		if accountResult.err == nil {
			sendEmail(accountResult.entry.Email, subject, event.describeToBuyer(toastytradeAddress))
		} else if accountResult.err == leveldb.ErrNotFound {
			//no email was found, so do nothing
		} else {
			log.Panic(accountResult.err)
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

func (event BurnablePaymentCreated) describeToSeller(BPAddress common.Address) []byte {
	return []byte(fmt.Sprintf("You have created a Toastytrade at %s.", BPAddress.Hex()))
}

func (event BurnablePaymentCreated) describeToBuyer(BPAddress common.Address) []byte {
	return []byte{}
}

func (event BurnablePaymentCommitted) describeToSeller(BPAddress common.Address) []byte {
	return []byte(fmt.Sprintf("A buyer has committed to your Toastytrade!\n\n"+
		"The buyer %s has committed to your Toastytrade at address %s. They should soon complete the deposit and mark it as complete. You will receive another notification once this happens.",
		event.Committer.Hex(), BPAddress.Hex()))
}

func (event BurnablePaymentCommitted) describeToBuyer(BPAddress common.Address) []byte {
	return []byte(fmt.Sprintf("You have committed to the Toastytrade at address %s.", BPAddress.Hex()))
}

func (event BurnablePaymentFundsBurned) describeToSeller(BPAddress common.Address) []byte {
	return []byte(fmt.Sprintf("You have burned %v ETH for the Toastytrade %s.", weiToEth(event.Amount), BPAddress.Hex()))
}

func (event BurnablePaymentFundsBurned) describeToBuyer(BPAddress common.Address) []byte {
	return []byte(fmt.Sprintf("The Seller has burned %v ETH for the Toastytrade %s.\n\n"+
		"This likely happened because the Seller did not see the money appear in his bank account by the deadline. For more information, see the post here[LINK NEEDED].",
		weiToEth(event.Amount), BPAddress.Hex()))
}

func (event BurnablePaymentFundsReleased) describeToSeller(BPAddress common.Address) []byte {
	return []byte(fmt.Sprintf("You released %v ETH to the Buyer for the Toastytrade %s.", weiToEth(event.Amount), BPAddress.Hex()))
}

func (event BurnablePaymentFundsReleased) describeToBuyer(BPAddress common.Address) []byte {
	return []byte(fmt.Sprintf("%v ETH has been released to you for the Toastytrade %s!\n\n"+
		"You should now see the ETH in your wallet.",
		weiToEth(event.Amount), BPAddress.Hex()))
}

func (event BurnablePaymentFundsRecovered) describeToSeller(BPAddress common.Address) []byte {
	return []byte(fmt.Sprintf("You have canceled the Toastytrade at %s. You should see the funds returned to your wallet.", BPAddress.Hex()))
}

func (event BurnablePaymentFundsRecovered) describeToBuyer(BPAddress common.Address) []byte {
	return []byte{}
}

func (event BurnablePaymentClaimStarted) describeToSeller(BPAddress common.Address) []byte {
	return []byte(fmt.Sprintf("The Buyer has marked the deposit as complete.\n\n" +
		"The autorelease timer is now counting down. Verify that you have received the money. If you have, you can either manually release the payment, or do nothing and wait for the autorelease to trigger.\n\n" +
		"If you have not received the money, this may be a troll or malicious buyer. If you can't verify the deposit has been made, you should deny the claim before the autorelease timer runs out, and contact the buyer." +
		"If this happens multiple times for the same toastytrade, you may have to burn the funds, to prevent a scam from profiting."))
}

func (event BurnablePaymentClaimStarted) describeToBuyer(BPAddress common.Address) []byte {
	return []byte(fmt.Sprintf("You have marked the deposit completed for the Toastytrade at %s.\n\n"+
		"The Seller has been notified to verify the deposit, and should then release the payment. Unless he denies the claim before the autorelease timer runs out, the ether will automatically be released to you.",
		BPAddress.Hex()))
}

func (event BurnablePaymentClaimCanceled) describeToSeller(BPAddress common.Address) []byte {
	return []byte(fmt.Sprintf("You have denied the Buyer's claim to the ether.\n\n" +
		"We recommend contacting the Buyer to resolve whatever issue caused you to deny the claim."))
}

func (event BurnablePaymentClaimCanceled) describeToBuyer(BPAddress common.Address) []byte {
	return []byte(fmt.Sprintf("The Seller has denied your claim for the Toastytrade at %s.\n\n"+
		"Unless you get a \"burn\" notification, your ether is still available. The Seller likely denied the claim because they haven't seen the deposited fiat. We recommend contacting the Seller and helping them to see evidence of the deposit.\n\n",
		BPAddress.Hex()))
}

func (event BurnablePaymentClaimTriggered) describeToSeller(BPAddress common.Address) []byte {
	return []byte(fmt.Sprintf("The Buyer of Toastytrade %s has triggered the autorelease.\n\n"+
		"The Toastytrade is now concluded; no further action is necessary.", BPAddress.Hex()))
}

func (event BurnablePaymentClaimTriggered) describeToBuyer(BPAddress common.Address) []byte {
	return []byte(fmt.Sprintf("You have triggered the autorelease for Toastytrade %s!\n\n"+
		"The ether contained in the Toastytrade has been released to your account.",
		BPAddress.Hex()))
}

func (event BurnablePaymentClosed) describeToSeller(BPAddress common.Address) []byte {
	return []byte(fmt.Sprintf("The Toastytrade at %s is now closed.\n\n"+
		"No further action is necessary.", BPAddress.Hex()))
}

func (event BurnablePaymentClosed) describeToBuyer(BPAddress common.Address) []byte {
	return []byte(fmt.Sprintf("The Toastytrade at %s is now closed.\n\n"+
		"No further action is necessary.", BPAddress.Hex()))
}

type autoreleaseWarningDescriber struct {
	autoreleaseAvailableTime time.Time
}

type autoreleaseAvailableDescriber struct {

}

func (d autoreleaseWarningDescriber) describeToSeller(BPAddress common.Address) []byte {
	return []byte(fmt.Sprintf("The Toastytrade at %s will be unlocked for the Buyer to autorelease %s.\n\n"+
		"If you've received the fiat deposit, you can ignore this message. But if you haven't, make sure to deny the Buyer's claim!",
		BPAddress.Hex(), humanize.Time(d.autoreleaseAvailableTime)))
}

func (d autoreleaseWarningDescriber) describeToBuyer(BPAddress common.Address) []byte {
	return []byte{}
}

func (d autoreleaseAvailableDescriber) describeToSeller(BPAddress common.Address) []byte {
	return []byte{}
}

func (d autoreleaseAvailableDescriber) describeToBuyer(BPAddress common.Address) []byte {
	return []byte(fmt.Sprintf("The Toastytrade at %s is now available for autorelease!\n\n"+
		"Return to the Toastytrade page here[LINK NEEDED] to withdraw your ether.",
		BPAddress.Hex()))
}

func (ttEntry toastytradeEntry) shouldSendAutoreleaseWarning() bool {
	if ttEntry.AutoreleaseWarningSent {
		return false
	}
	getAccountReqChan <- ttEntry.Seller
	accountResult := <- getAccountResChan
	if accountResult.err == leveldb.ErrNotFound {
		return false
	}	else if accountResult.err != nil {
		log.Panic(fmt.Sprintf("error getting account entry %s: ", ttEntry.Seller.Hex()), accountResult.err)
	}
	accountEntry := accountResult.entry
	warningIntervalSeconds := accountEntry.WarningIntervalMinutes*60
	return ttEntry.AutoreleaseTime.Int64() - int64(warningIntervalSeconds) < time.Now().Unix()
}

func (ttEntry toastytradeEntry) shouldSendAutoreleaseAvailable() bool {
	if ttEntry.AutoreleaseAvailableSent {
		return false
	}
	return ttEntry.AutoreleaseTime.Int64() < time.Now().Unix()
}

func (ttEntry toastytradeEntry) createAutoreleaseWarningDescriber() autoreleaseWarningDescriber {
	return autoreleaseWarningDescriber{time.Unix(ttEntry.AutoreleaseTime.Int64(), 0)}
}
