package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetNFTDatum2(addresses []string) ([]string, error) {
	results := []string{}

	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		return nil, err
	}

	for i, address := range addresses {
		addressParts := strings.Split(address, "/")
		addressHash := common.HexToAddress(addressParts[0])
		keyHash := common.HexToHash(addressParts[1])
		fmt.Println("Retrieving stored data for ", addressHash, " / ", keyHash)
		storedData, err := client.StorageAt(context.Background(), addressHash, keyHash, nil)
		if err != nil {
			return nil, err
		}

		results[i] = string(storedData)
	}
	return results, nil
}
