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
		"", "'.'", "'@'", "'action'", "'on'", "':'", "'{'", "','", "'}'", "'['",
		"']'", "'&'", "'$'", "'bind'", "'='", "'while'", "'if'", "'else'", "'del'",
		"'<-'", "'->'", "'<=>'", "'('", "')'", "'type'", "'in'", "'string'",
		"'integer'", "'float'", "'boolean'", "'raw'", "'function'", "'?'", "'-'",
		"'!'", "'*'", "'/'", "'+'", "'>'", "'<'", "'>='", "'<='", "'=='", "'!='",
		"'&&'", "'||'", "'++'", "'--'", "'constant'", "'true'", "'false'", "'null'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "STRING", "IDENTIFIER", "NEWLINE", "WHITESPACE", "INT",
	}
	staticData.ruleNames = []string{
		"prog", "definitions", "definition", "variable", "namedAction", "action",
		"actionParallel", "actionSequence", "actionAsync", "actionReference",
		"actionBind", "actionWhile", "actionCond", "actionIf", "actionIfElse",
		"actionLeaf", "actionAssign", "actionClear", "binding", "namedType",
		"type", "rawType", "typeMap", "typeList", "typeReference", "typeLeaf",
		"typeMapEntry", "namedExpr", "expr", "exprBrackets", "exprAlternative",
		"exprOperator1", "exprOperator2", "exprReference", "exprMap", "exprList",
		"exprMapEntry", "namedLiteral", "literal", "literalTerminal", "literalString",
		"literalInt", "literalBoolean", "literalFloat", "literalNull", "literalReference",
		"literalMap", "literalList", "literalMapEntry",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 56, 619, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 1, 0, 1, 0, 1, 0, 1, 1, 5, 1, 103, 8, 1, 10, 1, 12,
		1, 106, 9, 1, 1, 1, 1, 1, 5, 1, 110, 8, 1, 10, 1, 12, 1, 113, 9, 1, 5,
		1, 115, 8, 1, 10, 1, 12, 1, 118, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 124,
		8, 2, 1, 3, 1, 3, 1, 3, 3, 3, 129, 8, 3, 1, 3, 1, 3, 1, 3, 5, 3, 134, 8,
		3, 10, 3, 12, 3, 137, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 145,
		8, 4, 10, 4, 12, 4, 148, 9, 4, 1, 4, 1, 4, 5, 4, 152, 8, 4, 10, 4, 12,
		4, 155, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 165,
		8, 5, 1, 6, 1, 6, 5, 6, 169, 8, 6, 10, 6, 12, 6, 172, 9, 6, 1, 6, 1, 6,
		4, 6, 176, 8, 6, 11, 6, 12, 6, 177, 5, 6, 180, 8, 6, 10, 6, 12, 6, 183,
		9, 6, 1, 6, 5, 6, 186, 8, 6, 10, 6, 12, 6, 189, 9, 6, 1, 6, 1, 6, 1, 7,
		1, 7, 5, 7, 195, 8, 7, 10, 7, 12, 7, 198, 9, 7, 1, 7, 1, 7, 4, 7, 202,
		8, 7, 11, 7, 12, 7, 203, 5, 7, 206, 8, 7, 10, 7, 12, 7, 209, 9, 7, 1, 7,
		5, 7, 212, 8, 7, 10, 7, 12, 7, 215, 9, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 3, 10, 227, 8, 10, 1, 10, 1, 10, 3, 10,
		231, 8, 10, 1, 10, 1, 10, 3, 10, 235, 8, 10, 1, 10, 1, 10, 5, 10, 239,
		8, 10, 10, 10, 12, 10, 242, 9, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1,
		11, 5, 11, 250, 8, 11, 10, 11, 12, 11, 253, 9, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 3, 12, 259, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 265, 8, 13,
		10, 13, 12, 13, 268, 9, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 5,
		14, 276, 8, 14, 10, 14, 12, 14, 279, 9, 14, 1, 14, 1, 14, 5, 14, 283, 8,
		14, 10, 14, 12, 14, 286, 9, 14, 1, 14, 1, 14, 1, 14, 5, 14, 291, 8, 14,
		10, 14, 12, 14, 294, 9, 14, 1, 14, 1, 14, 1, 15, 1, 15, 3, 15, 300, 8,
		15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 320, 8,
		18, 10, 18, 12, 18, 323, 9, 18, 1, 18, 1, 18, 5, 18, 327, 8, 18, 10, 18,
		12, 18, 330, 9, 18, 1, 18, 1, 18, 3, 18, 334, 8, 18, 1, 18, 1, 18, 1, 18,
		5, 18, 339, 8, 18, 10, 18, 12, 18, 342, 9, 18, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 20, 1, 20, 1, 20, 3, 20, 352, 8, 20, 1, 20, 1, 20, 3, 20,
		356, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 362, 8, 21, 1, 22, 1, 22,
		5, 22, 366, 8, 22, 10, 22, 12, 22, 369, 9, 22, 1, 22, 1, 22, 4, 22, 373,
		8, 22, 11, 22, 12, 22, 374, 1, 22, 5, 22, 378, 8, 22, 10, 22, 12, 22, 381,
		9, 22, 1, 22, 5, 22, 384, 8, 22, 10, 22, 12, 22, 387, 9, 22, 1, 22, 1,
		22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 419, 8, 28,
		1, 28, 1, 28, 4, 28, 423, 8, 28, 11, 28, 12, 28, 424, 1, 28, 1, 28, 5,
		28, 429, 8, 28, 10, 28, 12, 28, 432, 9, 28, 1, 29, 1, 29, 1, 29, 1, 29,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 447,
		8, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 463, 8, 32, 1, 33, 1, 33, 1, 33,
		3, 33, 468, 8, 33, 1, 34, 1, 34, 5, 34, 472, 8, 34, 10, 34, 12, 34, 475,
		9, 34, 1, 34, 1, 34, 4, 34, 479, 8, 34, 11, 34, 12, 34, 480, 1, 34, 5,
		34, 484, 8, 34, 10, 34, 12, 34, 487, 9, 34, 1, 34, 5, 34, 490, 8, 34, 10,
		34, 12, 34, 493, 9, 34, 1, 34, 1, 34, 1, 35, 1, 35, 5, 35, 499, 8, 35,
		10, 35, 12, 35, 502, 9, 35, 1, 35, 1, 35, 4, 35, 506, 8, 35, 11, 35, 12,
		35, 507, 1, 35, 5, 35, 511, 8, 35, 10, 35, 12, 35, 514, 9, 35, 1, 35, 5,
		35, 517, 8, 35, 10, 35, 12, 35, 520, 9, 35, 1, 35, 1, 35, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 3,
		38, 536, 8, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 3, 39, 544, 8,
		39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 43,
		1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 5, 46, 563, 8, 46, 10,
		46, 12, 46, 566, 9, 46, 1, 46, 1, 46, 4, 46, 570, 8, 46, 11, 46, 12, 46,
		571, 1, 46, 5, 46, 575, 8, 46, 10, 46, 12, 46, 578, 9, 46, 1, 46, 5, 46,
		581, 8, 46, 10, 46, 12, 46, 584, 9, 46, 1, 46, 1, 46, 1, 47, 1, 47, 5,
		47, 590, 8, 47, 10, 47, 12, 47, 593, 9, 47, 1, 47, 1, 47, 4, 47, 597, 8,
		47, 11, 47, 12, 47, 598, 1, 47, 5, 47, 602, 8, 47, 10, 47, 12, 47, 605,
		9, 47, 1, 47, 5, 47, 608, 8, 47, 10, 47, 12, 47, 611, 9, 47, 1, 47, 1,
		47, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 0, 3, 6, 36, 56, 49, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78,
		80, 82, 84, 86, 88, 90, 92, 94, 96, 0, 9, 2, 0, 7, 7, 54, 54, 1, 0, 19,
		21, 1, 0, 26, 30, 1, 0, 35, 36, 2, 0, 33, 33, 37, 37, 1, 0, 38, 41, 1,
		0, 42, 43, 1, 0, 46, 47, 1, 0, 49, 50, 659, 0, 98, 1, 0, 0, 0, 2, 116,
		1, 0, 0, 0, 4, 123, 1, 0, 0, 0, 6, 128, 1, 0, 0, 0, 8, 138, 1, 0, 0, 0,
		10, 164, 1, 0, 0, 0, 12, 166, 1, 0, 0, 0, 14, 192, 1, 0, 0, 0, 16, 218,
		1, 0, 0, 0, 18, 221, 1, 0, 0, 0, 20, 224, 1, 0, 0, 0, 22, 245, 1, 0, 0,
		0, 24, 258, 1, 0, 0, 0, 26, 260, 1, 0, 0, 0, 28, 271, 1, 0, 0, 0, 30, 299,
		1, 0, 0, 0, 32, 301, 1, 0, 0, 0, 34, 305, 1, 0, 0, 0, 36, 333, 1, 0, 0,
		0, 38, 343, 1, 0, 0, 0, 40, 348, 1, 0, 0, 0, 42, 361, 1, 0, 0, 0, 44, 363,
		1, 0, 0, 0, 46, 390, 1, 0, 0, 0, 48, 394, 1, 0, 0, 0, 50, 397, 1, 0, 0,
		0, 52, 399, 1, 0, 0, 0, 54, 403, 1, 0, 0, 0, 56, 418, 1, 0, 0, 0, 58, 433,
		1, 0, 0, 0, 60, 437, 1, 0, 0, 0, 62, 446, 1, 0, 0, 0, 64, 462, 1, 0, 0,
		0, 66, 464, 1, 0, 0, 0, 68, 469, 1, 0, 0, 0, 70, 496, 1, 0, 0, 0, 72, 523,
		1, 0, 0, 0, 74, 527, 1, 0, 0, 0, 76, 535, 1, 0, 0, 0, 78, 543, 1, 0, 0,
		0, 80, 545, 1, 0, 0, 0, 82, 547, 1, 0, 0, 0, 84, 549, 1, 0, 0, 0, 86, 551,
		1, 0, 0, 0, 88, 555, 1, 0, 0, 0, 90, 557, 1, 0, 0, 0, 92, 560, 1, 0, 0,
		0, 94, 587, 1, 0, 0, 0, 96, 614, 1, 0, 0, 0, 98, 99, 3, 2, 1, 0, 99, 100,
		5, 0, 0, 1, 100, 1, 1, 0, 0, 0, 101, 103, 5, 54, 0, 0, 102, 101, 1, 0,
		0, 0, 103, 106, 1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0,
		105, 107, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 107, 111, 3, 4, 2, 0, 108,
		110, 5, 54, 0, 0, 109, 108, 1, 0, 0, 0, 110, 113, 1, 0, 0, 0, 111, 109,
		1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 115, 1, 0, 0, 0, 113, 111, 1, 0,
		0, 0, 114, 104, 1, 0, 0, 0, 115, 118, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0,
		116, 117, 1, 0, 0, 0, 117, 3, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0, 119, 124,
		3, 8, 4, 0, 120, 124, 3, 38, 19, 0, 121, 124, 3, 74, 37, 0, 122, 124, 3,
		54, 27, 0, 123, 119, 1, 0, 0, 0, 123, 120, 1, 0, 0, 0, 123, 121, 1, 0,
		0, 0, 123, 122, 1, 0, 0, 0, 124, 5, 1, 0, 0, 0, 125, 126, 6, 3, -1, 0,
		126, 129, 5, 53, 0, 0, 127, 129, 5, 2, 0, 0, 128, 125, 1, 0, 0, 0, 128,
		127, 1, 0, 0, 0, 129, 135, 1, 0, 0, 0, 130, 131, 10, 2, 0, 0, 131, 132,
		5, 1, 0, 0, 132, 134, 5, 53, 0, 0, 133, 130, 1, 0, 0, 0, 134, 137, 1, 0,
		0, 0, 135, 133, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 7, 1, 0, 0, 0, 137,
		135, 1, 0, 0, 0, 138, 139, 5, 3, 0, 0, 139, 140, 5, 53, 0, 0, 140, 141,
		5, 4, 0, 0, 141, 142, 3, 40, 20, 0, 142, 146, 5, 5, 0, 0, 143, 145, 5,
		54, 0, 0, 144, 143, 1, 0, 0, 0, 145, 148, 1, 0, 0, 0, 146, 144, 1, 0, 0,
		0, 146, 147, 1, 0, 0, 0, 147, 149, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 149,
		153, 3, 10, 5, 0, 150, 152, 5, 54, 0, 0, 151, 150, 1, 0, 0, 0, 152, 155,
		1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 9, 1, 0, 0,
		0, 155, 153, 1, 0, 0, 0, 156, 165, 3, 12, 6, 0, 157, 165, 3, 14, 7, 0,
		158, 165, 3, 16, 8, 0, 159, 165, 3, 18, 9, 0, 160, 165, 3, 20, 10, 0, 161,
		165, 3, 22, 11, 0, 162, 165, 3, 24, 12, 0, 163, 165, 3, 30, 15, 0, 164,
		156, 1, 0, 0, 0, 164, 157, 1, 0, 0, 0, 164, 158, 1, 0, 0, 0, 164, 159,
		1, 0, 0, 0, 164, 160, 1, 0, 0, 0, 164, 161, 1, 0, 0, 0, 164, 162, 1, 0,
		0, 0, 164, 163, 1, 0, 0, 0, 165, 11, 1, 0, 0, 0, 166, 181, 5, 6, 0, 0,
		167, 169, 5, 54, 0, 0, 168, 167, 1, 0, 0, 0, 169, 172, 1, 0, 0, 0, 170,
		168, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 173, 1, 0, 0, 0, 172, 170,
		1, 0, 0, 0, 173, 175, 3, 10, 5, 0, 174, 176, 7, 0, 0, 0, 175, 174, 1, 0,
		0, 0, 176, 177, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0,
		178, 180, 1, 0, 0, 0, 179, 170, 1, 0, 0, 0, 180, 183, 1, 0, 0, 0, 181,
		179, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182, 187, 1, 0, 0, 0, 183, 181,
		1, 0, 0, 0, 184, 186, 5, 54, 0, 0, 185, 184, 1, 0, 0, 0, 186, 189, 1, 0,
		0, 0, 187, 185, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0, 188, 190, 1, 0, 0, 0,
		189, 187, 1, 0, 0, 0, 190, 191, 5, 8, 0, 0, 191, 13, 1, 0, 0, 0, 192, 207,
		5, 9, 0, 0, 193, 195, 5, 54, 0, 0, 194, 193, 1, 0, 0, 0, 195, 198, 1, 0,
		0, 0, 196, 194, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 199, 1, 0, 0, 0,
		198, 196, 1, 0, 0, 0, 199, 201, 3, 10, 5, 0, 200, 202, 7, 0, 0, 0, 201,
		200, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0, 203, 204,
		1, 0, 0, 0, 204, 206, 1, 0, 0, 0, 205, 196, 1, 0, 0, 0, 206, 209, 1, 0,
		0, 0, 207, 205, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 213, 1, 0, 0, 0,
		209, 207, 1, 0, 0, 0, 210, 212, 5, 54, 0, 0, 211, 210, 1, 0, 0, 0, 212,
		215, 1, 0, 0, 0, 213, 211, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 216,
		1, 0, 0, 0, 215, 213, 1, 0, 0, 0, 216, 217, 5, 10, 0, 0, 217, 15, 1, 0,
		0, 0, 218, 219, 5, 11, 0, 0, 219, 220, 3, 10, 5, 0, 220, 17, 1, 0, 0, 0,
		221, 222, 5, 12, 0, 0, 222, 223, 5, 53, 0, 0, 223, 19, 1, 0, 0, 0, 224,
		226, 5, 13, 0, 0, 225, 227, 3, 36, 18, 0, 226, 225, 1, 0, 0, 0, 226, 227,
		1, 0, 0, 0, 227, 230, 1, 0, 0, 0, 228, 229, 5, 4, 0, 0, 229, 231, 3, 40,
		20, 0, 230, 228, 1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 234, 1, 0, 0, 0,
		232, 233, 5, 14, 0, 0, 233, 235, 3, 56, 28, 0, 234, 232, 1, 0, 0, 0, 234,
		235, 1, 0, 0, 0, 235, 236, 1, 0, 0, 0, 236, 240, 5, 5, 0, 0, 237, 239,
		5, 54, 0, 0, 238, 237, 1, 0, 0, 0, 239, 242, 1, 0, 0, 0, 240, 238, 1, 0,
		0, 0, 240, 241, 1, 0, 0, 0, 241, 243, 1, 0, 0, 0, 242, 240, 1, 0, 0, 0,
		243, 244, 3, 10, 5, 0, 244, 21, 1, 0, 0, 0, 245, 246, 5, 15, 0, 0, 246,
		247, 3, 56, 28, 0, 247, 251, 5, 5, 0, 0, 248, 250, 5, 54, 0, 0, 249, 248,
		1, 0, 0, 0, 250, 253, 1, 0, 0, 0, 251, 249, 1, 0, 0, 0, 251, 252, 1, 0,
		0, 0, 252, 254, 1, 0, 0, 0, 253, 251, 1, 0, 0, 0, 254, 255, 3, 10, 5, 0,
		255, 23, 1, 0, 0, 0, 256, 259, 3, 26, 13, 0, 257, 259, 3, 28, 14, 0, 258,
		256, 1, 0, 0, 0, 258, 257, 1, 0, 0, 0, 259, 25, 1, 0, 0, 0, 260, 261, 5,
		16, 0, 0, 261, 262, 3, 56, 28, 0, 262, 266, 5, 5, 0, 0, 263, 265, 5, 54,
		0, 0, 264, 263, 1, 0, 0, 0, 265, 268, 1, 0, 0, 0, 266, 264, 1, 0, 0, 0,
		266, 267, 1, 0, 0, 0, 267, 269, 1, 0, 0, 0, 268, 266, 1, 0, 0, 0, 269,
		270, 3, 10, 5, 0, 270, 27, 1, 0, 0, 0, 271, 272, 5, 16, 0, 0, 272, 273,
		3, 56, 28, 0, 273, 277, 5, 5, 0, 0, 274, 276, 5, 54, 0, 0, 275, 274, 1,
		0, 0, 0, 276, 279, 1, 0, 0, 0, 277, 275, 1, 0, 0, 0, 277, 278, 1, 0, 0,
		0, 278, 280, 1, 0, 0, 0, 279, 277, 1, 0, 0, 0, 280, 284, 3, 10, 5, 0, 281,
		283, 5, 54, 0, 0, 282, 281, 1, 0, 0, 0, 283, 286, 1, 0, 0, 0, 284, 282,
		1, 0, 0, 0, 284, 285, 1, 0, 0, 0, 285, 287, 1, 0, 0, 0, 286, 284, 1, 0,
		0, 0, 287, 288, 5, 17, 0, 0, 288, 292, 5, 5, 0, 0, 289, 291, 5, 54, 0,
		0, 290, 289, 1, 0, 0, 0, 291, 294, 1, 0, 0, 0, 292, 290, 1, 0, 0, 0, 292,
		293, 1, 0, 0, 0, 293, 295, 1, 0, 0, 0, 294, 292, 1, 0, 0, 0, 295, 296,
		3, 10, 5, 0, 296, 29, 1, 0, 0, 0, 297, 300, 3, 32, 16, 0, 298, 300, 3,
		34, 17, 0, 299, 297, 1, 0, 0, 0, 299, 298, 1, 0, 0, 0, 300, 31, 1, 0, 0,
		0, 301, 302, 3, 6, 3, 0, 302, 303, 5, 14, 0, 0, 303, 304, 3, 56, 28, 0,
		304, 33, 1, 0, 0, 0, 305, 306, 5, 18, 0, 0, 306, 307, 3, 6, 3, 0, 307,
		35, 1, 0, 0, 0, 308, 309, 6, 18, -1, 0, 309, 310, 3, 6, 3, 0, 310, 311,
		7, 1, 0, 0, 311, 312, 3, 6, 3, 0, 312, 334, 1, 0, 0, 0, 313, 314, 3, 76,
		38, 0, 314, 315, 5, 20, 0, 0, 315, 316, 3, 6, 3, 0, 316, 334, 1, 0, 0,
		0, 317, 321, 5, 22, 0, 0, 318, 320, 5, 54, 0, 0, 319, 318, 1, 0, 0, 0,
		320, 323, 1, 0, 0, 0, 321, 319, 1, 0, 0, 0, 321, 322, 1, 0, 0, 0, 322,
		324, 1, 0, 0, 0, 323, 321, 1, 0, 0, 0, 324, 328, 3, 36, 18, 0, 325, 327,
		5, 54, 0, 0, 326, 325, 1, 0, 0, 0, 327, 330, 1, 0, 0, 0, 328, 326, 1, 0,
		0, 0, 328, 329, 1, 0, 0, 0, 329, 331, 1, 0, 0, 0, 330, 328, 1, 0, 0, 0,
		331, 332, 5, 23, 0, 0, 332, 334, 1, 0, 0, 0, 333, 308, 1, 0, 0, 0, 333,
		313, 1, 0, 0, 0, 333, 317, 1, 0, 0, 0, 334, 340, 1, 0, 0, 0, 335, 336,
		10, 2, 0, 0, 336, 337, 5, 7, 0, 0, 337, 339, 3, 36, 18, 3, 338, 335, 1,
		0, 0, 0, 339, 342, 1, 0, 0, 0, 340, 338, 1, 0, 0, 0, 340, 341, 1, 0, 0,
		0, 341, 37, 1, 0, 0, 0, 342, 340, 1, 0, 0, 0, 343, 344, 5, 24, 0, 0, 344,
		345, 5, 53, 0, 0, 345, 346, 5, 5, 0, 0, 346, 347, 3, 40, 20, 0, 347, 39,
		1, 0, 0, 0, 348, 351, 3, 42, 21, 0, 349, 350, 5, 14, 0, 0, 350, 352, 3,
		76, 38, 0, 351, 349, 1, 0, 0, 0, 351, 352, 1, 0, 0, 0, 352, 355, 1, 0,
		0, 0, 353, 354, 5, 25, 0, 0, 354, 356, 3, 94, 47, 0, 355, 353, 1, 0, 0,
		0, 355, 356, 1, 0, 0, 0, 356, 41, 1, 0, 0, 0, 357, 362, 3, 44, 22, 0, 358,
		362, 3, 46, 23, 0, 359, 362, 3, 48, 24, 0, 360, 362, 3, 50, 25, 0, 361,
		357, 1, 0, 0, 0, 361, 358, 1, 0, 0, 0, 361, 359, 1, 0, 0, 0, 361, 360,
		1, 0, 0, 0, 362, 43, 1, 0, 0, 0, 363, 367, 5, 6, 0, 0, 364, 366, 5, 54,
		0, 0, 365, 364, 1, 0, 0, 0, 366, 369, 1, 0, 0, 0, 367, 365, 1, 0, 0, 0,
		367, 368, 1, 0, 0, 0, 368, 370, 1, 0, 0, 0, 369, 367, 1, 0, 0, 0, 370,
		379, 3, 52, 26, 0, 371, 373, 7, 0, 0, 0, 372, 371, 1, 0, 0, 0, 373, 374,
		1, 0, 0, 0, 374, 372, 1, 0, 0, 0, 374, 375, 1, 0, 0, 0, 375, 376, 1, 0,
		0, 0, 376, 378, 3, 52, 26, 0, 377, 372, 1, 0, 0, 0, 378, 381, 1, 0, 0,
		0, 379, 377, 1, 0, 0, 0, 379, 380, 1, 0, 0, 0, 380, 385, 1, 0, 0, 0, 381,
		379, 1, 0, 0, 0, 382, 384, 5, 54, 0, 0, 383, 382, 1, 0, 0, 0, 384, 387,
		1, 0, 0, 0, 385, 383, 1, 0, 0, 0, 385, 386, 1, 0, 0, 0, 386, 388, 1, 0,
		0, 0, 387, 385, 1, 0, 0, 0, 388, 389, 5, 8, 0, 0, 389, 45, 1, 0, 0, 0,
		390, 391, 5, 9, 0, 0, 391, 392, 5, 10, 0, 0, 392, 393, 3, 40, 20, 0, 393,
		47, 1, 0, 0, 0, 394, 395, 5, 12, 0, 0, 395, 396, 5, 53, 0, 0, 396, 49,
		1, 0, 0, 0, 397, 398, 7, 2, 0, 0, 398, 51, 1, 0, 0, 0, 399, 400, 5, 53,
		0, 0, 400, 401, 5, 5, 0, 0, 401, 402, 3, 40, 20, 0, 402, 53, 1, 0, 0, 0,
		403, 404, 5, 31, 0, 0, 404, 405, 5, 53, 0, 0, 405, 406, 5, 4, 0, 0, 406,
		407, 3, 40, 20, 0, 407, 408, 5, 5, 0, 0, 408, 409, 3, 56, 28, 0, 409, 55,
		1, 0, 0, 0, 410, 411, 6, 28, -1, 0, 411, 419, 3, 6, 3, 0, 412, 419, 3,
		76, 38, 0, 413, 419, 3, 68, 34, 0, 414, 419, 3, 70, 35, 0, 415, 419, 3,
		58, 29, 0, 416, 419, 3, 66, 33, 0, 417, 419, 3, 62, 31, 0, 418, 410, 1,
		0, 0, 0, 418, 412, 1, 0, 0, 0, 418, 413, 1, 0, 0, 0, 418, 414, 1, 0, 0,
		0, 418, 415, 1, 0, 0, 0, 418, 416, 1, 0, 0, 0, 418, 417, 1, 0, 0, 0, 419,
		430, 1, 0, 0, 0, 420, 422, 10, 2, 0, 0, 421, 423, 3, 64, 32, 0, 422, 421,
		1, 0, 0, 0, 423, 424, 1, 0, 0, 0, 424, 422, 1, 0, 0, 0, 424, 425, 1, 0,
		0, 0, 425, 429, 1, 0, 0, 0, 426, 427, 10, 1, 0, 0, 427, 429, 3, 60, 30,
		0, 428, 420, 1, 0, 0, 0, 428, 426, 1, 0, 0, 0, 429, 432, 1, 0, 0, 0, 430,
		428, 1, 0, 0, 0, 430, 431, 1, 0, 0, 0, 431, 57, 1, 0, 0, 0, 432, 430, 1,
		0, 0, 0, 433, 434, 5, 22, 0, 0, 434, 435, 3, 56, 28, 0, 435, 436, 5, 23,
		0, 0, 436, 59, 1, 0, 0, 0, 437, 438, 5, 32, 0, 0, 438, 439, 3, 56, 28,
		0, 439, 440, 5, 5, 0, 0, 440, 441, 3, 56, 28, 0, 441, 61, 1, 0, 0, 0, 442,
		443, 5, 33, 0, 0, 443, 447, 3, 56, 28, 0, 444, 445, 5, 34, 0, 0, 445, 447,
		3, 56, 28, 0, 446, 442, 1, 0, 0, 0, 446, 444, 1, 0, 0, 0, 447, 63, 1, 0,
		0, 0, 448, 449, 7, 3, 0, 0, 449, 463, 3, 56, 28, 0, 450, 451, 7, 4, 0,
		0, 451, 463, 3, 56, 28, 0, 452, 453, 7, 5, 0, 0, 453, 463, 3, 56, 28, 0,
		454, 455, 7, 6, 0, 0, 455, 463, 3, 56, 28, 0, 456, 457, 5, 44, 0, 0, 457,
		463, 3, 56, 28, 0, 458, 459, 5, 45, 0, 0, 459, 463, 3, 56, 28, 0, 460,
		461, 7, 7, 0, 0, 461, 463, 3, 56, 28, 0, 462, 448, 1, 0, 0, 0, 462, 450,
		1, 0, 0, 0, 462, 452, 1, 0, 0, 0, 462, 454, 1, 0, 0, 0, 462, 456, 1, 0,
		0, 0, 462, 458, 1, 0, 0, 0, 462, 460, 1, 0, 0, 0, 463, 65, 1, 0, 0, 0,
		464, 465, 5, 12, 0, 0, 465, 467, 5, 53, 0, 0, 466, 468, 3, 56, 28, 0, 467,
		466, 1, 0, 0, 0, 467, 468, 1, 0, 0, 0, 468, 67, 1, 0, 0, 0, 469, 473, 5,
		6, 0, 0, 470, 472, 5, 54, 0, 0, 471, 470, 1, 0, 0, 0, 472, 475, 1, 0, 0,
		0, 473, 471, 1, 0, 0, 0, 473, 474, 1, 0, 0, 0, 474, 476, 1, 0, 0, 0, 475,
		473, 1, 0, 0, 0, 476, 485, 3, 72, 36, 0, 477, 479, 7, 0, 0, 0, 478, 477,
		1, 0, 0, 0, 479, 480, 1, 0, 0, 0, 480, 478, 1, 0, 0, 0, 480, 481, 1, 0,
		0, 0, 481, 482, 1, 0, 0, 0, 482, 484, 3, 72, 36, 0, 483, 478, 1, 0, 0,
		0, 484, 487, 1, 0, 0, 0, 485, 483, 1, 0, 0, 0, 485, 486, 1, 0, 0, 0, 486,
		491, 1, 0, 0, 0, 487, 485, 1, 0, 0, 0, 488, 490, 5, 54, 0, 0, 489, 488,
		1, 0, 0, 0, 490, 493, 1, 0, 0, 0, 491, 489, 1, 0, 0, 0, 491, 492, 1, 0,
		0, 0, 492, 494, 1, 0, 0, 0, 493, 491, 1, 0, 0, 0, 494, 495, 5, 8, 0, 0,
		495, 69, 1, 0, 0, 0, 496, 500, 5, 9, 0, 0, 497, 499, 5, 54, 0, 0, 498,
		497, 1, 0, 0, 0, 499, 502, 1, 0, 0, 0, 500, 498, 1, 0, 0, 0, 500, 501,
		1, 0, 0, 0, 501, 503, 1, 0, 0, 0, 502, 500, 1, 0, 0, 0, 503, 512, 3, 56,
		28, 0, 504, 506, 7, 0, 0, 0, 505, 504, 1, 0, 0, 0, 506, 507, 1, 0, 0, 0,
		507, 505, 1, 0, 0, 0, 507, 508, 1, 0, 0, 0, 508, 509, 1, 0, 0, 0, 509,
		511, 3, 56, 28, 0, 510, 505, 1, 0, 0, 0, 511, 514, 1, 0, 0, 0, 512, 510,
		1, 0, 0, 0, 512, 513, 1, 0, 0, 0, 513, 518, 1, 0, 0, 0, 514, 512, 1, 0,
		0, 0, 515, 517, 5, 54, 0, 0, 516, 515, 1, 0, 0, 0, 517, 520, 1, 0, 0, 0,
		518, 516, 1, 0, 0, 0, 518, 519, 1, 0, 0, 0, 519, 521, 1, 0, 0, 0, 520,
		518, 1, 0, 0, 0, 521, 522, 5, 10, 0, 0, 522, 71, 1, 0, 0, 0, 523, 524,
		5, 53, 0, 0, 524, 525, 5, 5, 0, 0, 525, 526, 3, 56, 28, 0, 526, 73, 1,
		0, 0, 0, 527, 528, 5, 48, 0, 0, 528, 529, 5, 53, 0, 0, 529, 530, 5, 5,
		0, 0, 530, 531, 3, 76, 38, 0, 531, 75, 1, 0, 0, 0, 532, 536, 3, 78, 39,
		0, 533, 536, 3, 92, 46, 0, 534, 536, 3, 94, 47, 0, 535, 532, 1, 0, 0, 0,
		535, 533, 1, 0, 0, 0, 535, 534, 1, 0, 0, 0, 536, 77, 1, 0, 0, 0, 537, 544,
		3, 80, 40, 0, 538, 544, 3, 82, 41, 0, 539, 544, 3, 84, 42, 0, 540, 544,
		3, 86, 43, 0, 541, 544, 3, 88, 44, 0, 542, 544, 3, 90, 45, 0, 543, 537,
		1, 0, 0, 0, 543, 538, 1, 0, 0, 0, 543, 539, 1, 0, 0, 0, 543, 540, 1, 0,
		0, 0, 543, 541, 1, 0, 0, 0, 543, 542, 1, 0, 0, 0, 544, 79, 1, 0, 0, 0,
		545, 546, 5, 52, 0, 0, 546, 81, 1, 0, 0, 0, 547, 548, 5, 56, 0, 0, 548,
		83, 1, 0, 0, 0, 549, 550, 7, 8, 0, 0, 550, 85, 1, 0, 0, 0, 551, 552, 5,
		56, 0, 0, 552, 553, 5, 1, 0, 0, 553, 554, 5, 56, 0, 0, 554, 87, 1, 0, 0,
		0, 555, 556, 5, 51, 0, 0, 556, 89, 1, 0, 0, 0, 557, 558, 5, 12, 0, 0, 558,
		559, 5, 53, 0, 0, 559, 91, 1, 0, 0, 0, 560, 564, 5, 6, 0, 0, 561, 563,
		5, 54, 0, 0, 562, 561, 1, 0, 0, 0, 563, 566, 1, 0, 0, 0, 564, 562, 1, 0,
		0, 0, 564, 565, 1, 0, 0, 0, 565, 567, 1, 0, 0, 0, 566, 564, 1, 0, 0, 0,
		567, 576, 3, 96, 48, 0, 568, 570, 7, 0, 0, 0, 569, 568, 1, 0, 0, 0, 570,
		571, 1, 0, 0, 0, 571, 569, 1, 0, 0, 0, 571, 572, 1, 0, 0, 0, 572, 573,
		1, 0, 0, 0, 573, 575, 3, 96, 48, 0, 574, 569, 1, 0, 0, 0, 575, 578, 1,
		0, 0, 0, 576, 574, 1, 0, 0, 0, 576, 577, 1, 0, 0, 0, 577, 582, 1, 0, 0,
		0, 578, 576, 1, 0, 0, 0, 579, 581, 5, 54, 0, 0, 580, 579, 1, 0, 0, 0, 581,
		584, 1, 0, 0, 0, 582, 580, 1, 0, 0, 0, 582, 583, 1, 0, 0, 0, 583, 585,
		1, 0, 0, 0, 584, 582, 1, 0, 0, 0, 585, 586, 5, 8, 0, 0, 586, 93, 1, 0,
		0, 0, 587, 591, 5, 9, 0, 0, 588, 590, 5, 54, 0, 0, 589, 588, 1, 0, 0, 0,
		590, 593, 1, 0, 0, 0, 591, 589, 1, 0, 0, 0, 591, 592, 1, 0, 0, 0, 592,
		594, 1, 0, 0, 0, 593, 591, 1, 0, 0, 0, 594, 603, 3, 76, 38, 0, 595, 597,
		7, 0, 0, 0, 596, 595, 1, 0, 0, 0, 597, 598, 1, 0, 0, 0, 598, 596, 1, 0,
		0, 0, 598, 599, 1, 0, 0, 0, 599, 600, 1, 0, 0, 0, 600, 602, 3, 76, 38,
		0, 601, 596, 1, 0, 0, 0, 602, 605, 1, 0, 0, 0, 603, 601, 1, 0, 0, 0, 603,
		604, 1, 0, 0, 0, 604, 609, 1, 0, 0, 0, 605, 603, 1, 0, 0, 0, 606, 608,
		5, 54, 0, 0, 607, 606, 1, 0, 0, 0, 608, 611, 1, 0, 0, 0, 609, 607, 1, 0,
		0, 0, 609, 610, 1, 0, 0, 0, 610, 612, 1, 0, 0, 0, 611, 609, 1, 0, 0, 0,
		612, 613, 5, 10, 0, 0, 613, 95, 1, 0, 0, 0, 614, 615, 5, 53, 0, 0, 615,
		616, 5, 5, 0, 0, 616, 617, 3, 76, 38, 0, 617, 97, 1, 0, 0, 0, 64, 104,
		111, 116, 123, 128, 135, 146, 153, 164, 170, 177, 181, 187, 196, 203, 207,
		213, 226, 230, 234, 240, 251, 258, 266, 277, 284, 292, 299, 321, 328, 333,
		340, 351, 355, 361, 367, 374, 379, 385, 418, 424, 428, 430, 446, 462, 467,
		473, 480, 485, 491, 500, 507, 512, 518, 535, 543, 564, 571, 576, 582, 591,
		598, 603, 609,
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
	VyLangParserT__20      = 21
	VyLangParserT__21      = 22
	VyLangParserT__22      = 23
	VyLangParserT__23      = 24
	VyLangParserT__24      = 25
	VyLangParserT__25      = 26
	VyLangParserT__26      = 27
	VyLangParserT__27      = 28
	VyLangParserT__28      = 29
	VyLangParserT__29      = 30
	VyLangParserT__30      = 31
	VyLangParserT__31      = 32
	VyLangParserT__32      = 33
	VyLangParserT__33      = 34
	VyLangParserT__34      = 35
	VyLangParserT__35      = 36
	VyLangParserT__36      = 37
	VyLangParserT__37      = 38
	VyLangParserT__38      = 39
	VyLangParserT__39      = 40
	VyLangParserT__40      = 41
	VyLangParserT__41      = 42
	VyLangParserT__42      = 43
	VyLangParserT__43      = 44
	VyLangParserT__44      = 45
	VyLangParserT__45      = 46
	VyLangParserT__46      = 47
	VyLangParserT__47      = 48
	VyLangParserT__48      = 49
	VyLangParserT__49      = 50
	VyLangParserT__50      = 51
	VyLangParserSTRING     = 52
	VyLangParserIDENTIFIER = 53
	VyLangParserNEWLINE    = 54
	VyLangParserWHITESPACE = 55
	VyLangParserINT        = 56
)

