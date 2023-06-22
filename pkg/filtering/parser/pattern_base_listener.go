// Code generated from Pattern.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Pattern

import "github.com/antlr4-go/antlr/v4"

// BasePatternListener is a complete listener for a parse tree produced by PatternParser.
type BasePatternListener struct{}

var _ PatternListener = &BasePatternListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePatternListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePatternListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePatternListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePatternListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BasePatternListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BasePatternListener) ExitRoot(ctx *RootContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePatternListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePatternListener) ExitExpression(ctx *ExpressionContext) {}

// EnterOr is called when production or is entered.
func (s *BasePatternListener) EnterOr(ctx *OrContext) {}

// ExitOr is called when production or is exited.
func (s *BasePatternListener) ExitOr(ctx *OrContext) {}

// EnterAnd is called when production and is entered.
func (s *BasePatternListener) EnterAnd(ctx *AndContext) {}

// ExitAnd is called when production and is exited.
func (s *BasePatternListener) ExitAnd(ctx *AndContext) {}

// EnterAtom is called when production atom is entered.
func (s *BasePatternListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BasePatternListener) ExitAtom(ctx *AtomContext) {}

// EnterParentheses is called when production parentheses is entered.
func (s *BasePatternListener) EnterParentheses(ctx *ParenthesesContext) {}

// ExitParentheses is called when production parentheses is exited.
func (s *BasePatternListener) ExitParentheses(ctx *ParenthesesContext) {}

// EnterNot is called when production not is entered.
func (s *BasePatternListener) EnterNot(ctx *NotContext) {}

// ExitNot is called when production not is exited.
func (s *BasePatternListener) ExitNot(ctx *NotContext) {}

// EnterWildcard is called when production wildcard is entered.
func (s *BasePatternListener) EnterWildcard(ctx *WildcardContext) {}

// ExitWildcard is called when production wildcard is exited.
func (s *BasePatternListener) ExitWildcard(ctx *WildcardContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BasePatternListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BasePatternListener) ExitLiteral(ctx *LiteralContext) {}
