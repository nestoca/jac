package main

import (
	"fmt"
	"regexp"
	"strings"
)

type PatternFilter struct {
	expr *regexp.Regexp
}

func NewPatternFilter(pattern string) (*PatternFilter, error) {
	return NewPatternsFilter(strings.Split(pattern, ","))
}

func NewPatternsFilter(patterns []string) (*PatternFilter, error) {
	if len(patterns) == 0 {
		return &PatternFilter{expr: nil}, nil
	}
	if len(patterns) == 1 && patterns[0] == "" {
		return &PatternFilter{expr: nil}, nil
	}
	builder := strings.Builder{}
	for i, p := range patterns {
		if i > 0 {
			builder.WriteString("|")
		}
		builder.WriteString("^")
		builder.WriteString(strings.ReplaceAll(p, "*", ".*"))
		builder.WriteString("$")
	}
	expr, err := regexp.Compile(builder.String())
	if err != nil {
		return nil, fmt.Errorf("parsing patterns %q: %w", patterns, err)
	}
	return &PatternFilter{expr: expr}, nil
}

func (f *PatternFilter) Match(s string) bool {
	return f.expr == nil || f.expr.MatchString(s)
}