// VyLangParser rules.
const (
	VyLangParserRULE_prog             = 0
	VyLangParserRULE_definitions      = 1
	VyLangParserRULE_definition       = 2
	VyLangParserRULE_variable         = 3
	VyLangParserRULE_namedAction      = 4
	VyLangParserRULE_action           = 5
	VyLangParserRULE_actionParallel   = 6
	VyLangParserRULE_actionSequence   = 7
	VyLangParserRULE_actionAsync      = 8
	VyLangParserRULE_actionReference  = 9
	VyLangParserRULE_actionBind       = 10
	VyLangParserRULE_actionWhile      = 11
	VyLangParserRULE_actionCond       = 12
	VyLangParserRULE_actionIf         = 13
	VyLangParserRULE_actionIfElse     = 14
	VyLangParserRULE_actionLeaf       = 15
	VyLangParserRULE_actionAssign     = 16
	VyLangParserRULE_actionClear      = 17
	VyLangParserRULE_binding          = 18
	VyLangParserRULE_namedType        = 19
	VyLangParserRULE_type             = 20
	VyLangParserRULE_rawType          = 21
	VyLangParserRULE_typeMap          = 22
	VyLangParserRULE_typeList         = 23
	VyLangParserRULE_typeReference    = 24
	VyLangParserRULE_typeLeaf         = 25
	VyLangParserRULE_typeMapEntry     = 26
	VyLangParserRULE_namedExpr        = 27
	VyLangParserRULE_expr             = 28
	VyLangParserRULE_exprBrackets     = 29
	VyLangParserRULE_exprAlternative  = 30
	VyLangParserRULE_exprOperator1    = 31
	VyLangParserRULE_exprOperator2    = 32
	VyLangParserRULE_exprReference    = 33
	VyLangParserRULE_exprMap          = 34
	VyLangParserRULE_exprList         = 35
	VyLangParserRULE_exprMapEntry     = 36
	VyLangParserRULE_namedLiteral     = 37
	VyLangParserRULE_literal          = 38
	VyLangParserRULE_literalTerminal  = 39
	VyLangParserRULE_literalString    = 40
	VyLangParserRULE_literalInt       = 41
	VyLangParserRULE_literalBoolean   = 42
	VyLangParserRULE_literalFloat     = 43
	VyLangParserRULE_literalNull      = 44
	VyLangParserRULE_literalReference = 45
	VyLangParserRULE_literalMap       = 46
	VyLangParserRULE_literalList      = 47
	VyLangParserRULE_literalMapEntry  = 48
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) Definitions() IDefinitionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefinitionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefinitionsContext)
}

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(VyLangParserEOF, 0)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *VyLangParser) Prog() (localctx IProgContext) {
	this := p
	_ = this

	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, VyLangParserRULE_prog)

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
		p.SetState(98)
		p.Definitions()
	}
	{
		p.SetState(99)
		p.Match(VyLangParserEOF)
	}

	return localctx
}

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

