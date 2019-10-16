package response_parser

import (
	"reflect"
	"sort"
	"strconv"

	"github.com/jamieabc/finance_indicator/fetcher"
	"github.com/jamieabc/finance_indicator/finance_report"
)

func New(data fetcher.ResponseData) Parser {
	return &parserData{data: data}
}

const (
	Q1                = "03-31"
	Q2                = "06-30"
	Q3                = "09-30"
	Q4                = "12-31"
	balanceSheetCount = 60 // to decide data is bs or is
)

type parserData struct {
	data fetcher.ResponseData
}

type quarterlyData map[string]interface{}

func (p parserData) Parse() []finance_report.ReportData {
	result := make([]finance_report.ReportData, 0)
	categorized := categorizeData(p.data)
	sorted := sortString(sliceKeys(p.data.Date))
	for _, v := range sorted {
		result = append(result, toObject(categorized[v]))
	}

	return result
}

func categorizeData(d fetcher.ResponseData) map[string]quarterlyData {
	result := make(map[string]quarterlyData)

	for i, v := range d.Date {
		if _, ok := result[v]; !ok {
			result[v] = make(quarterlyData)
		}
		key := d.DataType[i]
		value := d.Value[i]
		result[v][key] = value
	}
	return result
}

func sliceKeys(s []string) []string {
	mapping := make(map[string]bool)
	for _, v := range s {
		if _, ok := mapping[v]; !ok {
			mapping[v] = true
		}
	}

	result := make([]string, 0)
	for k := range mapping {
		result = append(result, k)
	}
	return result
}

func sortString(s []string) sort.StringSlice {
	result := sort.StringSlice{}
	for _, str := range s {
		result = append(result, str)
	}
	sort.Sort(result)
	return result
}

func toObject(data quarterlyData) finance_report.ReportData {
	result := finance_report.ReportData{
		IncomeStatementData: finance_report.IncomeStatementData{},
		BalanceSheetData:    finance_report.BalanceSheetData{},
	}

	if isBalanceSheetData(data) {
		result.BalanceSheetData = toBalanceSheet(data)
	} else {
		result.IncomeStatementData = toIncomeStatement(data)
	}
	return result
}

func isBalanceSheetData(data quarterlyData) bool {
	count := keyCount(data)
	return count > balanceSheetCount
}

func keyCount(data quarterlyData) int {
	count := 0
	for range data {
		count++
	}
	return count
}

func toIncomeStatement(data quarterlyData) finance_report.IncomeStatementData {
	result := finance_report.IncomeStatementData{}
	elm := reflect.ValueOf(&result).Elem()
	typeOf := elm.Type()

	for i := 0; i < elm.NumField(); i++ {
		f := elm.Field(i)
		name := typeOf.Field(i).Name
		val, _ := data[name].(float64)
		f.SetFloat(val)
	}
	return result
}

func toBalanceSheet(data quarterlyData) finance_report.BalanceSheetData {
	result := finance_report.BalanceSheetData{}
	elm := reflect.ValueOf(&result).Elem()
	typeOf := elm.Type()

	for i := 0; i < elm.NumField(); i++ {
		f := elm.Field(i)
		name := typeOf.Field(i).Name
		val, _ := strconv.ParseInt(data[name].(string), 10, 32)
		f.SetInt(val)
	}
	return result
}
