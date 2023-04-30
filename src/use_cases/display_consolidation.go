package use_cases

import (
	"fmt"
	"investment-warlock/investor"
	"investment-warlock/utils"

	"github.com/jedib0t/go-pretty/v6/table"
)

func DisplayConsolidation(wallet investor.Wallet) {
	printConsolidationByKind(wallet)
	displayConsolidationSubmenu(wallet)
}

func printConsolidationByKind(wallet investor.Wallet) {
	utils.PrintWalletHeader(wallet)
	writer := utils.NewTableWriter()
	writer.AppendHeader(table.Row{"Tipo", "Quantidade", "Preço Médio", "R$ Patrimonio atual", "% Carteira", "% Variação", "R$ Total variação"})
	consolidation := investor.ConsolidateByKind(wallet)

	for _, consolidatedGroup := range consolidation {
		assetType := utils.TranslateKind(consolidatedGroup.Grouper)
		averagePrice := utils.CurrencyFormat(consolidatedGroup.AveragePrice)
		total := utils.CurrencyFormat(consolidatedGroup.CurrentAmount)
		percentage := utils.PercentageFormat(consolidatedGroup.WalletPercentage)
		variationPercentage := utils.PercentageFormat(consolidatedGroup.VariationPercentage)
		variationAmount := utils.CurrencyFormat(consolidatedGroup.VariationAmount)

		writer.AppendRow([]interface{}{
			assetType,
			consolidatedGroup.TotalQuantity,
			averagePrice,
			total,
			percentage,
			variationPercentage,
			variationAmount,
		})
	}

	writer.SortBy([]table.SortBy{
		{Name: "% Carteira", Mode: table.DscNumeric},
		{Name: "Total", Mode: table.DscNumeric},
		{Name: "Quantidade", Mode: table.DscNumeric},
	})

	writer.Render()
}

func displayConsolidationSubmenu(wallet investor.Wallet) {
	fmt.Println("")
	fmt.Println("1 - Exibir consolidação por ativo")
	fmt.Println("2 - Exibir consolidação setor de ativo")
	fmt.Println("3 - Voltar")
	fmt.Print("Opção: ")
	option := utils.ReadOption()
	switch option {
	case 1:
		fmt.Println("Selecionar apenas por classe de ativo: ")
		fmt.Println("1 - Fundos Imobiliários")
		fmt.Println("2 - Ações")
		fmt.Println("3 - Todos")
		filterOption := utils.ReadOption()
		filterType := translateTypeFromOption(filterOption)
		printConsolidationByAsset(wallet, filterType)
		break
	case 2:
		fmt.Println("Ainda não implementado")
	}
}

func printConsolidationByAsset(wallet investor.Wallet, filterType string) {
	utils.PrintWalletHeader(wallet)
	writer := utils.NewTableWriter()
	writer.AppendHeader(table.Row{"Ativo", "Tipo", "Quantidade", "Preço Médio", "Preço atual", "R$ Patrimonio atual", "% Carteira", "% Variação", "R$ Total variação"})
	for _, consolidatedAsset := range wallet.Consolidation {
		if filterType != "" && consolidatedAsset.Details != filterType {
			continue
		}

		assetType := utils.TranslateKind(consolidatedAsset.Details)
		quantity := consolidatedAsset.TotalQuantity
		averagePrice := utils.CurrencyFormat(consolidatedAsset.AveragePrice)
		total := utils.CurrencyFormat(consolidatedAsset.CurrentAmount)
		percentage := utils.PercentageFormat(consolidatedAsset.WalletPercentage)

		assetCurrentPrice := utils.CurrencyFormat(consolidatedAsset.AssetCurrentPrice)
		variationPercentage := utils.PercentageFormat(consolidatedAsset.VariationPercentage)
		variationAmount := utils.CurrencyFormat(consolidatedAsset.VariationAmount)

		writer.AppendRow([]interface{}{
			consolidatedAsset.Grouper,
			assetType,
			quantity,
			averagePrice,
			assetCurrentPrice,
			total,
			percentage,
			variationPercentage,
			variationAmount,
		})
	}

	writer.SortBy([]table.SortBy{
		{Name: "% Carteira", Mode: table.DscNumeric},
		{Name: "Total", Mode: table.DscNumeric},
		{Name: "Quantidade", Mode: table.DscNumeric},
	})

	writer.Render()
}

func translateTypeFromOption(option uint64) string {
	switch option {
	case 1:
		return "real_state"
	case 2:
		return "stock"
	}

	return ""
}
