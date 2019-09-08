package gen

import "go/ast"

type Import struct {
	Name string
	Path string
}

func PlainImport(path string) Import {
	return Import{Path: path}
}

func (pkg Import) S(name string) ast.Node {
	return &ast.SelectorExpr{
		X:   Ident(pkg.Name),
		Sel: Ident(name),
	}
}

func (pkg Import) AST() *ast.ImportSpec {
	return &ast.ImportSpec{
		Name: Ident(pkg.Name),
		Path: StringLit(pkg.Path),
	}
}
