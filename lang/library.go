package lang

import (
	"fmt"
	"github.com/Bitspark/go-vyze/lang/parser"
	"github.com/Bitspark/go-vyze/system"
	"github.com/Bitspark/go-vyze/util"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"strings"
)

type Library struct {
	listener listener

	Pipes     map[string]*Pipe
	Universes map[string]*system.Universe

	terms      *util.Stack[string]
	identPaths *util.Stack[string]
	identPath  string
	pipes      *util.Stack[*Pipe]
	univ       *system.Universe
	entryName  string
	entryNames *util.Stack[string]
	models     *util.Stack[*system.UniverseObjectInfo]
	relations  *util.Stack[*system.UniverseObjectInfo]
	context    string
	modifier   *bool
	modifiers  *util.Stack[bool]
	errors     []ParseError
}

type ParseError struct {
	Err    string `json:"error"`
	Path   string `json:"path"`
	Line   int    `json:"line"`
	Column int    `json:"column"`
}

func (pe ParseError) String() string {
	if pe.Line > 0 || pe.Column > 0 {
		if len(pe.Path) > 0 {
			return fmt.Sprintf("[l%d:%d, %s] %v", pe.Line, pe.Column, pe.Path, pe.Err)
		} else {
			return fmt.Sprintf("[l%d:%d] %v", pe.Line, pe.Column, pe.Err)
		}
	} else {
		if len(pe.Path) > 0 {
			return fmt.Sprintf("[%s] %v", pe.Path, pe.Err)
		} else {
			return fmt.Sprintf("%v", pe.Err)
		}
	}
}

func NewLibrary(univ *system.Universe) *Library {
	l := &Library{
		Pipes:      map[string]*Pipe{},
		terms:      util.NewStack[string](),
		pipes:      util.NewStack[*Pipe](),
		entryNames: util.NewStack[string](),
		models:     util.NewStack[*system.UniverseObjectInfo](),
		relations:  util.NewStack[*system.UniverseObjectInfo](),
		modifiers:  util.NewStack[bool](),
		univ:       univ,
	}
	l.listener = &vylangListener{Library: l}
	return l
}

func (l *Library) Parse(source string) []ParseError {
	if l.univ == nil {
		return []ParseError{{Err: fmt.Sprintf("require universe")}}
	}

	is := antlr.NewInputStream(source)
	lexer := parser.NewVyLangLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	vyParser := parser.NewVyLangParser(stream)
	vyParser.AddErrorListener(l.listener)

	antlr.ParseTreeWalkerDefault.Walk(l.listener, vyParser.Definitions())

	if l.terms.Size() != 0 {
		return []ParseError{{Err: fmt.Sprintf("have extra terms: %s", l.terms.String())}}
	}

	return l.errors
}

func (l *Library) ParsePipe(source string) (*Pipe, []ParseError) {
	if l.univ == nil {
		return nil, []ParseError{{Err: fmt.Sprintf("require universe")}}
	}

	is := antlr.NewInputStream(source)
	lexer := parser.NewVyLangLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	vyParser := parser.NewVyLangParser(stream)
	vyParser.AddErrorListener(l.listener)

	antlr.ParseTreeWalkerDefault.Walk(l.listener, vyParser.ContextPipe())

	if l.terms.Size() != 0 {
		return nil, []ParseError{{Err: fmt.Sprintf("have extra terms: %s", l.terms.String())}}
	}

	return l.pipes.Pop(), l.errors
}

type listener interface {
	parser.VyLangListener
	antlr.ErrorListener
}

var _ listener = vylangListener{}

type vylangListener struct {
	*Library
}

const (
	ErrTypeUniverse = 1 << iota
	ErrTypeAbort
	ErrTypeWarning
)

func (v vylangListener) addError(ctx antlr.ParserRuleContext, err error, errType int) {
	if errType == ErrTypeWarning {
		return
	}
	v.errors = append(v.errors, ParseError{
		Err:    err.Error(),
		Path:   v.entryNames.Join("."),
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
	})
}

func (v vylangListener) VisitTerminal(node antlr.TerminalNode) {
	t := node.GetText()
	if strings.TrimSpace(t) == "" {
		return
	}
	if v.identPaths != nil {
		v.identPaths.Push(t)
	} else {
		if len(t) == 1 && strings.ContainsAny(t, `:$",()[]{}`) {
			return
		}
		v.terms.Push(t)
	}
}

func (v vylangListener) VisitErrorNode(node antlr.ErrorNode) {
}

func (v vylangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
}

func (v vylangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
}

func (v vylangListener) EnterDefinitions(c *parser.DefinitionsContext) {
}

func (v vylangListener) EnterDefinition(c *parser.DefinitionContext) {
}

