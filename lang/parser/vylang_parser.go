// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // VyLang
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type VyLangParser struct {
	*antlr.BaseParser
}

var vylangParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func vylangParserInit() {
	staticData := &vylangParserStaticData
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
		"definitions", "definition", "namedPipe", "contextPipe", "pipe", "pipeProperty",
		"pipeNamedProperty", "pipeTerminal", "pipeField", "pipeFieldForward",
		"pipeFieldBackward", "pipeModified", "pipeMap", "pipeMapEntry", "pipeModifier",
		"pathModel", "pathRelation", "identPath", "sep",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 25, 186, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 5, 0, 40, 8, 0, 10, 0,
		12, 0, 43, 9, 0, 1, 0, 5, 0, 46, 8, 0, 10, 0, 12, 0, 49, 9, 0, 1, 0, 1,
		0, 5, 0, 53, 8, 0, 10, 0, 12, 0, 56, 9, 0, 5, 0, 58, 8, 0, 10, 0, 12, 0,
		61, 9, 0, 1, 0, 5, 0, 64, 8, 0, 10, 0, 12, 0, 67, 9, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1,
		4, 3, 4, 84, 8, 4, 1, 5, 1, 5, 3, 5, 88, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 8, 1, 8, 3, 8, 98, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 11, 3, 11, 109, 8, 11, 1, 11, 5, 11, 112, 8, 11,
		10, 11, 12, 11, 115, 9, 11, 1, 11, 1, 11, 1, 12, 1, 12, 3, 12, 121, 8,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 127, 8, 12, 10, 12, 12, 12, 130,
		9, 12, 3, 12, 132, 8, 12, 1, 12, 3, 12, 135, 8, 12, 1, 12, 1, 12, 1, 13,
		1, 13, 3, 13, 141, 8, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1,
		15, 3, 15, 150, 8, 15, 3, 15, 152, 8, 15, 3, 15, 154, 8, 15, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 162, 8, 16, 3, 16, 164, 8, 16, 3,
		16, 166, 8, 16, 1, 17, 1, 17, 1, 17, 5, 17, 171, 8, 17, 10, 17, 12, 17,
		174, 9, 17, 1, 18, 1, 18, 3, 18, 178, 8, 18, 1, 18, 1, 18, 3, 18, 182,
		8, 18, 3, 18, 184, 8, 18, 1, 18, 0, 0, 19, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 0, 1, 1, 0, 5, 11, 191, 0, 41,
		1, 0, 0, 0, 2, 70, 1, 0, 0, 0, 4, 72, 1, 0, 0, 0, 6, 76, 1, 0, 0, 0, 8,
		83, 1, 0, 0, 0, 10, 87, 1, 0, 0, 0, 12, 89, 1, 0, 0, 0, 14, 93, 1, 0, 0,
		0, 16, 97, 1, 0, 0, 0, 18, 99, 1, 0, 0, 0, 20, 103, 1, 0, 0, 0, 22, 108,
		1, 0, 0, 0, 24, 118, 1, 0, 0, 0, 26, 140, 1, 0, 0, 0, 28, 142, 1, 0, 0,
		0, 30, 153, 1, 0, 0, 0, 32, 165, 1, 0, 0, 0, 34, 167, 1, 0, 0, 0, 36, 183,
		1, 0, 0, 0, 38, 40, 5, 23, 0, 0, 39, 38, 1, 0, 0, 0, 40, 43, 1, 0, 0, 0,
		41, 39, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 59, 1, 0, 0, 0, 43, 41, 1,
		0, 0, 0, 44, 46, 5, 23, 0, 0, 45, 44, 1, 0, 0, 0, 46, 49, 1, 0, 0, 0, 47,
		45, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 50, 1, 0, 0, 0, 49, 47, 1, 0, 0,
		0, 50, 54, 3, 2, 1, 0, 51, 53, 5, 23, 0, 0, 52, 51, 1, 0, 0, 0, 53, 56,
		1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 58, 1, 0, 0, 0,
		56, 54, 1, 0, 0, 0, 57, 47, 1, 0, 0, 0, 58, 61, 1, 0, 0, 0, 59, 57, 1,
		0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 65, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 62,
		64, 5, 23, 0, 0, 63, 62, 1, 0, 0, 0, 64, 67, 1, 0, 0, 0, 65, 63, 1, 0,
		0, 0, 65, 66, 1, 0, 0, 0, 66, 68, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 68, 69,
		5, 0, 0, 1, 69, 1, 1, 0, 0, 0, 70, 71, 3, 4, 2, 0, 71, 3, 1, 0, 0, 0, 72,
		73, 5, 1, 0, 0, 73, 74, 5, 22, 0, 0, 74, 75, 3, 6, 3, 0, 75, 5, 1, 0, 0,
		0, 76, 77, 5, 2, 0, 0, 77, 78, 3, 30, 15, 0, 78, 79, 5, 3, 0, 0, 79, 80,
		3, 22, 11, 0, 80, 7, 1, 0, 0, 0, 81, 84, 3, 10, 5, 0, 82, 84, 3, 24, 12,
		0, 83, 81, 1, 0, 0, 0, 83, 82, 1, 0, 0, 0, 84, 9, 1, 0, 0, 0, 85, 88, 3,
		14, 7, 0, 86, 88, 3, 16, 8, 0, 87, 85, 1, 0, 0, 0, 87, 86, 1, 0, 0, 0,
		88, 11, 1, 0, 0, 0, 89, 90, 5, 22, 0, 0, 90, 91, 5, 4, 0, 0, 91, 92, 3,
		10, 5, 0, 92, 13, 1, 0, 0, 0, 93, 94, 7, 0, 0, 0, 94, 15, 1, 0, 0, 0, 95,
		98, 3, 18, 9, 0, 96, 98, 3, 20, 10, 0, 97, 95, 1, 0, 0, 0, 97, 96, 1, 0,
		0, 0, 98, 17, 1, 0, 0, 0, 99, 100, 3, 32, 16, 0, 100, 101, 5, 3, 0, 0,
		101, 102, 3, 22, 11, 0, 102, 19, 1, 0, 0, 0, 103, 104, 5, 12, 0, 0, 104,
		105, 3, 32, 16, 0, 105, 106, 3, 22, 11, 0, 106, 21, 1, 0, 0, 0, 107, 109,
		3, 28, 14, 0, 108, 107, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 113, 1,
		0, 0, 0, 110, 112, 5, 23, 0, 0, 111, 110, 1, 0, 0, 0, 112, 115, 1, 0, 0,
		0, 113, 111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 116, 1, 0, 0, 0, 115,
		113, 1, 0, 0, 0, 116, 117, 3, 8, 4, 0, 117, 23, 1, 0, 0, 0, 118, 120, 5,
		13, 0, 0, 119, 121, 5, 23, 0, 0, 120, 119, 1, 0, 0, 0, 120, 121, 1, 0,
		0, 0, 121, 131, 1, 0, 0, 0, 122, 128, 3, 26, 13, 0, 123, 124, 3, 36, 18,
		0, 124, 125, 3, 26, 13, 0, 125, 127, 1, 0, 0, 0, 126, 123, 1, 0, 0, 0,
		127, 130, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129,
		132, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 131, 122, 1, 0, 0, 0, 131, 132,
		1, 0, 0, 0, 132, 134, 1, 0, 0, 0, 133, 135, 5, 23, 0, 0, 134, 133, 1, 0,
		0, 0, 134, 135, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 137, 5, 14, 0, 0,
		137, 25, 1, 0, 0, 0, 138, 141, 3, 12, 6, 0, 139, 141, 3, 10, 5, 0, 140,
		138, 1, 0, 0, 0, 140, 139, 1, 0, 0, 0, 141, 27, 1, 0, 0, 0, 142, 143, 5,
		15, 0, 0, 143, 144, 5, 16, 0, 0, 144, 29, 1, 0, 0, 0, 145, 154, 5, 22,
		0, 0, 146, 151, 3, 34, 17, 0, 147, 149, 5, 17, 0, 0, 148, 150, 5, 22, 0,
		0, 149, 148, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 152, 1, 0, 0, 0, 151,
		147, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 154, 1, 0, 0, 0, 153, 145,
		1, 0, 0, 0, 153, 146, 1, 0, 0, 0, 154, 31, 1, 0, 0, 0, 155, 166, 5, 22,
		0, 0, 156, 157, 3, 34, 17, 0, 157, 158, 5, 18, 0, 0, 158, 163, 5, 22, 0,
		0, 159, 161, 5, 17, 0, 0, 160, 162, 5, 22, 0, 0, 161, 160, 1, 0, 0, 0,
		161, 162, 1, 0, 0, 0, 162, 164, 1, 0, 0, 0, 163, 159, 1, 0, 0, 0, 163,
		164, 1, 0, 0, 0, 164, 166, 1, 0, 0, 0, 165, 155, 1, 0, 0, 0, 165, 156,
		1, 0, 0, 0, 166, 33, 1, 0, 0, 0, 167, 172, 5, 22, 0, 0, 168, 169, 5, 19,
		0, 0, 169, 171, 5, 22, 0, 0, 170, 168, 1, 0, 0, 0, 171, 174, 1, 0, 0, 0,
		172, 170, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 35, 1, 0, 0, 0, 174, 172,
		1, 0, 0, 0, 175, 184, 5, 23, 0, 0, 176, 178, 5, 23, 0, 0, 177, 176, 1,
		0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 181, 5, 20, 0,
		0, 180, 182, 5, 23, 0, 0, 181, 180, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182,
		184, 1, 0, 0, 0, 183, 175, 1, 0, 0, 0, 183, 177, 1, 0, 0, 0, 184, 37, 1,
		0, 0, 0, 25, 41, 47, 54, 59, 65, 83, 87, 97, 108, 113, 120, 128, 131, 134,
		140, 149, 151, 153, 161, 163, 165, 172, 177, 181, 183,
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

// VyLangParserInit initializes any static state used to implement VyLangParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewVyLangParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func VyLangParserInit() {
	staticData := &vylangParserStaticData
	staticData.once.Do(vylangParserInit)
}

