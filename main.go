package main

import (
	"fmt"
	"os"

	"github.com/jamieabc/finance_indicator/finance_report"

	"github.com/jamieabc/finance_indicator/response_parser"
	"github.com/k0kubun/pp"

	dateRange "github.com/jamieabc/finance_indicator/date_rage"

	"github.com/jamieabc/finance_indicator/fetcher"
)

func main() {
	if len(os.Args) == 1 {
		showHelp()
		return
	}

	// TODO: merge two operations into one, abstraction is wrong, to many couplings
	respIS, err := fetcher.FetchIS(string(os.Args[1]), dateRange.NewRange(dateRange.FiveYear))
	if nil != err {
		fmt.Printf("fetch with error: %s\n", err)
		return
	}
	isParser := response_parser.New(respIS)
	is := isParser.Parse()

	respBS, err := fetcher.FetchBS(string(os.Args[1]), dateRange.NewRange(dateRange.FiveYear))
	if nil != err {
		fmt.Printf("fetch with error: %s\n", err)
		return
	}
	bsParser := response_parser.New(respBS)
	bs := bsParser.Parse()

	result := merge(is, bs)

	pp.Printf("result: %+v\n", result)
}

func showHelp() {
	fmt.Println("Usage: finance_indicator stockID")
}

func merge(is []finance_report.ReportData, bs []finance_report.ReportData) []finance_report.ReportData {
	result := make([]finance_report.ReportData, len(is))
	for i, v := range is {
		result[i] = v
		result[i].BalanceSheetData = bs[i].BalanceSheetData
	}
	return result
}
