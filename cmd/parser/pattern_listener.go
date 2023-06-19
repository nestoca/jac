// Code generated from Pattern.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Pattern
import "github.com/antlr4-go/antlr/v4"

// PatternListener is a complete listener for a parse tree produced by PatternParser.
type PatternListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterOr is called when entering the or production.
	EnterOr(c *OrContext)

	// EnterAnd is called when entering the and production.
	EnterAnd(c *AndContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterParentheses is called when entering the parentheses production.
	EnterParentheses(c *ParenthesesContext)

	// EnterNot is called when entering the not production.
	EnterNot(c *NotContext)

	// EnterWildcard is called when entering the wildcard production.
	EnterWildcard(c *WildcardContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitOr is called when exiting the or production.
	ExitOr(c *OrContext)

	// ExitAnd is called when exiting the and production.
	ExitAnd(c *AndContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitParentheses is called when exiting the parentheses production.
	ExitParentheses(c *ParenthesesContext)

	// ExitNot is called when exiting the not production.
	ExitNot(c *NotContext)

	// ExitWildcard is called when exiting the wildcard production.
	ExitWildcard(c *WildcardContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)
}
