package interpret

import (
	"fmt"
)

type TokenType int

const (
	NONE TokenType = iota
	WHITESPACE
	NUM
	DOT
	L_BRACKET
	R_BRACKET
	ADD
	SUB
	MUL
	DIV
	EXP
)

func (t TokenType) String() string {
	strings := [...]string{"NONE", "WHITESPACE", "NUM", ".", "(", ")", "ADD", "SUB", "MUL", "DIV", "EXP"}

	if t < NONE || t > EXP {
		return "UNKOWN"
	}

	return strings[t]
}

type Token struct {
	Type     TokenType
	HasValue bool
	Value    float64
}

func NewToken(_type TokenType, hasValue bool, value float64) *Token {
	return &Token{Type: _type, HasValue: hasValue, Value: value}
}

func (token Token) String() string {
	if token.HasValue {
		return fmt.Sprintf("%v", token.Value)
	} else {
		return fmt.Sprintf("%v", token.Type.String())
	}
}