func (s *DefinitionsContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(VyLangParserNEWLINE)
}

func (s *DefinitionsContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(VyLangParserNEWLINE, i)
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
	p.EnterRule(localctx, 2, VyLangParserRULE_definitions)
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
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&18295875650453512) != 0 {
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == VyLangParserNEWLINE {
			{
				p.SetState(101)
				p.Match(VyLangParserNEWLINE)
			}

			p.SetState(106)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(107)
			p.Definition()
		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(108)
					p.Match(VyLangParserNEWLINE)
				}

			}
			p.SetState(113)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
		}

		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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

func (s *DefinitionContext) NamedAction() INamedActionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamedActionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamedActionContext)
}

func (s *DefinitionContext) NamedType() INamedTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamedTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamedTypeContext)
}

func (s *DefinitionContext) NamedLiteral() INamedLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamedLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamedLiteralContext)
}

func (s *DefinitionContext) NamedExpr() INamedExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamedExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamedExprContext)
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
	p.EnterRule(localctx, 4, VyLangParserRULE_definition)

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

	p.SetState(123)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case VyLangParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.NamedAction()
		}

	case VyLangParserT__23:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(120)
			p.NamedType()
		}

	case VyLangParserT__47:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(121)
			p.NamedLiteral()
		}

	case VyLangParserT__30:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(122)
			p.NamedExpr()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VyLangParserIDENTIFIER, 0)
}

