package filtering

import "strings"

type PatternFilter struct {
	pattern Pattern
}

func NewPatternFilter(value string) (*PatternFilter, error) {
	if value == "" {
		return &PatternFilter{}, nil
	}
	pattern, err := Parse(value)
	if err != nil {
		return nil, err
	}
	return &PatternFilter{pattern: pattern}, nil
}

func NewPatternsFilter(values []string) (*PatternFilter, error) {
	return NewPatternFilter(strings.Join(values, ","))
}

func (f *PatternFilter) Match(values []string) bool {
	return f.pattern == nil || f.pattern.Match(values)
}
