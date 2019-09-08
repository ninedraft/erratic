package gen

import "go/ast"

type FuncBuilder struct {
	*ast.FuncDecl
}

func (builder FuncBuilder) Doc(doc string) FuncBuilder {
	builder.FuncDecl.Doc.List = append(builder.FuncDecl.Doc.List, Doc(doc).List...)
	return builder
}

func (builder FuncBuilder) Ret(exprs ...ast.Expr) FuncBuilder {
	builder.FuncDecl.Body.List = []ast.Stmt{Return(exprs...)}
	return builder
}

func (builder FuncBuilder) Returns(tt ...ast.Expr) FuncBuilder {
	var ret = make([]*ast.Field, 0, len(tt))
	for _, t := range tt {
		ret = append(ret, &ast.Field{
			Type: t,
		})
	}
	builder.FuncDecl.Type.Results.List = ret
	return builder
}

func (builder FuncBuilder) Body(stmts ...ast.Stmt) FuncBuilder {
	builder.FuncDecl.Body.List = append(builder.FuncDecl.Body.List, stmts...)
	return builder
}

func (builder FuncBuilder) Build() *ast.FuncDecl {
	return builder.FuncDecl
}

func Method(name, recvName string, recvType ast.Expr) FuncBuilder {
	return FuncBuilder{
		FuncDecl: &ast.FuncDecl{
			Name: Ident(name),
			Recv: &ast.FieldList{
				List: []*ast.Field{Field(recvName, nil, nil, recvType)},
			},
			Body: &ast.BlockStmt{},
		},
	}
}

func Return(exprs ...ast.Expr) *ast.ReturnStmt {
	return &ast.ReturnStmt{
		Results: exprs,
	}
}