// NewVyLangParser produces a new parser instance for the optional input antlr.TokenStream.
func NewVyLangParser(input antlr.TokenStream) *VyLangParser {
	VyLangParserInit()
	this := new(VyLangParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &vylangParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// VyLangParser tokens.
const (
	VyLangParserEOF        = antlr.TokenEOF
	VyLangParserT__0       = 1
	VyLangParserT__1       = 2
	VyLangParserT__2       = 3
	VyLangParserT__3       = 4
	VyLangParserT__4       = 5
	VyLangParserT__5       = 6
	VyLangParserT__6       = 7
	VyLangParserT__7       = 8
	VyLangParserT__8       = 9
	VyLangParserT__9       = 10
	VyLangParserT__10      = 11
	VyLangParserT__11      = 12
	VyLangParserT__12      = 13
	VyLangParserT__13      = 14
	VyLangParserT__14      = 15
	VyLangParserT__15      = 16
	VyLangParserT__16      = 17
	VyLangParserT__17      = 18
	VyLangParserT__18      = 19
	VyLangParserT__19      = 20
	VyLangParserSTRING     = 21
	VyLangParserIDENT      = 22
	VyLangParserNL         = 23
	VyLangParserWHITESPACE = 24
	VyLangParserDIGITS     = 25
)

// VyLangParser rules.
const (
	VyLangParserRULE_definitions       = 0
	VyLangParserRULE_definition        = 1
	VyLangParserRULE_namedPipe         = 2
	VyLangParserRULE_contextPipe       = 3
	VyLangParserRULE_pipe              = 4
	VyLangParserRULE_pipeProperty      = 5
	VyLangParserRULE_pipeNamedProperty = 6
	VyLangParserRULE_pipeTerminal      = 7
	VyLangParserRULE_pipeField         = 8
	VyLangParserRULE_pipeFieldForward  = 9
	VyLangParserRULE_pipeFieldBackward = 10
	VyLangParserRULE_pipeModified      = 11
	VyLangParserRULE_pipeMap           = 12
	VyLangParserRULE_pipeMapEntry      = 13
	VyLangParserRULE_pipeModifier      = 14
	VyLangParserRULE_pathModel         = 15
	VyLangParserRULE_pathRelation      = 16
	VyLangParserRULE_identPath         = 17
	VyLangParserRULE_sep               = 18
)

// IDefinitionsContext is an interface to support dynamic dispatch.
type IDefinitionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefinitionsContext differentiates from other interfaces.
	IsDefinitionsContext()
}

type DefinitionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionsContext() *DefinitionsContext {
	var p = new(DefinitionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_definitions
	return p
}

func (*DefinitionsContext) IsDefinitionsContext() {}

func NewDefinitionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionsContext {
	var p = new(DefinitionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_definitions

	return p
}

func (s *DefinitionsContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionsContext) EOF() antlr.TerminalNode {
	return s.GetToken(VyLangParserEOF, 0)
}

func (s *DefinitionsContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(VyLangParserNL)
}

func (s *DefinitionsContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(VyLangParserNL, i)
}

func (s *DefinitionsContext) AllDefinition() []IDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDefinitionContext); ok {
			tst[i] = t.(IDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *DefinitionsContext) Definition(i int) IDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefinitionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *DefinitionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterDefinitions(s)
	}
}

func (s *DefinitionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitDefinitions(s)
	}
}

