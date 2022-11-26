grammar VyLang;
prog:	definitions EOF ;

definitions: (NEWLINE* definition NEWLINE*)*;
definition: namedAction | namedType | namedLiteral | namedExpr;

// Variable

variable: IDENTIFIER | variable '.' IDENTIFIER | '@';

// Actions

namedAction: 'action' IDENTIFIER 'on' type ':' NEWLINE* action NEWLINE*;
action: actionParallel
      | actionSequence
      | actionAsync
      | actionReference
      | actionBind
      | actionWhile
      | actionCond
      | actionLeaf;
actionParallel: '{' (NEWLINE* action (NEWLINE|',')+)* NEWLINE* '}';
actionSequence: '[' (NEWLINE* action (NEWLINE|',')+)* NEWLINE* ']';
actionAsync: '&' action;
actionReference: '$' IDENTIFIER;
actionBind: 'bind' binding? ('on' type)? ('=' expr)? ':' NEWLINE* action;
actionWhile: 'while' expr ':' NEWLINE* action;
actionCond: actionIf | actionIfElse;
actionIf: 'if' expr ':' NEWLINE* action;
actionIfElse: 'if' expr ':' NEWLINE* action NEWLINE* 'else' ':' NEWLINE* action;
actionLeaf: actionAssign
          | actionClear;
actionAssign: variable '=' expr;
actionClear: 'del' variable;

binding: variable ('<-'|'->'|'<=>') variable
       | literal '->' variable
       | binding ',' binding
       | '(' NEWLINE* binding NEWLINE* ')';

// Types

namedType: 'type' IDENTIFIER ':' type;
type: rawType ('=' literal)? ('in' literalList)?;
rawType: typeMap
       | typeList
       | typeReference
       | typeLeaf;
typeMap: '{' NEWLINE* typeMapEntry ((NEWLINE|',')+ typeMapEntry)* NEWLINE* '}';
typeList: '[' ']' type;
typeReference: '$' IDENTIFIER;
typeLeaf: 'string' | 'integer' | 'float' | 'boolean' | 'raw';

typeMapEntry: IDENTIFIER ':' type;

// Expressions

namedExpr: 'function' IDENTIFIER 'on' type ':' expr;
expr: variable
    | literal
    | exprMap
    | exprList
    | exprBrackets
    | exprReference
    | exprOperator1
    | expr exprOperator2+
    | expr exprAlternative;
exprBrackets: '(' expr ')';
exprAlternative: '?' expr ':' expr;
exprOperator1: '-' expr
             | '!' expr;
exprOperator2: ('*'|'/') expr
             | ('+'|'-') expr
             | ('>'|'<'|'>='|'<=') expr
             | ('=='|'!=') expr
             | '&&' expr
             | '||' expr
             | ('++'|'--') expr;
exprReference: '$' IDENTIFIER expr?;
exprMap: '{' NEWLINE* exprMapEntry ((NEWLINE|',')+ exprMapEntry)* NEWLINE* '}';
exprList: '[' NEWLINE* expr ((NEWLINE|',')+ expr)* NEWLINE* ']';
exprMapEntry: IDENTIFIER ':' expr;

// Literals

namedLiteral: 'constant' IDENTIFIER ':' literal;
literal: literalTerminal
       | literalMap
       | literalList;
literalTerminal: literalString
               | literalInt
               | literalBoolean
               | literalFloat
               | literalNull
               | literalReference;
literalString: STRING;
literalInt: INT;
literalBoolean: 'true' | 'false';
literalFloat: INT '.' INT;
literalNull: 'null';
literalReference: '$' IDENTIFIER;
literalMap: '{' NEWLINE* literalMapEntry ((NEWLINE|',')+ literalMapEntry)* NEWLINE* '}';
literalList: '[' NEWLINE* literal ((NEWLINE|',')+ literal)* NEWLINE* ']';

literalMapEntry: IDENTIFIER ':' literal;

// Terminals

STRING: '"' ~["]* '"';
IDENTIFIER: [a-zA-Z]+[a-zA-Z0-9]*;
NEWLINE: [\r\n]+;
WHITESPACE: [\t ] -> skip;
INT: [0-9]+;
