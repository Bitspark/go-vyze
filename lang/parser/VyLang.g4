grammar VyLang;

definitions: NL* (NL* definition NL*)* NL*;
definition: namedPipe;

// Pipe

namedPipe: 'pipe' IDENT pipeContexualized;
pipeContexualized: 'on' IDENT '->' pipeModifier? NL* pipe;
pipe: pipeProperty
     | pipeMap;
pipeProperty: pipeTerminal
             | pipeField;
pipeTerminal: 'id' | 'name' | 'created' | 'data' | 'size' | 'user' | 'value' | IDENT;
pipeField: IDENT '->' pipeModifier? pipe;
pipeMap: '{' NL? (pipeMapEntry (sep pipeMapEntry)*)? NL? '}';
pipeMapEntry: IDENT ':' pipe
             | pipeProperty;
pipeModifier: '[' ']';

// Separator

sep: NL | (NL? ',' NL?);

// Terminals

STRING: '"' ~["]* '"';
IDENT: [a-zA-Z]+[a-zA-Z0-9_]*;
NL: [\r\n]+;
WHITESPACE: [\t ] -> skip;
DIGITS: [0-9]+;
