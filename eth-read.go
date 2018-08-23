package main

import (
  "log"
  "math/big"
  "context"
  "github.com/ethereum/go-ethereum"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/ethclient"
  "github.com/ethereum/go-ethereum/crypto"
	"github.com/SomniaStellarum/StellarUtilities/slog"
)

const ROPSTEN_FACTORY_ADDRESS = "0x649cc62BB1cF4F73827387D2B663d89570016673"

var cli *ethclient.Client

var (
  factoryAddress common.Address
  toastytradeFactory *ToastytradeFactory
)

var (
  CreatedTopic common.Hash
  FundsAddedTopic common.Hash
  PayerStatementTopic common.Hash
  WorkerStatementTopic common.Hash
  FundsRecoveredTopic common.Hash
  CommittedTopic common.Hash
  FundsBurnedTopic common.Hash
  FundsReleasedTopic common.Hash
  ClaimStartedTopic common.Hash
  ClaimCanceledTopic common.Hash
  ClaimTriggeredTopic common.Hash
  ClosedTopic common.Hash
  UnclosedTopic common.Hash
)

func initEthStuff() error {
  CreatedTopic = crypto.Keccak256Hash([]byte("Created(address,bool,address,uint256,uint256,string)"))
  // event Created(address indexed contractAddress, bool payerOpened, address creator, uint commitThreshold, uint autoreleaseInterval, string title);
  // event FundsAdded(address from, uint amount); //The payer has added funds to the BP.
  // event PayerStatement(string statement);
  // event WorkerStatement(string statement);
  // event FundsRecovered();
  // event Committed(address committer);
  // event FundsBurned(uint amount);
  // event FundsReleased(uint amount);
  // event ClaimStarted();
  // event ClaimCanceled();
  // event ClaimTriggered();
  // event Closed();
  // event Unclosed();

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

func ethReadLoop(fromBlock uint64) {
  queryStartBlock := big.NewInt(int64(fromBlock))

  ctx := context.Background()

  for {
    var query ethereum.FilterQuery
    query.FromBlock = queryStartBlock

    //First get all events (only one kind, ToastytradeCreated) from factory and process
    query.Addresses = []common.Address{factoryAddress}

    logs, err := cli.FilterLogs(ctx, query)
    if err != nil {
      log.Panic("error filtering for factory events", err)
    }

    var newToastytradeAddresses []common.Address
    latestBlock := fromBlock

    for _, log := range(logs) {
      event := new(ToastytradeFactoryToastytradeCreated)
      toastytradeFactory.ToastytradeFactoryFilterer.contract.UnpackLog(event, "ToastytradeCreated", log)

      newToastytradeAddresses = append(newToastytradeAddresses, event.ToastytradeAddress)

      if log.BlockNumber > latestBlock {
        latestBlock = log.BlockNumber
      }
    }

    ttAddressesFromDB, err := GetToastytradeAddresses()
    if err != nil {
      log.Panic("can't get Toastytrade addresses from db:", err)
    }
    toastytradeAddresses := append(ttAddressesFromDB, newToastytradeAddresses...)

    //create toastytradeContract instances to use their filterer in the next loop
    toastytradeContracts := make(map[common.Address]*BurnablePayment)

    for _, addr := range(toastytradeAddresses) {
      contract, err := NewBurnablePayment(addr, cli)
      if err != nil {
        log.Panic("error creating toastytrade @ ",addr.Hex())
      }

      toastytradeContracts[addr] = contract
    }

    //now deal with each new event from each toastytrade
    query.Addresses = toastytradeAddresses

    logs, err = cli.FilterLogs(ctx, query)
    if err != nil {
      log.Panic("error filtering for toastytrade events", err)
    }

    for _, log := range(logs) {
      addr := log.Address

      switch log.Topics[0] {
  		case CreatedTopic:
  			//decode event
  			event := new(BurnablePaymentCreated)
  			toastytradeContracts[addr].BurnablePaymentFilterer.contract.UnpackLog(event, "Created", log)

  			payerAddress := event.Creator // The ToastytradeFactory only creates BPs where the creator is the payer
        

  		}
    }
  }
}
