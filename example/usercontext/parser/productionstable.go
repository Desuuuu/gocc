// Code generated by gocc; DO NOT EDIT.

package parser

import "github.com/Desuuuu/gocc/example/usercontext/ast"

type (
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib, interface{}) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String:     `S' : Grammar	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Grammar : Words "...42..."	<< X[0], nil >>`,
		Id:         "Grammar",
		NTType:     1,
		Index:      1,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Words : capitalized	<<  >>`,
		Id:         "Words",
		NTType:     2,
		Index:      2,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Words : Words capitalized	<<  >>`,
		Id:         "Words",
		NTType:     2,
		Index:      3,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Words : lowercase	<< ast.NewIdentifier(X[0], C) >>`,
		Id:         "Words",
		NTType:     2,
		Index:      4,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewIdentifier(X[0], C)
		},
	},
	ProdTabEntry{
		String:     `Words : Words lowercase	<< ast.NewIdentifier(X[1], C) >>`,
		Id:         "Words",
		NTType:     2,
		Index:      5,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewIdentifier(X[1], C)
		},
	},
}
