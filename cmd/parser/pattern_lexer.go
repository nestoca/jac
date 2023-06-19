// Code generated from Pattern.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type PatternLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var PatternLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func patternlexerLexerInit() {
	staticData := &PatternLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "','", "'&'", "'('", "')'", "'!'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "WILDCARD", "LITERAL", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "WILDCARD", "LITERAL", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 8, 44, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 4, 5, 29, 8, 5, 11, 5, 12, 5, 30,
		1, 6, 4, 6, 34, 8, 6, 11, 6, 12, 6, 35, 1, 7, 4, 7, 39, 8, 7, 11, 7, 12,
		7, 40, 1, 7, 1, 7, 0, 0, 8, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7,
		15, 8, 1, 0, 3, 6, 0, 42, 42, 45, 45, 48, 57, 65, 90, 95, 95, 97, 122,
		5, 0, 45, 45, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32,
		32, 46, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 1, 17, 1, 0, 0, 0, 3, 19, 1, 0, 0, 0, 5, 21, 1, 0, 0, 0, 7,
		23, 1, 0, 0, 0, 9, 25, 1, 0, 0, 0, 11, 28, 1, 0, 0, 0, 13, 33, 1, 0, 0,
		0, 15, 38, 1, 0, 0, 0, 17, 18, 5, 44, 0, 0, 18, 2, 1, 0, 0, 0, 19, 20,
		5, 38, 0, 0, 20, 4, 1, 0, 0, 0, 21, 22, 5, 40, 0, 0, 22, 6, 1, 0, 0, 0,
		23, 24, 5, 41, 0, 0, 24, 8, 1, 0, 0, 0, 25, 26, 5, 33, 0, 0, 26, 10, 1,
		0, 0, 0, 27, 29, 7, 0, 0, 0, 28, 27, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 30,
		28, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 12, 1, 0, 0, 0, 32, 34, 7, 1, 0,
		0, 33, 32, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 35, 36,
		1, 0, 0, 0, 36, 14, 1, 0, 0, 0, 37, 39, 7, 2, 0, 0, 38, 37, 1, 0, 0, 0,
		39, 40, 1, 0, 0, 0, 40, 38, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 42, 1,
		0, 0, 0, 42, 43, 6, 7, 0, 0, 43, 16, 1, 0, 0, 0, 5, 0, 28, 30, 35, 40,
		1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// PatternLexerInit initializes any static state used to implement PatternLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewPatternLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func PatternLexerInit() {
	staticData := &PatternLexerLexerStaticData
	staticData.once.Do(patternlexerLexerInit)
}

// NewPatternLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewPatternLexer(input antlr.CharStream) *PatternLexer {
	PatternLexerInit()
	l := new(PatternLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &PatternLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Pattern.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// PatternLexer tokens.
const (
	PatternLexerT__0     = 1
	PatternLexerT__1     = 2
	PatternLexerT__2     = 3
	PatternLexerT__3     = 4
	PatternLexerT__4     = 5
	PatternLexerWILDCARD = 6
	PatternLexerLITERAL  = 7
	PatternLexerWS       = 8
)
