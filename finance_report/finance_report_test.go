package financeReport_test

import (
	"testing"

	financeReport "github.com/jamieabc/finance_indicator/finance_report"
	"github.com/stretchr/testify/assert"
)

var fixture financeReport.IncomeStatementData

func init() {
	fixture = financeReport.IncomeStatementData{
		ICostOfGoodsSold:                       75171000.0,
		IEPS:                                   2.32,
		IEXDF:                                  -12000.0,
		IFinancialCost:                         98000.0,
		IGrossProfit:                           251528000.0,
		IIncomeAfterTaxes:                      91233000.0,
		IManagementExp:                         31916000.0,
		INonControllingTotalConsolidatedProfit: 7609000.0,
		IOperatingExpenses:                     135036000.0,
		IOperatingIncome:                       116492000.0,
		IOtherComprehensiveIncome:              -10000.0,
		IOtherNonOperatingRevenue:              1543000.0,
		IOTHNOE:                                -4782000.0,
		IParentCompanyTotalConsolidatedProfit:  83614000.0,
		IPreTaxIncome:                          113155000.0,
		IRDExp:                                 17377000.0,
		IRevenue:                               326699000.0,
		ISellingExp:                            85616000.0,
		ITAX:                                   21922000.0,
		ITotalConsolidatedProfitForThePeriod:   91223000.0,
		ITotalNonoperatingIncomeAndExpense:     -3337000.0,
	}
}

func TestReporterIncomeStatement(t *testing.T) {
	r := financeReport.New(fixture)

	assert.Equal(t, 75171000.0, r.CostOfGoodsSold(), "wrong cogs")
	assert.Equal(t, 2.32, r.EPS(), "wrong eps")
	assert.Equal(t, 21922000.0, r.TAX(), "wrong tax")
}
