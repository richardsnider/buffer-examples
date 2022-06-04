package main

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var BuildDateVersionLinkerFlag string
var BuildCommitLinkerFlag string

type WeightRange struct {
	NFTAddress string
	Start      float64
	End        float64
}

type NFTData struct {
	Address    string
	Attributes map[string]float64
}

func main() {
	err := processNFTs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func processNFTs() error {
	fmt.Println("BuildDateVersionLinkerFlag: ", BuildDateVersionLinkerFlag)
	fmt.Println("BuildCommitLinkerFlag: ", BuildCommitLinkerFlag)

	nftStringList := os.Getenv("NFT_ADDRESSES")
	nftAddresses := strings.Split(nftStringList, ",")
	fmt.Println("Match NFTs: ", nftAddresses)

	// TODO: add configuration options to use ropsten test network
	realBlockChainData, getNFTDatumErr := GetNFTDatum(nftAddresses)
	if getNFTDatumErr != nil {
		return getNFTDatumErr
	}

	fmt.Println(realBlockChainData)
	nftDatum := []NFTData{} //TODO: format blockchain data to fit this struct

	winningNFTAddress, computeErr := computerWinningNFT(nftDatum)

	if computeErr != nil {
		return computeErr
	}

	fmt.Println("Winning NFT Address: ", winningNFTAddress)
	return saveResults(winningNFTAddress, "./results.json")
}

func GetNFTDatum(addresses []string) ([]string, error) {
	results := []string{}

	// client, ethClientErr := ethclient.Dial("http://127.0.0.1:7545")
	client, ethClientErr := ethclient.Dial("https://mainnet.infura.io/v3/8f2a16c4b1934e8bbe3331c9a2376108")
	defer client.Close()
	if ethClientErr != nil {
		return nil, ethClientErr
	}

	for _, address := range addresses {
		addressParts := strings.Split(address, "/")
		addressHash := common.HexToAddress(addressParts[0])
		decimalVersion, _ := strconv.ParseInt(addressParts[1], 16, 32)
		fmt.Println("decimal version: ", decimalVersion)
		fmt.Println(fmt.Sprint("%X", decimalVersion))
		keyHash := common.HexToHash(fmt.Sprint("%X", decimalVersion))
		fmt.Println("Retrieving stored data for: ", addressHash, " / ", keyHash)
		storedData, err := client.CodeAt(context.Background(), addressHash, nil)
		// storedData, err := client.StorageAt(context.Background(), addressHash, keyHash, nil)
		if err != nil {
			return nil, err
		}

		results = append(results, string(storedData))
	}
	return results, nil
}

func saveResults(results map[string]string, filepath string) error {
	jsonBytes, marshalErr := json.Marshal(results)
	if marshalErr != nil {
		return marshalErr
	}

	writeErr := os.WriteFile(filepath, jsonBytes, 0777)
	if writeErr != nil {
		return writeErr
	}

	return nil
}

func computerWinningNFT(nftDatum []NFTData) (map[string]string, error) {
	winningNFTAttributes := make(map[string]string)

	commonAttributes := nftDatum[0].Attributes // TODO: compute list of common attributes
	for attributeKeyName := range commonAttributes {
		attributeWeightTotal := float64(0)
		for _, nft := range nftDatum {
			attributeWeightTotal += nft.Attributes[attributeKeyName]
		}

		weightPercentiles := []float64{}
		for index, nft := range nftDatum {
			weightPercentiles[index] = nft.Attributes[attributeKeyName] / attributeWeightTotal
		}

		weightRanges := []WeightRange{}
		for index, weightPercentile := range weightPercentiles {
			previousRangeEnd := float64(0)
			if index > 0 {
				previousRangeEnd = weightRanges[index-1].End
			}

			weightRanges[index] = WeightRange{
				NFTAddress: nftDatum[index].Address,
				Start:      previousRangeEnd,
				End:        previousRangeEnd + weightPercentile,
			}
		}

		fmt.Println("NFT Weight Ranges for attribute ", attributeKeyName, ": ", weightRanges)

		num, err := rand.Int(rand.Reader, big.NewInt(100))
		if err != nil {
			return nil, err
		}

		randFloat := float64(num.Int64()) / (1 << 63)
		fmt.Println("crypto/rand result for ", attributeKeyName, ": ", randFloat)

		for _, weightRange := range weightRanges {
			if randFloat >= weightRange.Start && randFloat < weightRange.End {
				winningNFTAttributes[attributeKeyName] = weightRange.NFTAddress
			}
		}

		if winningNFTAttributes[attributeKeyName] == "" {
			return nil, errors.New("Failed to select a winning NFT address for attribute: " + attributeKeyName)
		}
	}

	return winningNFTAttributes, nil
}
