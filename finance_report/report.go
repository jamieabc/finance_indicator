package financeReport

type ReportData struct {
	IncomeStatementData
}

func New(is IncomeStatementData) Reporter {
	return ReportData{
		IncomeStatementData{
			ICostOfGoodsSold:                       is.ICostOfGoodsSold,
			IEPS:                                   is.IEPS,
			IEXDF:                                  is.IEXDF,
			IFinancialCost:                         is.IFinancialCost,
			IGrossProfit:                           is.IGrossProfit,
			IIncomeAfterTaxes:                      is.IIncomeAfterTaxes,
			IManagementExp:                         is.IManagementExp,
			INonControllingTotalConsolidatedProfit: is.INonControllingTotalConsolidatedProfit,
			IOperatingExpenses:                     is.IOperatingExpenses,
			IOperatingIncome:                       is.IOperatingIncome,
			IOtherComprehensiveIncome:              is.IOtherComprehensiveIncome,
			IOtherNonOperatingRevenue:              is.IOtherNonOperatingRevenue,
			IOTHNOE:                                is.IOTHNOE,
			IParentCompanyTotalConsolidatedProfit:  is.IParentCompanyTotalConsolidatedProfit,
			IPreTaxIncome:                          is.IPreTaxIncome,
			IRDExp:                                 is.IRDExp,
			IRevenue:                               is.IRevenue,
			ISellingExp:                            is.ISellingExp,
			ITAX:                                   is.ITAX,
			ITotalConsolidatedProfitForThePeriod:   is.ITotalConsolidatedProfitForThePeriod,
			ITotalNonoperatingIncomeAndExpense:     is.ITotalNonoperatingIncomeAndExpense,
		},
	}
}
