package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/victoryeo/golang-web3/example"
	"io/ioutil"
	"log"
	"math/big"
	"path/filepath"
	"strings"
)

func main() {
	/**
	 * Connecting to provider
	 */
	//client, err := ethclient.Dial("https://eth-rinkeby.alchemyapi.io/v2/<apikey>")
	client, err := ethclient.Dial("wss://eth-rinkeby.alchemyapi.io/v2/<apikey>")

	if err != nil {
		log.Fatal(err)
	}

	// with no 0x
	exampleAddress := "927Efa57F32bbD680eEd92C053D7D541bE0E7684"

	// with no 0x
	priv := "<privatekey>"

	key, err := crypto.HexToECDSA(priv)

	/**
	 * Connecting to contract at an address
	 */

	contractAddress := common.HexToAddress(exampleAddress)
	exampleClient, err := example.NewExample(contractAddress, client)
	fmt.Println("client:", exampleClient)

	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(key)
	fmt.Println("auth:", auth)

	// set the gas price manually
	gasPrice, err := client.SuggestGasPrice(context.Background())
	auth.GasPrice = gasPrice

	/**
	 * Calling contract method
	 */
	tx, err := exampleClient.Hello(auth, "hello")
	fmt.Println("tx:", tx)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Pending TX: 0x%x\n", tx.Hash())

	/**
	 * Events
	 */

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	var chLog = make(chan types.Log)
	ctx := context.Background()

	sub, err := client.SubscribeFilterLogs(ctx, query, chLog)

	if err != nil {
		log.Println("Subscribe:", err)
		return
	}

	abiPath, _ := filepath.Abs("./contracts/Example.abi")
	file, err := ioutil.ReadFile(abiPath)

	if err != nil {
		fmt.Println("Failed to read file:", err)
	}

	exampleAbi, err := abi.JSON(strings.NewReader(string(file)))
	if err != nil {
		fmt.Println("Invalid abi:", err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case log := <-chLog:
			var helloEvent struct {
				Name  string
				Count *big.Int
			}

			out, err := exampleAbi.Unpack("Hello", log.Data)

			if err != nil {
				fmt.Println("Failed to unpack:", err)
			}
			fmt.Println("After unpack:", out)

			fmt.Println("Contract:", log.Address.Hex())
			fmt.Println("Name:", out[0])
			fmt.Println("Count:", out[1])
			fmt.Println("HelloEvent:", helloEvent)
		}
	}

}
