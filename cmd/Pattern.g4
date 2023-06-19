grammar Pattern;

parse: expression EOF;

expression: or;

or: and (',' and)*;

and: atom ('+' atom)*;

atom: parentheses | not | wildcard | literal;

parentheses: '(' expression ')';

not: '!' expression;

wildcard: WILDCARD;

literal: LITERAL;

WILDCARD: ([a-zA-Z0-9_-] | '*')+;
LITERAL: [a-zA-Z0-9_-]+;
WS: [ \t\r\n]+ -> skip;
