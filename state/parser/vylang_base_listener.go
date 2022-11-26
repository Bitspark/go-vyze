// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // VyLang
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseVyLangListener is a complete listener for a parse tree produced by VyLangParser.
type BaseVyLangListener struct{}

var _ VyLangListener = &BaseVyLangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVyLangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVyLangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVyLangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVyLangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseVyLangListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseVyLangListener) ExitProg(ctx *ProgContext) {}

// EnterDefinitions is called when production definitions is entered.
func (s *BaseVyLangListener) EnterDefinitions(ctx *DefinitionsContext) {}

// ExitDefinitions is called when production definitions is exited.
func (s *BaseVyLangListener) ExitDefinitions(ctx *DefinitionsContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BaseVyLangListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BaseVyLangListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseVyLangListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseVyLangListener) ExitVariable(ctx *VariableContext) {}

// EnterNamedAction is called when production namedAction is entered.
func (s *BaseVyLangListener) EnterNamedAction(ctx *NamedActionContext) {}

// ExitNamedAction is called when production namedAction is exited.
func (s *BaseVyLangListener) ExitNamedAction(ctx *NamedActionContext) {}

// EnterAction is called when production action is entered.
func (s *BaseVyLangListener) EnterAction(ctx *ActionContext) {}

// ExitAction is called when production action is exited.
func (s *BaseVyLangListener) ExitAction(ctx *ActionContext) {}

// EnterActionParallel is called when production actionParallel is entered.
func (s *BaseVyLangListener) EnterActionParallel(ctx *ActionParallelContext) {}

// ExitActionParallel is called when production actionParallel is exited.
func (s *BaseVyLangListener) ExitActionParallel(ctx *ActionParallelContext) {}

// EnterActionSequence is called when production actionSequence is entered.
func (s *BaseVyLangListener) EnterActionSequence(ctx *ActionSequenceContext) {}

// ExitActionSequence is called when production actionSequence is exited.
func (s *BaseVyLangListener) ExitActionSequence(ctx *ActionSequenceContext) {}

// EnterActionAsync is called when production actionAsync is entered.
func (s *BaseVyLangListener) EnterActionAsync(ctx *ActionAsyncContext) {}

// ExitActionAsync is called when production actionAsync is exited.
func (s *BaseVyLangListener) ExitActionAsync(ctx *ActionAsyncContext) {}

// EnterActionReference is called when production actionReference is entered.
func (s *BaseVyLangListener) EnterActionReference(ctx *ActionReferenceContext) {}

// ExitActionReference is called when production actionReference is exited.
func (s *BaseVyLangListener) ExitActionReference(ctx *ActionReferenceContext) {}

// EnterActionBind is called when production actionBind is entered.
func (s *BaseVyLangListener) EnterActionBind(ctx *ActionBindContext) {}

// ExitActionBind is called when production actionBind is exited.
func (s *BaseVyLangListener) ExitActionBind(ctx *ActionBindContext) {}

// EnterActionWhile is called when production actionWhile is entered.
func (s *BaseVyLangListener) EnterActionWhile(ctx *ActionWhileContext) {}

// ExitActionWhile is called when production actionWhile is exited.
func (s *BaseVyLangListener) ExitActionWhile(ctx *ActionWhileContext) {}

// EnterActionCond is called when production actionCond is entered.
func (s *BaseVyLangListener) EnterActionCond(ctx *ActionCondContext) {}

// ExitActionCond is called when production actionCond is exited.
func (s *BaseVyLangListener) ExitActionCond(ctx *ActionCondContext) {}

// EnterActionIf is called when production actionIf is entered.
func (s *BaseVyLangListener) EnterActionIf(ctx *ActionIfContext) {}

// ExitActionIf is called when production actionIf is exited.
func (s *BaseVyLangListener) ExitActionIf(ctx *ActionIfContext) {}

// EnterActionIfElse is called when production actionIfElse is entered.
func (s *BaseVyLangListener) EnterActionIfElse(ctx *ActionIfElseContext) {}

// ExitActionIfElse is called when production actionIfElse is exited.
func (s *BaseVyLangListener) ExitActionIfElse(ctx *ActionIfElseContext) {}

// EnterActionLeaf is called when production actionLeaf is entered.
func (s *BaseVyLangListener) EnterActionLeaf(ctx *ActionLeafContext) {}

// ExitActionLeaf is called when production actionLeaf is exited.
func (s *BaseVyLangListener) ExitActionLeaf(ctx *ActionLeafContext) {}

// EnterActionAssign is called when production actionAssign is entered.
func (s *BaseVyLangListener) EnterActionAssign(ctx *ActionAssignContext) {}

// ExitActionAssign is called when production actionAssign is exited.
func (s *BaseVyLangListener) ExitActionAssign(ctx *ActionAssignContext) {}

