// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type VyLangLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var vylanglexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func vylanglexerLexerInit() {
	staticData := &vylanglexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'pipe'", "'on'", "'->'", "':'", "'id'", "'name'", "'created'",
		"'data'", "'size'", "'user'", "'value'", "'<-'", "'{'", "'}'", "'['",
		"']'", "'/'", "'#'", "'.'", "','",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "STRING", "IDENT", "NL", "WHITESPACE", "DIGITS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "STRING", "IDENT", "NL", "WHITESPACE", "DIGITS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 25, 154, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20,
		1, 20, 5, 20, 123, 8, 20, 10, 20, 12, 20, 126, 9, 20, 1, 20, 1, 20, 1,
		21, 4, 21, 131, 8, 21, 11, 21, 12, 21, 132, 1, 21, 5, 21, 136, 8, 21, 10,
		21, 12, 21, 139, 9, 21, 1, 22, 4, 22, 142, 8, 22, 11, 22, 12, 22, 143,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 4, 24, 151, 8, 24, 11, 24, 12, 24, 152,
		0, 0, 25, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19,
		10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37,
		19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 1, 0, 6, 1, 0, 34,
		34, 2, 0, 65, 90, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0,
		10, 10, 13, 13, 2, 0, 9, 9, 32, 32, 1, 0, 48, 57, 158, 0, 1, 1, 0, 0, 0,
		0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0,
		0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0,
		0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0,
		0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1,
		0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41,
		1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0,
		49, 1, 0, 0, 0, 1, 51, 1, 0, 0, 0, 3, 56, 1, 0, 0, 0, 5, 59, 1, 0, 0, 0,
		7, 62, 1, 0, 0, 0, 9, 64, 1, 0, 0, 0, 11, 67, 1, 0, 0, 0, 13, 72, 1, 0,
		0, 0, 15, 80, 1, 0, 0, 0, 17, 85, 1, 0, 0, 0, 19, 90, 1, 0, 0, 0, 21, 95,
		1, 0, 0, 0, 23, 101, 1, 0, 0, 0, 25, 104, 1, 0, 0, 0, 27, 106, 1, 0, 0,
		0, 29, 108, 1, 0, 0, 0, 31, 110, 1, 0, 0, 0, 33, 112, 1, 0, 0, 0, 35, 114,
		1, 0, 0, 0, 37, 116, 1, 0, 0, 0, 39, 118, 1, 0, 0, 0, 41, 120, 1, 0, 0,
		0, 43, 130, 1, 0, 0, 0, 45, 141, 1, 0, 0, 0, 47, 145, 1, 0, 0, 0, 49, 150,
		1, 0, 0, 0, 51, 52, 5, 112, 0, 0, 52, 53, 5, 105, 0, 0, 53, 54, 5, 112,
		0, 0, 54, 55, 5, 101, 0, 0, 55, 2, 1, 0, 0, 0, 56, 57, 5, 111, 0, 0, 57,
		58, 5, 110, 0, 0, 58, 4, 1, 0, 0, 0, 59, 60, 5, 45, 0, 0, 60, 61, 5, 62,
		0, 0, 61, 6, 1, 0, 0, 0, 62, 63, 5, 58, 0, 0, 63, 8, 1, 0, 0, 0, 64, 65,
		5, 105, 0, 0, 65, 66, 5, 100, 0, 0, 66, 10, 1, 0, 0, 0, 67, 68, 5, 110,
		0, 0, 68, 69, 5, 97, 0, 0, 69, 70, 5, 109, 0, 0, 70, 71, 5, 101, 0, 0,
		71, 12, 1, 0, 0, 0, 72, 73, 5, 99, 0, 0, 73, 74, 5, 114, 0, 0, 74, 75,
		5, 101, 0, 0, 75, 76, 5, 97, 0, 0, 76, 77, 5, 116, 0, 0, 77, 78, 5, 101,
		0, 0, 78, 79, 5, 100, 0, 0, 79, 14, 1, 0, 0, 0, 80, 81, 5, 100, 0, 0, 81,
		82, 5, 97, 0, 0, 82, 83, 5, 116, 0, 0, 83, 84, 5, 97, 0, 0, 84, 16, 1,
		0, 0, 0, 85, 86, 5, 115, 0, 0, 86, 87, 5, 105, 0, 0, 87, 88, 5, 122, 0,
		0, 88, 89, 5, 101, 0, 0, 89, 18, 1, 0, 0, 0, 90, 91, 5, 117, 0, 0, 91,
		92, 5, 115, 0, 0, 92, 93, 5, 101, 0, 0, 93, 94, 5, 114, 0, 0, 94, 20, 1,
		0, 0, 0, 95, 96, 5, 118, 0, 0, 96, 97, 5, 97, 0, 0, 97, 98, 5, 108, 0,
		0, 98, 99, 5, 117, 0, 0, 99, 100, 5, 101, 0, 0, 100, 22, 1, 0, 0, 0, 101,
		102, 5, 60, 0, 0, 102, 103, 5, 45, 0, 0, 103, 24, 1, 0, 0, 0, 104, 105,
		5, 123, 0, 0, 105, 26, 1, 0, 0, 0, 106, 107, 5, 125, 0, 0, 107, 28, 1,
		0, 0, 0, 108, 109, 5, 91, 0, 0, 109, 30, 1, 0, 0, 0, 110, 111, 5, 93, 0,
		0, 111, 32, 1, 0, 0, 0, 112, 113, 5, 47, 0, 0, 113, 34, 1, 0, 0, 0, 114,
		115, 5, 35, 0, 0, 115, 36, 1, 0, 0, 0, 116, 117, 5, 46, 0, 0, 117, 38,
		1, 0, 0, 0, 118, 119, 5, 44, 0, 0, 119, 40, 1, 0, 0, 0, 120, 124, 5, 34,
		0, 0, 121, 123, 8, 0, 0, 0, 122, 121, 1, 0, 0, 0, 123, 126, 1, 0, 0, 0,
		124, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 127, 1, 0, 0, 0, 126,
		124, 1, 0, 0, 0, 127, 128, 5, 34, 0, 0, 128, 42, 1, 0, 0, 0, 129, 131,
		7, 1, 0, 0, 130, 129, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 130, 1, 0,
		0, 0, 132, 133, 1, 0, 0, 0, 133, 137, 1, 0, 0, 0, 134, 136, 7, 2, 0, 0,
		135, 134, 1, 0, 0, 0, 136, 139, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 137,
		138, 1, 0, 0, 0, 138, 44, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 140, 142, 7,
		3, 0, 0, 141, 140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 141, 1, 0, 0,
		0, 143, 144, 1, 0, 0, 0, 144, 46, 1, 0, 0, 0, 145, 146, 7, 4, 0, 0, 146,
		147, 1, 0, 0, 0, 147, 148, 6, 23, 0, 0, 148, 48, 1, 0, 0, 0, 149, 151,
		7, 5, 0, 0, 150, 149, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 150, 1, 0,
		0, 0, 152, 153, 1, 0, 0, 0, 153, 50, 1, 0, 0, 0, 6, 0, 124, 132, 137, 143,
		152, 1, 6, 0, 0,
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

