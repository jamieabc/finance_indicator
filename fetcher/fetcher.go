package fetcher

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	dateRange "github.com/jamieabc/finance_indicator/date_rage"
)

const (
	url         = "http://finmindapi.servebeer.com/api/data"
	contentType = "application/json"
)

type dataKey int

const (
	IncomeStatement = iota
	BalanceSheet
	Dividend
)

var keyMapping map[dataKey]string

type parameter struct {
	Data    string `json:"dataset"`
	StockID string `json:"stock_id"`
	Date    string `json:"date"`
}

type dataType struct {
	Date     []string  `json:"date"`
	StockID  []string  `json:"stock_id"`
	DataType []string  `json:"type"`
	Value    []float64 `json:"value"`
}

type incomeStatement struct {
	Data dataType `json:"data"`
}

func init() {
	keyMapping = make(map[dataKey]string)
	keyMapping[IncomeStatement] = "FinancialStatements"
	keyMapping[BalanceSheet] = "BalanceSheet"
	keyMapping[Dividend] = "StockDividend"
}

func Fetch(stockID string, r dateRange.Ranger) (interface{}, error) {
	param := parameter{
		Data:    keyMapping[IncomeStatement],
		StockID: stockID,
		Date:    r.String(),
	}
	params, err := json.Marshal(param)
	if nil != err {
		return nil, err
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(params))
	if nil != err {
		return nil, err
	}
	request.Header.Set("Content-Type", contentType)
	client := &http.Client{}
	resp, err := client.Do(request)
	if nil != err {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var i incomeStatement
	fmt.Printf("body: %v\n", string(body))
	err = json.Unmarshal(body, &i)
	if nil != err {
		fmt.Printf("unmarshal error: %s\n", err)
		return nil, err
	}
	fmt.Printf("body: %v\n", i)
	return nil, nil
}