func (s *VariableContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *VyLangParser) Variable() (localctx IVariableContext) {
	return p.variable(0)
}

func (p *VyLangParser) variable(_p int) (localctx IVariableContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewVariableContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IVariableContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, VyLangParserRULE_variable, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
	p.SetState(128)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case VyLangParserIDENTIFIER:
		{
			p.SetState(126)
			p.Match(VyLangParserIDENTIFIER)
		}

	case VyLangParserT__1:
		{
			p.SetState(127)
			p.Match(VyLangParserT__1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewVariableContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, VyLangParserRULE_variable)
			p.SetState(130)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(131)
				p.Match(VyLangParserT__0)
			}
			{
				p.SetState(132)
				p.Match(VyLangParserIDENTIFIER)
			}

		}
		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// INamedActionContext is an interface to support dynamic dispatch.
type INamedActionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamedActionContext differentiates from other interfaces.
	IsNamedActionContext()
}

type NamedActionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamedActionContext() *NamedActionContext {
	var p = new(NamedActionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_namedAction
	return p
}

func (*NamedActionContext) IsNamedActionContext() {}

func NewNamedActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamedActionContext {
	var p = new(NamedActionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_namedAction

	return p
}

func (s *NamedActionContext) GetParser() antlr.Parser { return s.parser }

func (s *NamedActionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VyLangParserIDENTIFIER, 0)
}

func (s *NamedActionContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *NamedActionContext) Action_() IActionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionContext)
}

func (s *NamedActionContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(VyLangParserNEWLINE)
}

func (s *NamedActionContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(VyLangParserNEWLINE, i)
}

func (s *NamedActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedActionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamedActionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterNamedAction(s)
	}
}

func (s *NamedActionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitNamedAction(s)
	}
}

func (p *VyLangParser) NamedAction() (localctx INamedActionContext) {
	this := p
	_ = this

	localctx = NewNamedActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, VyLangParserRULE_namedAction)
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
		p.SetState(138)
		p.Match(VyLangParserT__2)
	}
	{
		p.SetState(139)
		p.Match(VyLangParserIDENTIFIER)
	}
	{
		p.SetState(140)
		p.Match(VyLangParserT__3)
	}
	{
		p.SetState(141)
		p.Type_()
	}
	{
		p.SetState(142)
		p.Match(VyLangParserT__4)
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == VyLangParserNEWLINE {
		{
			p.SetState(143)
			p.Match(VyLangParserNEWLINE)
		}

		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(149)
		p.Action_()
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(150)
				p.Match(VyLangParserNEWLINE)
			}

		}
		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	return localctx
}

// IActionContext is an interface to support dynamic dispatch.
type IActionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActionContext differentiates from other interfaces.
	IsActionContext()
}

type ActionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionContext() *ActionContext {
	var p = new(ActionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_action
	return p
}

func (*ActionContext) IsActionContext() {}

func NewActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionContext {
	var p = new(ActionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_action

	return p
}

func (s *ActionContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionContext) ActionParallel() IActionParallelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionParallelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionParallelContext)
}

func (s *ActionContext) ActionSequence() IActionSequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionSequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionSequenceContext)
}

func (s *ActionContext) ActionAsync() IActionAsyncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionAsyncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionAsyncContext)
}

func (s *ActionContext) ActionReference() IActionReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionReferenceContext)
}

func (s *ActionContext) ActionBind() IActionBindContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionBindContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionBindContext)
}

func (s *ActionContext) ActionWhile() IActionWhileContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionWhileContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionWhileContext)
}

func (s *ActionContext) ActionCond() IActionCondContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionCondContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionCondContext)
}

func (s *ActionContext) ActionLeaf() IActionLeafContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionLeafContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionLeafContext)
}

func (s *ActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterAction(s)
	}
}

func (s *ActionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitAction(s)
	}
}

func (p *VyLangParser) Action_() (localctx IActionContext) {
	this := p
	_ = this

	localctx = NewActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, VyLangParserRULE_action)

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

	p.SetState(164)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case VyLangParserT__5:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(156)
			p.ActionParallel()
		}

	case VyLangParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(157)
			p.ActionSequence()
		}

	case VyLangParserT__10:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(158)
			p.ActionAsync()
		}

	case VyLangParserT__11:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(159)
			p.ActionReference()
		}

	case VyLangParserT__12:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(160)
			p.ActionBind()
		}

	case VyLangParserT__14:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(161)
			p.ActionWhile()
		}

	case VyLangParserT__15:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(162)
			p.ActionCond()
		}

	case VyLangParserT__1, VyLangParserT__17, VyLangParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(163)
			p.ActionLeaf()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IActionParallelContext is an interface to support dynamic dispatch.
type IActionParallelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActionParallelContext differentiates from other interfaces.
	IsActionParallelContext()
}

type ActionParallelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionParallelContext() *ActionParallelContext {
	var p = new(ActionParallelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_actionParallel
	return p
}

func (*ActionParallelContext) IsActionParallelContext() {}

func NewActionParallelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionParallelContext {
	var p = new(ActionParallelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_actionParallel

	return p
}

func (s *ActionParallelContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionParallelContext) AllAction_() []IActionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IActionContext); ok {
			len++
		}
	}

	tst := make([]IActionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IActionContext); ok {
			tst[i] = t.(IActionContext)
			i++
		}
	}

	return tst
}

func (s *ActionParallelContext) Action_(i int) IActionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionContext); ok {
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

	return t.(IActionContext)
}

func (s *ActionParallelContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(VyLangParserNEWLINE)
}

func (s *ActionParallelContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(VyLangParserNEWLINE, i)
}

func (s *ActionParallelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionParallelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionParallelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterActionParallel(s)
	}
}

func (s *ActionParallelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitActionParallel(s)
	}
}

func (p *VyLangParser) ActionParallel() (localctx IActionParallelContext) {
	this := p
	_ = this

	localctx = NewActionParallelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, VyLangParserRULE_actionParallel)
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
		p.SetState(166)
		p.Match(VyLangParserT__5)
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(170)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == VyLangParserNEWLINE {
				{
					p.SetState(167)
					p.Match(VyLangParserNEWLINE)
				}

				p.SetState(172)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(173)
				p.Action_()
			}
			p.SetState(175)
			p.GetErrorHandler().Sync(p)
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(174)
						_la = p.GetTokenStream().LA(1)

						if !(_la == VyLangParserT__6 || _la == VyLangParserNEWLINE) {
							p.GetErrorHandler().RecoverInline(p)
						} else {
							p.GetErrorHandler().ReportMatch(p)
							p.Consume()
						}
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

				p.SetState(177)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
			}

		}
		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == VyLangParserNEWLINE {
		{
			p.SetState(184)
			p.Match(VyLangParserNEWLINE)
		}

		p.SetState(189)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(190)
		p.Match(VyLangParserT__7)
	}

	return localctx
}

// IActionSequenceContext is an interface to support dynamic dispatch.
type IActionSequenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActionSequenceContext differentiates from other interfaces.
	IsActionSequenceContext()
}

type ActionSequenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionSequenceContext() *ActionSequenceContext {
	var p = new(ActionSequenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_actionSequence
	return p
}

func (*ActionSequenceContext) IsActionSequenceContext() {}

func NewActionSequenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionSequenceContext {
	var p = new(ActionSequenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_actionSequence

	return p
}

func (s *ActionSequenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionSequenceContext) AllAction_() []IActionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IActionContext); ok {
			len++
		}
	}

	tst := make([]IActionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IActionContext); ok {
			tst[i] = t.(IActionContext)
			i++
		}
	}

	return tst
}

func (s *ActionSequenceContext) Action_(i int) IActionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionContext); ok {
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

	return t.(IActionContext)
}

func (s *ActionSequenceContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(VyLangParserNEWLINE)
}

func (s *ActionSequenceContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(VyLangParserNEWLINE, i)
}

func (s *ActionSequenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionSequenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionSequenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterActionSequence(s)
	}
}

func (s *ActionSequenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitActionSequence(s)
	}
}

func (p *VyLangParser) ActionSequence() (localctx IActionSequenceContext) {
	this := p
	_ = this

	localctx = NewActionSequenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, VyLangParserRULE_actionSequence)
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
		p.SetState(192)
		p.Match(VyLangParserT__8)
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(196)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == VyLangParserNEWLINE {
				{
					p.SetState(193)
					p.Match(VyLangParserNEWLINE)
				}

				p.SetState(198)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(199)
				p.Action_()
			}
			p.SetState(201)
			p.GetErrorHandler().Sync(p)
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(200)
						_la = p.GetTokenStream().LA(1)

						if !(_la == VyLangParserT__6 || _la == VyLangParserNEWLINE) {
							p.GetErrorHandler().RecoverInline(p)
						} else {
							p.GetErrorHandler().ReportMatch(p)
							p.Consume()
						}
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

				p.SetState(203)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
			}

		}
		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == VyLangParserNEWLINE {
		{
			p.SetState(210)
			p.Match(VyLangParserNEWLINE)
		}

		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(216)
		p.Match(VyLangParserT__9)
	}

	return localctx
}

// IActionAsyncContext is an interface to support dynamic dispatch.
type IActionAsyncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActionAsyncContext differentiates from other interfaces.
	IsActionAsyncContext()
}

type ActionAsyncContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionAsyncContext() *ActionAsyncContext {
	var p = new(ActionAsyncContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_actionAsync
	return p
}

func (*ActionAsyncContext) IsActionAsyncContext() {}

func NewActionAsyncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionAsyncContext {
	var p = new(ActionAsyncContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_actionAsync

	return p
}

func (s *ActionAsyncContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionAsyncContext) Action_() IActionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionContext)
}

func (s *ActionAsyncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionAsyncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionAsyncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterActionAsync(s)
	}
}

func (s *ActionAsyncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitActionAsync(s)
	}
}

func (p *VyLangParser) ActionAsync() (localctx IActionAsyncContext) {
	this := p
	_ = this

	localctx = NewActionAsyncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, VyLangParserRULE_actionAsync)

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
		p.SetState(218)
		p.Match(VyLangParserT__10)
	}
	{
		p.SetState(219)
		p.Action_()
	}

	return localctx
}

// IActionReferenceContext is an interface to support dynamic dispatch.
type IActionReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActionReferenceContext differentiates from other interfaces.
	IsActionReferenceContext()
}

type ActionReferenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionReferenceContext() *ActionReferenceContext {
	var p = new(ActionReferenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_actionReference
	return p
}

func (*ActionReferenceContext) IsActionReferenceContext() {}

func NewActionReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionReferenceContext {
	var p = new(ActionReferenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_actionReference

	return p
}

func (s *ActionReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionReferenceContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VyLangParserIDENTIFIER, 0)
}

func (s *ActionReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterActionReference(s)
	}
}

func (s *ActionReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitActionReference(s)
	}
}

func (p *VyLangParser) ActionReference() (localctx IActionReferenceContext) {
	this := p
	_ = this

	localctx = NewActionReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, VyLangParserRULE_actionReference)

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
		p.SetState(221)
		p.Match(VyLangParserT__11)
	}
	{
		p.SetState(222)
		p.Match(VyLangParserIDENTIFIER)
	}

	return localctx
}

// IActionBindContext is an interface to support dynamic dispatch.
type IActionBindContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActionBindContext differentiates from other interfaces.
	IsActionBindContext()
}

type ActionBindContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionBindContext() *ActionBindContext {
	var p = new(ActionBindContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_actionBind
	return p
}

func (*ActionBindContext) IsActionBindContext() {}

func NewActionBindContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionBindContext {
	var p = new(ActionBindContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_actionBind

	return p
}

func (s *ActionBindContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionBindContext) Action_() IActionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionContext)
}

func (s *ActionBindContext) Binding() IBindingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindingContext)
}

func (s *ActionBindContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ActionBindContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ActionBindContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(VyLangParserNEWLINE)
}

func (s *ActionBindContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(VyLangParserNEWLINE, i)
}

func (s *ActionBindContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionBindContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionBindContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterActionBind(s)
	}
}

func (s *ActionBindContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitActionBind(s)
	}
}

func (p *VyLangParser) ActionBind() (localctx IActionBindContext) {
	this := p
	_ = this

	localctx = NewActionBindContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, VyLangParserRULE_actionBind)
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
		p.SetState(224)
		p.Match(VyLangParserT__12)
	}
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&89509042598187588) != 0 {
		{
			p.SetState(225)
			p.binding(0)
		}

	}
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == VyLangParserT__3 {
		{
			p.SetState(228)
			p.Match(VyLangParserT__3)
		}
		{
			p.SetState(229)
			p.Type_()
		}

	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == VyLangParserT__13 {
		{
			p.SetState(232)
			p.Match(VyLangParserT__13)
		}
		{
			p.SetState(233)
			p.expr(0)
		}

	}
	{
		p.SetState(236)
		p.Match(VyLangParserT__4)
	}
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == VyLangParserNEWLINE {
		{
			p.SetState(237)
			p.Match(VyLangParserNEWLINE)
		}

		p.SetState(242)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(243)
		p.Action_()
	}

	return localctx
}

// IActionWhileContext is an interface to support dynamic dispatch.
type IActionWhileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActionWhileContext differentiates from other interfaces.
	IsActionWhileContext()
}

type ActionWhileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionWhileContext() *ActionWhileContext {
	var p = new(ActionWhileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_actionWhile
	return p
}

func (*ActionWhileContext) IsActionWhileContext() {}

func NewActionWhileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionWhileContext {
	var p = new(ActionWhileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_actionWhile

	return p
}

func (s *ActionWhileContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionWhileContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ActionWhileContext) Action_() IActionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionContext)
}