// VyLangLexerInit initializes any static state used to implement VyLangLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewVyLangLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func VyLangLexerInit() {
	staticData := &vylanglexerLexerStaticData
	staticData.once.Do(vylanglexerLexerInit)
}

// NewVyLangLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewVyLangLexer(input antlr.CharStream) *VyLangLexer {
	VyLangLexerInit()
	l := new(VyLangLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &vylanglexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "VyLang.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// VyLangLexer tokens.
const (
	VyLangLexerT__0       = 1
	VyLangLexerT__1       = 2
	VyLangLexerT__2       = 3
	VyLangLexerT__3       = 4
	VyLangLexerT__4       = 5
	VyLangLexerT__5       = 6
	VyLangLexerT__6       = 7
	VyLangLexerT__7       = 8
	VyLangLexerT__8       = 9
	VyLangLexerT__9       = 10
	VyLangLexerT__10      = 11
	VyLangLexerT__11      = 12
	VyLangLexerT__12      = 13
	VyLangLexerT__13      = 14
	VyLangLexerT__14      = 15
	VyLangLexerT__15      = 16
	VyLangLexerT__16      = 17
	VyLangLexerT__17      = 18
	VyLangLexerT__18      = 19
	VyLangLexerT__19      = 20
	VyLangLexerSTRING     = 21
	VyLangLexerIDENT      = 22
	VyLangLexerNL         = 23
	VyLangLexerWHITESPACE = 24
	VyLangLexerDIGITS     = 25
)
