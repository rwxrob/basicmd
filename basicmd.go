package basicmd

import (
	"io"
)

type AST struct {
}

func ParseAST(in io.Reader) *AST {
	ast := new(AST)
	// TODO
	return ast
}

func Validate(in io.Reader) bool {
	// TODO
	return false
}