func (v vylangListener) EnterNamedPipe(c *parser.NamedPipeContext) {
}

func (v vylangListener) EnterPipe(c *parser.PipeContext) {
	if v.modifier != nil {
		v.modifiers.Push(*v.modifier)
		v.modifier = nil
	}

	v.entryNames.Push(v.identPath)

	if v.context == "model" {
		v.terms.Pop() // Should be "->"
		model := v.univ.GetModel(v.identPath, v.univ.Name)
		if model == nil {
			v.addError(c, fmt.Errorf("model '%s' not found", v.identPath), ErrTypeUniverse)
		}
		v.models.Push(model)
		v.identPaths = nil
	} else if v.context == "forward" {
		v.terms.Pop() // Should be "->"
		if v.models.Value() == nil {
			v.addError(c, fmt.Errorf("invalid model"), ErrTypeWarning)
			v.relations.Push(nil)
			v.models.Push(nil)
		} else {
			rel := v.univ.GetModel(fmt.Sprintf("%s#%s/", v.models.Value().Mapping.String(), v.identPath), v.univ.Name)
			var model *system.UniverseObjectInfo
			if rel == nil {
				v.addError(c, fmt.Errorf("relation '%s' not found", v.identPath), ErrTypeUniverse)
			} else {
				model = v.univ.GetModel(v.univ.GetTarget(rel.Mapping).String(), v.univ.Name)
				if model == nil {
					v.addError(c, fmt.Errorf("model '%s' not found", v.univ.GetOrigin(rel.Mapping).String()), ErrTypeUniverse)
				}
			}
			v.relations.Push(rel)
			v.models.Push(model)
		}
		v.identPaths = nil
	} else if v.context == "backward" {
		v.terms.Pop() // Should be "<-"
		if v.models.Value() == nil {
			v.addError(c, fmt.Errorf("invalid model"), ErrTypeWarning)
			v.relations.Push(nil)
			v.models.Push(nil)
		} else {
			rel := v.univ.GetModel(fmt.Sprintf("%s/", v.identPath), v.univ.Name)
			var model *system.UniverseObjectInfo
			if rel == nil {
				v.addError(c, fmt.Errorf("relation '%s' not found", v.identPath), ErrTypeUniverse)
			} else {
				model = v.univ.GetModel(v.univ.GetOrigin(rel.Mapping).String(), v.univ.Name)
				if model == nil {
					v.addError(c, fmt.Errorf("model '%s' not found", v.univ.GetOrigin(rel.Mapping).String()), ErrTypeUniverse)
				}
			}
			v.relations.Push(rel)
			v.models.Push(model)
		}
		v.identPaths = nil
	}
	v.context = ""

	//v.entryNames.Push(v.entryName)
	v.entryName = ""
}

func (v vylangListener) EnterPipeTerminal(c *parser.PipeTerminalContext) {
}

func (v vylangListener) EnterPipeField(c *parser.PipeFieldContext) {
}

func (v vylangListener) EnterPipeMap(c *parser.PipeMapContext) {
	p := &Pipe{}
	v.pipes.Push(p)
	p.Node = &system.Node{
		Type: system.NodeTypeMap,
		Map: &system.MapNode{
			Entries: []system.MapNodeEntry{},
		},
	}
}

func (v vylangListener) EnterPipeMapEntry(c *parser.PipeMapEntryContext) {
}

func (v vylangListener) EnterPipeModifier(c *parser.PipeModifierContext) {
}

func (v vylangListener) EnterSep(c *parser.SepContext) {
}

func (v vylangListener) ExitDefinitions(c *parser.DefinitionsContext) {
	v.terms.Pop() // Should be "<EOF>"
}

func (v vylangListener) ExitDefinition(c *parser.DefinitionContext) {
}

func (v vylangListener) ExitNamedPipe(c *parser.NamedPipeContext) {
	name := v.terms.Pop()
	v.terms.Pop() // Should be "pipe"
	v.Library.Pipes[name] = v.pipes.Pop()
}

func (v vylangListener) ExitPipe(c *parser.PipeContext) {
	v.entryName = v.entryNames.Pop()
}

