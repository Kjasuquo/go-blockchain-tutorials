package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	client, err := ethclient.Dial(
		"https://eth-sepolia.g.alchemy.com/v2/API-KEY",
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected Successfully")

	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Latest Block:", blockNumber)

	// GET BALANCE

	address := common.HexToAddress(
		"0x3BD82d57f63ef1CDB5D177A6B149312B4d80D390",
	)

	balance, err := client.BalanceAt(
		context.Background(),
		address,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("wei: ", balance)

	ethValue := new(big.Float)
	ethValue.SetString(balance.String())

	eth := new(big.Float).Quo(
		ethValue,
		big.NewFloat(1e18),
	)

	fmt.Println("ETH:", eth)
}
