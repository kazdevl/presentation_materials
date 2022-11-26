package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

func main() {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "target/target.go", nil, 0)
	if err != nil {
		/* エラー処理 */
		log.Fatal(err)
	}
	cfg := &types.Config{Importer: importer.Default()}
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
	}
	pkg, err := cfg.Check("", fset, []*ast.File{file}, info)
	if err != nil {
		/* エラー処理 */
		log.Fatal(err)
	}
	fmt.Println("package is", pkg.Path())
	for e, t := range info.Types {
		// Collect the underlying types.
		fmt.Printf("type=%+v, underlying=%+v\n", t.Type.String(), t.Type.Underlying())
		// Collect structs to determine the fields of a receiver.
		if v, ok := t.Type.(*types.Struct); ok {
			if v.NumFields() == 0 {
				continue
			}
			fmt.Printf("struct=%+v, e=%+v\n", v.Field(0).Type().Underlying(), e)
		}
		fmt.Println("-----------------")
	}
}
