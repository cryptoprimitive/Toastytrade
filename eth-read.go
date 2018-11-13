package main

import (
	"context"
	"github.com/SomniaStellarum/StellarUtilities/slog"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"time"
)

const ROPSTEN_FACTORY_ADDRESS = "0x649cc62BB1cF4F73827387D2B663d89570016673"

var cli *ethclient.Client

var (
	factoryAddress     common.Address
	toastytradeFactory *ToastytradeFactory
)

var (
	CreatedTopic         common.Hash
	FundsAddedTopic      common.Hash
	PayerStatementTopic  common.Hash
	WorkerStatementTopic common.Hash
	FundsRecoveredTopic  common.Hash
	CommittedTopic       common.Hash
	FundsBurnedTopic     common.Hash
	FundsReleasedTopic   common.Hash
	ClaimStartedTopic    common.Hash
	ClaimCanceledTopic   common.Hash
	ClaimTriggeredTopic  common.Hash
	ClosedTopic          common.Hash
	UnclosedTopic        common.Hash
)

func mainLoop(fromBlock uint64) {
	slog.DebugPrint("starting eth read loop")
	ctx := context.Background()

	lastBlockScanned := big.NewInt(int64(fromBlock) - 1)
	var queryStartBlock, queryEndBlock big.Int

	for {
		slog.DebugPrint("loopin")
		//Set the beginning of the queried range to the block just after the last block scanned
		//queryStartBlock = lastBlockScanned + 1
		queryStartBlock.Add(lastBlockScanned, big.NewInt(1))

		//Get latest block
		block, err := cli.HeaderByNumber(ctx, nil)
		if err != nil {
			log.Panic("error getting block number: ", err)
		}

		//Set the end of the range to the latest block
		//(We could use 'latest', but this way we have finer control, to make sure we aren't re-scanning blocks)
		queryEndBlock = *block.Number

		if queryStartBlock.Cmp(&queryEndBlock) == 1 { // If queryStartBlock > queryEndBlock
			//Indicates there has been no new block since last loop. Sleep a bit and continue.
			slog.DebugPrint("No new blocks, sleepin")
			time.Sleep(time.Minute)
			continue
		}

		slog.DebugPrint("Scanning block range: ", queryStartBlock, queryEndBlock)

		var factoryQuery ethereum.FilterQuery
		factoryQuery.FromBlock = &queryStartBlock
		factoryQuery.ToBlock = &queryEndBlock

		//This range will be used for both filter calls that follow in this loop.

		//Get all events ToastytradeCreated events from Factory contract, and extract address of new BPs

		factoryQuery.Addresses = []common.Address{factoryAddress}

		logs, err := cli.FilterLogs(ctx, factoryQuery)
		if err != nil {
			log.Panic("error filtering for factory events", err)
		}

		var newToastytradeAddresses []common.Address

		for _, log := range logs {
			event := new(ToastytradeFactoryToastytradeCreated)
			toastytradeFactory.ToastytradeFactoryFilterer.contract.UnpackLog(event, "ToastytradeCreated", log)

			newToastytradeAddresses = append(newToastytradeAddresses, event.ToastytradeAddress)
		}
		slog.DebugPrint("TTs from Factory: ", len(newToastytradeAddresses))

		//Additionally fetch all already-cached BP addresses from DB

		getAllToastytradeAddressesReqChan <- true
		result := <-getAllToastytradeAddressesResChan
		if result.err != nil {
			log.Panic(result.err)
		}
		ttAddressesFromDB := result.addresses

		slog.DebugPrint("TTs from DB: ", len(ttAddressesFromDB))

		//Concacenate the DB's adddresses and those from the newly-created BPs.
		toastytradeAddresses := append(ttAddressesFromDB, newToastytradeAddresses...)

		//create toastytradeContract instances to use their filterer in the next loop
		toastytradeContracts := make(map[common.Address]*BurnablePayment)
		for _, addr := range toastytradeAddresses {
			contract, err := NewBurnablePayment(addr, cli)
			if err != nil {
				log.Panic("error creating toastytrade @ ", addr.Hex())
			}

			toastytradeContracts[addr] = contract
		}

		//now deal with each new event from each toastytrade
		var ttQuery ethereum.FilterQuery
		ttQuery.Addresses = toastytradeAddresses
		ttQuery.FromBlock = &queryStartBlock
		ttQuery.ToBlock = &queryEndBlock

		logs, err = cli.FilterLogs(ctx, ttQuery)
		if err != nil {
			log.Panic("error filtering for toastytrade events: ", err)
		}

		for _, thisLog := range logs {
			addr := thisLog.Address

			switch thisLog.Topics[0] {
			case CreatedTopic:
				slog.DebugPrint("Created event found")
				//decode event
				event := new(BurnablePaymentCreated)
				toastytradeContracts[addr].BurnablePaymentFilterer.contract.UnpackLog(event, "Created", thisLog)

				ttEntry := new(toastytradeEntry)
				ttEntry.Seller = event.Creator

				state, _, _, _, balance, commitThreshold, _, _, _, autoreleaseInterval, _, err := toastytradeContracts[addr].BurnablePaymentCaller.GetFullState(nil)
				if err != nil {
					log.Panic("error calling bp.getFullState: ", err)
				}

				ttEntry.State = state
				ttEntry.Balance = *balance
				ttEntry.CommitThreshold = *commitThreshold
				ttEntry.AutoreleaseInterval = *autoreleaseInterval

				createdUpdate := new(createdUpdate)
				createdUpdate.ttAddr = addr
				createdUpdate.entry = ttEntry

				createdUpdateChan <- createdUpdate

				err = notifyParties(addr, event)
				if err != nil {
					log.Panic(err)
				}

			case FundsRecoveredTopic:
				slog.DebugPrint("FundsRecovered topic found")

				event := new(BurnablePaymentFundsRecovered)
				toastytradeContracts[addr].BurnablePaymentFilterer.contract.UnpackLog(event, "FundsRecovered", thisLog)

				fundsRecoveredUpdate := new(noArgUpdate)
				fundsRecoveredUpdate.ttAddr = addr

				fundsRecoveredUpdateChan <- fundsRecoveredUpdate

				err := notifyParties(addr, event)
				if err != nil {
					log.Panic(err)
				}

			case CommittedTopic:
				slog.DebugPrint("Committed event found")

				event := new(BurnablePaymentCommitted)
				toastytradeContracts[addr].BurnablePaymentFilterer.contract.UnpackLog(event, "Committed", thisLog)

				committedUpdate := new(committedUpdate)
				committedUpdate.ttAddr = addr
				committedUpdate.buyerAddr = event.Committer

				committedUpdateChan <- committedUpdate

				err := notifyParties(addr, event)
				if err != nil {
					log.Panic(err)
				}

			case FundsBurnedTopic:
				slog.DebugPrint("FundsBurned event found")

				event := new(BurnablePaymentFundsBurned)
				toastytradeContracts[addr].BurnablePaymentFilterer.contract.UnpackLog(event, "FundsBurned", thisLog)

				fundsBurnedUpdate := new(fundsBurnedUpdate)
				fundsBurnedUpdate.ttAddr = addr
				fundsBurnedUpdate.amount = *event.Amount

				fundsBurnedUpdateChan <- fundsBurnedUpdate

				err := notifyParties(addr, event)
				if err != nil {
					log.Panic(err)
				}

			case FundsReleasedTopic:
				slog.DebugPrint("FundsReleased event found")

				event := new(BurnablePaymentFundsReleased)
				toastytradeContracts[addr].BurnablePaymentFilterer.contract.UnpackLog(event, "FundsReleased", thisLog)

				fundsReleasedUpdate := new(fundsReleasedUpdate)
				fundsReleasedUpdate.ttAddr = addr
				fundsReleasedUpdate.amount = *event.Amount

				fundsReleasedUpdateChan <- fundsReleasedUpdate

				err := notifyParties(addr, event)
				if err != nil {
					log.Panic(err)
				}

			case ClaimStartedTopic:
				slog.DebugPrint("ClaimStarted event found")

				event := new(BurnablePaymentClaimStarted)
				toastytradeContracts[addr].BurnablePaymentFilterer.contract.UnpackLog(event, "ClaimStarted", thisLog)

				claimStartedUpdate := new(noArgUpdate)
				claimStartedUpdate.ttAddr = addr

				claimStartedUpdateChan <- claimStartedUpdate

				err := notifyParties(addr, event)
				if err != nil {
					log.Panic(err)
				}

			case ClaimCanceledTopic:
				slog.DebugPrint("ClaimCanceled event found")

				event := new(BurnablePaymentClaimCanceled)
				toastytradeContracts[addr].BurnablePaymentFilterer.contract.UnpackLog(event, "ClaimCanceled", thisLog)

				claimCanceledUpdate := new(noArgUpdate)
				claimCanceledUpdate.ttAddr = addr

				claimCanceledUpdateChan <- claimCanceledUpdate

				err := notifyParties(addr, event)
				if err != nil {
					log.Panic(err)
				}

			case ClaimTriggeredTopic:
				slog.DebugPrint("ClaimTriggered event found")

				event := new(BurnablePaymentClaimTriggered)
				toastytradeContracts[addr].BurnablePaymentFilterer.contract.UnpackLog(event, "ClaimTriggered", thisLog)

				claimTriggeredUpdate := new(noArgUpdate)
				claimTriggeredUpdate.ttAddr = addr

				claimTriggeredUpdateChan <- claimTriggeredUpdate

				err := notifyParties(addr, event)
				if err != nil {
					log.Panic(err)
				}

			case ClosedTopic:
				slog.DebugPrint("Closed event found")

				event := new(BurnablePaymentClosed)
				toastytradeContracts[addr].BurnablePaymentFilterer.contract.UnpackLog(event, "Closed", thisLog)

				closedUpdate := new(noArgUpdate)
				closedUpdate.ttAddr = addr

				closedUpdateChan <- closedUpdate

				err := notifyParties(addr, event)
				if err != nil {
					log.Panic(err)
				}

			}
		}

		lastBlockScanned.Set(&queryEndBlock)

		slog.DebugPrint("sleepin")

		time.Sleep(time.Minute)
	}
}
