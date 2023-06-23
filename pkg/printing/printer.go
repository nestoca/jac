package printing

import (
	"github.com/TwiN/go-color"
	"github.com/nestoca/jac/pkg/live"
	"regexp"
	"strings"
)

type Printer struct {
	opts    *PrintOpts
	catalog *live.Catalog
}

const recursionLimit = 50

func NewPrinter(opts *PrintOpts, catalog *live.Catalog) *Printer {
	return &Printer{
		opts:    opts,
		catalog: catalog,
	}
}

var wordBoundariesRegex = regexp.MustCompile(`(\w+|\W+)`)

func highlight(text string) string {
	// All words/delimiters are colorized individually because
	// colorizing the whole string as a whole does not display
	// correctly when cells wrap in the table library.
	tokens := wordBoundariesRegex.FindAllString(text, -1)
	for i, token := range tokens {
		tokens[i] = color.Colorize(color.Yellow, token)
	}
	return strings.Join(tokens, "")
}

func highlightAll(texts []string) (highlighted []string) {
	for _, text := range texts {
		highlighted = append(highlighted, highlight(text))
	}
	return
}
