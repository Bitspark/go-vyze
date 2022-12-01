grammar VyLang;

definitions: NL* (NL* definition NL*)* NL* EOF;
definition: namedPipe;

// Pipe

namedPipe: 'pipe' IDENT contextPipe;
contextPipe: 'on' pathModel '->' pipeModified;
pipe: pipeProperty | pipeMap;
pipeProperty: pipeTerminal | pipeField;
pipeNamedProperty: IDENT ':' pipeProperty;
pipeTerminal: '@id' | '@name' | '@created' | '@size' | '@user' | '@base64' | '@hex' | '@string' | '@integer' | '@float' | '@boolean' | '@auto' | '@value';
pipeField: pipeFieldForward | pipeFieldBackward;
pipeFieldForward: pathRelation '->' pipeModified;
pipeFieldBackward: '<-' pathRelation pipeModified;
pipeModified: pipeModifier? NL* pipe;
pipeMap: '{' NL? (pipeMapEntry (sep pipeMapEntry)*)? NL? '}';
pipeMapEntry: pipeNamedProperty | pipeProperty;
pipeModifier: '[' ']';
pathModel: IDENT | identPath ('/' IDENT?)?;
pathRelation: IDENT | identPath '#' IDENT ('/' IDENT?)?;
identPath: IDENT ('.' IDENT)*;

// Separator

sep: NL | (NL? ',' NL?);

// Terminals

STRING: '"' ~["]* '"';
IDENT: [a-zA-Z]+[a-zA-Z0-9_]*;
NL: [\r\n]+;
WHITESPACE: [\t ] -> skip;
DIGITS: [0-9]+;
