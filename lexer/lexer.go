package lexer

import "github.com/HlufD/my_interpreter/token"

/*
*	input        string
*	position     int  // current position in input string (points to current char)
*	readPosition int  // current reading position in input strng (after current char) next to position if posi = 1 ,readp = 2
*	ch           byte // current char under examination in the string
 */

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

/**
* ! this method takes a char which is being read contuues reading the charachter untill it encouter a non alphabetic character except _ (underscore) -> then if it gets non alphabetic char it stops and retuen the slice of the input from the first position
to the index of non alphanumeric char
* l.readChar() // evertime called it updates the readPosiion and position = readPosiion-1
* when the cha bescome " "non letter of  "" this loop breaks
* readPosition-> char after the " "
* position -> at at the char " "
*/

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	}
	return l.input[l.readPosition]
}

/**
* ! this method reterns a single token bosed on lexter ch properety which is the  byte form of current charatcter beign read
* ! it calls the readChar() method to get the next ch to be tokinzed
 */

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch   // the firat one or the fist char
			l.readChar() // we update or ch,positio and readPosition
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}

		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch   // the firat one or the fist char
			l.readChar() // we update or ch,positio and readPosition
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '>':
		tok = newToken(token.GT, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

/*
*	l.readChar() //method that changes postion ,readPosition and ch of l // this intializes
*	// postion : 0
*	// ch :"l"
*	// readPostion : 1
 */

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// helpers

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}
