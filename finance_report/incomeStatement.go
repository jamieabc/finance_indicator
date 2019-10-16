package financeReport

type IncomeStatementData struct {
	ICostOfGoodsSold                       float64
	IEPS                                   float64
	IEXDF                                  float64
	IFinancialCost                         float64
	IGrossProfit                           float64
	IIncomeAfterTaxes                      float64
	IManagementExp                         float64
	INonControllingTotalConsolidatedProfit float64
	IOperatingExpenses                     float64
	IOperatingIncome                       float64
	IOtherComprehensiveIncome              float64
	IOtherNonOperatingRevenue              float64
	IOTHNOE                                float64
	IParentCompanyTotalConsolidatedProfit  float64
	IPreTaxIncome                          float64
	IRDExp                                 float64
	IRevenue                               float64
	ISellingExp                            float64
	ITAX                                   float64
	ITotalConsolidatedProfitForThePeriod   float64
	ITotalNonoperatingIncomeAndExpense     float64
}

func (i IncomeStatementData) CostOfGoodsSold() float64 {
	return i.ICostOfGoodsSold
}

func (i IncomeStatementData) EPS() float64 {
	return i.IEPS
}

func (i IncomeStatementData) EXDF() float64 {
	return i.IEXDF
}

func (i IncomeStatementData) FinancialCost() float64 {
	return i.IFinancialCost
}

func (i IncomeStatementData) GrossProfit() float64 {
	return i.IGrossProfit
}

func (i IncomeStatementData) IncomeAfterTaxes() float64 {
	return i.IIncomeAfterTaxes
}

func (i IncomeStatementData) ManagementExp() float64 {
	return i.IManagementExp
}

func (i IncomeStatementData) NonControllingTotalConsolidatedProfit() float64 {
	return i.INonControllingTotalConsolidatedProfit
}

func (i IncomeStatementData) OperatingExpenses() float64 {
	return i.IOperatingExpenses
}

func (i IncomeStatementData) OperatingIncome() float64 {
	return i.IOperatingIncome
}

func (i IncomeStatementData) OtherComprehensiveIncome() float64 {
	return i.IOtherComprehensiveIncome
}

func (i IncomeStatementData) OtherNonOperatingRevenue() float64 {
	return i.IOtherNonOperatingRevenue
}

func (i IncomeStatementData) OTHNOE() float64 {
	return i.IOTHNOE
}

func (i IncomeStatementData) ParentCompanyTotalConsolidatedProfit() float64 {
	return i.IParentCompanyTotalConsolidatedProfit
}

func (i IncomeStatementData) PreTaxIncome() float64 {
	return i.IPreTaxIncome
}

func (i IncomeStatementData) RDExp() float64 {
	return i.IRDExp
}

func (i IncomeStatementData) Revenue() float64 {
	return i.IRevenue
}

func (i IncomeStatementData) SellingExp() float64 {
	return i.ISellingExp
}

func (i IncomeStatementData) TAX() float64 {
	return i.ITAX
}

func (i IncomeStatementData) TotalConsolidatedProfitForThePeriod() float64 {
	return i.ITotalConsolidatedProfitForThePeriod
}

func (i IncomeStatementData) TotalNonoperatingIncomeAndExpense() float64 {
	return i.ITotalNonoperatingIncomeAndExpense
}
