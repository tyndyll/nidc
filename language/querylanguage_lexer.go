// Code generated from QueryLanguage.g4 by ANTLR 4.7.1. DO NOT EDIT.

package language

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 19, 107,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17,
	3, 18, 3, 18, 3, 19, 6, 19, 94, 10, 19, 13, 19, 14, 19, 95, 3, 20, 6, 20,
	99, 10, 20, 13, 20, 14, 20, 100, 3, 21, 6, 21, 104, 10, 21, 13, 21, 14,
	21, 105, 2, 2, 22, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10,
	19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 2, 33, 2, 35, 2, 37,
	17, 39, 18, 41, 19, 3, 2, 6, 3, 2, 50, 59, 6, 2, 50, 59, 67, 92, 97, 97,
	99, 124, 7, 2, 34, 34, 50, 59, 67, 92, 97, 97, 99, 124, 5, 2, 11, 12, 15,
	15, 34, 34, 2, 106, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2,
	2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2,
	2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3,
	2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 37,
	3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 3, 43, 3, 2, 2, 2, 5,
	45, 3, 2, 2, 2, 7, 47, 3, 2, 2, 2, 9, 49, 3, 2, 2, 2, 11, 51, 3, 2, 2,
	2, 13, 53, 3, 2, 2, 2, 15, 55, 3, 2, 2, 2, 17, 59, 3, 2, 2, 2, 19, 62,
	3, 2, 2, 2, 21, 65, 3, 2, 2, 2, 23, 69, 3, 2, 2, 2, 25, 71, 3, 2, 2, 2,
	27, 73, 3, 2, 2, 2, 29, 79, 3, 2, 2, 2, 31, 86, 3, 2, 2, 2, 33, 88, 3,
	2, 2, 2, 35, 90, 3, 2, 2, 2, 37, 93, 3, 2, 2, 2, 39, 98, 3, 2, 2, 2, 41,
	103, 3, 2, 2, 2, 43, 44, 7, 46, 2, 2, 44, 4, 3, 2, 2, 2, 45, 46, 7, 42,
	2, 2, 46, 6, 3, 2, 2, 2, 47, 48, 7, 125, 2, 2, 48, 8, 3, 2, 2, 2, 49, 50,
	7, 127, 2, 2, 50, 10, 3, 2, 2, 2, 51, 52, 7, 43, 2, 2, 52, 12, 3, 2, 2,
	2, 53, 54, 7, 60, 2, 2, 54, 14, 3, 2, 2, 2, 55, 56, 7, 62, 2, 2, 56, 57,
	7, 47, 2, 2, 57, 58, 7, 93, 2, 2, 58, 16, 3, 2, 2, 2, 59, 60, 7, 95, 2,
	2, 60, 61, 7, 47, 2, 2, 61, 18, 3, 2, 2, 2, 62, 63, 7, 47, 2, 2, 63, 64,
	7, 93, 2, 2, 64, 20, 3, 2, 2, 2, 65, 66, 7, 95, 2, 2, 66, 67, 7, 47, 2,
	2, 67, 68, 7, 64, 2, 2, 68, 22, 3, 2, 2, 2, 69, 70, 7, 48, 2, 2, 70, 24,
	3, 2, 2, 2, 71, 72, 7, 36, 2, 2, 72, 26, 3, 2, 2, 2, 73, 74, 7, 79, 2,
	2, 74, 75, 7, 67, 2, 2, 75, 76, 7, 86, 2, 2, 76, 77, 7, 69, 2, 2, 77, 78,
	7, 74, 2, 2, 78, 28, 3, 2, 2, 2, 79, 80, 7, 84, 2, 2, 80, 81, 7, 71, 2,
	2, 81, 82, 7, 86, 2, 2, 82, 83, 7, 87, 2, 2, 83, 84, 7, 84, 2, 2, 84, 85,
	7, 80, 2, 2, 85, 30, 3, 2, 2, 2, 86, 87, 9, 2, 2, 2, 87, 32, 3, 2, 2, 2,
	88, 89, 9, 3, 2, 2, 89, 34, 3, 2, 2, 2, 90, 91, 9, 4, 2, 2, 91, 36, 3,
	2, 2, 2, 92, 94, 9, 5, 2, 2, 93, 92, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95,
	93, 3, 2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 38, 3, 2, 2, 2, 97, 99, 5, 31,
	16, 2, 98, 97, 3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 100,
	101, 3, 2, 2, 2, 101, 40, 3, 2, 2, 2, 102, 104, 5, 33, 17, 2, 103, 102,
	3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 103, 3, 2, 2, 2, 105, 106, 3, 2,
	2, 2, 106, 42, 3, 2, 2, 2, 6, 2, 95, 100, 105, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "','", "'('", "'{'", "'}'", "')'", "':'", "'<-['", "']-'", "'-['",
	"']->'", "'.'", "'\"'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "Match", "Return",
	"WS", "NumberText", "AnyText",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "Match", "Return", "NUMBER_CLASSES", "UNICODE_SYMBOL_CLASSES",
	"UNICODE_SYMBOL_CLASSES_WITH_SPACES", "WS", "NumberText", "AnyText",
}

type QueryLanguageLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewQueryLanguageLexer(input antlr.CharStream) *QueryLanguageLexer {

	l := new(QueryLanguageLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "QueryLanguage.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// QueryLanguageLexer tokens.
const (
	QueryLanguageLexerT__0       = 1
	QueryLanguageLexerT__1       = 2
	QueryLanguageLexerT__2       = 3
	QueryLanguageLexerT__3       = 4
	QueryLanguageLexerT__4       = 5
	QueryLanguageLexerT__5       = 6
	QueryLanguageLexerT__6       = 7
	QueryLanguageLexerT__7       = 8
	QueryLanguageLexerT__8       = 9
	QueryLanguageLexerT__9       = 10
	QueryLanguageLexerT__10      = 11
	QueryLanguageLexerT__11      = 12
	QueryLanguageLexerMatch      = 13
	QueryLanguageLexerReturn     = 14
	QueryLanguageLexerWS         = 15
	QueryLanguageLexerNumberText = 16
	QueryLanguageLexerAnyText    = 17
)
