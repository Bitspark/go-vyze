package state

import (
	"errors"
	"fmt"
	"github.com/Bitspark/go-vyze/state/parser"
	"github.com/Bitspark/go-vyze/util"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"os"
	"path"
	"strconv"
	"strings"
)

type Parser struct {
	lib      *Library
	listener *vylangListener
}

func NewParser(lib *Library) *Parser {
	return &Parser{
		lib: lib,
	}
}

func (c *Parser) ParseDir(dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, f := range files {
		chPath := path.Join(dir, f.Name())
		if f.IsDir() {
			if err := c.ParseDir(chPath); err != nil {
				return fmt.Errorf("loading program from %s: %v", chPath, err)
			}
		} else {
			if !strings.HasSuffix(f.Name(), ".vy") {
				continue
			}
			if err := c.ParseFile(chPath); err != nil {
				return fmt.Errorf("loading program from %s: %v", chPath, err)
			}
		}
	}

	return nil
}

func (c *Parser) ParseFile(file string) error {
	sourceBts, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	return c.ParseSource(string(sourceBts))
}

func (c *Parser) ParseSource(source string) error {
	vyInputStream := antlr.NewInputStream(source)

	vyLexer := parser.NewVyLangLexer(vyInputStream)
	vyStream := antlr.NewCommonTokenStream(vyLexer, antlr.TokenDefaultChannel)
	vyParser := parser.NewVyLangParser(vyStream)

	c.listener = &vylangListener{
		lib:     c.lib,
		parser:  vyParser,
		terms:   util.NewStack[string](),
		types:   util.NewStack[*Type](),
		exprs:   util.NewStack[*Expression](),
		actions: util.NewStack[*Action](),
	}

	antlr.ParseTreeWalkerDefault.Walk(c.listener, vyParser.Prog())

	if len(c.listener.errs) > 0 {
		errStr := ""
		for _, le := range c.listener.errs {
			errStr += le.String() + "\n"
		}
		return errors.New(errStr)
	}

	return nil
}

var _ parser.VyLangListener = &vylangListener{}

type LangError struct {
	Ctx    antlr.ParserRuleContext
	Err    error
	parser antlr.Parser
}

func (err LangError) String() string {
	return fmt.Sprintf("line %d: %s (rule %s)",
		err.Ctx.GetStart().GetLine(),
		err.Err.Error(),
		err.parser.GetRuleNames()[err.Ctx.GetRuleIndex()],
	)
}

type vylangListener struct {
	lib     *Library
	parser  antlr.Parser
	terms   *util.Stack[string]
	types   *util.Stack[*Type]
	exprs   *util.Stack[*Expression]
	actions *util.Stack[*Action]
	errs    []LangError
}

// Functions

func (v *vylangListener) PushError(c antlr.ParserRuleContext, err error) {
	v.errs = append(v.errs, LangError{
		Ctx:    c,
		Err:    err,
		parser: v.parser,
	})
}

// Visitors: Types

func (v *vylangListener) EnterNamedType(c *parser.NamedTypeContext) {
}

func (v *vylangListener) ExitNamedType(c *parser.NamedTypeContext) {
	name := v.terms.Pop()
	v.terms.Pop() // should be "type"
	t := v.types.Pop()
	t.Name = name
	err := v.lib.addType(t)
	if err != nil {
		v.PushError(c, err)
	}
}

func (v *vylangListener) EnterType(c *parser.TypeContext) {
	v.types.Push(&Type{})
}

func (v *vylangListener) ExitType(c *parser.TypeContext) {
	cld := c.GetChildCount()
	if cld == 1 {
		return
	}
	if cld == 3 {
		op := v.terms.Pop()
		expr := v.exprs.Pop()
		if op == "=" {
			v.types.Value().Initial = expr.Value
			return
		}
		if op == "in" {
			if expr.valType.ListOf == nil {
				v.PushError(c, fmt.Errorf("require a list"))
				return
			}
			v.types.Value().Options = expr.Value.([]any)
			return
		}
		return
	}
	if cld == 5 {
		_ = v.terms.Pop()
		expr := v.exprs.Pop()
		if expr.valType.ListOf == nil {
			v.PushError(c, fmt.Errorf("require a list"))
			return
		}
		v.types.Value().Options = expr.Value.([]any)

		_ = v.terms.Pop()
		expr = v.exprs.Pop()
		v.types.Value().Initial = expr.Value
		return
	}
}

func (v *vylangListener) EnterTypeMap(c *parser.TypeMapContext) {
	v.types.Value().MapOf = map[string]*Type{}
}