func (v vylangListener) ExitPipeTerminal(c *parser.PipeTerminalContext) {
	t := v.terms.Pop()[1:]
	p := v.pipes.Value()
	v.entryName = t

	switch t {
	case string(system.FieldTypeID):
		p.Node = &system.Node{
			Type:  system.NodeTypeValue,
			Value: &system.ValueNode{Field: system.FieldTypeID, Format: system.FormatTypeHex},
		}
	case string(system.FieldTypeName):
		p.Node = &system.Node{
			Type:  system.NodeTypeValue,
			Value: &system.ValueNode{Field: system.FieldTypeName, Format: system.FormatTypeString},
		}
	case string(system.FieldTypeSize):
		p.Node = &system.Node{
			Type:  system.NodeTypeValue,
			Value: &system.ValueNode{Field: system.FieldTypeSize, Format: system.FormatTypeInteger},
		}
	case string(system.FieldTypeCreated):
		p.Node = &system.Node{
			Type:  system.NodeTypeValue,
			Value: &system.ValueNode{Field: system.FieldTypeCreated, Format: system.FormatTypeInteger},
		}
	case string(system.FieldTypeUser):
		p.Node = &system.Node{
			Type:  system.NodeTypeValue,
			Value: &system.ValueNode{Field: system.FieldTypeUser, Format: system.FormatTypeHex},
		}
	case string(system.FormatTypeBase64):
		p.Node = &system.Node{
			Type:  system.NodeTypeValue,
			Value: &system.ValueNode{Field: system.FieldTypeData, Format: system.FormatTypeBase64},
		}
	case string(system.FormatTypeHex):
		p.Node = &system.Node{
			Type:  system.NodeTypeValue,
			Value: &system.ValueNode{Field: system.FieldTypeData, Format: system.FormatTypeHex},
		}
	case string(system.FormatTypeString):
		p.Node = &system.Node{
			Type:  system.NodeTypeValue,
			Value: &system.ValueNode{Field: system.FieldTypeData, Format: system.FormatTypeString},
		}
	case string(system.FormatTypeInteger):
		p.Node = &system.Node{
			Type:  system.NodeTypeValue,
			Value: &system.ValueNode{Field: system.FieldTypeData, Format: system.FormatTypeInteger},
		}
	case string(system.FormatTypeFloat):
		p.Node = &system.Node{
			Type:  system.NodeTypeValue,
			Value: &system.ValueNode{Field: system.FieldTypeData, Format: system.FormatTypeFloat},
		}
	case string(system.FormatTypeBoolean):
		p.Node = &system.Node{
			Type:  system.NodeTypeValue,
			Value: &system.ValueNode{Field: system.FieldTypeData, Format: system.FormatTypeBoolean},
		}
	case "auto", "value": // TODO: Eventually remove value
		field := system.FieldTypeData
		format := system.FormatTypeBase64
		if v.univ.HasAbstract(v.models.Value().Mapping, system.ParseUniverseObjectIdentifier("data.@string")) {
			format = system.FormatTypeString
		} else if v.univ.HasAbstract(v.models.Value().Mapping, system.ParseUniverseObjectIdentifier("data.@integer")) {
			format = system.FormatTypeInteger
		} else if v.univ.HasAbstract(v.models.Value().Mapping, system.ParseUniverseObjectIdentifier("data.@float")) {
			format = system.FormatTypeFloat
		} else if v.univ.HasAbstract(v.models.Value().Mapping, system.ParseUniverseObjectIdentifier("data.@boolean")) {
			format = system.FormatTypeBoolean
		} else if v.univ.HasAbstract(v.models.Value().Mapping, system.ParseUniverseObjectIdentifier("data.@data")) {
			format = system.FormatTypeBase64
		} else {
			field = system.FieldTypeID
			format = system.FormatTypeHex
		}
		p.Node = &system.Node{
			Type:  system.NodeTypeValue,
			Value: &system.ValueNode{Field: field, Format: format},
		}
	}
}

func (v vylangListener) ExitPipeField(c *parser.PipeFieldContext) {
}

func (v vylangListener) ExitPipeMap(c *parser.PipeMapContext) {
}

func (v vylangListener) ExitPipeMapEntry(c *parser.PipeMapEntryContext) {
	pe := v.pipes.Pop()
	pm := v.pipes.Value()

	if pm == nil {
		v.addError(c, fmt.Errorf("invalid pipe"), ErrTypeAbort)
		return
	}

	if pm.Node == nil || pm.Node.Map == nil {
		v.addError(c, fmt.Errorf("invalid map node"), ErrTypeAbort)
		return
	}

	if pe.Node == nil {
		v.addError(c, fmt.Errorf("invalid entry node"), ErrTypeAbort)
		return
	}

	pm.Node.Map.Entries = append(pm.Node.Map.Entries, system.MapNodeEntry{
		Name: v.entryName,
		Node: *pe.Node,
	})
}

func (v vylangListener) ExitPipeModifier(c *parser.PipeModifierContext) {
	t := true
	v.modifier = &t
}

func (v vylangListener) ExitSep(c *parser.SepContext) {
}

func (v vylangListener) EnterContextPipe(c *parser.ContextPipeContext) {
	v.context = "model"
}

