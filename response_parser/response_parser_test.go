package response_parser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamieabc/finance_indicator/response_parser"

	"github.com/jamieabc/finance_indicator/fetcher"
)

var fixtureBS fetcher.ResponseData
var fixtureIS fetcher.ResponseData

func init() {
	fixtureBS = fetcher.ResponseData{
		Date: []string{
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
		},
		StockID: []string{
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
		},
		DataType: []string{
			"AccountsPayable",
			"AccountsPayable_per",
			"AccountsReceivableNet",
			"AccountsReceivableNet_per",
			"CapitalStock",
			"CapitalStock_per",
			"CapitalSurplus",
			"CapitalSurplus_per",
			"CapitalSurplusAdditionalPaidInCapital",
			"CapitalSurplusAdditionalPaidInCapital_per",
			"CashAndCashEquivalents",
			"CashAndCashEquivalents_per",
			"CurrentAssets",
			"CurrentAssets_per",
			"CurrentLiabilities",
			"CurrentLiabilities_per",
			"CurrentProvisions",
			"CurrentProvisions_per",
			"DeferredTaxAssets",
			"DeferredTaxAssets_per",
			"Equity",
			"Equity_per",
			"EquityAttributableToOwnersOfParent",
			"EquityAttributableToOwnersOfParent_per",
			"EquivalentIssueSharesOfAdvanceReceiptsForOrdinaryShare",
			"IntangibleAssets",
			"IntangibleAssets_per",
			"Inventories",
			"Inventories_per",
			"LegalReserve",
			"LegalReserve_per",
			"Liabilities",
			"Liabilities_per",
			"LongtermBorrowings",
			"LongtermBorrowings_per",
			"NoncontrollingInterests",
			"NoncontrollingInterests_per",
			"NoncurrentAssets",
			"NoncurrentAssets_per",
			"NoncurrentLiabilities",
			"NoncurrentLiabilities_per",
			"NumberOfSharesInEntityHeldByEntityAndByItsSubsidiaries",
			"OrdinaryShare",
			"OrdinaryShare_per",
			"OtherCurrentAssets",
			"OtherCurrentAssets_per",
			"OtherCurrentLiabilities",
			"OtherCurrentLiabilities_per",
			"OtherEquityInterest",
			"OtherEquityInterest_per",
			"OtherNoncurrentAssets",
			"OtherNoncurrentAssets_per",
			"OtherNoncurrentLiabilities",
			"OtherNoncurrentLiabilities_per",
			"OtherPayables",
			"OtherPayables_per",
			"OtherReceivable",
			"OtherReceivable_per",
			"PropertyPlantAndEquipment",
			"PropertyPlantAndEquipment_per",
			"RetainedEarnings",
			"RetainedEarnings_per",
			"UnappropriatedRetainedEarningsAaccumulatedDeficit",
			"UnappropriatedRetainedEarningsAaccumulatedDeficit_per",
		},
		Value: []interface{}{
			"3930000",
			"0",
			"80493000",
			"4",
			"377524000",
			"20",
			"348273000",
			"19",
			"348135000",
			"19",
			"983563000",
			"53",
			"1261730000",
			"68",
			"805194000",
			"44",
			"150000",
			"0",
			"9232000",
			"1",
			"927948000",
			"50",
			"679469000",
			"37",
			"0",
			"29987000",
			"2",
			"11944000",
			"1",
			"71497000",
			"4",
			"920541000",
			"50",
			"0",
			"0",
			"248479000",
			"13",
			"586759000",
			"32",
			"115347000",
			"6",
			"1391000",
			"360210000",
			"19",
			"3414000",
			"0",
			"14221000",
			"1",
			"-1006000",
			"0",
			"214808000",
			"12",
			"8285000",
			"0",
			"351218000",
			"19",
			"6923000",
			"0",
			"74825000",
			"4",
			"249468000",
			"14",
			"177875000",
			"10",
		},
	}

	fixtureIS = fetcher.ResponseData{
		Date: []string{
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
			"2019-06-30",
		},
		StockID: []string{
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
			"8489",
		},
		DataType: []string{
			"CostOfGoodsSold",
			"EPS",
			"EXDF",
			"FinancialCost",
			"GrossProfit",
			"IncomeAfterTaxes",
			"ManagementExp",
			"NonControllingTotalConsolidatedProfit",
			"OperatingExpenses",
			"OperatingIncome",
			"OtherComprehensiveIncome",
			"OtherNonOperatingRevenue",
			"OTHNOE",
			"ParentCompanyTotalConsolidatedProfit",
			"PreTaxIncome",
			"RDExp",
			"Revenue",
			"SellingExp",
			"TAX",
			"TotalConsolidatedProfitForThePeriod",
			"TotalNonoperatingIncomeAndExpense",
		},
		Value: []interface{}{
			156796000.0,
			0.46,
			304000.0,
			911000.0,
			217822000.0,
			32460000.0,
			43282000.0,
			16168000.0,
			163784000.0,
			54038000.0,
			-337000.0,
			3747000.0,
			-4424000.0,
			15955000.0,
			52450000.0,
			16697000.0,
			374618000.0,
			103631000.0,
			19990000.0,
			32123000.0,
			-1588000.0,
		},
	}
}

func TestParseWhenBalanceSheet(t *testing.T) {
	p := response_parser.New(fixtureBS)
	assert.NotNil(t, p, "nil interface")

	result := p.Parse()
	assert.Equal(t, 2, result[0].Quarter, "wrong quarter")
	assert.Equal(t, 2019, result[0].Year, "wrong year")
	assert.Equal(t, 1, len(result), "wrong array size")
	assert.Equal(t, 10, result[0].BalanceSheetData.UnappropriatedRetainedEarningsAaccumulatedDeficit_per, "wrong result")
	assert.Equal(t, 177875000, result[0].BalanceSheetData.UnappropriatedRetainedEarningsAaccumulatedDeficit, "wrong result")
	assert.Equal(t, 3930000, result[0].BalanceSheetData.AccountsPayable, "wrong result")
}

func TestParseWhenIncomeStatement(t *testing.T) {
	p := response_parser.New(fixtureIS)
	assert.NotNil(t, p, "nil interface")

	result := p.Parse()
	assert.Equal(t, 2, result[0].Quarter, "wrong quarter")
	assert.Equal(t, 2019, result[0].Year, "wrong year")
	assert.Equal(t, 1, len(result), "wrong array size")
	assert.Equal(t, -1588000.0, result[0].IncomeStatementData.TotalNonoperatingIncomeAndExpense, "wrong result")
	assert.Equal(t, 32123000.0, result[0].IncomeStatementData.TotalConsolidatedProfitForThePeriod, "wrong result")
	assert.Equal(t, 156796000.0, result[0].IncomeStatementData.CostOfGoodsSold, "wrong result")
}
