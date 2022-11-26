// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // VyLang
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// VyLangListener is a complete listener for a parse tree produced by VyLangParser.
type VyLangListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterDefinitions is called when entering the definitions production.
	EnterDefinitions(c *DefinitionsContext)

	// EnterDefinition is called when entering the definition production.
	EnterDefinition(c *DefinitionContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterNamedAction is called when entering the namedAction production.
	EnterNamedAction(c *NamedActionContext)

	// EnterAction is called when entering the action production.
	EnterAction(c *ActionContext)

	// EnterActionParallel is called when entering the actionParallel production.
	EnterActionParallel(c *ActionParallelContext)

	// EnterActionSequence is called when entering the actionSequence production.
	EnterActionSequence(c *ActionSequenceContext)

	// EnterActionAsync is called when entering the actionAsync production.
	EnterActionAsync(c *ActionAsyncContext)

	// EnterActionReference is called when entering the actionReference production.
	EnterActionReference(c *ActionReferenceContext)

	// EnterActionBind is called when entering the actionBind production.
	EnterActionBind(c *ActionBindContext)

	// EnterActionWhile is called when entering the actionWhile production.
	EnterActionWhile(c *ActionWhileContext)

	// EnterActionCond is called when entering the actionCond production.
	EnterActionCond(c *ActionCondContext)

	// EnterActionIf is called when entering the actionIf production.
	EnterActionIf(c *ActionIfContext)

	// EnterActionIfElse is called when entering the actionIfElse production.
	EnterActionIfElse(c *ActionIfElseContext)

	// EnterActionLeaf is called when entering the actionLeaf production.
	EnterActionLeaf(c *ActionLeafContext)

	// EnterActionAssign is called when entering the actionAssign production.
	EnterActionAssign(c *ActionAssignContext)

	// EnterActionClear is called when entering the actionClear production.
	EnterActionClear(c *ActionClearContext)

	// EnterBinding is called when entering the binding production.
	EnterBinding(c *BindingContext)

	// EnterNamedType is called when entering the namedType production.
	EnterNamedType(c *NamedTypeContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterRawType is called when entering the rawType production.
	EnterRawType(c *RawTypeContext)

	// EnterTypeMap is called when entering the typeMap production.
	EnterTypeMap(c *TypeMapContext)

	// EnterTypeList is called when entering the typeList production.
	EnterTypeList(c *TypeListContext)

	// EnterTypeReference is called when entering the typeReference production.
	EnterTypeReference(c *TypeReferenceContext)

	// EnterTypeLeaf is called when entering the typeLeaf production.
	EnterTypeLeaf(c *TypeLeafContext)

	// EnterTypeMapEntry is called when entering the typeMapEntry production.
	EnterTypeMapEntry(c *TypeMapEntryContext)

	// EnterNamedExpr is called when entering the namedExpr production.
	EnterNamedExpr(c *NamedExprContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterExprBrackets is called when entering the exprBrackets production.
	EnterExprBrackets(c *ExprBracketsContext)

	// EnterExprAlternative is called when entering the exprAlternative production.
	EnterExprAlternative(c *ExprAlternativeContext)

	// EnterExprOperator1 is called when entering the exprOperator1 production.
	EnterExprOperator1(c *ExprOperator1Context)

	// EnterExprOperator2 is called when entering the exprOperator2 production.
	EnterExprOperator2(c *ExprOperator2Context)

	// EnterExprReference is called when entering the exprReference production.
	EnterExprReference(c *ExprReferenceContext)

	// EnterExprMap is called when entering the exprMap production.
	EnterExprMap(c *ExprMapContext)

	// EnterExprList is called when entering the exprList production.
	EnterExprList(c *ExprListContext)

	// EnterExprMapEntry is called when entering the exprMapEntry production.
	EnterExprMapEntry(c *ExprMapEntryContext)

	// EnterNamedLiteral is called when entering the namedLiteral production.
	EnterNamedLiteral(c *NamedLiteralContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterLiteralTerminal is called when entering the literalTerminal production.
	EnterLiteralTerminal(c *LiteralTerminalContext)

	// EnterLiteralString is called when entering the literalString production.
	EnterLiteralString(c *LiteralStringContext)

	// EnterLiteralInt is called when entering the literalInt production.
	EnterLiteralInt(c *LiteralIntContext)

	// EnterLiteralBoolean is called when entering the literalBoolean production.
	EnterLiteralBoolean(c *LiteralBooleanContext)

	// EnterLiteralFloat is called when entering the literalFloat production.
	EnterLiteralFloat(c *LiteralFloatContext)

	// EnterLiteralNull is called when entering the literalNull production.
	EnterLiteralNull(c *LiteralNullContext)

	// EnterLiteralReference is called when entering the literalReference production.
	EnterLiteralReference(c *LiteralReferenceContext)

	// EnterLiteralMap is called when entering the literalMap production.
	EnterLiteralMap(c *LiteralMapContext)

	// EnterLiteralList is called when entering the literalList production.
	EnterLiteralList(c *LiteralListContext)

	// EnterLiteralMapEntry is called when entering the literalMapEntry production.
	EnterLiteralMapEntry(c *LiteralMapEntryContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitDefinitions is called when exiting the definitions production.
	ExitDefinitions(c *DefinitionsContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitNamedAction is called when exiting the namedAction production.
	ExitNamedAction(c *NamedActionContext)

	// ExitAction is called when exiting the action production.
	ExitAction(c *ActionContext)

	// ExitActionParallel is called when exiting the actionParallel production.
	ExitActionParallel(c *ActionParallelContext)

	// ExitActionSequence is called when exiting the actionSequence production.
	ExitActionSequence(c *ActionSequenceContext)

	// ExitActionAsync is called when exiting the actionAsync production.
	ExitActionAsync(c *ActionAsyncContext)

	// ExitActionReference is called when exiting the actionReference production.
	ExitActionReference(c *ActionReferenceContext)

	// ExitActionBind is called when exiting the actionBind production.
	ExitActionBind(c *ActionBindContext)

	// ExitActionWhile is called when exiting the actionWhile production.
	ExitActionWhile(c *ActionWhileContext)

	// ExitActionCond is called when exiting the actionCond production.
	ExitActionCond(c *ActionCondContext)

	// ExitActionIf is called when exiting the actionIf production.
	ExitActionIf(c *ActionIfContext)

	// ExitActionIfElse is called when exiting the actionIfElse production.
	ExitActionIfElse(c *ActionIfElseContext)

	// ExitActionLeaf is called when exiting the actionLeaf production.
	ExitActionLeaf(c *ActionLeafContext)

	// ExitActionAssign is called when exiting the actionAssign production.
	ExitActionAssign(c *ActionAssignContext)

	// ExitActionClear is called when exiting the actionClear production.
	ExitActionClear(c *ActionClearContext)

	// ExitBinding is called when exiting the binding production.
	ExitBinding(c *BindingContext)

	// ExitNamedType is called when exiting the namedType production.
	ExitNamedType(c *NamedTypeContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitRawType is called when exiting the rawType production.
	ExitRawType(c *RawTypeContext)

	// ExitTypeMap is called when exiting the typeMap production.
	ExitTypeMap(c *TypeMapContext)

	// ExitTypeList is called when exiting the typeList production.
	ExitTypeList(c *TypeListContext)

	// ExitTypeReference is called when exiting the typeReference production.
	ExitTypeReference(c *TypeReferenceContext)

	// ExitTypeLeaf is called when exiting the typeLeaf production.
	ExitTypeLeaf(c *TypeLeafContext)

	// ExitTypeMapEntry is called when exiting the typeMapEntry production.
	ExitTypeMapEntry(c *TypeMapEntryContext)

	// ExitNamedExpr is called when exiting the namedExpr production.
	ExitNamedExpr(c *NamedExprContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitExprBrackets is called when exiting the exprBrackets production.
	ExitExprBrackets(c *ExprBracketsContext)

	// ExitExprAlternative is called when exiting the exprAlternative production.
	ExitExprAlternative(c *ExprAlternativeContext)

	// ExitExprOperator1 is called when exiting the exprOperator1 production.
	ExitExprOperator1(c *ExprOperator1Context)

	// ExitExprOperator2 is called when exiting the exprOperator2 production.
	ExitExprOperator2(c *ExprOperator2Context)

	// ExitExprReference is called when exiting the exprReference production.
	ExitExprReference(c *ExprReferenceContext)

	// ExitExprMap is called when exiting the exprMap production.
	ExitExprMap(c *ExprMapContext)

	// ExitExprList is called when exiting the exprList production.
	ExitExprList(c *ExprListContext)

	// ExitExprMapEntry is called when exiting the exprMapEntry production.
	ExitExprMapEntry(c *ExprMapEntryContext)

	// ExitNamedLiteral is called when exiting the namedLiteral production.
	ExitNamedLiteral(c *NamedLiteralContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitLiteralTerminal is called when exiting the literalTerminal production.
	ExitLiteralTerminal(c *LiteralTerminalContext)

	// ExitLiteralString is called when exiting the literalString production.
	ExitLiteralString(c *LiteralStringContext)

	// ExitLiteralInt is called when exiting the literalInt production.
	ExitLiteralInt(c *LiteralIntContext)

	// ExitLiteralBoolean is called when exiting the literalBoolean production.
	ExitLiteralBoolean(c *LiteralBooleanContext)

	// ExitLiteralFloat is called when exiting the literalFloat production.
	ExitLiteralFloat(c *LiteralFloatContext)

	// ExitLiteralNull is called when exiting the literalNull production.
	ExitLiteralNull(c *LiteralNullContext)

	// ExitLiteralReference is called when exiting the literalReference production.
	ExitLiteralReference(c *LiteralReferenceContext)

	// ExitLiteralMap is called when exiting the literalMap production.
	ExitLiteralMap(c *LiteralMapContext)

	// ExitLiteralList is called when exiting the literalList production.
	ExitLiteralList(c *LiteralListContext)

	// ExitLiteralMapEntry is called when exiting the literalMapEntry production.
	ExitLiteralMapEntry(c *LiteralMapEntryContext)
}
