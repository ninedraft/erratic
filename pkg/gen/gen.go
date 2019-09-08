package gen

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	"sort"
	"strings"
)

// Ident returns an ident with nil object and provided name.
func Ident(name string) *ast.Ident {
	return &ast.Ident{Name: name}
}

// StructTypeDecl returns a struct type declaration with provided fields declarations.
func StructTypeDecl(name string, doc *ast.CommentGroup, fields ...*ast.Field) *ast.GenDecl {
	var spec = &ast.TypeSpec{
		Name: Ident(name),
		Doc:  doc,
		Type: &ast.StructType{
			Fields: &ast.FieldList{
				List: fields,
			},
		},
	}
	return &ast.GenDecl{
		Specs: []ast.Spec{spec},
	}
}

// StructLit returns a composite literal, which represents a struct literal with provided type name.
func StructLit(name string, fields ...*ast.KeyValueExpr) *ast.CompositeLit {
	var fieldsExpr = make([]ast.Expr, 0, len(fields))
	for _, field := range fields {
		fieldsExpr = append(fieldsExpr, field)
	}
	return &ast.CompositeLit{
		Type: Ident(name),
		Elts: fieldsExpr,
	}
}

// FieldExpr returns a key-value expression for struct literal.
func FieldExpr(name string, value ast.Expr) *ast.KeyValueExpr {
	return &ast.KeyValueExpr{
		Key:   Ident(name),
		Value: value,
	}
}

// Doc returns a comment group, created from provided text, splitted by "\n".
func Doc(text string) *ast.CommentGroup {
	var lines = strings.SplitN(text, "\n", -1)
	return DocLines(lines)
}

// DocLines returns a comment group, created from provided lines.
func DocLines(lines []string) *ast.CommentGroup {
	var comments = make([]*ast.Comment, 0, len(lines))
	for _, line := range lines {
		comments = append(comments, &ast.Comment{
			Text: line,
		})
	}
	return &ast.CommentGroup{
		List: comments,
	}
}

// Tags represents canonical struct field tags key-value pairs.
// Example: `json:"field-name" toml:"field"`
type Tags map[string]string

// Lit returns a basic literal of kind STRING with canonicaly encoded
// key-value pairs.
func (tags Tags) Lit() *ast.BasicLit {
	if len(tags) > 0 {
		var tagLit = bytes.NewBufferString("`")
		for key, val := range tags {
			_, _ = fmt.Fprintf(tagLit, " %s: %q", key, val)
		}
		_, _ = tagLit.WriteString("`")
		return &ast.BasicLit{
			Kind:  token.STRING,
			Value: tagLit.String(),
		}
	}
	return nil
}

// DefaultTags returns default tags for json and flags keys.
func DefaultTags(fieldName string) Tags {
	return Tags{
		"json": ToKebabCase(fieldName),
		"flag": ToSnakeCase(fieldName),
	}
}

// Field returns a field declaration with provided name, doc, tags and type.
func Field(name string, doc *ast.CommentGroup, tags *ast.BasicLit, ft ast.Expr) *ast.Field {
	return &ast.Field{
		Names: []*ast.Ident{Ident(name)},
		Doc:   doc,
		Type:  ft,
		Tag:   tags,
	}
}

// Fields contains represents key-value field value declarations.
type Fields map[string]ast.Expr

// Names returns fields sorted names.
func (fields Fields) Names() []string {
	var names = make([]string, 0, len(fields))
	for name := range fields {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// AST returns slice of key-value field experessions, sorted by names.
func (fields Fields) AST() []*ast.KeyValueExpr {
	var kv = make([]*ast.KeyValueExpr, 0, len(fields))
	for _, name := range fields.Names() {
		kv = append(kv, FieldExpr(name, fields[name]))
	}
	return kv
}

func Var(name string, doc *ast.CommentGroup, value ast.Expr) *ast.GenDecl {
	return &ast.GenDecl{
		Doc: doc,
		Tok: token.VAR,
		Specs: []ast.Spec{
			&ast.ValueSpec{
				Names:  []*ast.Ident{Ident(name)},
				Values: []ast.Expr{value},
			},
		},
	}
}
