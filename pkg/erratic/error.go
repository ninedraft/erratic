package erratic

import (
	"go/ast"

	"github.com/ninedraft/erratic/pkg/gen"
)

type Error struct {
	Name     string  `json:"name"`
	Doc      string  `json:"doc"`
	Code     int     `json:"code"`
	Message  string  `json:"message"`
	UnwrapOn string  `json:"unwrap_on"`
	Fields   []Field `json:"fields" toml:"field"`
}

func (err *Error) Declarations() []ast.Decl {
	return []ast.Decl{
		err.MessageTemplate(),
		err.Type(),
		err.JSONhelper(),
		err.MethodError(),
	}
}

func (err *Error) MessageTemplate() *ast.GenDecl {
	return gen.Var("_tmpl"+err.Name, gen.Doc(""), nil)
}

func (err *Error) Type() *ast.GenDecl {
	var fields = make([]*ast.Field, 0, len(err.Fields))
	for _, field := range err.Fields {
		fields = append(fields, field.Decl())
	}
	return gen.StructTypeDecl(err.Name, gen.Doc(err.Name), fields...)
}

func (err *Error) JSONhelper() *ast.GenDecl {
	var fields = make([]*ast.Field, 0, len(err.Fields)+2)
	fields = append(fields,
		gen.Field("ErrorCode",
			gen.Doc("ErrorCode contains represents unique error ID."),
			gen.DefaultTags("ErrorCode").Lit(),
			gen.Int64(),
		),
		gen.Field("Message",
			gen.Doc("Message contains text error representation."),
			gen.DefaultTags("Message").Lit(),
			gen.String(),
		),
	)
	for _, field := range err.Fields {
		fields = append(fields, field.Decl())
	}
	return gen.StructTypeDecl("_json"+err.Name, nil, fields...)
}

func (err *Error) MethodError() *ast.FuncDecl {
	return gen.Method("Error", "err", gen.Ident(err.Name)).
		Returns(gen.Ident(err.Name)).
		Ret(gen.StringLit("uniplemented")).
		Build()
}

type Field struct {
	Name    string            `json:"name"`
	Type    string            `json:"type"`
	Tags    map[string]string `json:"tags"`
	Default interface{}       `json:"default"`
	Doc     string            `json:"doc"`
}

func (field *Field) Decl() *ast.Field {
	return gen.Field(field.Name,
		gen.Doc(field.Doc),
		gen.Tags(field.Tags).Lit(),
		gen.Ident(field.Type),
	)
}
