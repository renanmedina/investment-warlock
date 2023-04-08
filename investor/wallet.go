package investor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Wallet struct {
	Id            int64         `json:"id"`
	Name          string        `json:"name"`
	Client        Client        `json:"client"`
	Transactions  []Transaction `json:"transactions"`
	Consolidation map[string]ConsolidatedAsset
}

func (w Wallet) Total() float64 {
	total := 0.0
	for _, transaction := range w.Transactions {
		total += transaction.Total()
	}

	return total
}

func (w Wallet) Consolidate() Wallet {
	consolidationMap := make(map[string]ConsolidatedAsset)

	for _, transaction := range w.Transactions {
		asset := transaction.Asset
		consolidator, alreadyOnMap := consolidationMap[asset.Ticker]

		if alreadyOnMap {
			consolidator.Add(transaction)
			continue
		}

		consolidated := NewConsolidatedAsset(asset, transaction.TotalWithoutTaxes(), transaction.Quantity, transaction.AssetPrice())
		consolidationMap[asset.Ticker] = consolidated
	}

	walletTotal := w.Total()
	for assetTicket, consolidation := range consolidationMap {
		consolidation.PercentageOf(walletTotal)
		consolidationMap[assetTicket] = consolidation
	}

	w.Consolidation = consolidationMap
	return w
}

func BuildWalletFromJsonFile(filepath string) Wallet {
	jsonFile, err := os.Open(filepath)

	if err != nil {
		fmt.Println("Erro ao ler o arquivo")
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var wallet Wallet

	json.Unmarshal(byteValue, &wallet)

	wallet = wallet.Consolidate()
	return wallet
}
