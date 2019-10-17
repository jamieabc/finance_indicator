package response_parser

import "github.com/jamieabc/finance_indicator/raw_data"

type Parser interface {
	Parse() []raw_data.ReportData
}
