package raw_data

// ReportData - data
type ReportData struct {
	BalanceSheetData
	IncomeStatementData
	Quarter int
	Year    int
}

func New(is IncomeStatementData, bs BalanceSheetData) ReportData {
	return ReportData{
		IncomeStatementData: is,
		BalanceSheetData:    bs,
	}
}
