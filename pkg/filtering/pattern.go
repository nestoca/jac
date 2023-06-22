package filtering

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/nestoca/jac/pkg/filtering/parser"
	"regexp"
	"strings"
)

type PatternBuilder struct {
	*parser.BasePatternListener
	stack  []Pattern
	errors []error
}

func (b *PatternBuilder) pushPattern(pattern Pattern) {
	b.stack = append(b.stack, pattern)
}

func (b *PatternBuilder) popPattern() Pattern {
	length := len(b.stack)
	if length == 0 {
		return nil
	}
	pattern := b.stack[length-1]
	b.stack = b.stack[:length-1]
	return pattern
}

func (b *PatternBuilder) peekPattern() Pattern {
	length := len(b.stack)
	if length == 0 {
		return nil
	}
	return b.stack[length-1]
}

func (b *PatternBuilder) addError(err error) {
	b.errors = append(b.errors, err)
}

func (b *PatternBuilder) EnterRoot(_ *parser.RootContext) {
	b.pushPattern(&AndPattern{})
}

func (b *PatternBuilder) EnterOr(_ *parser.OrContext) {
	or := &OrPattern{}
	b.peekPattern().(CompositePattern).AddPattern(or)
	b.pushPattern(or)
}

func (b *PatternBuilder) ExitOr(_ *parser.OrContext) {
	b.popPattern()
}

func (b *PatternBuilder) EnterAnd(_ *parser.AndContext) {
	and := &AndPattern{}
	b.peekPattern().(CompositePattern).AddPattern(and)
	b.pushPattern(and)
}

func (b *PatternBuilder) ExitAnd(_ *parser.AndContext) {
	b.popPattern()
}

func (b *PatternBuilder) EnterNot(_ *parser.NotContext) {
	not := &NotPattern{}
	b.peekPattern().(CompositePattern).AddPattern(not)
	b.pushPattern(not)
}

func (b *PatternBuilder) ExitNot(_ *parser.NotContext) {
	b.popPattern()
}

func (b *PatternBuilder) EnterWildcard(ctx *parser.WildcardContext) {
	wildcard, err := NewWildcardPattern(ctx.GetText())
	if err != nil {
		b.addError(err)
		return
	}
	b.peekPattern().(CompositePattern).AddPattern(wildcard)
}

func (b *PatternBuilder) EnterLiteral(ctx *parser.LiteralContext) {
	identifier := &LiteralPattern{
		Value: ctx.GetText(),
	}
	b.peekPattern().(CompositePattern).AddPattern(identifier)
}

func Parse(value string) (Pattern, error) {
	input := antlr.NewInputStream(value)
	lexer := parser.NewPatternLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewPatternParser(stream)
	p.AddErrorListener(antlr.NewConsoleErrorListener())
	p.BuildParseTrees = true
	builder := &PatternBuilder{}
	antlr.ParseTreeWalkerDefault.Walk(builder, p.Root())
	if len(builder.errors) > 0 {
		return nil, fmt.Errorf("parsing pattern %q: %v", value, builder.errors)
	}
	return builder.popPattern(), nil
}

type CompositePattern interface {
	AddPattern(pattern Pattern)
}

type Pattern interface {
	Match(values []string) bool
}

// *** AndPattern ***

type AndPattern struct {
	Filters []Pattern
}

func (f *AndPattern) AddPattern(pattern Pattern) {
	f.Filters = append(f.Filters, pattern)
}

func (f *AndPattern) Match(values []string) bool {
	for _, filter := range f.Filters {
		if !filter.Match(values) {
			return false
		}
	}
	return true
}

// *** OrPattern ***

type OrPattern struct {
	Filters []Pattern
}

func (f *OrPattern) AddPattern(pattern Pattern) {
	f.Filters = append(f.Filters, pattern)
}

func (f *OrPattern) Match(values []string) bool {
	for _, filter := range f.Filters {
		if filter.Match(values) {
			return true
		}
	}
	return false
}

// *** NotPattern ***

type NotPattern struct {
	Filter Pattern
}

func (f *NotPattern) AddPattern(pattern Pattern) {
	if f.Filter != nil {
		panic("not pattern already has a filter")
	}
	f.Filter = pattern
}

func (f *NotPattern) Match(values []string) bool {
	return !f.Filter.Match(values)
}

// *** WildcardPattern ***

type WildcardPattern struct {
	regex *regexp.Regexp
}

func NewWildcardPattern(value string) (*WildcardPattern, error) {
	quoted := regexp.QuoteMeta(value)
	expr, err := regexp.Compile("^" + strings.ReplaceAll(quoted, "\\*", ".*") + "$")
	if err != nil {
		return nil, fmt.Errorf("parsing wildcard expression %q: %w", value, err)
	}

	return &WildcardPattern{
		regex: expr,
	}, nil
}

func (f *WildcardPattern) Match(values []string) bool {
	for _, value := range values {
		if f.regex.MatchString(value) {
			return true
		}
	}
	return false
}

// *** LiteralPattern ***

type LiteralPattern struct {
	Value string
}

func (f *LiteralPattern) Match(values []string) bool {
	for _, value := range values {
		if value == f.Value {
			return true
		}
	}
	return false
}