func (v vylangListener) EnterIdentPath(c *parser.IdentPathContext) {
}

func (v vylangListener) ExitContextPipe(c *parser.ContextPipeContext) {
	v.terms.Pop() // Should be "on"
}

func (v vylangListener) ExitIdentPath(c *parser.IdentPathContext) {
}

func (v vylangListener) EnterPipeNamedProperty(c *parser.PipeNamedPropertyContext) {
}

func (v vylangListener) EnterPipeProperty(c *parser.PipePropertyContext) {
	p := &Pipe{}
	v.pipes.Push(p)
}

func (v vylangListener) ExitPipeNamedProperty(c *parser.PipeNamedPropertyContext) {
	v.entryName = v.terms.Pop()
}

func (v vylangListener) ExitPipeProperty(c *parser.PipePropertyContext) {
}

func (v vylangListener) EnterPipeFieldForward(c *parser.PipeFieldForwardContext) {
	v.context = "forward"
}

func (v vylangListener) EnterPipeFieldBackward(c *parser.PipeFieldBackwardContext) {
	v.context = "backward"
}

func (v vylangListener) ExitPipeFieldForward(c *parser.PipeFieldForwardContext) {
	rel := v.relations.Pop()
	relIdent := ""
	if rel == nil {
		v.addError(c, fmt.Errorf("invalid relation"), ErrTypeWarning)
	} else {
		relIdent = rel.Mapping.String()
	}
	pn := v.pipes.Pop()
	if pn == nil || pn.Node == nil {
		v.addError(c, fmt.Errorf("invalid node"), ErrTypeAbort)
		return
	}
	pf := v.pipes.Value()
	envType := system.EnvironmentTypePrimitive
	nd := *pn.Node
	mdf := v.modifiers.Pop()
	if mdf {
		envType = system.EnvironmentTypeList
		nd = system.Node{
			Type: system.NodeTypeList,
			List: &system.ListNode{
				Entry: *pn.Node,
			},
		}
	}
	pf.Node = &system.Node{
		Type: system.NodeTypeRelation,
		Relation: &system.RelationNode{
			Type:     envType,
			Relation: relIdent,
			Node:     nd,
		},
	}
}

func (v vylangListener) ExitPipeFieldBackward(c *parser.PipeFieldBackwardContext) {
	rel := v.relations.Pop()
	relIdent := ""
	if rel == nil {
		v.addError(c, fmt.Errorf("invalid relation"), ErrTypeWarning)
	} else {
		relIdent = rel.Mapping.String()
	}
	pn := v.pipes.Pop()
	if pn == nil || pn.Node == nil {
		v.addError(c, fmt.Errorf("invalid node"), ErrTypeAbort)
		return
	}
	pf := v.pipes.Value()
	envType := system.EnvironmentTypePrimitive
	nd := *pn.Node
	mdf := v.modifiers.Pop()
	if mdf {
		envType = system.EnvironmentTypeList
		nd = system.Node{
			Type: system.NodeTypeList,
			List: &system.ListNode{
				Entry: *pn.Node,
			},
		}
	}
	pf.Node = &system.Node{
		Type: system.NodeTypeRelation,
		Relation: &system.RelationNode{
			Type:     envType,
			Relation: relIdent,
			Node:     nd,
			Reverse:  true,
		},
	}
}

func (v vylangListener) EnterPipeModified(c *parser.PipeModifiedContext) {
	f := false
	v.modifier = &f
}

func (v vylangListener) ExitPipeModified(c *parser.PipeModifiedContext) {
	model := v.models.Pop()
	if model == nil {
		v.addError(c, fmt.Errorf("invalid model"), ErrTypeWarning)
	} else {
		pipe := v.pipes.Value()
		if pipe == nil {
			v.addError(c, fmt.Errorf("invalid pipe"), ErrTypeWarning)
		} else {
			pipe.Model = *model
		}
	}
}

func (v vylangListener) EnterPathModel(c *parser.PathModelContext) {
	v.identPaths = util.NewStack[string]()
}

func (v vylangListener) EnterPathRelation(c *parser.PathRelationContext) {
	v.identPaths = util.NewStack[string]()
}

func (v vylangListener) ExitPathModel(c *parser.PathModelContext) {
	v.identPath = strings.Join(v.identPaths.Empty(), "")
	v.identPaths = nil
}

func (v vylangListener) ExitPathRelation(c *parser.PathRelationContext) {
	v.identPath = strings.Join(v.identPaths.Empty(), "")
	v.identPaths = nil
}

func (v vylangListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	v.errors = append(v.errors, ParseError{
		Err:    msg,
		Line:   line,
		Column: column,
	})
}

func (v vylangListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

func (v vylangListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

func (v vylangListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
}