func (s *ActionWhileContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(VyLangParserNEWLINE)
}

func (s *ActionWhileContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(VyLangParserNEWLINE, i)
}

func (s *ActionWhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionWhileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionWhileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterActionWhile(s)
	}
}

func (s *ActionWhileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitActionWhile(s)
	}
}

func (p *VyLangParser) ActionWhile() (localctx IActionWhileContext) {
	this := p
	_ = this

	localctx = NewActionWhileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, VyLangParserRULE_actionWhile)
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
		p.SetState(245)
		p.Match(VyLangParserT__14)
	}
	{
		p.SetState(246)
		p.expr(0)
	}
	{
		p.SetState(247)
		p.Match(VyLangParserT__4)
	}
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == VyLangParserNEWLINE {
		{
			p.SetState(248)
			p.Match(VyLangParserNEWLINE)
		}

		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(254)
		p.Action_()
	}

	return localctx
}

// IActionCondContext is an interface to support dynamic dispatch.
type IActionCondContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActionCondContext differentiates from other interfaces.
	IsActionCondContext()
}

type ActionCondContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionCondContext() *ActionCondContext {
	var p = new(ActionCondContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_actionCond
	return p
}

func (*ActionCondContext) IsActionCondContext() {}

func NewActionCondContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionCondContext {
	var p = new(ActionCondContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_actionCond

	return p
}

func (s *ActionCondContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionCondContext) ActionIf() IActionIfContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionIfContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionIfContext)
}

func (s *ActionCondContext) ActionIfElse() IActionIfElseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionIfElseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionIfElseContext)
}

func (s *ActionCondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionCondContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionCondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterActionCond(s)
	}
}

func (s *ActionCondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitActionCond(s)
	}
}

func (p *VyLangParser) ActionCond() (localctx IActionCondContext) {
	this := p
	_ = this

	localctx = NewActionCondContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, VyLangParserRULE_actionCond)

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

	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(256)
			p.ActionIf()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(257)
			p.ActionIfElse()
		}

	}

	return localctx
}

// IActionIfContext is an interface to support dynamic dispatch.
type IActionIfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActionIfContext differentiates from other interfaces.
	IsActionIfContext()
}

type ActionIfContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionIfContext() *ActionIfContext {
	var p = new(ActionIfContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_actionIf
	return p
}

func (*ActionIfContext) IsActionIfContext() {}

func NewActionIfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionIfContext {
	var p = new(ActionIfContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_actionIf

	return p
}

func (s *ActionIfContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionIfContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ActionIfContext) Action_() IActionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionContext)
}

func (s *ActionIfContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(VyLangParserNEWLINE)
}

func (s *ActionIfContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(VyLangParserNEWLINE, i)
}

func (s *ActionIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionIfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionIfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterActionIf(s)
	}
}

func (s *ActionIfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitActionIf(s)
	}
}

func (p *VyLangParser) ActionIf() (localctx IActionIfContext) {
	this := p
	_ = this

	localctx = NewActionIfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, VyLangParserRULE_actionIf)
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
		p.SetState(260)
		p.Match(VyLangParserT__15)
	}
	{
		p.SetState(261)
		p.expr(0)
	}
	{
		p.SetState(262)
		p.Match(VyLangParserT__4)
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == VyLangParserNEWLINE {
		{
			p.SetState(263)
			p.Match(VyLangParserNEWLINE)
		}

		p.SetState(268)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(269)
		p.Action_()
	}

	return localctx
}

// IActionIfElseContext is an interface to support dynamic dispatch.
type IActionIfElseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActionIfElseContext differentiates from other interfaces.
	IsActionIfElseContext()
}

type ActionIfElseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionIfElseContext() *ActionIfElseContext {
	var p = new(ActionIfElseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_actionIfElse
	return p
}

func (*ActionIfElseContext) IsActionIfElseContext() {}

func NewActionIfElseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionIfElseContext {
	var p = new(ActionIfElseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_actionIfElse

	return p
}

func (s *ActionIfElseContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionIfElseContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ActionIfElseContext) AllAction_() []IActionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IActionContext); ok {
			len++
		}
	}

	tst := make([]IActionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IActionContext); ok {
			tst[i] = t.(IActionContext)
			i++
		}
	}

	return tst
}

func (s *ActionIfElseContext) Action_(i int) IActionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionContext); ok {
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

	return t.(IActionContext)
}

func (s *ActionIfElseContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(VyLangParserNEWLINE)
}

func (s *ActionIfElseContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(VyLangParserNEWLINE, i)
}

func (s *ActionIfElseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionIfElseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionIfElseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterActionIfElse(s)
	}
}

func (s *ActionIfElseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitActionIfElse(s)
	}
}

func (p *VyLangParser) ActionIfElse() (localctx IActionIfElseContext) {
	this := p
	_ = this

	localctx = NewActionIfElseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, VyLangParserRULE_actionIfElse)
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
		p.SetState(271)
		p.Match(VyLangParserT__15)
	}
	{
		p.SetState(272)
		p.expr(0)
	}
	{
		p.SetState(273)
		p.Match(VyLangParserT__4)
	}
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == VyLangParserNEWLINE {
		{
			p.SetState(274)
			p.Match(VyLangParserNEWLINE)
		}

		p.SetState(279)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(280)
		p.Action_()
	}
	p.SetState(284)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == VyLangParserNEWLINE {
		{
			p.SetState(281)
			p.Match(VyLangParserNEWLINE)
		}

		p.SetState(286)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(287)
		p.Match(VyLangParserT__16)
	}
	{
		p.SetState(288)
		p.Match(VyLangParserT__4)
	}
	p.SetState(292)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == VyLangParserNEWLINE {
		{
			p.SetState(289)
			p.Match(VyLangParserNEWLINE)
		}

		p.SetState(294)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(295)
		p.Action_()
	}

	return localctx
}

// IActionLeafContext is an interface to support dynamic dispatch.
type IActionLeafContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActionLeafContext differentiates from other interfaces.
	IsActionLeafContext()
}

type ActionLeafContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionLeafContext() *ActionLeafContext {
	var p = new(ActionLeafContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_actionLeaf
	return p
}

func (*ActionLeafContext) IsActionLeafContext() {}

func NewActionLeafContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionLeafContext {
	var p = new(ActionLeafContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_actionLeaf

	return p
}

func (s *ActionLeafContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionLeafContext) ActionAssign() IActionAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionAssignContext)
}

func (s *ActionLeafContext) ActionClear() IActionClearContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionClearContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionClearContext)
}

func (s *ActionLeafContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionLeafContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionLeafContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterActionLeaf(s)
	}
}

func (s *ActionLeafContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitActionLeaf(s)
	}
}

func (p *VyLangParser) ActionLeaf() (localctx IActionLeafContext) {
	this := p
	_ = this

	localctx = NewActionLeafContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, VyLangParserRULE_actionLeaf)

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

	p.SetState(299)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case VyLangParserT__1, VyLangParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(297)
			p.ActionAssign()
		}

	case VyLangParserT__17:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(298)
			p.ActionClear()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IActionAssignContext is an interface to support dynamic dispatch.
type IActionAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActionAssignContext differentiates from other interfaces.
	IsActionAssignContext()
}

type ActionAssignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionAssignContext() *ActionAssignContext {
	var p = new(ActionAssignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_actionAssign
	return p
}

func (*ActionAssignContext) IsActionAssignContext() {}

func NewActionAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionAssignContext {
	var p = new(ActionAssignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_actionAssign

	return p
}

func (s *ActionAssignContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionAssignContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ActionAssignContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ActionAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionAssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterActionAssign(s)
	}
}

func (s *ActionAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitActionAssign(s)
	}
}

func (p *VyLangParser) ActionAssign() (localctx IActionAssignContext) {
	this := p
	_ = this

	localctx = NewActionAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, VyLangParserRULE_actionAssign)

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
		p.SetState(301)
		p.variable(0)
	}
	{
		p.SetState(302)
		p.Match(VyLangParserT__13)
	}
	{
		p.SetState(303)
		p.expr(0)
	}

	return localctx
}

// IActionClearContext is an interface to support dynamic dispatch.
type IActionClearContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActionClearContext differentiates from other interfaces.
	IsActionClearContext()
}

type ActionClearContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionClearContext() *ActionClearContext {
	var p = new(ActionClearContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_actionClear
	return p
}

func (*ActionClearContext) IsActionClearContext() {}

func NewActionClearContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionClearContext {
	var p = new(ActionClearContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_actionClear

	return p
}

func (s *ActionClearContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionClearContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ActionClearContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionClearContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionClearContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterActionClear(s)
	}
}

func (s *ActionClearContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitActionClear(s)
	}
}

func (p *VyLangParser) ActionClear() (localctx IActionClearContext) {
	this := p
	_ = this

	localctx = NewActionClearContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, VyLangParserRULE_actionClear)

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
		p.SetState(305)
		p.Match(VyLangParserT__17)
	}
	{
		p.SetState(306)
		p.variable(0)
	}

	return localctx
}

// IBindingContext is an interface to support dynamic dispatch.
type IBindingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBindingContext differentiates from other interfaces.
	IsBindingContext()
}

type BindingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindingContext() *BindingContext {
	var p = new(BindingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_binding
	return p
}

func (*BindingContext) IsBindingContext() {}

func NewBindingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindingContext {
	var p = new(BindingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_binding

	return p
}

func (s *BindingContext) GetParser() antlr.Parser { return s.parser }

func (s *BindingContext) AllVariable() []IVariableContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableContext); ok {
			len++
		}
	}

	tst := make([]IVariableContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableContext); ok {
			tst[i] = t.(IVariableContext)
			i++
		}
	}

	return tst
}

func (s *BindingContext) Variable(i int) IVariableContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
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

	return t.(IVariableContext)
}

func (s *BindingContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *BindingContext) AllBinding() []IBindingContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBindingContext); ok {
			len++
		}
	}

	tst := make([]IBindingContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBindingContext); ok {
			tst[i] = t.(IBindingContext)
			i++
		}
	}

	return tst
}

func (s *BindingContext) Binding(i int) IBindingContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindingContext); ok {
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

	return t.(IBindingContext)
}

func (s *BindingContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(VyLangParserNEWLINE)
}

func (s *BindingContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(VyLangParserNEWLINE, i)
}

func (s *BindingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BindingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterBinding(s)
	}
}

func (s *BindingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitBinding(s)
	}
}

func (p *VyLangParser) Binding() (localctx IBindingContext) {
	return p.binding(0)
}

func (p *VyLangParser) binding(_p int) (localctx IBindingContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewBindingContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBindingContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 36
	p.EnterRecursionRule(localctx, 36, VyLangParserRULE_binding, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
	p.SetState(333)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case VyLangParserT__1, VyLangParserIDENTIFIER:
		{
			p.SetState(309)
			p.variable(0)
		}
		{
			p.SetState(310)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3670016) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(311)
			p.variable(0)
		}

	case VyLangParserT__5, VyLangParserT__8, VyLangParserT__11, VyLangParserT__48, VyLangParserT__49, VyLangParserT__50, VyLangParserSTRING, VyLangParserINT:
		{
			p.SetState(313)
			p.Literal()
		}
		{
			p.SetState(314)
			p.Match(VyLangParserT__19)
		}
		{
			p.SetState(315)
			p.variable(0)
		}

	case VyLangParserT__21:
		{
			p.SetState(317)
			p.Match(VyLangParserT__21)
		}
		p.SetState(321)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == VyLangParserNEWLINE {
			{
				p.SetState(318)
				p.Match(VyLangParserNEWLINE)
			}

			p.SetState(323)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(324)
			p.binding(0)
		}
		p.SetState(328)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == VyLangParserNEWLINE {
			{
				p.SetState(325)
				p.Match(VyLangParserNEWLINE)
			}

			p.SetState(330)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(331)
			p.Match(VyLangParserT__22)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(340)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewBindingContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, VyLangParserRULE_binding)
			p.SetState(335)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(336)
				p.Match(VyLangParserT__6)
			}
			{
				p.SetState(337)
				p.binding(3)
			}

		}
		p.SetState(342)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())
	}

	return localctx
}

// INamedTypeContext is an interface to support dynamic dispatch.
type INamedTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamedTypeContext differentiates from other interfaces.
	IsNamedTypeContext()
}

type NamedTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamedTypeContext() *NamedTypeContext {
	var p = new(NamedTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_namedType
	return p
}

func (*NamedTypeContext) IsNamedTypeContext() {}

func NewNamedTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamedTypeContext {
	var p = new(NamedTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_namedType

	return p
}

func (s *NamedTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *NamedTypeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VyLangParserIDENTIFIER, 0)
}

func (s *NamedTypeContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *NamedTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamedTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterNamedType(s)
	}
}

func (s *NamedTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitNamedType(s)
	}
}

func (p *VyLangParser) NamedType() (localctx INamedTypeContext) {
	this := p
	_ = this

	localctx = NewNamedTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, VyLangParserRULE_namedType)

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
		p.SetState(343)
		p.Match(VyLangParserT__23)
	}
	{
		p.SetState(344)
		p.Match(VyLangParserIDENTIFIER)
	}
	{
		p.SetState(345)
		p.Match(VyLangParserT__4)
	}
	{
		p.SetState(346)
		p.Type_()
	}

	return localctx
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_type
	return p
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) RawType() IRawTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRawTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRawTypeContext)
}

