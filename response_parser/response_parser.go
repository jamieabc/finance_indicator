package response_parser

import "github.com/jamieabc/finance_indicator/finance_report"

type Parser interface {
	Parse() []finance_report.ReportData
}