// EnterActionClear is called when production actionClear is entered.
func (s *BaseVyLangListener) EnterActionClear(ctx *ActionClearContext) {}

// ExitActionClear is called when production actionClear is exited.
func (s *BaseVyLangListener) ExitActionClear(ctx *ActionClearContext) {}

// EnterBinding is called when production binding is entered.
func (s *BaseVyLangListener) EnterBinding(ctx *BindingContext) {}

// ExitBinding is called when production binding is exited.
func (s *BaseVyLangListener) ExitBinding(ctx *BindingContext) {}

// EnterNamedType is called when production namedType is entered.
func (s *BaseVyLangListener) EnterNamedType(ctx *NamedTypeContext) {}

// ExitNamedType is called when production namedType is exited.
func (s *BaseVyLangListener) ExitNamedType(ctx *NamedTypeContext) {}

// EnterType is called when production type is entered.
func (s *BaseVyLangListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseVyLangListener) ExitType(ctx *TypeContext) {}

// EnterRawType is called when production rawType is entered.
func (s *BaseVyLangListener) EnterRawType(ctx *RawTypeContext) {}

// ExitRawType is called when production rawType is exited.
func (s *BaseVyLangListener) ExitRawType(ctx *RawTypeContext) {}

// EnterTypeMap is called when production typeMap is entered.
func (s *BaseVyLangListener) EnterTypeMap(ctx *TypeMapContext) {}

// ExitTypeMap is called when production typeMap is exited.
func (s *BaseVyLangListener) ExitTypeMap(ctx *TypeMapContext) {}

// EnterTypeList is called when production typeList is entered.
func (s *BaseVyLangListener) EnterTypeList(ctx *TypeListContext) {}

// ExitTypeList is called when production typeList is exited.
func (s *BaseVyLangListener) ExitTypeList(ctx *TypeListContext) {}

// EnterTypeReference is called when production typeReference is entered.
func (s *BaseVyLangListener) EnterTypeReference(ctx *TypeReferenceContext) {}

// ExitTypeReference is called when production typeReference is exited.
func (s *BaseVyLangListener) ExitTypeReference(ctx *TypeReferenceContext) {}

// EnterTypeLeaf is called when production typeLeaf is entered.
func (s *BaseVyLangListener) EnterTypeLeaf(ctx *TypeLeafContext) {}

// ExitTypeLeaf is called when production typeLeaf is exited.
func (s *BaseVyLangListener) ExitTypeLeaf(ctx *TypeLeafContext) {}

// EnterTypeMapEntry is called when production typeMapEntry is entered.
func (s *BaseVyLangListener) EnterTypeMapEntry(ctx *TypeMapEntryContext) {}

// ExitTypeMapEntry is called when production typeMapEntry is exited.
func (s *BaseVyLangListener) ExitTypeMapEntry(ctx *TypeMapEntryContext) {}

// EnterNamedExpr is called when production namedExpr is entered.
func (s *BaseVyLangListener) EnterNamedExpr(ctx *NamedExprContext) {}

// ExitNamedExpr is called when production namedExpr is exited.
func (s *BaseVyLangListener) ExitNamedExpr(ctx *NamedExprContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseVyLangListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseVyLangListener) ExitExpr(ctx *ExprContext) {}

// EnterExprBrackets is called when production exprBrackets is entered.
func (s *BaseVyLangListener) EnterExprBrackets(ctx *ExprBracketsContext) {}

// ExitExprBrackets is called when production exprBrackets is exited.
func (s *BaseVyLangListener) ExitExprBrackets(ctx *ExprBracketsContext) {}

// EnterExprAlternative is called when production exprAlternative is entered.
func (s *BaseVyLangListener) EnterExprAlternative(ctx *ExprAlternativeContext) {}

// ExitExprAlternative is called when production exprAlternative is exited.
func (s *BaseVyLangListener) ExitExprAlternative(ctx *ExprAlternativeContext) {}

// EnterExprOperator1 is called when production exprOperator1 is entered.
func (s *BaseVyLangListener) EnterExprOperator1(ctx *ExprOperator1Context) {}

// ExitExprOperator1 is called when production exprOperator1 is exited.
func (s *BaseVyLangListener) ExitExprOperator1(ctx *ExprOperator1Context) {}

// EnterExprOperator2 is called when production exprOperator2 is entered.
func (s *BaseVyLangListener) EnterExprOperator2(ctx *ExprOperator2Context) {}

// ExitExprOperator2 is called when production exprOperator2 is exited.
func (s *BaseVyLangListener) ExitExprOperator2(ctx *ExprOperator2Context) {}

// EnterExprReference is called when production exprReference is entered.
func (s *BaseVyLangListener) EnterExprReference(ctx *ExprReferenceContext) {}

