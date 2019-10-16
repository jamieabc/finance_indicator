package main

import (
	"fmt"
	"os"

	dateRange "github.com/jamieabc/finance_indicator/date_rage"

	"github.com/jamieabc/finance_indicator/fetcher"
)

func main() {
	if len(os.Args) == 1 {
		showHelp()
		return
	}

	_, _ = fetcher.Fetch(string(os.Args[1]), dateRange.NewRange(dateRange.ThisYear))
}

func showHelp() {
	fmt.Println("Usage: finance_indicator stockID")
}
