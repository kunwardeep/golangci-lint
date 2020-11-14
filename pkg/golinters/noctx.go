package golinters

import (
	"github.com/sonatard/noctx"
	"golang.org/x/tools/go/analysis"

	"github.com/kunwardeep/golangci-lint/pkg/golinters/goanalysis"
)

func NewNoctx() *goanalysis.Linter {
	analyzers := []*analysis.Analyzer{
		noctx.Analyzer,
	}

	return goanalysis.NewLinter(
		"noctx",
		"noctx finds sending http request without context.Context",
		analyzers,
		nil,
	).WithLoadMode(goanalysis.LoadModeTypesInfo)
}