func (v *vylangListener) ExitTypeMap(c *parser.TypeMapContext) {
}

func (v *vylangListener) EnterTypeMapEntry(c *parser.TypeMapEntryContext) {
}

func (v *vylangListener) ExitTypeMapEntry(c *parser.TypeMapEntryContext) {
	entryName := v.terms.Pop()
	entryType := v.types.Pop()
	v.types.Value().MapOf[entryName] = entryType
}

func (v *vylangListener) EnterTypeList(c *parser.TypeListContext) {
	if v.types.Size() == 0 {
		v.types.Push(&Type{})
	}
}

func (v *vylangListener) ExitTypeList(c *parser.TypeListContext) {
	entryType := v.types.Pop()
	v.types.Value().ListOf = entryType
}

func (v *vylangListener) EnterTypeLeaf(c *parser.TypeLeafContext) {
}

func (v *vylangListener) ExitTypeLeaf(c *parser.TypeLeafContext) {
	t := v.terms.Pop()
	err := v.types.Value().Leaf.FromString(t)
	if err != nil {
		v.PushError(c, err)
	}
}

func (v *vylangListener) EnterTypeReference(c *parser.TypeReferenceContext) {
}

func (v *vylangListener) ExitTypeReference(c *parser.TypeReferenceContext) {
	refName := v.terms.Pop()
	v.types.Value().Reference = refName
}

// Visitors: Expressions

func (v *vylangListener) EnterExpr(c *parser.ExprContext) {
}

func (v *vylangListener) ExitExpr(c *parser.ExprContext) {
}

func (v *vylangListener) EnterExprOperator2(c *parser.ExprOperator2Context) {
	v.exprs.Push(&Expression{valType: &Type{}})
}

func (v *vylangListener) ExitExprOperator2(c *parser.ExprOperator2Context) {
	op := v.terms.Pop()
	expr1 := v.exprs.Pop()
	expr2 := v.exprs.Pop()

	v.exprs.Value().Operation = op
	v.exprs.Value().Children = []*Expression{expr1, expr2}
}

func (v *vylangListener) EnterLiteral(c *parser.LiteralContext) {
	v.exprs.Push(&Expression{valType: &Type{}})
}

func (v *vylangListener) ExitLiteral(c *parser.LiteralContext) {
	t := v.exprs.Pop()
	l := v.exprs.Value()
	if l != nil && l.valType.ListOf != nil {
		var lst []any
		if l.Value == nil {
			lst = []any{}
		} else {
			lst = l.Value.([]any)
		}
		lst = append(lst, t.Value)
		l.Value = lst
		return
	}
	v.exprs.Push(t)
}

func (v *vylangListener) EnterLiteralMap(c *parser.LiteralMapContext) {
	v.exprs.Value().valType.MapOf = map[string]*Type{}
}

func (v *vylangListener) ExitLiteralMap(c *parser.LiteralMapContext) {
}

func (v *vylangListener) EnterLiteralMapEntry(c *parser.LiteralMapEntryContext) {
}

func (v *vylangListener) ExitLiteralMapEntry(c *parser.LiteralMapEntryContext) {
	entryName := v.terms.Pop()
	entryExpr := v.exprs.Pop()
	l := v.exprs.Value()
	if l != nil && l.valType.MapOf != nil {
		var mp map[string]any
		if l.Value == nil {
			mp = map[string]any{}
		} else {
			mp = l.Value.(map[string]any)
		}
		mp[entryName] = entryExpr.Value
		l.Value = mp
		return
	}
}

func (v *vylangListener) EnterLiteralList(c *parser.LiteralListContext) {
	v.exprs.Push(&Expression{valType: &Type{rt{ListOf: &Type{}}}})
}

func (v *vylangListener) ExitLiteralList(c *parser.LiteralListContext) {
}

func (v *vylangListener) EnterLiteralString(c *parser.LiteralStringContext) {
	v.exprs.Value().valType = &Type{rt{Leaf: LeafString}}
}

func (v *vylangListener) ExitLiteralString(c *parser.LiteralStringContext) {
	val := v.terms.Pop()
	v.exprs.Value().Value = val[1 : len(val)-1]
}

func (v *vylangListener) EnterLiteralBoolean(c *parser.LiteralBooleanContext) {
	v.exprs.Value().valType = &Type{rt{Leaf: LeafBoolean}}
}

func (v *vylangListener) ExitLiteralBoolean(c *parser.LiteralBooleanContext) {
	val := v.terms.Pop()
	v.exprs.Value().Value = val == "true"
}

