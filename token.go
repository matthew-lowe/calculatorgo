package main

import (
	"fmt"
)

type TokenType int

const (
	NONE TokenType = iota
	WHITESPACE
	NUM
	ADD
	SUB
	MUL
	DIV
)

type Token struct {
	Type     TokenType
	HasValue bool
	Value    float64
}

func NewToken(_type TokenType, hasValue bool, value float64) *Token {
	return &Token{Type: _type, HasValue: hasValue, Value: value}
}

func (token Token) ToString() string {
	if token.HasValue {
		return fmt.Sprintf("Token(type: %v, value: %v)", token.Type, token.Value)
	} else {
		return fmt.Sprintf("Token(type: %v)", token.Type)
	}
}
