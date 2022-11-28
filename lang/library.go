package lang

import (
	"errors"
	"fmt"
	"github.com/Bitspark/go-vyze/lang/parser"
	"github.com/Bitspark/go-vyze/system"
	"github.com/Bitspark/go-vyze/util"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"log"
	"strings"
)

type Library struct {
	listener parser.VyLangListener

	namedPipes map[string]*Pipe
	universes  map[string]*system.Universe

	terms      *util.Stack[string]
	identPaths *util.Stack[string]
	identPath  string
	pipes      *util.Stack[*Pipe]
	univ       *system.Universe
	entryName  string
	models     *util.Stack[*system.UniverseObjectInfo]
	relations  *util.Stack[*system.UniverseObjectInfo]
	context    string
	modifier   bool
}

func NewLibrary(univ *system.Universe) *Library {
	l := &Library{
		namedPipes: map[string]*Pipe{},
		terms:      util.NewStack[string](),
		pipes:      util.NewStack[*Pipe](),
		models:     util.NewStack[*system.UniverseObjectInfo](),
		relations:  util.NewStack[*system.UniverseObjectInfo](),
		univ:       univ,
	}
	l.listener = &vylangListener{Library: l}
	return l
}

func (l *Library) Parse(source string) error {
	is := antlr.NewInputStream(source)
	lexer := parser.NewVyLangLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	vyParser := parser.NewVyLangParser(stream)

	antlr.ParseTreeWalkerDefault.Walk(l.listener, vyParser.Definitions())

	if l.terms.Size() != 0 {
		return fmt.Errorf("have extra terms: %s", l.terms.String())
	}

	return nil
}

func (l *Library) ParsePipe(source string) (*Pipe, error) {
	is := antlr.NewInputStream(source)
	lexer := parser.NewVyLangLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	vyParser := parser.NewVyLangParser(stream)

	antlr.ParseTreeWalkerDefault.Walk(l.listener, vyParser.ContextPipe())

	if l.terms.Size() != 0 {
		return nil, fmt.Errorf("have extra terms: %s", l.terms.String())
	}
	if l.pipes.Size() != 1 {
		return nil, errors.New("invalid pipe")
	}

	return l.pipes.Pop(), nil
}

var _ parser.VyLangListener = &vylangListener{}

type vylangListener struct {
	*Library
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
	if v.context == "model" {
		v.terms.Pop() // Should be "->"
		model := v.univ.GetModel(v.identPath, v.univ.Name)
		if model == nil {
			log.Printf("model not found: %s", v.identPath)
		}
		v.models.Push(model)
		v.identPaths = nil
	} else if v.context == "forward" {
		v.terms.Pop() // Should be "->"
		rel := v.univ.GetModel(fmt.Sprintf("%s#%s/", v.models.Value().Mapping.String(), v.identPath), v.univ.Name)
		if rel == nil {
			log.Printf("relation not found: %s", v.identPath)
		}
		v.relations.Push(rel)
		model := v.univ.GetModel(v.univ.GetTarget(rel.Mapping).String(), v.univ.Name)
		if model == nil {
			log.Printf("model not found: %s", v.univ.GetOrigin(rel.Mapping).String())
		}
		v.models.Push(model)
		v.identPaths = nil
	} else if v.context == "backward" {
		v.terms.Pop() // Should be "<-"
		rel := v.univ.GetModel(fmt.Sprintf("%s/", v.identPath), v.univ.Name)
		if rel == nil {
			log.Printf("relation not found: %s", v.identPath)
		}
		v.relations.Push(rel)
		model := v.univ.GetModel(v.univ.GetOrigin(rel.Mapping).String(), v.univ.Name)
		if model == nil {
			log.Printf("model not found: %s", v.univ.GetOrigin(rel.Mapping).String())
		}
		v.models.Push(model)
		v.identPaths = nil
	}
	v.context = ""
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
}

func (v vylangListener) ExitDefinition(c *parser.DefinitionContext) {
}

func (v vylangListener) ExitNamedPipe(c *parser.NamedPipeContext) {
}

func (v vylangListener) ExitPipe(c *parser.PipeContext) {
}

func (v vylangListener) ExitPipeTerminal(c *parser.PipeTerminalContext) {
	t := v.terms.Pop()
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
	case string(system.FieldTypeData):
		p.Node = &system.Node{
			Type:  system.NodeTypeValue,
			Value: &system.ValueNode{Field: system.FieldTypeData, Format: system.FormatTypeBase64},
		}
	case "value":
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

	pm.Node.Map.Entries = append(pm.Node.Map.Entries, system.MapNodeEntry{
		Name: v.entryName,
		Node: *pe.Node,
	})
}

func (v vylangListener) ExitPipeModifier(c *parser.PipeModifierContext) {
	v.modifier = true
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
	v.entryName = v.identPath
	pn := v.pipes.Pop()
	pf := v.pipes.Value()
	envType := system.EnvironmentTypePrimitive
	if v.modifier {
		envType = system.EnvironmentTypeList
	}
	pf.Node = &system.Node{
		Type: system.NodeTypeRelation,
		Relation: &system.RelationNode{
			Type:     envType,
			Relation: rel.Mapping.String(),
			Node:     *pn.Node,
		},
	}
	v.modifier = false
}

func (v vylangListener) ExitPipeFieldBackward(c *parser.PipeFieldBackwardContext) {
	rel := v.relations.Pop()
	v.identPaths = nil
	v.entryName = v.identPath
	pn := v.pipes.Pop()
	pf := v.pipes.Value()
	envType := system.EnvironmentTypePrimitive
	if v.modifier {
		envType = system.EnvironmentTypeList
	}
	pf.Node = &system.Node{
		Type: system.NodeTypeRelation,
		Relation: &system.RelationNode{
			Type:     envType,
			Relation: rel.Mapping.String(),
			Node:     *pn.Node,
			Reverse:  true,
		},
	}
	v.modifier = false
}

func (v vylangListener) EnterPipeModified(c *parser.PipeModifiedContext) {
}

func (v vylangListener) ExitPipeModified(c *parser.PipeModifiedContext) {
	v.models.Pop()
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
