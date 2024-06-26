package astx

import (
	"fmt"
	"testing"

	"github.com/Desuuuu/gocc/example/errorrecovery/ast"
	"github.com/Desuuuu/gocc/example/errorrecovery/errors"
	"github.com/Desuuuu/gocc/example/errorrecovery/lexer"
	"github.com/Desuuuu/gocc/example/errorrecovery/parser"
)

func TestFail(t *testing.T) {
	sml, err := test([]byte("a b ; d e f"))
	if err != nil {
		t.Fail()
	}
	fmt.Print("output: [\n")
	for _, s := range sml {
		switch sym := s.(type) {
		case *errors.Error:
			fmt.Printf("%s\n", sym)
		default:
			fmt.Printf("\t%v\n", sym)
		}
	}
	fmt.Println("]")
}

func test(src []byte) (astree ast.StmtList, err error) {
	fmt.Printf("input: %s\n", src)
	s := lexer.NewLexer(src)
	p := parser.NewParser()
	a, err := p.Parse(s)
	if err == nil {
		astree = a.(ast.StmtList)
	}
	return
}
