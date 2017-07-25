// Code generated by gocc; DO NOT EDIT.

package lexer

import (
	// "fmt"
	"io/ioutil"
	"unicode/utf8"

	// "github.com/goccmack/gocc/example/mail/util"
	"github.com/goccmack/gocc/example/mail/token"
)

const (
	NoState    = -1
	NumStates  = 15
	NumSymbols = 33
)

type Lexer struct {
	src    []byte
	pos    int
	line   int
	column int
}

func NewLexer(src []byte) *Lexer {
	lexer := &Lexer{
		src:    src,
		pos:    0,
		line:   1,
		column: 1,
	}
	return lexer
}

func NewLexerFile(fpath string) (*Lexer, error) {
	src, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	return NewLexer(src), nil
}

func (l *Lexer) Scan() (tok *token.Token) {
	// fmt.Printf("Lexer.Scan() pos=%d\n", l.pos)
	tok = new(token.Token)
	if l.pos >= len(l.src) {
		tok.Type = token.EOF
		tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = l.pos, l.line, l.column
		return
	}
	start, startLine, startColumn, end := l.pos, l.line, l.column, 0
	tok.Type = token.INVALID
	state, rune1, size := 0, rune(-1), 0
	for state != -1 {
		// fmt.Printf("\tpos=%d, line=%d, col=%d, state=%d\n", l.pos, l.line, l.column, state)
		if l.pos >= len(l.src) {
			rune1 = -1
		} else {
			rune1, size = utf8.DecodeRune(l.src[l.pos:])
			l.pos += size
		}
		// Production start
		if rune1 != -1 {
			state = TransTab[state](rune1)
		} else {
			state = -1
		}
		// Production end

		// Debug start
		// nextState := -1
		// if rune1 != -1 {
		// 	nextState = TransTab[state](rune1)
		// }
		// fmt.Printf("\tS%d, : tok=%s, rune == %s(%x), next state == %d\n", state, token.TokMap.Id(tok.Type), util.RuneToString(rune1), rune1, nextState)
		// fmt.Printf("\t\tpos=%d, size=%d, start=%d, end=%d\n", l.pos, size, start, end)
		// if nextState != -1 {
		// 	fmt.Printf("\t\taction:%s\n", ActTab[nextState].String())
		// }
		// state = nextState
		// Debug end

		if state != -1 {

			switch rune1 {
			case '\n':
				l.line++
				l.column = 1
			case '\r':
				l.column = 1
			case '\t':
				l.column += 4
			default:
				l.column++
			}

			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				// fmt.Printf("\t Accept(%s), %s(%d)\n", string(act), token.TokMap.Id(tok), tok)
				end = l.pos
			case ActTab[state].Ignore != "":
				// fmt.Printf("\t Ignore(%s)\n", string(act))
				start, startLine, startColumn = l.pos, l.line, l.column
				state = 0
				if start >= len(l.src) {
					tok.Type = token.EOF
				}

			}
		} else {
			if tok.Type == token.INVALID {
				end = l.pos
			}
		}
	}
	if end > start {
		l.pos = end
		tok.Lit = l.src[start:end]
	} else {
		tok.Lit = []byte{}
	}
	tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = start, startLine, startColumn
	// fmt.Printf("Token at %s: %s \"%s\"\n", tok.String(), token.TokMap.Id(tok.Type), tok.Lit)

	return
}

func (l *Lexer) Reset() {
	l.pos = 0
}

/*
Lexer symbols:
0: '@'
1: '!'
2: '#'
3: '$'
4: '%'
5: '&'
6: '''
7: '*'
8: '+'
9: '-'
10: '/'
11: '='
12: '?'
13: '^'
14: '_'
15: '`'
16: '{'
17: '|'
18: '}'
19: '~'
20: '.'
21: '\'
22: '"'
23: '"'
24: '\t'
25: '\n'
26: '\r'
27: ' '
28: 'A'-'Z'
29: 'a'-'z'
30: '0'-'9'
31: \u0100-\U0010ffff
32: .
*/
