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

const (
	responseOK = 200
)

var keyMapping map[dataKey]string

type parameter struct {
	Data    string `json:"dataset"`
	StockID string `json:"stock_id"`
	Date    string `json:"date"`
}

type ResponseData struct {
	Date     []string  `json:"date"`
	StockID  []string  `json:"stock_id"`
	DataType []string  `json:"type"`
	Value    []float64 `json:"value"`
}

type response struct {
	Data   ResponseData `json:"data"`
	Status int          `json:"status"`
}

func init() {
	keyMapping = make(map[dataKey]string)
	keyMapping[IncomeStatement] = "FinancialStatements"
	keyMapping[BalanceSheet] = "BalanceSheet"
	keyMapping[Dividend] = "StockDividend"
}

// TODO: too long function
// Fetch - fetch data from server
func Fetch(stockID string, r dateRange.Ranger) (ResponseData, error) {
	param := parameter{
		Data:    keyMapping[IncomeStatement],
		StockID: stockID,
		Date:    r.String(),
	}
	params, err := json.Marshal(param)
	if nil != err {
		return ResponseData{}, err
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(params))
	if nil != err {
		return ResponseData{}, err
	}
	request.Header.Set("Content-Type", contentType)
	client := &http.Client{}
	resp, err := client.Do(request)
	if nil != err {
		return ResponseData{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		fmt.Printf("read response body with error: %s\n", err)
		return ResponseData{}, err
	}

	var rs response
	err = json.Unmarshal(body, &rs)
	if nil != err {
		fmt.Printf("unmarshal error: %s\n", err)
		return ResponseData{}, err
	}

	if rs.Status != responseOK {
		msg := fmt.Errorf("incorrect response status %d", rs.Status)
		fmt.Println(msg)
		return ResponseData{}, msg
	}

	return rs.Data, nil
}
