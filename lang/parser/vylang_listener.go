// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // VyLang
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// VyLangListener is a complete listener for a parse tree produced by VyLangParser.
type VyLangListener interface {
	antlr.ParseTreeListener

	// EnterDefinitions is called when entering the definitions production.
	EnterDefinitions(c *DefinitionsContext)

	// EnterDefinition is called when entering the definition production.
	EnterDefinition(c *DefinitionContext)

	// EnterNamedPipe is called when entering the namedPipe production.
	EnterNamedPipe(c *NamedPipeContext)

	// EnterContextPipe is called when entering the contextPipe production.
	EnterContextPipe(c *ContextPipeContext)

	// EnterPipe is called when entering the pipe production.
	EnterPipe(c *PipeContext)

	// EnterPipeProperty is called when entering the pipeProperty production.
	EnterPipeProperty(c *PipePropertyContext)

	// EnterPipeNamedProperty is called when entering the pipeNamedProperty production.
	EnterPipeNamedProperty(c *PipeNamedPropertyContext)

	// EnterPipeTerminal is called when entering the pipeTerminal production.
	EnterPipeTerminal(c *PipeTerminalContext)

	// EnterPipeField is called when entering the pipeField production.
	EnterPipeField(c *PipeFieldContext)

	// EnterPipeFieldForward is called when entering the pipeFieldForward production.
	EnterPipeFieldForward(c *PipeFieldForwardContext)

	// EnterPipeFieldBackward is called when entering the pipeFieldBackward production.
	EnterPipeFieldBackward(c *PipeFieldBackwardContext)

	// EnterPipeModified is called when entering the pipeModified production.
	EnterPipeModified(c *PipeModifiedContext)

	// EnterPipeMap is called when entering the pipeMap production.
	EnterPipeMap(c *PipeMapContext)

	// EnterPipeMapEntry is called when entering the pipeMapEntry production.
	EnterPipeMapEntry(c *PipeMapEntryContext)

	// EnterPipeModifier is called when entering the pipeModifier production.
	EnterPipeModifier(c *PipeModifierContext)

	// EnterPathModel is called when entering the pathModel production.
	EnterPathModel(c *PathModelContext)

	// EnterPathRelation is called when entering the pathRelation production.
	EnterPathRelation(c *PathRelationContext)

	// EnterIdentPath is called when entering the identPath production.
	EnterIdentPath(c *IdentPathContext)

	// EnterSep is called when entering the sep production.
	EnterSep(c *SepContext)

	// ExitDefinitions is called when exiting the definitions production.
	ExitDefinitions(c *DefinitionsContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitNamedPipe is called when exiting the namedPipe production.
	ExitNamedPipe(c *NamedPipeContext)

	// ExitContextPipe is called when exiting the contextPipe production.
	ExitContextPipe(c *ContextPipeContext)

	// ExitPipe is called when exiting the pipe production.
	ExitPipe(c *PipeContext)

	// ExitPipeProperty is called when exiting the pipeProperty production.
	ExitPipeProperty(c *PipePropertyContext)

	// ExitPipeNamedProperty is called when exiting the pipeNamedProperty production.
	ExitPipeNamedProperty(c *PipeNamedPropertyContext)

	// ExitPipeTerminal is called when exiting the pipeTerminal production.
	ExitPipeTerminal(c *PipeTerminalContext)

	// ExitPipeField is called when exiting the pipeField production.
	ExitPipeField(c *PipeFieldContext)

	// ExitPipeFieldForward is called when exiting the pipeFieldForward production.
	ExitPipeFieldForward(c *PipeFieldForwardContext)

	// ExitPipeFieldBackward is called when exiting the pipeFieldBackward production.
	ExitPipeFieldBackward(c *PipeFieldBackwardContext)

	// ExitPipeModified is called when exiting the pipeModified production.
	ExitPipeModified(c *PipeModifiedContext)

	// ExitPipeMap is called when exiting the pipeMap production.
	ExitPipeMap(c *PipeMapContext)

	// ExitPipeMapEntry is called when exiting the pipeMapEntry production.
	ExitPipeMapEntry(c *PipeMapEntryContext)

	// ExitPipeModifier is called when exiting the pipeModifier production.
	ExitPipeModifier(c *PipeModifierContext)

	// ExitPathModel is called when exiting the pathModel production.
	ExitPathModel(c *PathModelContext)

	// ExitPathRelation is called when exiting the pathRelation production.
	ExitPathRelation(c *PathRelationContext)

	// ExitIdentPath is called when exiting the identPath production.
	ExitIdentPath(c *IdentPathContext)

	// ExitSep is called when exiting the sep production.
	ExitSep(c *SepContext)
}