func (p *VyLangParser) Definitions() (localctx IDefinitionsContext) {
	this := p
	_ = this

	localctx = NewDefinitionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, VyLangParserRULE_definitions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(38)
				p.Match(VyLangParserNL)
			}

		}
		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(47)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == VyLangParserNL {
				{
					p.SetState(44)
					p.Match(VyLangParserNL)
				}

				p.SetState(49)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(50)
				p.Definition()
			}
			p.SetState(54)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

			for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				if _alt == 1 {
					{
						p.SetState(51)
						p.Match(VyLangParserNL)
					}

				}
				p.SetState(56)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
			}

		}
		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == VyLangParserNL {
		{
			p.SetState(62)
			p.Match(VyLangParserNL)
		}

		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(68)
		p.Match(VyLangParserEOF)
	}

	return localctx
}

// IDefinitionContext is an interface to support dynamic dispatch.
type IDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefinitionContext differentiates from other interfaces.
	IsDefinitionContext()
}

type DefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionContext() *DefinitionContext {
	var p = new(DefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_definition
	return p
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) NamedPipe() INamedPipeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamedPipeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamedPipeContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterDefinition(s)
	}
}

func (s *DefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitDefinition(s)
	}
}

func (p *VyLangParser) Definition() (localctx IDefinitionContext) {
	this := p
	_ = this

	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, VyLangParserRULE_definition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.NamedPipe()
	}

	return localctx
}

