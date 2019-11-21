// based on: https://github.com/cosenmarco/yabdd
grammar Feature;

feature: comment* featureHeader featureBody (NewLine+ | EOF)?;

featureHeader: (Space | NewLine)* tags? Feature Space* content NewLine+;

featureBody: background? scenario+;

background: (Space | NewLine)* tags* Background Space* content? (NewLine | EOF) blockBody (NewLine | EOF);

blockBody: (given | when | or | then | and | example)*;

scenario: (Space | NewLine)* tags* Space* Scenario Space* content (NewLine | EOF) blockBody;

// Annotations

tags: (Space | NewLine)* At anyText value? Space* NewLine;

anyText: .*?;

value: LBracket content RBracket;

// Keywords

given: (Space | NewLine)* Given step;

when: (Space | NewLine)* When step;

or: (Space | NewLine)* Or step;

and: (Space | NewLine)* And step;

then: (Space | NewLine)* Then step;

example: (Space | NewLine)* Examples step;

// Steps and data tables

step: Space* stepContent;

stepContent: stepText table? (NewLine+ | EOF);

stepText: (contentNoQuotes | Space | parameter)*;

table: tableHeader row*;

tableHeader: Space* Pipe cell+ (NewLine+ | EOF);

row: Space* Pipe cell+ (NewLine+ | EOF);

cell: Space* contentNoPipes Pipe;

parameter: Quote anyText Quote;

// Common

contentNoQuotes: (Char|LBracket) (Char|LBracket|RBracket|At|Pipe|Space)*;

contentNoPipes: (Char|LBracket) (Char|LBracket|RBracket|At|Quote|Space)*;

content: (Char|LBracket) (Char|LBracket|RBracket|At|Quote|Pipe|Space)*;

comment: '#' Space* commentText NewLine;

commentText: IDENTIFIER ':' Space* commentValue;

commentValue: IDENTIFIER;

IDENTIFIER: LetterOrDigit LetterOrDigit*;

EmptyLine: NewLine Space+ (NewLine | EOF) -> skip;

And: 'And ' | '而且' | '并且' | '同时' ;
Or: 'Or ' | '或者';
Given: 'Given ' | '假如' | '假设' | '假定' ;
When: 'When ' | '当';
Then: 'Then ' | '那么';
Examples: 'Example ' | '例子';
Background: 'Background' | '背景' [ ] Space? COLON;
Scenario: 'Scenario' | '场景' | '剧本' Space? COLON;
Feature: 'Feature' | '功能' Space? COLON;

Space : [ \t];
NewLine : '\r\n' | '\n';
Quote: '"';
LBracket: '(';
RBracket: ')';
At: '@';
Pipe: '|';
Char: ~[ \t\r\n()]+?;

fragment COLON: ':';

fragment LetterOrDigit
    : Letter
    | [0-9]
    ;

fragment Letter
    : [a-zA-Z$_-] // these are the "java letters" below 0x7F
    | ~[\u0000-\u007F\uD800-\uDBFF] // covers all characters above 0x7F which are not a surrogate
    | [\uD800-\uDBFF] [\uDC00-\uDFFF] // covers UTF-16 surrogate pairs encodings for U+10000 to U+10FFFF
    ;
