package raw_data

type IncomeStatementData struct {
	// 營業收入合計
	Revenue float64

	// 營業成本合計
	CostOfGoodsSold float64

	// 營業毛利（毛損）淨額
	GrossProfit float64

	// 推銷費用
	SellingExp float64

	// 管理費用
	ManagementExp float64

	// 研究發展費用
	RDExp float64

	// 營業費用合計
	OperatingExpenses float64

	// 營業利益（損失）
	OperatingIncome float64

	// 其他利益及損失淨額
	OTHNOE float64

	// 採用權益法認列之關聯企業及合資損益之份額淨額
	ASSO float64

	// 本期綜合損益總額
	TotalConsolidatedProfitForThePeriod float64

	// 其他綜合損益（淨額）
	OtherComprehensiveIncome float64

	// 母公司業主（綜合損益）
	ParentCompanyTotalConsolidatedProfit float64

	// 非控制權益（綜合損益）
	NonControllingTotalConsolidatedProfit float64

	// 其他收入
	OtherNonOperatingRevenue float64

	// 財務成本淨額
	FinancialCost float64

	// 避險工具之損益
	HedgingInstrumentProfit float64

	// 營業外收入及支出合計
	TotalNonOperatingRevenue float64

	// 備供出售金融資產未實現損益
	UNFI float64

	// 國外營運機構財務報表換算之兌換差額
	EXDF float64

	// 稅前淨利（淨損）
	PreTaxIncome float64

	// 本期淨利（淨損）
	IncomeAfterTaxes float64

	// 所得稅費用（利益）合計
	TAX float64

	// 基本每股盈餘
	EPS float64

	// 營業外收入及支出合計
	TotalNonoperatingIncomeAndExpense float64
}
