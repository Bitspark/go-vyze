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
		"'data'", "'size'", "'user'", "'value'", "'string'", "'integer'", "'float'",
		"'boolean'", "'<-'", "'{'", "'}'", "'['", "']'", "'/'", "'#'", "'.'",
		"','",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "STRING", "IDENT", "NL", "WHITESPACE",
		"DIGITS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "STRING",
		"IDENT", "NL", "WHITESPACE", "DIGITS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 29, 191, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21,
		1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 5, 24, 160, 8, 24, 10, 24, 12,
		24, 163, 9, 24, 1, 24, 1, 24, 1, 25, 4, 25, 168, 8, 25, 11, 25, 12, 25,
		169, 1, 25, 5, 25, 173, 8, 25, 10, 25, 12, 25, 176, 9, 25, 1, 26, 4, 26,
		179, 8, 26, 11, 26, 12, 26, 180, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 4,
		28, 188, 8, 28, 11, 28, 12, 28, 189, 0, 0, 29, 1, 1, 3, 2, 5, 3, 7, 4,
		9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23,
		47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 1, 0, 6, 1, 0, 34, 34,
		2, 0, 65, 90, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 10,
		10, 13, 13, 2, 0, 9, 9, 32, 32, 1, 0, 48, 57, 195, 0, 1, 1, 0, 0, 0, 0,
		3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0,
		11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0,
		0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0,
		0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0,
		0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1,
		0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49,
		1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0,
		57, 1, 0, 0, 0, 1, 59, 1, 0, 0, 0, 3, 64, 1, 0, 0, 0, 5, 67, 1, 0, 0, 0,
		7, 70, 1, 0, 0, 0, 9, 72, 1, 0, 0, 0, 11, 75, 1, 0, 0, 0, 13, 80, 1, 0,
		0, 0, 15, 88, 1, 0, 0, 0, 17, 93, 1, 0, 0, 0, 19, 98, 1, 0, 0, 0, 21, 103,
		1, 0, 0, 0, 23, 109, 1, 0, 0, 0, 25, 116, 1, 0, 0, 0, 27, 124, 1, 0, 0,
		0, 29, 130, 1, 0, 0, 0, 31, 138, 1, 0, 0, 0, 33, 141, 1, 0, 0, 0, 35, 143,
		1, 0, 0, 0, 37, 145, 1, 0, 0, 0, 39, 147, 1, 0, 0, 0, 41, 149, 1, 0, 0,
		0, 43, 151, 1, 0, 0, 0, 45, 153, 1, 0, 0, 0, 47, 155, 1, 0, 0, 0, 49, 157,
		1, 0, 0, 0, 51, 167, 1, 0, 0, 0, 53, 178, 1, 0, 0, 0, 55, 182, 1, 0, 0,
		0, 57, 187, 1, 0, 0, 0, 59, 60, 5, 112, 0, 0, 60, 61, 5, 105, 0, 0, 61,
		62, 5, 112, 0, 0, 62, 63, 5, 101, 0, 0, 63, 2, 1, 0, 0, 0, 64, 65, 5, 111,
		0, 0, 65, 66, 5, 110, 0, 0, 66, 4, 1, 0, 0, 0, 67, 68, 5, 45, 0, 0, 68,
		69, 5, 62, 0, 0, 69, 6, 1, 0, 0, 0, 70, 71, 5, 58, 0, 0, 71, 8, 1, 0, 0,
		0, 72, 73, 5, 105, 0, 0, 73, 74, 5, 100, 0, 0, 74, 10, 1, 0, 0, 0, 75,
		76, 5, 110, 0, 0, 76, 77, 5, 97, 0, 0, 77, 78, 5, 109, 0, 0, 78, 79, 5,
		101, 0, 0, 79, 12, 1, 0, 0, 0, 80, 81, 5, 99, 0, 0, 81, 82, 5, 114, 0,
		0, 82, 83, 5, 101, 0, 0, 83, 84, 5, 97, 0, 0, 84, 85, 5, 116, 0, 0, 85,
		86, 5, 101, 0, 0, 86, 87, 5, 100, 0, 0, 87, 14, 1, 0, 0, 0, 88, 89, 5,
		100, 0, 0, 89, 90, 5, 97, 0, 0, 90, 91, 5, 116, 0, 0, 91, 92, 5, 97, 0,
		0, 92, 16, 1, 0, 0, 0, 93, 94, 5, 115, 0, 0, 94, 95, 5, 105, 0, 0, 95,
		96, 5, 122, 0, 0, 96, 97, 5, 101, 0, 0, 97, 18, 1, 0, 0, 0, 98, 99, 5,
		117, 0, 0, 99, 100, 5, 115, 0, 0, 100, 101, 5, 101, 0, 0, 101, 102, 5,
		114, 0, 0, 102, 20, 1, 0, 0, 0, 103, 104, 5, 118, 0, 0, 104, 105, 5, 97,
		0, 0, 105, 106, 5, 108, 0, 0, 106, 107, 5, 117, 0, 0, 107, 108, 5, 101,
		0, 0, 108, 22, 1, 0, 0, 0, 109, 110, 5, 115, 0, 0, 110, 111, 5, 116, 0,
		0, 111, 112, 5, 114, 0, 0, 112, 113, 5, 105, 0, 0, 113, 114, 5, 110, 0,
		0, 114, 115, 5, 103, 0, 0, 115, 24, 1, 0, 0, 0, 116, 117, 5, 105, 0, 0,
		117, 118, 5, 110, 0, 0, 118, 119, 5, 116, 0, 0, 119, 120, 5, 101, 0, 0,
		120, 121, 5, 103, 0, 0, 121, 122, 5, 101, 0, 0, 122, 123, 5, 114, 0, 0,
		123, 26, 1, 0, 0, 0, 124, 125, 5, 102, 0, 0, 125, 126, 5, 108, 0, 0, 126,
		127, 5, 111, 0, 0, 127, 128, 5, 97, 0, 0, 128, 129, 5, 116, 0, 0, 129,
		28, 1, 0, 0, 0, 130, 131, 5, 98, 0, 0, 131, 132, 5, 111, 0, 0, 132, 133,
		5, 111, 0, 0, 133, 134, 5, 108, 0, 0, 134, 135, 5, 101, 0, 0, 135, 136,
		5, 97, 0, 0, 136, 137, 5, 110, 0, 0, 137, 30, 1, 0, 0, 0, 138, 139, 5,
		60, 0, 0, 139, 140, 5, 45, 0, 0, 140, 32, 1, 0, 0, 0, 141, 142, 5, 123,
		0, 0, 142, 34, 1, 0, 0, 0, 143, 144, 5, 125, 0, 0, 144, 36, 1, 0, 0, 0,
		145, 146, 5, 91, 0, 0, 146, 38, 1, 0, 0, 0, 147, 148, 5, 93, 0, 0, 148,
		40, 1, 0, 0, 0, 149, 150, 5, 47, 0, 0, 150, 42, 1, 0, 0, 0, 151, 152, 5,
		35, 0, 0, 152, 44, 1, 0, 0, 0, 153, 154, 5, 46, 0, 0, 154, 46, 1, 0, 0,
		0, 155, 156, 5, 44, 0, 0, 156, 48, 1, 0, 0, 0, 157, 161, 5, 34, 0, 0, 158,
		160, 8, 0, 0, 0, 159, 158, 1, 0, 0, 0, 160, 163, 1, 0, 0, 0, 161, 159,
		1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 164, 1, 0, 0, 0, 163, 161, 1, 0,
		0, 0, 164, 165, 5, 34, 0, 0, 165, 50, 1, 0, 0, 0, 166, 168, 7, 1, 0, 0,
		167, 166, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 169,
		170, 1, 0, 0, 0, 170, 174, 1, 0, 0, 0, 171, 173, 7, 2, 0, 0, 172, 171,
		1, 0, 0, 0, 173, 176, 1, 0, 0, 0, 174, 172, 1, 0, 0, 0, 174, 175, 1, 0,
		0, 0, 175, 52, 1, 0, 0, 0, 176, 174, 1, 0, 0, 0, 177, 179, 7, 3, 0, 0,
		178, 177, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 178, 1, 0, 0, 0, 180,
		181, 1, 0, 0, 0, 181, 54, 1, 0, 0, 0, 182, 183, 7, 4, 0, 0, 183, 184, 1,
		0, 0, 0, 184, 185, 6, 27, 0, 0, 185, 56, 1, 0, 0, 0, 186, 188, 7, 5, 0,
		0, 187, 186, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 187, 1, 0, 0, 0, 189,
		190, 1, 0, 0, 0, 190, 58, 1, 0, 0, 0, 6, 0, 161, 169, 174, 180, 189, 1,
		6, 0, 0,
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
	VyLangLexerT__20      = 21
	VyLangLexerT__21      = 22
	VyLangLexerT__22      = 23
	VyLangLexerT__23      = 24
	VyLangLexerSTRING     = 25
	VyLangLexerIDENT      = 26
	VyLangLexerNL         = 27
	VyLangLexerWHITESPACE = 28
	VyLangLexerDIGITS     = 29
)
