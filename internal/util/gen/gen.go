package gen

import (
	"github.com/Desuuuu/gocc/internal/util/gen/golang"
)

func Gen(outDir string) {
	golang.GenRune(outDir)
	golang.GenLitConv(outDir)
}
