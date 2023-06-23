package printing

type Printer struct {
	yaml        bool
	tree        bool
	showNames   bool
	showAll     bool
	isFiltering bool
}

const recursionLimit = 50

func NewPrinter(yaml, tree, showNames, showAll, isFiltering bool) *Printer {
	return &Printer{yaml, tree, showNames, showAll, isFiltering}
}

func highlight(text string) string {
	return "\033[33m\033[1m" + text + "\033[0m"
}

func highlightAll(texts []string) (highlighted []string) {
	for _, text := range texts {
		highlighted = append(highlighted, highlight(text))
	}
	return
}