func (s *TypeContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *TypeContext) LiteralList() ILiteralListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralListContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *VyLangParser) Type_() (localctx ITypeContext) {
	this := p
	_ = this

	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, VyLangParserRULE_type)

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
		p.SetState(348)
		p.RawType()
	}
	p.SetState(351)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(349)
			p.Match(VyLangParserT__13)
		}
		{
			p.SetState(350)
			p.Literal()
		}

	}
	p.SetState(355)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(353)
			p.Match(VyLangParserT__24)
		}
		{
			p.SetState(354)
			p.LiteralList()
		}

	}

	return localctx
}

// IRawTypeContext is an interface to support dynamic dispatch.
type IRawTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRawTypeContext differentiates from other interfaces.
	IsRawTypeContext()
}

type RawTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRawTypeContext() *RawTypeContext {
	var p = new(RawTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_rawType
	return p
}

func (*RawTypeContext) IsRawTypeContext() {}

func NewRawTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RawTypeContext {
	var p = new(RawTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_rawType

	return p
}

func (s *RawTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *RawTypeContext) TypeMap() ITypeMapContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeMapContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeMapContext)
}

func (s *RawTypeContext) TypeList() ITypeListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeListContext)
}

func (s *RawTypeContext) TypeReference() ITypeReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeReferenceContext)
}

func (s *RawTypeContext) TypeLeaf() ITypeLeafContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeLeafContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeLeafContext)
}

func (s *RawTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RawTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RawTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterRawType(s)
	}
}

func (s *RawTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitRawType(s)
	}
}

func (p *VyLangParser) RawType() (localctx IRawTypeContext) {
	this := p
	_ = this

	localctx = NewRawTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, VyLangParserRULE_rawType)

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

	p.SetState(361)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case VyLangParserT__5:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(357)
			p.TypeMap()
		}

	case VyLangParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(358)
			p.TypeList()
		}

	case VyLangParserT__11:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(359)
			p.TypeReference()
		}

	case VyLangParserT__25, VyLangParserT__26, VyLangParserT__27, VyLangParserT__28, VyLangParserT__29:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(360)
			p.TypeLeaf()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITypeMapContext is an interface to support dynamic dispatch.
type ITypeMapContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeMapContext differentiates from other interfaces.
	IsTypeMapContext()
}

type TypeMapContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeMapContext() *TypeMapContext {
	var p = new(TypeMapContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_typeMap
	return p
}

func (*TypeMapContext) IsTypeMapContext() {}

func NewTypeMapContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeMapContext {
	var p = new(TypeMapContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_typeMap

	return p
}

func (s *TypeMapContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeMapContext) AllTypeMapEntry() []ITypeMapEntryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeMapEntryContext); ok {
			len++
		}
	}

	tst := make([]ITypeMapEntryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeMapEntryContext); ok {
			tst[i] = t.(ITypeMapEntryContext)
			i++
		}
	}

	return tst
}

func (s *TypeMapContext) TypeMapEntry(i int) ITypeMapEntryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeMapEntryContext); ok {
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

	return t.(ITypeMapEntryContext)
}

func (s *TypeMapContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(VyLangParserNEWLINE)
}

func (s *TypeMapContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(VyLangParserNEWLINE, i)
}

func (s *TypeMapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeMapContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeMapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterTypeMap(s)
	}
}

func (s *TypeMapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitTypeMap(s)
	}
}

func (p *VyLangParser) TypeMap() (localctx ITypeMapContext) {
	this := p
	_ = this

	localctx = NewTypeMapContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, VyLangParserRULE_typeMap)
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
		p.SetState(363)
		p.Match(VyLangParserT__5)
	}
	p.SetState(367)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == VyLangParserNEWLINE {
		{
			p.SetState(364)
			p.Match(VyLangParserNEWLINE)
		}

		p.SetState(369)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(370)
		p.TypeMapEntry()
	}
	p.SetState(379)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(372)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == VyLangParserT__6 || _la == VyLangParserNEWLINE {
				{
					p.SetState(371)
					_la = p.GetTokenStream().LA(1)

					if !(_la == VyLangParserT__6 || _la == VyLangParserNEWLINE) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

				p.SetState(374)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(376)
				p.TypeMapEntry()
			}

		}
		p.SetState(381)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext())
	}
	p.SetState(385)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == VyLangParserNEWLINE {
		{
			p.SetState(382)
			p.Match(VyLangParserNEWLINE)
		}

		p.SetState(387)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(388)
		p.Match(VyLangParserT__7)
	}

	return localctx
}

// ITypeListContext is an interface to support dynamic dispatch.
type ITypeListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeListContext differentiates from other interfaces.
	IsTypeListContext()
}

type TypeListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeListContext() *TypeListContext {
	var p = new(TypeListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_typeList
	return p
}

func (*TypeListContext) IsTypeListContext() {}

func NewTypeListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeListContext {
	var p = new(TypeListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_typeList

	return p
}

func (s *TypeListContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeListContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *TypeListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterTypeList(s)
	}
}

func (s *TypeListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitTypeList(s)
	}
}

func (p *VyLangParser) TypeList() (localctx ITypeListContext) {
	this := p
	_ = this

	localctx = NewTypeListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, VyLangParserRULE_typeList)

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
		p.SetState(390)
		p.Match(VyLangParserT__8)
	}
	{
		p.SetState(391)
		p.Match(VyLangParserT__9)
	}
	{
		p.SetState(392)
		p.Type_()
	}

	return localctx
}

// ITypeReferenceContext is an interface to support dynamic dispatch.
type ITypeReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeReferenceContext differentiates from other interfaces.
	IsTypeReferenceContext()
}

type TypeReferenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeReferenceContext() *TypeReferenceContext {
	var p = new(TypeReferenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_typeReference
	return p
}

func (*TypeReferenceContext) IsTypeReferenceContext() {}

func NewTypeReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeReferenceContext {
	var p = new(TypeReferenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_typeReference

	return p
}

func (s *TypeReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeReferenceContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VyLangParserIDENTIFIER, 0)
}

func (s *TypeReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterTypeReference(s)
	}
}

func (s *TypeReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitTypeReference(s)
	}
}

func (p *VyLangParser) TypeReference() (localctx ITypeReferenceContext) {
	this := p
	_ = this

	localctx = NewTypeReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, VyLangParserRULE_typeReference)

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
		p.SetState(394)
		p.Match(VyLangParserT__11)
	}
	{
		p.SetState(395)
		p.Match(VyLangParserIDENTIFIER)
	}

	return localctx
}

// ITypeLeafContext is an interface to support dynamic dispatch.
type ITypeLeafContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeLeafContext differentiates from other interfaces.
	IsTypeLeafContext()
}

type TypeLeafContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeLeafContext() *TypeLeafContext {
	var p = new(TypeLeafContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_typeLeaf
	return p
}

func (*TypeLeafContext) IsTypeLeafContext() {}

func NewTypeLeafContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeLeafContext {
	var p = new(TypeLeafContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_typeLeaf

	return p
}

func (s *TypeLeafContext) GetParser() antlr.Parser { return s.parser }
func (s *TypeLeafContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeLeafContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeLeafContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterTypeLeaf(s)
	}
}

func (s *TypeLeafContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitTypeLeaf(s)
	}
}

func (p *VyLangParser) TypeLeaf() (localctx ITypeLeafContext) {
	this := p
	_ = this

	localctx = NewTypeLeafContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, VyLangParserRULE_typeLeaf)
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
		p.SetState(397)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2080374784) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITypeMapEntryContext is an interface to support dynamic dispatch.
type ITypeMapEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeMapEntryContext differentiates from other interfaces.
	IsTypeMapEntryContext()
}

type TypeMapEntryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeMapEntryContext() *TypeMapEntryContext {
	var p = new(TypeMapEntryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_typeMapEntry
	return p
}

func (*TypeMapEntryContext) IsTypeMapEntryContext() {}

func NewTypeMapEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeMapEntryContext {
	var p = new(TypeMapEntryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_typeMapEntry

	return p
}

func (s *TypeMapEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeMapEntryContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VyLangParserIDENTIFIER, 0)
}

func (s *TypeMapEntryContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *TypeMapEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeMapEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeMapEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterTypeMapEntry(s)
	}
}

func (s *TypeMapEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitTypeMapEntry(s)
	}
}

func (p *VyLangParser) TypeMapEntry() (localctx ITypeMapEntryContext) {
	this := p
	_ = this

	localctx = NewTypeMapEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, VyLangParserRULE_typeMapEntry)

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
		p.SetState(399)
		p.Match(VyLangParserIDENTIFIER)
	}
	{
		p.SetState(400)
		p.Match(VyLangParserT__4)
	}
	{
		p.SetState(401)
		p.Type_()
	}

	return localctx
}

// INamedExprContext is an interface to support dynamic dispatch.
type INamedExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamedExprContext differentiates from other interfaces.
	IsNamedExprContext()
}

type NamedExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamedExprContext() *NamedExprContext {
	var p = new(NamedExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_namedExpr
	return p
}

func (*NamedExprContext) IsNamedExprContext() {}

func NewNamedExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamedExprContext {
	var p = new(NamedExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_namedExpr

	return p
}

func (s *NamedExprContext) GetParser() antlr.Parser { return s.parser }

func (s *NamedExprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VyLangParserIDENTIFIER, 0)
}

func (s *NamedExprContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *NamedExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NamedExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamedExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterNamedExpr(s)
	}
}

func (s *NamedExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitNamedExpr(s)
	}
}

func (p *VyLangParser) NamedExpr() (localctx INamedExprContext) {
	this := p
	_ = this

	localctx = NewNamedExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, VyLangParserRULE_namedExpr)

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
		p.SetState(403)
		p.Match(VyLangParserT__30)
	}
	{
		p.SetState(404)
		p.Match(VyLangParserIDENTIFIER)
	}
	{
		p.SetState(405)
		p.Match(VyLangParserT__3)
	}
	{
		p.SetState(406)
		p.Type_()
	}
	{
		p.SetState(407)
		p.Match(VyLangParserT__4)
	}
	{
		p.SetState(408)
		p.expr(0)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ExprContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ExprContext) ExprMap() IExprMapContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprMapContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprMapContext)
}

func (s *ExprContext) ExprList() IExprListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *ExprContext) ExprBrackets() IExprBracketsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprBracketsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprBracketsContext)
}

func (s *ExprContext) ExprReference() IExprReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprReferenceContext)
}

func (s *ExprContext) ExprOperator1() IExprOperator1Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprOperator1Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprOperator1Context)
}

func (s *ExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) AllExprOperator2() []IExprOperator2Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprOperator2Context); ok {
			len++
		}
	}

	tst := make([]IExprOperator2Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprOperator2Context); ok {
			tst[i] = t.(IExprOperator2Context)
			i++
		}
	}

	return tst
}

func (s *ExprContext) ExprOperator2(i int) IExprOperator2Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprOperator2Context); ok {
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

	return t.(IExprOperator2Context)
}

func (s *ExprContext) ExprAlternative() IExprAlternativeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprAlternativeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprAlternativeContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *VyLangParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *VyLangParser) expr(_p int) (localctx IExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 56
	p.EnterRecursionRule(localctx, 56, VyLangParserRULE_expr, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
	p.SetState(418)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(411)
			p.variable(0)
		}

	case 2:
		{
			p.SetState(412)
			p.Literal()
		}

	case 3:
		{
			p.SetState(413)
			p.ExprMap()
		}

	case 4:
		{
			p.SetState(414)
			p.ExprList()
		}

	case 5:
		{
			p.SetState(415)
			p.ExprBrackets()
		}

	case 6:
		{
			p.SetState(416)
			p.ExprReference()
		}

	case 7:
		{
			p.SetState(417)
			p.ExprOperator1()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(430)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(428)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, VyLangParserRULE_expr)
				p.SetState(420)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				p.SetState(422)
				p.GetErrorHandler().Sync(p)
				_alt = 1
				for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					switch _alt {
					case 1:
						{
							p.SetState(421)
							p.ExprOperator2()
						}

					default:
						panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					}

					p.SetState(424)
					p.GetErrorHandler().Sync(p)
					_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext())
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, VyLangParserRULE_expr)
				p.SetState(426)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(427)
					p.ExprAlternative()
				}

			}

		}
		p.SetState(432)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext())
	}

	return localctx
}

// IExprBracketsContext is an interface to support dynamic dispatch.
type IExprBracketsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprBracketsContext differentiates from other interfaces.
	IsExprBracketsContext()
}

type ExprBracketsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprBracketsContext() *ExprBracketsContext {
	var p = new(ExprBracketsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_exprBrackets
	return p
}

func (*ExprBracketsContext) IsExprBracketsContext() {}

func NewExprBracketsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprBracketsContext {
	var p = new(ExprBracketsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_exprBrackets

	return p
}

func (s *ExprBracketsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprBracketsContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprBracketsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprBracketsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprBracketsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterExprBrackets(s)
	}
}

func (s *ExprBracketsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitExprBrackets(s)
	}
}

func (p *VyLangParser) ExprBrackets() (localctx IExprBracketsContext) {
	this := p
	_ = this

	localctx = NewExprBracketsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, VyLangParserRULE_exprBrackets)

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
		p.SetState(433)
		p.Match(VyLangParserT__21)
	}
	{
		p.SetState(434)
		p.expr(0)
	}
	{
		p.SetState(435)
		p.Match(VyLangParserT__22)
	}

	return localctx
}

// IExprAlternativeContext is an interface to support dynamic dispatch.
type IExprAlternativeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprAlternativeContext differentiates from other interfaces.
	IsExprAlternativeContext()
}

type ExprAlternativeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprAlternativeContext() *ExprAlternativeContext {
	var p = new(ExprAlternativeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_exprAlternative
	return p
}

func (*ExprAlternativeContext) IsExprAlternativeContext() {}

func NewExprAlternativeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprAlternativeContext {
	var p = new(ExprAlternativeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_exprAlternative

	return p
}

func (s *ExprAlternativeContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprAlternativeContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprAlternativeContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ExprAlternativeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprAlternativeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprAlternativeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterExprAlternative(s)
	}
}

func (s *ExprAlternativeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitExprAlternative(s)
	}
}