// INamedPipeContext is an interface to support dynamic dispatch.
type INamedPipeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamedPipeContext differentiates from other interfaces.
	IsNamedPipeContext()
}

type NamedPipeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamedPipeContext() *NamedPipeContext {
	var p = new(NamedPipeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_namedPipe
	return p
}

func (*NamedPipeContext) IsNamedPipeContext() {}

func NewNamedPipeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamedPipeContext {
	var p = new(NamedPipeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_namedPipe

	return p
}

func (s *NamedPipeContext) GetParser() antlr.Parser { return s.parser }

func (s *NamedPipeContext) IDENT() antlr.TerminalNode {
	return s.GetToken(VyLangParserIDENT, 0)
}

func (s *NamedPipeContext) ContextPipe() IContextPipeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContextPipeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContextPipeContext)
}

func (s *NamedPipeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedPipeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamedPipeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterNamedPipe(s)
	}
}

func (s *NamedPipeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitNamedPipe(s)
	}
}

func (p *VyLangParser) NamedPipe() (localctx INamedPipeContext) {
	this := p
	_ = this

	localctx = NewNamedPipeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, VyLangParserRULE_namedPipe)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Match(VyLangParserT__0)
	}
	{
		p.SetState(73)
		p.Match(VyLangParserIDENT)
	}
	{
		p.SetState(74)
		p.ContextPipe()
	}

	return localctx
}

// IContextPipeContext is an interface to support dynamic dispatch.
type IContextPipeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContextPipeContext differentiates from other interfaces.
	IsContextPipeContext()
}

type ContextPipeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContextPipeContext() *ContextPipeContext {
	var p = new(ContextPipeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_contextPipe
	return p
}

func (*ContextPipeContext) IsContextPipeContext() {}

func NewContextPipeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContextPipeContext {
	var p = new(ContextPipeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_contextPipe

	return p
}

func (s *ContextPipeContext) GetParser() antlr.Parser { return s.parser }

func (s *ContextPipeContext) PathModel() IPathModelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathModelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathModelContext)
}

func (s *ContextPipeContext) PipeModified() IPipeModifiedContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPipeModifiedContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPipeModifiedContext)
}

func (s *ContextPipeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContextPipeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContextPipeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterContextPipe(s)
	}
}

func (s *ContextPipeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitContextPipe(s)
	}
}

func (p *VyLangParser) ContextPipe() (localctx IContextPipeContext) {
	this := p
	_ = this

	localctx = NewContextPipeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, VyLangParserRULE_contextPipe)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		p.Match(VyLangParserT__1)
	}
	{
		p.SetState(77)
		p.PathModel()
	}
	{
		p.SetState(78)
		p.Match(VyLangParserT__2)
	}
	{
		p.SetState(79)
		p.PipeModified()
	}

	return localctx
}

// IPipeContext is an interface to support dynamic dispatch.
type IPipeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPipeContext differentiates from other interfaces.
	IsPipeContext()
}

type PipeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipeContext() *PipeContext {
	var p = new(PipeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_pipe
	return p
}

func (*PipeContext) IsPipeContext() {}

func NewPipeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipeContext {
	var p = new(PipeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_pipe

	return p
}

func (s *PipeContext) GetParser() antlr.Parser { return s.parser }

func (s *PipeContext) PipeProperty() IPipePropertyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPipePropertyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPipePropertyContext)
}

func (s *PipeContext) PipeMap() IPipeMapContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPipeMapContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPipeMapContext)
}

func (s *PipeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterPipe(s)
	}
}

func (s *PipeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitPipe(s)
	}
}

func (p *VyLangParser) Pipe() (localctx IPipeContext) {
	this := p
	_ = this

	localctx = NewPipeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, VyLangParserRULE_pipe)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(83)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case VyLangParserT__4, VyLangParserT__5, VyLangParserT__6, VyLangParserT__7, VyLangParserT__8, VyLangParserT__9, VyLangParserT__10, VyLangParserT__11, VyLangParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(81)
			p.PipeProperty()
		}

	case VyLangParserT__12:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(82)
			p.PipeMap()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPipePropertyContext is an interface to support dynamic dispatch.
type IPipePropertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPipePropertyContext differentiates from other interfaces.
	IsPipePropertyContext()
}

type PipePropertyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipePropertyContext() *PipePropertyContext {
	var p = new(PipePropertyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_pipeProperty
	return p
}

func (*PipePropertyContext) IsPipePropertyContext() {}

func NewPipePropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipePropertyContext {
	var p = new(PipePropertyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_pipeProperty

	return p
}

func (s *PipePropertyContext) GetParser() antlr.Parser { return s.parser }

func (s *PipePropertyContext) PipeTerminal() IPipeTerminalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPipeTerminalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPipeTerminalContext)
}

