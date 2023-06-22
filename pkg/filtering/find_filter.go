package filtering

import "strings"

type FindFilter struct {
	text string
}

func NewFindFilter(text string) *FindFilter {
	return &FindFilter{text: strings.ToLower(text)}
}

func (f *FindFilter) Match(values []string) bool {
	for _, value := range values {
		if strings.Contains(strings.ToLower(value), f.text) {
			return true
		}
	}
	return false
}
