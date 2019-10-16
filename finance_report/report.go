package finance_report

func New(is IncomeStatementData, bs BalanceSheetData) ReportData {
	return ReportData{
		IncomeStatementData: is,
		BalanceSheetData:    bs,
	}
}