func (p *VyLangParser) ExprAlternative() (localctx IExprAlternativeContext) {
	this := p
	_ = this

	localctx = NewExprAlternativeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, VyLangParserRULE_exprAlternative)

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
		p.SetState(437)
		p.Match(VyLangParserT__31)
	}
	{
		p.SetState(438)
		p.expr(0)
	}
	{
		p.SetState(439)
		p.Match(VyLangParserT__4)
	}
	{
		p.SetState(440)
		p.expr(0)
	}

	return localctx
}

// IExprOperator1Context is an interface to support dynamic dispatch.
type IExprOperator1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprOperator1Context differentiates from other interfaces.
	IsExprOperator1Context()
}

type ExprOperator1Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprOperator1Context() *ExprOperator1Context {
	var p = new(ExprOperator1Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_exprOperator1
	return p
}

func (*ExprOperator1Context) IsExprOperator1Context() {}

func NewExprOperator1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprOperator1Context {
	var p = new(ExprOperator1Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_exprOperator1

	return p
}

func (s *ExprOperator1Context) GetParser() antlr.Parser { return s.parser }

func (s *ExprOperator1Context) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprOperator1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprOperator1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprOperator1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterExprOperator1(s)
	}
}

func (s *ExprOperator1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitExprOperator1(s)
	}
}

func (p *VyLangParser) ExprOperator1() (localctx IExprOperator1Context) {
	this := p
	_ = this

	localctx = NewExprOperator1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, VyLangParserRULE_exprOperator1)

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

	p.SetState(446)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case VyLangParserT__32:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(442)
			p.Match(VyLangParserT__32)
		}
		{
			p.SetState(443)
			p.expr(0)
		}

	case VyLangParserT__33:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(444)
			p.Match(VyLangParserT__33)
		}
		{
			p.SetState(445)
			p.expr(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExprOperator2Context is an interface to support dynamic dispatch.
type IExprOperator2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprOperator2Context differentiates from other interfaces.
	IsExprOperator2Context()
}

type ExprOperator2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprOperator2Context() *ExprOperator2Context {
	var p = new(ExprOperator2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_exprOperator2
	return p
}

func (*ExprOperator2Context) IsExprOperator2Context() {}

func NewExprOperator2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprOperator2Context {
	var p = new(ExprOperator2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_exprOperator2

	return p
}

func (s *ExprOperator2Context) GetParser() antlr.Parser { return s.parser }

func (s *ExprOperator2Context) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprOperator2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprOperator2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprOperator2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterExprOperator2(s)
	}
}

func (s *ExprOperator2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitExprOperator2(s)
	}
}

func (p *VyLangParser) ExprOperator2() (localctx IExprOperator2Context) {
	this := p
	_ = this

	localctx = NewExprOperator2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, VyLangParserRULE_exprOperator2)
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

	p.SetState(462)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case VyLangParserT__34, VyLangParserT__35:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(448)
			_la = p.GetTokenStream().LA(1)

			if !(_la == VyLangParserT__34 || _la == VyLangParserT__35) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(449)
			p.expr(0)
		}

	case VyLangParserT__32, VyLangParserT__36:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(450)
			_la = p.GetTokenStream().LA(1)

			if !(_la == VyLangParserT__32 || _la == VyLangParserT__36) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(451)
			p.expr(0)
		}

	case VyLangParserT__37, VyLangParserT__38, VyLangParserT__39, VyLangParserT__40:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(452)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4123168604160) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(453)
			p.expr(0)
		}

	case VyLangParserT__41, VyLangParserT__42:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(454)
			_la = p.GetTokenStream().LA(1)

			if !(_la == VyLangParserT__41 || _la == VyLangParserT__42) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(455)
			p.expr(0)
		}

	case VyLangParserT__43:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(456)
			p.Match(VyLangParserT__43)
		}
		{
			p.SetState(457)
			p.expr(0)
		}

	case VyLangParserT__44:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(458)
			p.Match(VyLangParserT__44)
		}
		{
			p.SetState(459)
			p.expr(0)
		}

	case VyLangParserT__45, VyLangParserT__46:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(460)
			_la = p.GetTokenStream().LA(1)

			if !(_la == VyLangParserT__45 || _la == VyLangParserT__46) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(461)
			p.expr(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExprReferenceContext is an interface to support dynamic dispatch.
type IExprReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprReferenceContext differentiates from other interfaces.
	IsExprReferenceContext()
}

type ExprReferenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprReferenceContext() *ExprReferenceContext {
	var p = new(ExprReferenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_exprReference
	return p
}

func (*ExprReferenceContext) IsExprReferenceContext() {}

func NewExprReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprReferenceContext {
	var p = new(ExprReferenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_exprReference

	return p
}

func (s *ExprReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprReferenceContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VyLangParserIDENTIFIER, 0)
}

func (s *ExprReferenceContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterExprReference(s)
	}
}

func (s *ExprReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitExprReference(s)
	}
}

func (p *VyLangParser) ExprReference() (localctx IExprReferenceContext) {
	this := p
	_ = this

	localctx = NewExprReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, VyLangParserRULE_exprReference)

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
		p.SetState(464)
		p.Match(VyLangParserT__11)
	}
	{
		p.SetState(465)
		p.Match(VyLangParserIDENTIFIER)
	}
	p.SetState(467)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(466)
			p.expr(0)
		}

	}

	return localctx
}

// IExprMapContext is an interface to support dynamic dispatch.
type IExprMapContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprMapContext differentiates from other interfaces.
	IsExprMapContext()
}

type ExprMapContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprMapContext() *ExprMapContext {
	var p = new(ExprMapContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_exprMap
	return p
}

func (*ExprMapContext) IsExprMapContext() {}

func NewExprMapContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprMapContext {
	var p = new(ExprMapContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_exprMap

	return p
}

func (s *ExprMapContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprMapContext) AllExprMapEntry() []IExprMapEntryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprMapEntryContext); ok {
			len++
		}
	}

	tst := make([]IExprMapEntryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprMapEntryContext); ok {
			tst[i] = t.(IExprMapEntryContext)
			i++
		}
	}

	return tst
}

func (s *ExprMapContext) ExprMapEntry(i int) IExprMapEntryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprMapEntryContext); ok {
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

	return t.(IExprMapEntryContext)
}

func (s *ExprMapContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(VyLangParserNEWLINE)
}

func (s *ExprMapContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(VyLangParserNEWLINE, i)
}

func (s *ExprMapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprMapContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprMapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterExprMap(s)
	}
}

func (s *ExprMapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitExprMap(s)
	}
}

func (p *VyLangParser) ExprMap() (localctx IExprMapContext) {
	this := p
	_ = this

	localctx = NewExprMapContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, VyLangParserRULE_exprMap)
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
		p.SetState(469)
		p.Match(VyLangParserT__5)
	}
	p.SetState(473)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == VyLangParserNEWLINE {
		{
			p.SetState(470)
			p.Match(VyLangParserNEWLINE)
		}

		p.SetState(475)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(476)
		p.ExprMapEntry()
	}
	p.SetState(485)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(478)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == VyLangParserT__6 || _la == VyLangParserNEWLINE {
				{
					p.SetState(477)
					_la = p.GetTokenStream().LA(1)

					if !(_la == VyLangParserT__6 || _la == VyLangParserNEWLINE) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

				p.SetState(480)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(482)
				p.ExprMapEntry()
			}

		}
		p.SetState(487)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext())
	}
	p.SetState(491)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == VyLangParserNEWLINE {
		{
			p.SetState(488)
			p.Match(VyLangParserNEWLINE)
		}

		p.SetState(493)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(494)
		p.Match(VyLangParserT__7)
	}

	return localctx
}

// IExprListContext is an interface to support dynamic dispatch.
type IExprListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprListContext differentiates from other interfaces.
	IsExprListContext()
}

type ExprListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprListContext() *ExprListContext {
	var p = new(ExprListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_exprList
	return p
}

func (*ExprListContext) IsExprListContext() {}

func NewExprListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprListContext {
	var p = new(ExprListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_exprList

	return p
}

func (s *ExprListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprListContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprListContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ExprListContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(VyLangParserNEWLINE)
}

func (s *ExprListContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(VyLangParserNEWLINE, i)
}

func (s *ExprListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterExprList(s)
	}
}

func (s *ExprListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitExprList(s)
	}
}

func (p *VyLangParser) ExprList() (localctx IExprListContext) {
	this := p
	_ = this

	localctx = NewExprListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, VyLangParserRULE_exprList)
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
		p.SetState(496)
		p.Match(VyLangParserT__8)
	}
	p.SetState(500)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == VyLangParserNEWLINE {
		{
			p.SetState(497)
			p.Match(VyLangParserNEWLINE)
		}

		p.SetState(502)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(503)
		p.expr(0)
	}
	p.SetState(512)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 52, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(505)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == VyLangParserT__6 || _la == VyLangParserNEWLINE {
				{
					p.SetState(504)
					_la = p.GetTokenStream().LA(1)

					if !(_la == VyLangParserT__6 || _la == VyLangParserNEWLINE) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

				p.SetState(507)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(509)
				p.expr(0)
			}

		}
		p.SetState(514)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 52, p.GetParserRuleContext())
	}
	p.SetState(518)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == VyLangParserNEWLINE {
		{
			p.SetState(515)
			p.Match(VyLangParserNEWLINE)
		}

		p.SetState(520)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(521)
		p.Match(VyLangParserT__9)
	}

	return localctx
}

// IExprMapEntryContext is an interface to support dynamic dispatch.
type IExprMapEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprMapEntryContext differentiates from other interfaces.
	IsExprMapEntryContext()
}

type ExprMapEntryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprMapEntryContext() *ExprMapEntryContext {
	var p = new(ExprMapEntryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_exprMapEntry
	return p
}

func (*ExprMapEntryContext) IsExprMapEntryContext() {}

func NewExprMapEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprMapEntryContext {
	var p = new(ExprMapEntryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_exprMapEntry

	return p
}

func (s *ExprMapEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprMapEntryContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VyLangParserIDENTIFIER, 0)
}

func (s *ExprMapEntryContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprMapEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprMapEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprMapEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterExprMapEntry(s)
	}
}

func (s *ExprMapEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitExprMapEntry(s)
	}
}

func (p *VyLangParser) ExprMapEntry() (localctx IExprMapEntryContext) {
	this := p
	_ = this

	localctx = NewExprMapEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, VyLangParserRULE_exprMapEntry)

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
		p.SetState(523)
		p.Match(VyLangParserIDENTIFIER)
	}
	{
		p.SetState(524)
		p.Match(VyLangParserT__4)
	}
	{
		p.SetState(525)
		p.expr(0)
	}

	return localctx
}

// INamedLiteralContext is an interface to support dynamic dispatch.
type INamedLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamedLiteralContext differentiates from other interfaces.
	IsNamedLiteralContext()
}

type NamedLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamedLiteralContext() *NamedLiteralContext {
	var p = new(NamedLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_namedLiteral
	return p
}

func (*NamedLiteralContext) IsNamedLiteralContext() {}

func NewNamedLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamedLiteralContext {
	var p = new(NamedLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_namedLiteral

	return p
}

func (s *NamedLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *NamedLiteralContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VyLangParserIDENTIFIER, 0)
}

func (s *NamedLiteralContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *NamedLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamedLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterNamedLiteral(s)
	}
}

func (s *NamedLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitNamedLiteral(s)
	}
}

func (p *VyLangParser) NamedLiteral() (localctx INamedLiteralContext) {
	this := p
	_ = this

	localctx = NewNamedLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, VyLangParserRULE_namedLiteral)

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
		p.SetState(527)
		p.Match(VyLangParserT__47)
	}
	{
		p.SetState(528)
		p.Match(VyLangParserIDENTIFIER)
	}
	{
		p.SetState(529)
		p.Match(VyLangParserT__4)
	}
	{
		p.SetState(530)
		p.Literal()
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) LiteralTerminal() ILiteralTerminalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralTerminalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralTerminalContext)
}

func (s *LiteralContext) LiteralMap() ILiteralMapContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralMapContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralMapContext)
}

func (s *LiteralContext) LiteralList() ILiteralListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralListContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *VyLangParser) Literal() (localctx ILiteralContext) {
	this := p
	_ = this

	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, VyLangParserRULE_literal)

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

	p.SetState(535)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case VyLangParserT__11, VyLangParserT__48, VyLangParserT__49, VyLangParserT__50, VyLangParserSTRING, VyLangParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(532)
			p.LiteralTerminal()
		}

	case VyLangParserT__5:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(533)
			p.LiteralMap()
		}

	case VyLangParserT__8:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(534)
			p.LiteralList()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILiteralTerminalContext is an interface to support dynamic dispatch.
type ILiteralTerminalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralTerminalContext differentiates from other interfaces.
	IsLiteralTerminalContext()
}

type LiteralTerminalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralTerminalContext() *LiteralTerminalContext {
	var p = new(LiteralTerminalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_literalTerminal
	return p
}

func (*LiteralTerminalContext) IsLiteralTerminalContext() {}

func NewLiteralTerminalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralTerminalContext {
	var p = new(LiteralTerminalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_literalTerminal

	return p
}

func (s *LiteralTerminalContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralTerminalContext) LiteralString() ILiteralStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralStringContext)
}

func (s *LiteralTerminalContext) LiteralInt() ILiteralIntContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralIntContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralIntContext)
}

func (s *LiteralTerminalContext) LiteralBoolean() ILiteralBooleanContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralBooleanContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralBooleanContext)
}

func (s *LiteralTerminalContext) LiteralFloat() ILiteralFloatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralFloatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralFloatContext)
}

func (s *LiteralTerminalContext) LiteralNull() ILiteralNullContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralNullContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralNullContext)
}

func (s *LiteralTerminalContext) LiteralReference() ILiteralReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralReferenceContext)
}

