package ast

import (
	"testing"

	"github.com/HlufD/my_interpreter/token"
)

func TestString(t *testing.T) {
	//test := "let myVar = anotherVar"
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name:  &Identifier{Token: token.Token{Type: token.IDENT, Literal: "myVar"}, Value: "myVar"},
				Value: &Identifier{Token: token.Token{Type: token.IDENT, Literal: "anotherVar"}, Value: "anotherVar"},
			},
		},
	}
	if program.String() != "let myVar=anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}

}
