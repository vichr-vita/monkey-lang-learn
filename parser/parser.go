package parser

import (
	"fmt"

	"vichr.me/go/monkey-lang-interpreter/ast"
	"vichr.me/go/monkey-lang-interpreter/lexer"
	"vichr.me/go/monkey-lang-interpreter/token"
)

type Parser struct {
	lexer     *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []error
}

func NewParser(lexer *lexer.Lexer) *Parser {
	p := &Parser{lexer: lexer, errors: []error{}}

	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Errors() []error {
	return p.errors
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIsType(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Errorf("expected next token to be %s got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	statement := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	for !p.curTokenIsType(token.SEMICOLON) {
		p.nextToken()
	}

	return statement
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for !p.curTokenIsType(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) curTokenIsType(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIsType(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIsType(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}