func (v *vylangListener) EnterLiteralInt(c *parser.LiteralIntContext) {
	v.exprs.Value().valType = &Type{rt{Leaf: LeafInteger}}
}

func (v *vylangListener) ExitLiteralInt(c *parser.LiteralIntContext) {
	val := v.terms.Pop()
	valInt64, _ := strconv.ParseInt(val, 10, 64)
	v.exprs.Value().Value = int(valInt64)
}

func (v *vylangListener) EnterLiteralFloat(c *parser.LiteralFloatContext) {
	v.exprs.Value().valType = &Type{rt{Leaf: LeafFloat}}
}

func (v *vylangListener) ExitLiteralFloat(c *parser.LiteralFloatContext) {
	val := v.terms.Pop()
	valFloat64, _ := strconv.ParseFloat(val, 64)
	v.exprs.Value().Value = valFloat64
}

func (v *vylangListener) EnterVariable(c *parser.VariableContext) {
	v.exprs.Push(&Expression{valType: &Type{}})
}

func (v *vylangListener) ExitVariable(c *parser.VariableContext) {
	varName := v.terms.Pop()
	v.exprs.Value().Variable = varName
}

// Visitors: Actions

func (v *vylangListener) EnterNamedAction(c *parser.NamedActionContext) {
}

func (v *vylangListener) ExitNamedAction(c *parser.NamedActionContext) {
	actionType := v.types.Pop()
	v.terms.Pop() // should be "in"
	actionName := v.terms.Pop()
	v.terms.Pop() // should be "action"
	a := v.actions.Pop()
	a.Type = actionType
	a.Name = actionName
	err := v.lib.addAction(a)
	if err != nil {
		v.PushError(c, err)
	}
}

func (v *vylangListener) EnterAction(c *parser.ActionContext) {
	v.actions.Push(&Action{})
}

func (v *vylangListener) ExitAction(c *parser.ActionContext) {
	t := v.actions.Pop()
	l := v.actions.Value()
	if l != nil && l.Sequence != nil {
		l.Sequence = append(l.Sequence, l)
		return
	}
	if l != nil && l.Parallel != nil {
		l.Parallel = append(l.Parallel, l)
		return
	}
	v.actions.Push(t)
}

func (v *vylangListener) EnterActionSequence(c *parser.ActionSequenceContext) {
	v.actions.Value().Sequence = []*Action{}
}

func (v *vylangListener) ExitActionSequence(c *parser.ActionSequenceContext) {
}

func (v *vylangListener) EnterActionParallel(c *parser.ActionParallelContext) {
	v.actions.Value().Parallel = []*Action{}
}

func (v *vylangListener) ExitActionParallel(c *parser.ActionParallelContext) {
}

func (v *vylangListener) EnterActionIf(c *parser.ActionIfContext) {
}

func (v *vylangListener) ExitActionIf(c *parser.ActionIfContext) {
	v.terms.Pop() // should be "if"

	cond := v.exprs.Pop()
	then := v.actions.Pop()

	v.actions.Value().If = &ActionIf{
		Expression: cond,
		Then:       then,
	}
}

func (v *vylangListener) EnterActionIfElse(c *parser.ActionIfElseContext) {
}

func (v *vylangListener) ExitActionIfElse(c *parser.ActionIfElseContext) {
	v.terms.Pop() // should be "else"
	v.terms.Pop() // should be "if"

	cond := v.exprs.Pop()
	ifElse := v.actions.Pop()
	ifThen := v.actions.Pop()

	v.actions.Value().If = &ActionIf{
		Expression: cond,
		Then:       ifThen,
		Else:       ifElse,
	}
}

func (v *vylangListener) EnterActionAssign(c *parser.ActionAssignContext) {
}

func (v *vylangListener) ExitActionAssign(c *parser.ActionAssignContext) {
	v.terms.Pop() // should be "="
	action := v.actions.Value()
	expression := v.exprs.Pop()
	target := v.exprs.Pop()
	action.Target = target.Variable
	action.Expression = expression
}

// Visitors: Nodes

func (v *vylangListener) VisitTerminal(node antlr.TerminalNode) {
	t := node.GetText()
	if strings.TrimSpace(t) == "" {
		return
	}
	if len(t) == 1 && strings.ContainsAny(t, `:$",()[]{}`) {
		return
	}
	v.terms.Push(t)
}

func (v *vylangListener) VisitErrorNode(node antlr.ErrorNode) {
}

// Stubs