func (s *PipePropertyContext) PipeField() IPipeFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPipeFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPipeFieldContext)
}

func (s *PipePropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipePropertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipePropertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterPipeProperty(s)
	}
}

func (s *PipePropertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitPipeProperty(s)
	}
}

func (p *VyLangParser) PipeProperty() (localctx IPipePropertyContext) {
	this := p
	_ = this

	localctx = NewPipePropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, VyLangParserRULE_pipeProperty)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(87)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case VyLangParserT__4, VyLangParserT__5, VyLangParserT__6, VyLangParserT__7, VyLangParserT__8, VyLangParserT__9, VyLangParserT__10:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(85)
			p.PipeTerminal()
		}

	case VyLangParserT__11, VyLangParserIDENT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(86)
			p.PipeField()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPipeNamedPropertyContext is an interface to support dynamic dispatch.
type IPipeNamedPropertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPipeNamedPropertyContext differentiates from other interfaces.
	IsPipeNamedPropertyContext()
}

type PipeNamedPropertyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipeNamedPropertyContext() *PipeNamedPropertyContext {
	var p = new(PipeNamedPropertyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_pipeNamedProperty
	return p
}

func (*PipeNamedPropertyContext) IsPipeNamedPropertyContext() {}

func NewPipeNamedPropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipeNamedPropertyContext {
	var p = new(PipeNamedPropertyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_pipeNamedProperty

	return p
}

func (s *PipeNamedPropertyContext) GetParser() antlr.Parser { return s.parser }

func (s *PipeNamedPropertyContext) IDENT() antlr.TerminalNode {
	return s.GetToken(VyLangParserIDENT, 0)
}

func (s *PipeNamedPropertyContext) PipeProperty() IPipePropertyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPipePropertyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPipePropertyContext)
}

func (s *PipeNamedPropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipeNamedPropertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipeNamedPropertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterPipeNamedProperty(s)
	}
}

func (s *PipeNamedPropertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitPipeNamedProperty(s)
	}
}

func (p *VyLangParser) PipeNamedProperty() (localctx IPipeNamedPropertyContext) {
	this := p
	_ = this

	localctx = NewPipeNamedPropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, VyLangParserRULE_pipeNamedProperty)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.Match(VyLangParserIDENT)
	}
	{
		p.SetState(90)
		p.Match(VyLangParserT__3)
	}
	{
		p.SetState(91)
		p.PipeProperty()
	}

	return localctx
}

// IPipeTerminalContext is an interface to support dynamic dispatch.
type IPipeTerminalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPipeTerminalContext differentiates from other interfaces.
	IsPipeTerminalContext()
}

type PipeTerminalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipeTerminalContext() *PipeTerminalContext {
	var p = new(PipeTerminalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_pipeTerminal
	return p
}

func (*PipeTerminalContext) IsPipeTerminalContext() {}

func NewPipeTerminalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipeTerminalContext {
	var p = new(PipeTerminalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_pipeTerminal

	return p
}

func (s *PipeTerminalContext) GetParser() antlr.Parser { return s.parser }
func (s *PipeTerminalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipeTerminalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipeTerminalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterPipeTerminal(s)
	}
}

func (s *PipeTerminalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitPipeTerminal(s)
	}
}

func (p *VyLangParser) PipeTerminal() (localctx IPipeTerminalContext) {
	this := p
	_ = this

	localctx = NewPipeTerminalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, VyLangParserRULE_pipeTerminal)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4064) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IPipeFieldContext is an interface to support dynamic dispatch.
type IPipeFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPipeFieldContext differentiates from other interfaces.
	IsPipeFieldContext()
}

type PipeFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipeFieldContext() *PipeFieldContext {
	var p = new(PipeFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_pipeField
	return p
}

func (*PipeFieldContext) IsPipeFieldContext() {}

func NewPipeFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipeFieldContext {
	var p = new(PipeFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_pipeField

	return p
}

func (s *PipeFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *PipeFieldContext) PipeFieldForward() IPipeFieldForwardContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPipeFieldForwardContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPipeFieldForwardContext)
}

func (s *PipeFieldContext) PipeFieldBackward() IPipeFieldBackwardContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPipeFieldBackwardContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPipeFieldBackwardContext)
}

func (s *PipeFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipeFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipeFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterPipeField(s)
	}
}

func (s *PipeFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitPipeField(s)
	}
}

func (p *VyLangParser) PipeField() (localctx IPipeFieldContext) {
	this := p
	_ = this

	localctx = NewPipeFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, VyLangParserRULE_pipeField)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(97)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case VyLangParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(95)
			p.PipeFieldForward()
		}

	case VyLangParserT__11:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(96)
			p.PipeFieldBackward()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPipeFieldForwardContext is an interface to support dynamic dispatch.
type IPipeFieldForwardContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPipeFieldForwardContext differentiates from other interfaces.
	IsPipeFieldForwardContext()
}

type PipeFieldForwardContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipeFieldForwardContext() *PipeFieldForwardContext {
	var p = new(PipeFieldForwardContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_pipeFieldForward
	return p
}

func (*PipeFieldForwardContext) IsPipeFieldForwardContext() {}

func NewPipeFieldForwardContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipeFieldForwardContext {
	var p = new(PipeFieldForwardContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_pipeFieldForward

	return p
}

func (s *PipeFieldForwardContext) GetParser() antlr.Parser { return s.parser }

func (s *PipeFieldForwardContext) PathRelation() IPathRelationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathRelationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathRelationContext)
}

func (s *PipeFieldForwardContext) PipeModified() IPipeModifiedContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPipeModifiedContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPipeModifiedContext)
}

func (s *PipeFieldForwardContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipeFieldForwardContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipeFieldForwardContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterPipeFieldForward(s)
	}
}

func (s *PipeFieldForwardContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitPipeFieldForward(s)
	}
}

func (p *VyLangParser) PipeFieldForward() (localctx IPipeFieldForwardContext) {
	this := p
	_ = this

	localctx = NewPipeFieldForwardContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, VyLangParserRULE_pipeFieldForward)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.PathRelation()
	}
	{
		p.SetState(100)
		p.Match(VyLangParserT__2)
	}
	{
		p.SetState(101)
		p.PipeModified()
	}

	return localctx
}

// IPipeFieldBackwardContext is an interface to support dynamic dispatch.
type IPipeFieldBackwardContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPipeFieldBackwardContext differentiates from other interfaces.
	IsPipeFieldBackwardContext()
}

type PipeFieldBackwardContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipeFieldBackwardContext() *PipeFieldBackwardContext {
	var p = new(PipeFieldBackwardContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_pipeFieldBackward
	return p
}

func (*PipeFieldBackwardContext) IsPipeFieldBackwardContext() {}

func NewPipeFieldBackwardContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipeFieldBackwardContext {
	var p = new(PipeFieldBackwardContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_pipeFieldBackward

	return p
}

func (s *PipeFieldBackwardContext) GetParser() antlr.Parser { return s.parser }

func (s *PipeFieldBackwardContext) PathRelation() IPathRelationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathRelationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathRelationContext)
}

func (s *PipeFieldBackwardContext) PipeModified() IPipeModifiedContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPipeModifiedContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPipeModifiedContext)
}

func (s *PipeFieldBackwardContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipeFieldBackwardContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipeFieldBackwardContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterPipeFieldBackward(s)
	}
}

func (s *PipeFieldBackwardContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitPipeFieldBackward(s)
	}
}

func (p *VyLangParser) PipeFieldBackward() (localctx IPipeFieldBackwardContext) {
	this := p
	_ = this

	localctx = NewPipeFieldBackwardContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, VyLangParserRULE_pipeFieldBackward)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Match(VyLangParserT__11)
	}
	{
		p.SetState(104)
		p.PathRelation()
	}
	{
		p.SetState(105)
		p.PipeModified()
	}

	return localctx
}

// IPipeModifiedContext is an interface to support dynamic dispatch.
type IPipeModifiedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPipeModifiedContext differentiates from other interfaces.
	IsPipeModifiedContext()
}

type PipeModifiedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipeModifiedContext() *PipeModifiedContext {
	var p = new(PipeModifiedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_pipeModified
	return p
}

func (*PipeModifiedContext) IsPipeModifiedContext() {}

func NewPipeModifiedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipeModifiedContext {
	var p = new(PipeModifiedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_pipeModified

	return p
}

func (s *PipeModifiedContext) GetParser() antlr.Parser { return s.parser }

func (s *PipeModifiedContext) Pipe() IPipeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPipeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPipeContext)
}

func (s *PipeModifiedContext) PipeModifier() IPipeModifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPipeModifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPipeModifierContext)
}

func (s *PipeModifiedContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(VyLangParserNL)
}

func (s *PipeModifiedContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(VyLangParserNL, i)
}

func (s *PipeModifiedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipeModifiedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipeModifiedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterPipeModified(s)
	}
}

func (s *PipeModifiedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitPipeModified(s)
	}
}

func (p *VyLangParser) PipeModified() (localctx IPipeModifiedContext) {
	this := p
	_ = this

	localctx = NewPipeModifiedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, VyLangParserRULE_pipeModified)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == VyLangParserT__14 {
		{
			p.SetState(107)
			p.PipeModifier()
		}

	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == VyLangParserNL {
		{
			p.SetState(110)
			p.Match(VyLangParserNL)
		}

		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(116)
		p.Pipe()
	}

	return localctx
}

// IPipeMapContext is an interface to support dynamic dispatch.
type IPipeMapContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPipeMapContext differentiates from other interfaces.
	IsPipeMapContext()
}

type PipeMapContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipeMapContext() *PipeMapContext {
	var p = new(PipeMapContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_pipeMap
	return p
}

func (*PipeMapContext) IsPipeMapContext() {}

func NewPipeMapContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipeMapContext {
	var p = new(PipeMapContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_pipeMap

	return p
}

func (s *PipeMapContext) GetParser() antlr.Parser { return s.parser }

func (s *PipeMapContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(VyLangParserNL)
}

func (s *PipeMapContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(VyLangParserNL, i)
}

func (s *PipeMapContext) AllPipeMapEntry() []IPipeMapEntryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPipeMapEntryContext); ok {
			len++
		}
	}

	tst := make([]IPipeMapEntryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPipeMapEntryContext); ok {
			tst[i] = t.(IPipeMapEntryContext)
			i++
		}
	}

	return tst
}

func (s *PipeMapContext) PipeMapEntry(i int) IPipeMapEntryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPipeMapEntryContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPipeMapEntryContext)
}

func (s *PipeMapContext) AllSep() []ISepContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISepContext); ok {
			len++
		}
	}

	tst := make([]ISepContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISepContext); ok {
			tst[i] = t.(ISepContext)
			i++
		}
	}

	return tst
}

func (s *PipeMapContext) Sep(i int) ISepContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISepContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISepContext)
}

func (s *PipeMapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipeMapContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipeMapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterPipeMap(s)
	}
}

func (s *PipeMapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitPipeMap(s)
	}
}

func (p *VyLangParser) PipeMap() (localctx IPipeMapContext) {
	this := p
	_ = this

	localctx = NewPipeMapContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, VyLangParserRULE_pipeMap)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.Match(VyLangParserT__12)
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(119)
			p.Match(VyLangParserNL)
		}

	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4202464) != 0 {
		{
			p.SetState(122)
			p.PipeMapEntry()
		}
		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(123)
					p.Sep()
				}
				{
					p.SetState(124)
					p.PipeMapEntry()
				}

			}
			p.SetState(130)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
		}

	}
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == VyLangParserNL {
		{
			p.SetState(133)
			p.Match(VyLangParserNL)
		}

	}
	{
		p.SetState(136)
		p.Match(VyLangParserT__13)
	}

	return localctx
}

// IPipeMapEntryContext is an interface to support dynamic dispatch.
type IPipeMapEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPipeMapEntryContext differentiates from other interfaces.
	IsPipeMapEntryContext()
}

type PipeMapEntryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipeMapEntryContext() *PipeMapEntryContext {
	var p = new(PipeMapEntryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_pipeMapEntry
	return p
}

func (*PipeMapEntryContext) IsPipeMapEntryContext() {}

func NewPipeMapEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipeMapEntryContext {
	var p = new(PipeMapEntryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_pipeMapEntry

	return p
}

func (s *PipeMapEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *PipeMapEntryContext) PipeNamedProperty() IPipeNamedPropertyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPipeNamedPropertyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPipeNamedPropertyContext)
}

func (s *PipeMapEntryContext) PipeProperty() IPipePropertyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPipePropertyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPipePropertyContext)
}

func (s *PipeMapEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipeMapEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipeMapEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterPipeMapEntry(s)
	}
}

func (s *PipeMapEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitPipeMapEntry(s)
	}
}

func (p *VyLangParser) PipeMapEntry() (localctx IPipeMapEntryContext) {
	this := p
	_ = this

	localctx = NewPipeMapEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, VyLangParserRULE_pipeMapEntry)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(138)
			p.PipeNamedProperty()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(139)
			p.PipeProperty()
		}

	}

	return localctx
}

// IPipeModifierContext is an interface to support dynamic dispatch.
type IPipeModifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPipeModifierContext differentiates from other interfaces.
	IsPipeModifierContext()
}

type PipeModifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipeModifierContext() *PipeModifierContext {
	var p = new(PipeModifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_pipeModifier
	return p
}

func (*PipeModifierContext) IsPipeModifierContext() {}

func NewPipeModifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipeModifierContext {
	var p = new(PipeModifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_pipeModifier

	return p
}

func (s *PipeModifierContext) GetParser() antlr.Parser { return s.parser }
func (s *PipeModifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipeModifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipeModifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterPipeModifier(s)
	}
}

func (s *PipeModifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitPipeModifier(s)
	}
}

func (p *VyLangParser) PipeModifier() (localctx IPipeModifierContext) {
	this := p
	_ = this

	localctx = NewPipeModifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, VyLangParserRULE_pipeModifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		p.Match(VyLangParserT__14)
	}
	{
		p.SetState(143)
		p.Match(VyLangParserT__15)
	}

	return localctx
}

// IPathModelContext is an interface to support dynamic dispatch.
type IPathModelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPathModelContext differentiates from other interfaces.
	IsPathModelContext()
}

type PathModelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathModelContext() *PathModelContext {
	var p = new(PathModelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_pathModel
	return p
}

func (*PathModelContext) IsPathModelContext() {}

func NewPathModelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathModelContext {
	var p = new(PathModelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_pathModel

	return p
}

func (s *PathModelContext) GetParser() antlr.Parser { return s.parser }

func (s *PathModelContext) IDENT() antlr.TerminalNode {
	return s.GetToken(VyLangParserIDENT, 0)
}

func (s *PathModelContext) IdentPath() IIdentPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentPathContext)
}

func (s *PathModelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathModelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PathModelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterPathModel(s)
	}
}

func (s *PathModelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitPathModel(s)
	}
}

func (p *VyLangParser) PathModel() (localctx IPathModelContext) {
	this := p
	_ = this

	localctx = NewPathModelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, VyLangParserRULE_pathModel)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(145)
			p.Match(VyLangParserIDENT)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(146)
			p.IdentPath()
		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == VyLangParserT__16 {
			{
				p.SetState(147)
				p.Match(VyLangParserT__16)
			}
			p.SetState(149)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == VyLangParserIDENT {
				{
					p.SetState(148)
					p.Match(VyLangParserIDENT)
				}

			}

		}

	}

	return localctx
}

// IPathRelationContext is an interface to support dynamic dispatch.
type IPathRelationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPathRelationContext differentiates from other interfaces.
	IsPathRelationContext()
}

type PathRelationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathRelationContext() *PathRelationContext {
	var p = new(PathRelationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_pathRelation
	return p
}

func (*PathRelationContext) IsPathRelationContext() {}

func NewPathRelationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathRelationContext {
	var p = new(PathRelationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_pathRelation

	return p
}

func (s *PathRelationContext) GetParser() antlr.Parser { return s.parser }

func (s *PathRelationContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(VyLangParserIDENT)
}

func (s *PathRelationContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(VyLangParserIDENT, i)
}

func (s *PathRelationContext) IdentPath() IIdentPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentPathContext)
}

func (s *PathRelationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathRelationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PathRelationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterPathRelation(s)
	}
}

func (s *PathRelationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitPathRelation(s)
	}
}

func (p *VyLangParser) PathRelation() (localctx IPathRelationContext) {
	this := p
	_ = this

	localctx = NewPathRelationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, VyLangParserRULE_pathRelation)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(155)
			p.Match(VyLangParserIDENT)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(156)
			p.IdentPath()
		}
		{
			p.SetState(157)
			p.Match(VyLangParserT__17)
		}
		{
			p.SetState(158)
			p.Match(VyLangParserIDENT)
		}
		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == VyLangParserT__16 {
			{
				p.SetState(159)
				p.Match(VyLangParserT__16)
			}
			p.SetState(161)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(160)
					p.Match(VyLangParserIDENT)
				}

			}

		}

	}

	return localctx
}

// IIdentPathContext is an interface to support dynamic dispatch.
type IIdentPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentPathContext differentiates from other interfaces.
	IsIdentPathContext()
}

type IdentPathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentPathContext() *IdentPathContext {
	var p = new(IdentPathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_identPath
	return p
}

func (*IdentPathContext) IsIdentPathContext() {}

func NewIdentPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentPathContext {
	var p = new(IdentPathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_identPath

	return p
}

func (s *IdentPathContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentPathContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(VyLangParserIDENT)
}

func (s *IdentPathContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(VyLangParserIDENT, i)
}

func (s *IdentPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterIdentPath(s)
	}
}

func (s *IdentPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitIdentPath(s)
	}
}

func (p *VyLangParser) IdentPath() (localctx IIdentPathContext) {
	this := p
	_ = this

	localctx = NewIdentPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, VyLangParserRULE_identPath)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Match(VyLangParserIDENT)
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == VyLangParserT__18 {
		{
			p.SetState(168)
			p.Match(VyLangParserT__18)
		}
		{
			p.SetState(169)
			p.Match(VyLangParserIDENT)
		}

		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISepContext is an interface to support dynamic dispatch.
type ISepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSepContext differentiates from other interfaces.
	IsSepContext()
}

type SepContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySepContext() *SepContext {
	var p = new(SepContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_sep
	return p
}

func (*SepContext) IsSepContext() {}

func NewSepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SepContext {
	var p = new(SepContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_sep

	return p
}

func (s *SepContext) GetParser() antlr.Parser { return s.parser }

func (s *SepContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(VyLangParserNL)
}

func (s *SepContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(VyLangParserNL, i)
}

func (s *SepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterSep(s)
	}
}

func (s *SepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitSep(s)
	}
}

func (p *VyLangParser) Sep() (localctx ISepContext) {
	this := p
	_ = this

	localctx = NewSepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, VyLangParserRULE_sep)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(175)
			p.Match(VyLangParserNL)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == VyLangParserNL {
			{
				p.SetState(176)
				p.Match(VyLangParserNL)
			}

		}
		{
			p.SetState(179)
			p.Match(VyLangParserT__19)
		}
		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == VyLangParserNL {
			{
				p.SetState(180)
				p.Match(VyLangParserNL)
			}

		}

	}

	return localctx
}
