package erratic

import (
	"go/ast"

	"github.com/ninedraft/erratic/pkg/gen"
)

var pkgTemplate = &gen.Import{
	Name: "template",
	Path: "text/template",
}

type Package struct {
	Name   string  `json:"package" yaml:"package" toml:"package"`
	Errors []Error `json:"errors" toml:"error"`
}

func (pkg *Package) AST() *ast.File {
	var decls = make([]ast.Decl, 0, len(new(Error).Declarations())*len(pkg.Errors))
	for _, err := range pkg.Errors {
		decls = append(decls, err.Declarations()...)
	}
	return &ast.File{
		Name:    gen.Ident(pkg.Name),
		Imports: pkg.Imports(),
		Decls:   decls,
	}
}

func (pkg *Package) Imports() []*ast.ImportSpec {
	return []*ast.ImportSpec{
		pkgTemplate.AST(),
	}
}
