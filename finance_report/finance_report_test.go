package financeReport_test

import (
	"testing"

	financeReport "github.com/jamieabc/finance_indicator/finance_report"
	"github.com/stretchr/testify/assert"
)

var fixtureIS financeReport.IncomeStatementData
var fixtureBS financeReport.BalanceSheetData

func init() {
	fixtureIS = financeReport.IncomeStatementData{
		CostOfGoodsSold:                       75171000.0,
		EPS:                                   2.32,
		EXDF:                                  -12000.0,
		FinancialCost:                         98000.0,
		GrossProfit:                           251528000.0,
		IncomeAfterTaxes:                      91233000.0,
		ManagementExp:                         31916000.0,
		NonControllingTotalConsolidatedProfit: 7609000.0,
		OperatingExpenses:                     135036000.0,
		OperatingIncome:                       116492000.0,
		OtherComprehensiveIncome:              -10000.0,
		OtherNonOperatingRevenue:              1543000.0,
		OTHNOE:                                -4782000.0,
		ParentCompanyTotalConsolidatedProfit:  83614000.0,
		PreTaxIncome:                          113155000.0,
		RDExp:                                 17377000.0,
		Revenue:                               326699000.0,
		SellingExp:                            85616000.0,
		TAX:                                   21922000.0,
		TotalConsolidatedProfitForThePeriod:   91223000.0,
		TotalNonoperatingIncomeAndExpense:     -3337000.0,
	}

	fixtureBS = financeReport.BalanceSheetData{
		AccountsPayable:                           3654000,
		AccountsReceivableNet:                     88719000,
		AccountsReceivableNet_per:                 6,
		CapitalStock:                              300175008,
		CapitalStock_per:                          21,
		CapitalSurplus:                            348272992,
		CapitalSurplus_per:                        24,
		CapitalSurplusAdditionalPaidInCapital:     348135008,
		CapitalSurplusAdditionalPaidInCapital_per: 24,
		CashAndCashEquivalents:                    925836992,
		CashAndCashEquivalents_per:                65,
		CurrentAssets:                             1246294016,
		CurrentAssets_per:                         87,
		CurrentLiabilities:                        367108000,
		CurrentLiabilities_per:                    26,
		DeferredTaxAssets:                         11090000,
		DeferredTaxAssets_per:                     1,
		Equity:                                    1034286976,
		Equity_per:                                72,
		EquityAttributableToOwnersOfParent:        977667968,
		EquityAttributableToOwnersOfParent_per:    68,
		IntangibleAssets:                          12018000,
		IntangibleAssets_per:                      1,
		Inventories:                               25431000,
		Inventories_per:                           2,
		LegalReserve:                              11508000,
		LegalReserve_per:                          1,
		Liabilities:                               394116992,
		Liabilities_per:                           28,
		LongtermBorrowings:                        18057000,
		LongtermBorrowings_per:                    1,
		NoncontrollingInterests:                   56619000,
		NoncontrollingInterests_per:               4,
		NoncurrentAssets:                          182110000,
		NoncurrentAssets_per:                      13,
		NoncurrentLiabilities:                     27009000,
		NoncurrentLiabilities_per:                 2,
		OrdinaryShare:                             300175008,
		OrdinaryShare_per:                         21,
		OtherCurrentAssets:                        33539000,
		OtherCurrentAssets_per:                    2,
		OtherCurrentLiabilities:                   6536000,
		OtherEquityInterest:                       -17000,
		OtherNoncurrentAssets:                     73614000,
		OtherNoncurrentAssets_per:                 5,
		OtherNoncurrentLiabilities:                5862000,
		OtherPayables:                             90678000,
		OtherPayables_per:                         5986000,
		OtherReceivable_per:                       85388000,
		PropertyPlantAndEquipment:                 6,
		PropertyPlantAndEquipment_per:             329236992,
		RetainedEarnings:                          23,
		RetainedEarnings_per:                      317728992,
		UnappropriatedRetainedEarningsAaccumulatedDeficit:     22,
		UnappropriatedRetainedEarningsAaccumulatedDeficit_per: 3126000,
	}
}

func TestNew(t *testing.T) {
	r := financeReport.New(fixtureIS, fixtureBS)

	assert.Equal(t, 75171000.0, r.CostOfGoodsSold, "wrong IS")
	assert.Equal(t, 2.32, r.EPS, "wrong IS")
	assert.Equal(t, 17377000.0, r.RDExp, "wrong IS")
	assert.Equal(t, 21922000.0, r.TAX, "wrong IS")
	assert.Equal(t, 90678000, r.OtherPayables, "wrong BS")
	assert.Equal(t, 23, r.RetainedEarnings, "wrong BS")
}
