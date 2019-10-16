package finance_report

type IncomeStatementData struct {
	Revenue                               float64 // 營業收入合計
	CostOfGoodsSold                       float64 // 營業成本合計
	GrossProfit                           float64 // 營業毛利（毛損）淨額
	SellingExp                            float64 // 推銷費用
	ManagementExp                         float64 // 管理費用
	RDExp                                 float64 // 研究發展費用
	OperatingExpenses                     float64 // 營業費用合計
	OperatingIncome                       float64 // 營業利益（損失）
	OTHNOE                                float64 // 其他利益及損失淨額
	ASSO                                  float64 // 採用權益法認列之關聯企業及合資損益之份額淨額
	TotalConsolidatedProfitForThePeriod   float64 // 本期綜合損益總額
	OtherComprehensiveIncome              float64 // 其他綜合損益（淨額）
	ParentCompanyTotalConsolidatedProfit  float64 // 母公司業主（綜合損益）
	NonControllingTotalConsolidatedProfit float64 // 非控制權益（綜合損益）
	OtherNonOperatingRevenue              float64 // 其他收入
	FinancialCost                         float64 // 財務成本淨額
	HedgingInstrumentProfit               float64 // 避險工具之損益
	TotalNonOperatingRevenue              float64 // 營業外收入及支出合計
	UNFI                                  float64 // 備供出售金融資產未實現損益
	EXDF                                  float64 // 國外營運機構財務報表換算之兌換差額
	PreTaxIncome                          float64 // 稅前淨利（淨損）
	IncomeAfterTaxes                      float64 // 本期淨利（淨損）
	TAX                                   float64 // 所得稅費用（利益）合計
	EPS                                   float64 // 基本每股盈餘
}
