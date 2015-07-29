package main

import (
	"fmt"
	// "github.com/scgolang/sc"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
)

func usage(w io.Writer) {
	fmt.Fprintf(w, "Usage: %s SYNTHDEF_FILE\n", os.Args[0])
}

func main() {
	if len(os.Args) != 2 {
		usage(os.Stderr)
		os.Exit(1)
	}
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, os.Args[1], nil, parser.AllErrors)
	if err != nil {
		panic(err)
	}
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			fmt.Printf("found function %s\n", x.Name.Name)
			fmt.Printf("data           %v\n", x.Name.Obj.Data)
			return false
		}
		return true
	})
}
