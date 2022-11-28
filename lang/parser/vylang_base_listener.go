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

// EnterDefinitions is called when production definitions is entered.
func (s *BaseVyLangListener) EnterDefinitions(ctx *DefinitionsContext) {}

// ExitDefinitions is called when production definitions is exited.
func (s *BaseVyLangListener) ExitDefinitions(ctx *DefinitionsContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BaseVyLangListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BaseVyLangListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterNamedPipe is called when production namedPipe is entered.
func (s *BaseVyLangListener) EnterNamedPipe(ctx *NamedPipeContext) {}

// ExitNamedPipe is called when production namedPipe is exited.
func (s *BaseVyLangListener) ExitNamedPipe(ctx *NamedPipeContext) {}

// EnterContextPipe is called when production contextPipe is entered.
func (s *BaseVyLangListener) EnterContextPipe(ctx *ContextPipeContext) {}

// ExitContextPipe is called when production contextPipe is exited.
func (s *BaseVyLangListener) ExitContextPipe(ctx *ContextPipeContext) {}

// EnterPipe is called when production pipe is entered.
func (s *BaseVyLangListener) EnterPipe(ctx *PipeContext) {}

// ExitPipe is called when production pipe is exited.
func (s *BaseVyLangListener) ExitPipe(ctx *PipeContext) {}

// EnterPipeProperty is called when production pipeProperty is entered.
func (s *BaseVyLangListener) EnterPipeProperty(ctx *PipePropertyContext) {}

// ExitPipeProperty is called when production pipeProperty is exited.
func (s *BaseVyLangListener) ExitPipeProperty(ctx *PipePropertyContext) {}

// EnterPipeNamedProperty is called when production pipeNamedProperty is entered.
func (s *BaseVyLangListener) EnterPipeNamedProperty(ctx *PipeNamedPropertyContext) {}

// ExitPipeNamedProperty is called when production pipeNamedProperty is exited.
func (s *BaseVyLangListener) ExitPipeNamedProperty(ctx *PipeNamedPropertyContext) {}

// EnterPipeTerminal is called when production pipeTerminal is entered.
func (s *BaseVyLangListener) EnterPipeTerminal(ctx *PipeTerminalContext) {}

// ExitPipeTerminal is called when production pipeTerminal is exited.
func (s *BaseVyLangListener) ExitPipeTerminal(ctx *PipeTerminalContext) {}

// EnterPipeField is called when production pipeField is entered.
func (s *BaseVyLangListener) EnterPipeField(ctx *PipeFieldContext) {}

// ExitPipeField is called when production pipeField is exited.
func (s *BaseVyLangListener) ExitPipeField(ctx *PipeFieldContext) {}

// EnterPipeFieldForward is called when production pipeFieldForward is entered.
func (s *BaseVyLangListener) EnterPipeFieldForward(ctx *PipeFieldForwardContext) {}

// ExitPipeFieldForward is called when production pipeFieldForward is exited.
func (s *BaseVyLangListener) ExitPipeFieldForward(ctx *PipeFieldForwardContext) {}

// EnterPipeFieldBackward is called when production pipeFieldBackward is entered.
func (s *BaseVyLangListener) EnterPipeFieldBackward(ctx *PipeFieldBackwardContext) {}

// ExitPipeFieldBackward is called when production pipeFieldBackward is exited.
func (s *BaseVyLangListener) ExitPipeFieldBackward(ctx *PipeFieldBackwardContext) {}

// EnterPipeModified is called when production pipeModified is entered.
func (s *BaseVyLangListener) EnterPipeModified(ctx *PipeModifiedContext) {}

// ExitPipeModified is called when production pipeModified is exited.
func (s *BaseVyLangListener) ExitPipeModified(ctx *PipeModifiedContext) {}

// EnterPipeMap is called when production pipeMap is entered.
func (s *BaseVyLangListener) EnterPipeMap(ctx *PipeMapContext) {}

// ExitPipeMap is called when production pipeMap is exited.
func (s *BaseVyLangListener) ExitPipeMap(ctx *PipeMapContext) {}

// EnterPipeMapEntry is called when production pipeMapEntry is entered.
func (s *BaseVyLangListener) EnterPipeMapEntry(ctx *PipeMapEntryContext) {}

// ExitPipeMapEntry is called when production pipeMapEntry is exited.
func (s *BaseVyLangListener) ExitPipeMapEntry(ctx *PipeMapEntryContext) {}

// EnterPipeModifier is called when production pipeModifier is entered.
func (s *BaseVyLangListener) EnterPipeModifier(ctx *PipeModifierContext) {}

// ExitPipeModifier is called when production pipeModifier is exited.
func (s *BaseVyLangListener) ExitPipeModifier(ctx *PipeModifierContext) {}

// EnterPathModel is called when production pathModel is entered.
func (s *BaseVyLangListener) EnterPathModel(ctx *PathModelContext) {}

// ExitPathModel is called when production pathModel is exited.
func (s *BaseVyLangListener) ExitPathModel(ctx *PathModelContext) {}

// EnterPathRelation is called when production pathRelation is entered.
func (s *BaseVyLangListener) EnterPathRelation(ctx *PathRelationContext) {}

// ExitPathRelation is called when production pathRelation is exited.
func (s *BaseVyLangListener) ExitPathRelation(ctx *PathRelationContext) {}

// EnterIdentPath is called when production identPath is entered.
func (s *BaseVyLangListener) EnterIdentPath(ctx *IdentPathContext) {}

// ExitIdentPath is called when production identPath is exited.
func (s *BaseVyLangListener) ExitIdentPath(ctx *IdentPathContext) {}

// EnterSep is called when production sep is entered.
func (s *BaseVyLangListener) EnterSep(ctx *SepContext) {}

// ExitSep is called when production sep is exited.
func (s *BaseVyLangListener) ExitSep(ctx *SepContext) {}