// ExitExprReference is called when production exprReference is exited.
func (s *BaseVyLangListener) ExitExprReference(ctx *ExprReferenceContext) {}

// EnterExprMap is called when production exprMap is entered.
func (s *BaseVyLangListener) EnterExprMap(ctx *ExprMapContext) {}

// ExitExprMap is called when production exprMap is exited.
func (s *BaseVyLangListener) ExitExprMap(ctx *ExprMapContext) {}

// EnterExprList is called when production exprList is entered.
func (s *BaseVyLangListener) EnterExprList(ctx *ExprListContext) {}

// ExitExprList is called when production exprList is exited.
func (s *BaseVyLangListener) ExitExprList(ctx *ExprListContext) {}

// EnterExprMapEntry is called when production exprMapEntry is entered.
func (s *BaseVyLangListener) EnterExprMapEntry(ctx *ExprMapEntryContext) {}

// ExitExprMapEntry is called when production exprMapEntry is exited.
func (s *BaseVyLangListener) ExitExprMapEntry(ctx *ExprMapEntryContext) {}

// EnterNamedLiteral is called when production namedLiteral is entered.
func (s *BaseVyLangListener) EnterNamedLiteral(ctx *NamedLiteralContext) {}

// ExitNamedLiteral is called when production namedLiteral is exited.
func (s *BaseVyLangListener) ExitNamedLiteral(ctx *NamedLiteralContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseVyLangListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseVyLangListener) ExitLiteral(ctx *LiteralContext) {}

// EnterLiteralTerminal is called when production literalTerminal is entered.
func (s *BaseVyLangListener) EnterLiteralTerminal(ctx *LiteralTerminalContext) {}

// ExitLiteralTerminal is called when production literalTerminal is exited.
func (s *BaseVyLangListener) ExitLiteralTerminal(ctx *LiteralTerminalContext) {}

// EnterLiteralString is called when production literalString is entered.
func (s *BaseVyLangListener) EnterLiteralString(ctx *LiteralStringContext) {}

// ExitLiteralString is called when production literalString is exited.
func (s *BaseVyLangListener) ExitLiteralString(ctx *LiteralStringContext) {}

// EnterLiteralInt is called when production literalInt is entered.
func (s *BaseVyLangListener) EnterLiteralInt(ctx *LiteralIntContext) {}

// ExitLiteralInt is called when production literalInt is exited.
func (s *BaseVyLangListener) ExitLiteralInt(ctx *LiteralIntContext) {}

// EnterLiteralBoolean is called when production literalBoolean is entered.
func (s *BaseVyLangListener) EnterLiteralBoolean(ctx *LiteralBooleanContext) {}

// ExitLiteralBoolean is called when production literalBoolean is exited.
func (s *BaseVyLangListener) ExitLiteralBoolean(ctx *LiteralBooleanContext) {}

// EnterLiteralFloat is called when production literalFloat is entered.
func (s *BaseVyLangListener) EnterLiteralFloat(ctx *LiteralFloatContext) {}

// ExitLiteralFloat is called when production literalFloat is exited.
func (s *BaseVyLangListener) ExitLiteralFloat(ctx *LiteralFloatContext) {}

// EnterLiteralNull is called when production literalNull is entered.
func (s *BaseVyLangListener) EnterLiteralNull(ctx *LiteralNullContext) {}

// ExitLiteralNull is called when production literalNull is exited.
func (s *BaseVyLangListener) ExitLiteralNull(ctx *LiteralNullContext) {}

// EnterLiteralReference is called when production literalReference is entered.
func (s *BaseVyLangListener) EnterLiteralReference(ctx *LiteralReferenceContext) {}

// ExitLiteralReference is called when production literalReference is exited.
func (s *BaseVyLangListener) ExitLiteralReference(ctx *LiteralReferenceContext) {}

// EnterLiteralMap is called when production literalMap is entered.
func (s *BaseVyLangListener) EnterLiteralMap(ctx *LiteralMapContext) {}

// ExitLiteralMap is called when production literalMap is exited.
func (s *BaseVyLangListener) ExitLiteralMap(ctx *LiteralMapContext) {}

// EnterLiteralList is called when production literalList is entered.
func (s *BaseVyLangListener) EnterLiteralList(ctx *LiteralListContext) {}

// ExitLiteralList is called when production literalList is exited.
func (s *BaseVyLangListener) ExitLiteralList(ctx *LiteralListContext) {}

// EnterLiteralMapEntry is called when production literalMapEntry is entered.
func (s *BaseVyLangListener) EnterLiteralMapEntry(ctx *LiteralMapEntryContext) {}

// ExitLiteralMapEntry is called when production literalMapEntry is exited.
func (s *BaseVyLangListener) ExitLiteralMapEntry(ctx *LiteralMapEntryContext) {}