func (s *LiteralTerminalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralTerminalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralTerminalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterLiteralTerminal(s)
	}
}

func (s *LiteralTerminalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitLiteralTerminal(s)
	}
}

func (p *VyLangParser) LiteralTerminal() (localctx ILiteralTerminalContext) {
	this := p
	_ = this

	localctx = NewLiteralTerminalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, VyLangParserRULE_literalTerminal)

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

	p.SetState(543)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 55, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(537)
			p.LiteralString()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(538)
			p.LiteralInt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(539)
			p.LiteralBoolean()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(540)
			p.LiteralFloat()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(541)
			p.LiteralNull()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(542)
			p.LiteralReference()
		}

	}

	return localctx
}

// ILiteralStringContext is an interface to support dynamic dispatch.
type ILiteralStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralStringContext differentiates from other interfaces.
	IsLiteralStringContext()
}

type LiteralStringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralStringContext() *LiteralStringContext {
	var p = new(LiteralStringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_literalString
	return p
}

func (*LiteralStringContext) IsLiteralStringContext() {}

func NewLiteralStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralStringContext {
	var p = new(LiteralStringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_literalString

	return p
}

func (s *LiteralStringContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralStringContext) STRING() antlr.TerminalNode {
	return s.GetToken(VyLangParserSTRING, 0)
}

func (s *LiteralStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralStringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterLiteralString(s)
	}
}

func (s *LiteralStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitLiteralString(s)
	}
}

func (p *VyLangParser) LiteralString() (localctx ILiteralStringContext) {
	this := p
	_ = this

	localctx = NewLiteralStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, VyLangParserRULE_literalString)

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
		p.SetState(545)
		p.Match(VyLangParserSTRING)
	}

	return localctx
}

// ILiteralIntContext is an interface to support dynamic dispatch.
type ILiteralIntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralIntContext differentiates from other interfaces.
	IsLiteralIntContext()
}

type LiteralIntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralIntContext() *LiteralIntContext {
	var p = new(LiteralIntContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_literalInt
	return p
}

func (*LiteralIntContext) IsLiteralIntContext() {}

func NewLiteralIntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralIntContext {
	var p = new(LiteralIntContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_literalInt

	return p
}

func (s *LiteralIntContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralIntContext) INT() antlr.TerminalNode {
	return s.GetToken(VyLangParserINT, 0)
}

func (s *LiteralIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralIntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralIntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterLiteralInt(s)
	}
}

func (s *LiteralIntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitLiteralInt(s)
	}
}

func (p *VyLangParser) LiteralInt() (localctx ILiteralIntContext) {
	this := p
	_ = this

	localctx = NewLiteralIntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, VyLangParserRULE_literalInt)

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
		p.SetState(547)
		p.Match(VyLangParserINT)
	}

	return localctx
}

// ILiteralBooleanContext is an interface to support dynamic dispatch.
type ILiteralBooleanContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralBooleanContext differentiates from other interfaces.
	IsLiteralBooleanContext()
}

type LiteralBooleanContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralBooleanContext() *LiteralBooleanContext {
	var p = new(LiteralBooleanContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_literalBoolean
	return p
}

func (*LiteralBooleanContext) IsLiteralBooleanContext() {}

func NewLiteralBooleanContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralBooleanContext {
	var p = new(LiteralBooleanContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_literalBoolean

	return p
}

func (s *LiteralBooleanContext) GetParser() antlr.Parser { return s.parser }
func (s *LiteralBooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralBooleanContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralBooleanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterLiteralBoolean(s)
	}
}

func (s *LiteralBooleanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitLiteralBoolean(s)
	}
}

func (p *VyLangParser) LiteralBoolean() (localctx ILiteralBooleanContext) {
	this := p
	_ = this

	localctx = NewLiteralBooleanContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, VyLangParserRULE_literalBoolean)
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
		p.SetState(549)
		_la = p.GetTokenStream().LA(1)

		if !(_la == VyLangParserT__48 || _la == VyLangParserT__49) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILiteralFloatContext is an interface to support dynamic dispatch.
type ILiteralFloatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralFloatContext differentiates from other interfaces.
	IsLiteralFloatContext()
}

type LiteralFloatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralFloatContext() *LiteralFloatContext {
	var p = new(LiteralFloatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_literalFloat
	return p
}

func (*LiteralFloatContext) IsLiteralFloatContext() {}

func NewLiteralFloatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralFloatContext {
	var p = new(LiteralFloatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_literalFloat

	return p
}

func (s *LiteralFloatContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralFloatContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(VyLangParserINT)
}

func (s *LiteralFloatContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(VyLangParserINT, i)
}

func (s *LiteralFloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralFloatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralFloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterLiteralFloat(s)
	}
}

func (s *LiteralFloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitLiteralFloat(s)
	}
}

func (p *VyLangParser) LiteralFloat() (localctx ILiteralFloatContext) {
	this := p
	_ = this

	localctx = NewLiteralFloatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, VyLangParserRULE_literalFloat)

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
		p.SetState(551)
		p.Match(VyLangParserINT)
	}
	{
		p.SetState(552)
		p.Match(VyLangParserT__0)
	}
	{
		p.SetState(553)
		p.Match(VyLangParserINT)
	}

	return localctx
}

// ILiteralNullContext is an interface to support dynamic dispatch.
type ILiteralNullContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralNullContext differentiates from other interfaces.
	IsLiteralNullContext()
}

type LiteralNullContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralNullContext() *LiteralNullContext {
	var p = new(LiteralNullContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_literalNull
	return p
}

func (*LiteralNullContext) IsLiteralNullContext() {}

func NewLiteralNullContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralNullContext {
	var p = new(LiteralNullContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_literalNull

	return p
}

func (s *LiteralNullContext) GetParser() antlr.Parser { return s.parser }
func (s *LiteralNullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralNullContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralNullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterLiteralNull(s)
	}
}

func (s *LiteralNullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitLiteralNull(s)
	}
}

func (p *VyLangParser) LiteralNull() (localctx ILiteralNullContext) {
	this := p
	_ = this

	localctx = NewLiteralNullContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, VyLangParserRULE_literalNull)

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
		p.SetState(555)
		p.Match(VyLangParserT__50)
	}

	return localctx
}

// ILiteralReferenceContext is an interface to support dynamic dispatch.
type ILiteralReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralReferenceContext differentiates from other interfaces.
	IsLiteralReferenceContext()
}

type LiteralReferenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralReferenceContext() *LiteralReferenceContext {
	var p = new(LiteralReferenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_literalReference
	return p
}

func (*LiteralReferenceContext) IsLiteralReferenceContext() {}

func NewLiteralReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralReferenceContext {
	var p = new(LiteralReferenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_literalReference

	return p
}

func (s *LiteralReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralReferenceContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VyLangParserIDENTIFIER, 0)
}

func (s *LiteralReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterLiteralReference(s)
	}
}

func (s *LiteralReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitLiteralReference(s)
	}
}

func (p *VyLangParser) LiteralReference() (localctx ILiteralReferenceContext) {
	this := p
	_ = this

	localctx = NewLiteralReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, VyLangParserRULE_literalReference)

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
		p.SetState(557)
		p.Match(VyLangParserT__11)
	}
	{
		p.SetState(558)
		p.Match(VyLangParserIDENTIFIER)
	}

	return localctx
}

// ILiteralMapContext is an interface to support dynamic dispatch.
type ILiteralMapContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralMapContext differentiates from other interfaces.
	IsLiteralMapContext()
}

type LiteralMapContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralMapContext() *LiteralMapContext {
	var p = new(LiteralMapContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_literalMap
	return p
}

func (*LiteralMapContext) IsLiteralMapContext() {}

func NewLiteralMapContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralMapContext {
	var p = new(LiteralMapContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_literalMap

	return p
}

func (s *LiteralMapContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralMapContext) AllLiteralMapEntry() []ILiteralMapEntryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILiteralMapEntryContext); ok {
			len++
		}
	}

	tst := make([]ILiteralMapEntryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILiteralMapEntryContext); ok {
			tst[i] = t.(ILiteralMapEntryContext)
			i++
		}
	}

	return tst
}

func (s *LiteralMapContext) LiteralMapEntry(i int) ILiteralMapEntryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralMapEntryContext); ok {
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

	return t.(ILiteralMapEntryContext)
}

func (s *LiteralMapContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(VyLangParserNEWLINE)
}

func (s *LiteralMapContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(VyLangParserNEWLINE, i)
}

func (s *LiteralMapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralMapContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralMapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterLiteralMap(s)
	}
}

func (s *LiteralMapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitLiteralMap(s)
	}
}

func (p *VyLangParser) LiteralMap() (localctx ILiteralMapContext) {
	this := p
	_ = this

	localctx = NewLiteralMapContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, VyLangParserRULE_literalMap)
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
		p.SetState(560)
		p.Match(VyLangParserT__5)
	}
	p.SetState(564)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == VyLangParserNEWLINE {
		{
			p.SetState(561)
			p.Match(VyLangParserNEWLINE)
		}

		p.SetState(566)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(567)
		p.LiteralMapEntry()
	}
	p.SetState(576)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 58, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(569)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == VyLangParserT__6 || _la == VyLangParserNEWLINE {
				{
					p.SetState(568)
					_la = p.GetTokenStream().LA(1)

					if !(_la == VyLangParserT__6 || _la == VyLangParserNEWLINE) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

				p.SetState(571)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(573)
				p.LiteralMapEntry()
			}

		}
		p.SetState(578)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 58, p.GetParserRuleContext())
	}
	p.SetState(582)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == VyLangParserNEWLINE {
		{
			p.SetState(579)
			p.Match(VyLangParserNEWLINE)
		}

		p.SetState(584)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(585)
		p.Match(VyLangParserT__7)
	}

	return localctx
}

// ILiteralListContext is an interface to support dynamic dispatch.
type ILiteralListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralListContext differentiates from other interfaces.
	IsLiteralListContext()
}

type LiteralListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralListContext() *LiteralListContext {
	var p = new(LiteralListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_literalList
	return p
}

func (*LiteralListContext) IsLiteralListContext() {}

func NewLiteralListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralListContext {
	var p = new(LiteralListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_literalList

	return p
}

func (s *LiteralListContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralListContext) AllLiteral() []ILiteralContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILiteralContext); ok {
			len++
		}
	}

	tst := make([]ILiteralContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILiteralContext); ok {
			tst[i] = t.(ILiteralContext)
			i++
		}
	}

	return tst
}

func (s *LiteralListContext) Literal(i int) ILiteralContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
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

	return t.(ILiteralContext)
}

func (s *LiteralListContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(VyLangParserNEWLINE)
}

func (s *LiteralListContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(VyLangParserNEWLINE, i)
}

func (s *LiteralListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterLiteralList(s)
	}
}

func (s *LiteralListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitLiteralList(s)
	}
}

func (p *VyLangParser) LiteralList() (localctx ILiteralListContext) {
	this := p
	_ = this

	localctx = NewLiteralListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, VyLangParserRULE_literalList)
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
		p.SetState(587)
		p.Match(VyLangParserT__8)
	}
	p.SetState(591)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == VyLangParserNEWLINE {
		{
			p.SetState(588)
			p.Match(VyLangParserNEWLINE)
		}

		p.SetState(593)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(594)
		p.Literal()
	}
	p.SetState(603)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 62, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(596)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == VyLangParserT__6 || _la == VyLangParserNEWLINE {
				{
					p.SetState(595)
					_la = p.GetTokenStream().LA(1)

					if !(_la == VyLangParserT__6 || _la == VyLangParserNEWLINE) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

				p.SetState(598)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(600)
				p.Literal()
			}

		}
		p.SetState(605)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 62, p.GetParserRuleContext())
	}
	p.SetState(609)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == VyLangParserNEWLINE {
		{
			p.SetState(606)
			p.Match(VyLangParserNEWLINE)
		}

		p.SetState(611)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(612)
		p.Match(VyLangParserT__9)
	}

	return localctx
}

// ILiteralMapEntryContext is an interface to support dynamic dispatch.
type ILiteralMapEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralMapEntryContext differentiates from other interfaces.
	IsLiteralMapEntryContext()
}

type LiteralMapEntryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralMapEntryContext() *LiteralMapEntryContext {
	var p = new(LiteralMapEntryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VyLangParserRULE_literalMapEntry
	return p
}

func (*LiteralMapEntryContext) IsLiteralMapEntryContext() {}

func NewLiteralMapEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralMapEntryContext {
	var p = new(LiteralMapEntryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VyLangParserRULE_literalMapEntry

	return p
}

func (s *LiteralMapEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralMapEntryContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VyLangParserIDENTIFIER, 0)
}

func (s *LiteralMapEntryContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralMapEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralMapEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralMapEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.EnterLiteralMapEntry(s)
	}
}

func (s *LiteralMapEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VyLangListener); ok {
		listenerT.ExitLiteralMapEntry(s)
	}
}

func (p *VyLangParser) LiteralMapEntry() (localctx ILiteralMapEntryContext) {
	this := p
	_ = this

	localctx = NewLiteralMapEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, VyLangParserRULE_literalMapEntry)

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
		p.SetState(614)
		p.Match(VyLangParserIDENTIFIER)
	}
	{
		p.SetState(615)
		p.Match(VyLangParserT__4)
	}
	{
		p.SetState(616)
		p.Literal()
	}

	return localctx
}

func (p *VyLangParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 3:
		var t *VariableContext = nil
		if localctx != nil {
			t = localctx.(*VariableContext)
		}
		return p.Variable_Sempred(t, predIndex)

	case 18:
		var t *BindingContext = nil
		if localctx != nil {
			t = localctx.(*BindingContext)
		}
		return p.Binding_Sempred(t, predIndex)

	case 28:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *VyLangParser) Variable_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *VyLangParser) Binding_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *VyLangParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
