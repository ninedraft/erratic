package gen

import (
	"go/ast"
	"go/token"
	"strconv"
)

// String returns a string identifier.
func String() *ast.Ident {
	return Ident("string")
}

func StringLit(str string) *ast.BasicLit {
	return &ast.BasicLit{Kind: token.STRING, Value: strconv.Quote(str)}
}

// Error returns a string identifier.
func Error() *ast.Ident {
	return Ident("error")
}

// Int returns a string identifier.
func Int() *ast.Ident {
	return Ident("int")
}

// Int8 returns a string identifier.
func Int8() *ast.Ident {
	return Ident("int8")
}

// Int16 returns a string identifier.
func Int16() *ast.Ident {
	return Ident("int16")
}

// Int32 returns a string identifier.
func Int32() *ast.Ident {
	return Ident("int32")
}

// Int64 returns a string identifier.
func Int64() *ast.Ident {
	return Ident("int64")
}

// Uint returns a string identifier.
func Uint() *ast.Ident {
	return Ident("uint")
}

// Uint8 returns a string identifier.
func Uint8() *ast.Ident {
	return Ident("uint8")
}

// Uint16 returns a string identifier.
func Uint16() *ast.Ident {
	return Ident("uint16")
}

// Uint32 returns a string identifier.
func Uint32() *ast.Ident {
	return Ident("uint32")
}

// Uint64 returns a string identifier.
func Uint64() *ast.Ident {
	return Ident("uint64")
}
