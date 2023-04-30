package repositories

import (
	"investment-warlock/investor"
	"investment-warlock/utils"
)

func SaveTransaction(transaction investor.Transaction) investor.Transaction {
	db := utils.GetDatabase()

	// try creating if fails probably exists then updated (should improve this in the future)
	if _, errCreate := db.Create("transactions", transaction.ToMap()); errCreate != nil {
		if _, errUpdate := db.Change(transaction.Id, transaction.ToMap()); errUpdate != nil {
			panic(errUpdate)
		}
	}

	return transaction
}
