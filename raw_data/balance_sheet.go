package raw_data

type BalanceSheetData struct {
	//	現金及約當現金
	CashAndCashEquivalents int

	//	透過損益按公允價值衡量之金融資產－流動
	CurrentFinancialAssetsAtFairvalueThroughProfitOrLoss int

	//	備供出售金融資產－流動淨額
	CurrentAvailableForSaleFinancialAssets int

	//	持有至到期日金融資產－流動淨額
	CurrentHeldToMaturityFinancialAssets int

	//	應收帳款淨額
	AccountsReceivableNet int

	//	應收帳款－關係人淨額
	AccountsReceivableDuefromRelatedPartiesNet int

	//	其他應收款－關係人淨額
	OtherReceivablesDueFromRelatedParties int

	//	存貨
	Inventories int

	//	其他流動資產
	OtherCurrentAssets int

	//	流動資產合計
	CurrentAssets int

	//	採用權益法之投資淨額
	InvestmentAccountedForUsingEquityMethod int

	//	不動產、廠房及設備
	PropertyPlantAndEquipment int

	//	無形資產
	IntangibleAssets int

	//	遞延所得稅資產
	DeferredTaxAssets int

	//	其他非流動資產
	OtherNoncurrentAssets int

	//	非流動資產合計
	NoncurrentAssets int

	//	短期借款
	ShorttermBorrowings int

	//	透過損益按公允價值衡量之金融負債－流動
	CurrentFinancialLiabilitiesAtFairValueThroughProfitOrLoss int

	//	避險之衍生金融負債－流動
	CurrentDerivativeFinancialLiabilitiesForHedging int

	//	應付帳款
	AccountsPayable int

	//	應付帳款－關係人
	AccountsPayableToRelatedParties int

	//	其他應付款
	OtherPayables int

	//	當期所得稅負債
	CurrentTaxLiabilities int

	//	負債準備－流動
	CurrentProvisions int

	//	其他流動負債
	OtherCurrentLiabilities int

	//	流動負債合計
	CurrentLiabilities int

	//	應付公司債
	BondsPayable int

	//	長期借款
	LongtermBorrowings int

	//	其他非流動負債
	OtherNoncurrentLiabilities int

	//	非流動負債合計
	NoncurrentLiabilities int

	//	負債總額
	Liabilities int

	//	普通股股本
	OrdinaryShare int

	//	股本合計
	CapitalStock int

	//	資本公積－發行溢價
	CapitalSurplusAdditionalPaidInCapital int

	//	資本公積－取得或處分子公司股權價格與帳面價值差額
	CapitalSurplusChangesInOwnershipInterestsInSubsidiaries int

	//	資本公積－受贈資產
	CapitalSurplusDonatedAssetsReceived int

	//	資本公積－採用權益法認列關聯企業及合資股權淨值之變動數
	CapitalSurplusChangesInEquityOfAssociatesAndJointVenturesAccountedForUsingEquityMethod int

	//	資本公積－合併溢額
	CapitalSurplusNetAssetsFromMerger int

	//	資本公積合計
	CapitalSurplus int

	//	法定盈餘公積
	LegalReserve int

	//	未分配盈餘
	UnappropriatedRetainedEarningsAaccumulatedDeficit int

	//	保留盈餘合計
	RetainedEarnings int

	//	其他權益合計
	OtherEquityInterest int

	//	歸屬於母公司業主之權益合計
	EquityAttributableToOwnersOfParent int

	//	非控制權益
	NoncontrollingInterests int

	//	權益總額
	Equity int

	//	預收股款（權益項下）之約當發行股數
	EquivalentIssueSharesOfAdvanceReceiptsForOrdinaryShare int

	//	母公司暨子公司所持有之母公司庫藏股股數
	NumberOfSharesInEntityHeldByEntityAndByItsSubsidiaries int

	//	備供出售金融資產－非流動淨額
	NonCurrentAvailableForSaleFinancialAssets int

	//	應收建造合約款
	ConstructionContractReceivable int

	//	其他應收款淨額
	OtherReceivable int

	// 本期所得稅資產
	CurrentIncomeTaxAssets int

	AccountsReceivableNet_per                             int
	CapitalStock_per                                      int
	CapitalSurplus_per                                    int
	CapitalSurplusAdditionalPaidInCapital_per             int
	CashAndCashEquivalents_per                            int
	CurrentAssets_per                                     int
	CurrentLiabilities_per                                int
	DeferredTaxAssets_per                                 int
	Equity_per                                            int
	EquityAttributableToOwnersOfParent_per                int
	IntangibleAssets_per                                  int
	Inventories_per                                       int
	LegalReserve_per                                      int
	Liabilities_per                                       int
	LongtermBorrowings_per                                int
	NoncontrollingInterests_per                           int
	NoncurrentAssets_per                                  int
	NoncurrentLiabilities_per                             int
	OrdinaryShare_per                                     int
	OtherCurrentAssets_per                                int
	OtherNoncurrentAssets_per                             int
	OtherPayables_per                                     int
	OtherReceivable_per                                   int
	PropertyPlantAndEquipment_per                         int
	RetainedEarnings_per                                  int
	UnappropriatedRetainedEarningsAaccumulatedDeficit_per int
}