func (v *vylangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
}

func (v *vylangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
}

func (v *vylangListener) EnterProg(c *parser.ProgContext) {
}

func (v *vylangListener) EnterDefinitions(c *parser.DefinitionsContext) {
}

func (v *vylangListener) EnterDefinition(c *parser.DefinitionContext) {
}

func (v *vylangListener) EnterActionAsync(c *parser.ActionAsyncContext) {
}

func (v *vylangListener) EnterActionReference(c *parser.ActionReferenceContext) {
}

func (v *vylangListener) EnterActionBind(c *parser.ActionBindContext) {
}

func (v *vylangListener) EnterActionCond(c *parser.ActionCondContext) {
}

func (v *vylangListener) EnterActionLeaf(c *parser.ActionLeafContext) {
}

func (v *vylangListener) EnterActionClear(c *parser.ActionClearContext) {
}

func (v *vylangListener) EnterBinding(c *parser.BindingContext) {
}

func (v *vylangListener) EnterRawType(c *parser.RawTypeContext) {
}

func (v *vylangListener) EnterNamedExpr(c *parser.NamedExprContext) {
}

func (v *vylangListener) EnterExprReference(c *parser.ExprReferenceContext) {
}

func (v *vylangListener) EnterExprMap(c *parser.ExprMapContext) {
}

func (v *vylangListener) EnterExprList(c *parser.ExprListContext) {
}

func (v *vylangListener) EnterExprMapEntry(c *parser.ExprMapEntryContext) {
}

func (v *vylangListener) EnterNamedLiteral(c *parser.NamedLiteralContext) {
}

func (v *vylangListener) EnterLiteralTerminal(c *parser.LiteralTerminalContext) {
}

func (v *vylangListener) EnterLiteralNull(c *parser.LiteralNullContext) {
}

func (v *vylangListener) EnterLiteralReference(c *parser.LiteralReferenceContext) {
}

func (v *vylangListener) ExitProg(c *parser.ProgContext) {
}

func (v *vylangListener) ExitDefinitions(c *parser.DefinitionsContext) {
}

func (v *vylangListener) ExitDefinition(c *parser.DefinitionContext) {
}

func (v *vylangListener) ExitActionAsync(c *parser.ActionAsyncContext) {
}

func (v *vylangListener) ExitActionReference(c *parser.ActionReferenceContext) {
}

func (v *vylangListener) ExitActionBind(c *parser.ActionBindContext) {
}

func (v *vylangListener) ExitActionCond(c *parser.ActionCondContext) {
}

func (v *vylangListener) ExitActionLeaf(c *parser.ActionLeafContext) {
}

func (v *vylangListener) ExitActionClear(c *parser.ActionClearContext) {
}

func (v *vylangListener) ExitBinding(c *parser.BindingContext) {
}

func (v *vylangListener) ExitRawType(c *parser.RawTypeContext) {
}

func (v *vylangListener) ExitNamedExpr(c *parser.NamedExprContext) {
}

func (v *vylangListener) ExitExprReference(c *parser.ExprReferenceContext) {
}

func (v *vylangListener) ExitExprMap(c *parser.ExprMapContext) {
}

func (v *vylangListener) ExitExprList(c *parser.ExprListContext) {
}

func (v *vylangListener) ExitExprMapEntry(c *parser.ExprMapEntryContext) {
}

func (v *vylangListener) ExitNamedLiteral(c *parser.NamedLiteralContext) {
}

func (v *vylangListener) ExitLiteralTerminal(c *parser.LiteralTerminalContext) {
}

func (v *vylangListener) ExitLiteralNull(c *parser.LiteralNullContext) {
}

func (v *vylangListener) ExitLiteralReference(c *parser.LiteralReferenceContext) {
}

func (v *vylangListener) EnterActionWhile(c *parser.ActionWhileContext) {
}

func (v *vylangListener) ExitActionWhile(c *parser.ActionWhileContext) {
}

func (v *vylangListener) EnterExprAlternative(c *parser.ExprAlternativeContext) {
}

func (v *vylangListener) EnterExprOperator1(c *parser.ExprOperator1Context) {
}

func (v *vylangListener) ExitExprAlternative(c *parser.ExprAlternativeContext) {
}

func (v *vylangListener) ExitExprOperator1(c *parser.ExprOperator1Context) {
}

func (v *vylangListener) EnterExprBrackets(c *parser.ExprBracketsContext) {
}

func (v *vylangListener) ExitExprBrackets(c *parser.ExprBracketsContext) {
}
