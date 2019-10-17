package response_parser

import (
	"reflect"
	"sort"
	"strconv"

	"github.com/jamieabc/finance_indicator/fetcher"
	"github.com/jamieabc/finance_indicator/raw_data"
)

func New(data fetcher.ResponseData) Parser {
	return &parserData{data: data}
}

const (
	Q1              = "03-31"
	Q2              = "06-30"
	Q3              = "09-30"
	Q4              = "12-31"
	monthStartIndex = 5 // 2019-06-30
	balanceSheetKey = "AccountsPayable"
)

type parserData struct {
	data fetcher.ResponseData
}

type quarterlyData map[string]interface{}

func (p parserData) Parse() []raw_data.ReportData {
	result := make([]raw_data.ReportData, 0)
	categorized := categorizeData(p.data)
	sorted := sortString(sliceKeys(p.data.Date))
	for _, v := range sorted {
		result = append(result, toObject(v, categorized[v]))
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

func toObject(date string, data quarterlyData) raw_data.ReportData {
	year, _ := strconv.ParseInt(date[0:4], 10, 32)
	result := raw_data.ReportData{
		IncomeStatementData: raw_data.IncomeStatementData{},
		BalanceSheetData:    raw_data.BalanceSheetData{},
		Quarter:             quarter(date),
		Year:                int(year),
	}

	if isBalanceSheetData(data) {
		result.BalanceSheetData = toBalanceSheet(data)
	} else {
		result.IncomeStatementData = toIncomeStatement(data)
	}
	return result
}

func quarter(date string) int {
	switch date[monthStartIndex:] {
	case Q1:
		return 1
	case Q2:
		return 2
	case Q3:
		return 3
	case Q4:
		return 4
	default:
		return 0
	}
}

func isBalanceSheetData(data quarterlyData) bool {
	if _, ok := data[balanceSheetKey]; !ok {
		return false
	}
	return true
}

func toIncomeStatement(data quarterlyData) raw_data.IncomeStatementData {
	result := raw_data.IncomeStatementData{}
	elm := reflect.ValueOf(&result).Elem()
	typeOf := elm.Type()

	for i := 0; i < elm.NumField(); i++ {
		f := elm.Field(i)
		name := typeOf.Field(i).Name
		if val, ok := data[name]; ok {
			v, _ := val.(float64)
			f.SetFloat(v)
		}
	}
	return result
}

func toBalanceSheet(data quarterlyData) raw_data.BalanceSheetData {
	result := raw_data.BalanceSheetData{}
	elm := reflect.ValueOf(&result).Elem()
	typeOf := elm.Type()

	for i := 0; i < elm.NumField(); i++ {
		f := elm.Field(i)
		name := typeOf.Field(i).Name
		if val, ok := data[name]; ok {
			v, _ := strconv.ParseInt(val.(string), 10, 32)
			f.SetInt(v)
		}
	}
	return result
}
