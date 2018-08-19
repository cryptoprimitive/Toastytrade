package main

import (
  "fmt"
  "github.com/ethereum/go-ethereum/common"
)

type eventDescriber interface {
  bodyToSeller(BPAddress common.Address) []byte
  bodyToBuyer(BPAddress common.Address) []byte
  subjectToSeller(BPAddress common.Address) string
  subjectToBuyer(BPAddress common.Address) string
}

func notifyParties(toastytradeAddress common.Address, event *eventDescriber) error {
  //Get Buyer/Seller eth addresses from DB, using toastytradeAddress
  //Get Buyer/Seller email addresses from DB, using eth addresses
  //Send off both emails
  return nil
}

func sendEmail(to string, subject string, body []byte) error {
  return nil
}

//No need to define new structs; we'll use the same ones defined in BurnablePayment.go

func (event BurnablePaymentCommitted) subjectToSeller(BPAddress common.Address) string {
  return fmt.Sprintf("A buyer has committed to your Toastytrade %s", BPAddress)
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
