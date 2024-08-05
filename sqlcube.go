package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// StructField represents a field in a struct
type StructField struct {
	FieldName string
	FieldType string
}

// StructInfo stores the struct name and its fields
type StructInfo struct {
	StructName string
	Fields     []StructField
}

func (s *StructInfo) Hash() string {
	hasher := sha256.New()
	var fieldStrings []string
	for _, field := range s.Fields {
		fieldStrings = append(fieldStrings, field.FieldName+"|"+field.FieldType)
	}
	joinedFields := strings.Join(fieldStrings, ",")

	hasher.Write([]byte(joinedFields))
	return string(hasher.Sum(nil))
}

// changeStructAndPrint parses a Go file and extracts structs and their fields and
func changeStructAndPrint(filename string, targetDir string, mp map[string]string, debug bool, fullRewrite bool) (map[string]string, error) {
	if debug {
		fmt.Println(filename)
	}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return mp, err
	}

	var hasChanged = fullRewrite

	ast.Inspect(node, func(n ast.Node) bool {
		// Check for struct type definitions
		typ, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}

		structType, ok := typ.Type.(*ast.StructType)
		if !ok {
			return true
		}

		structInfo := StructInfo{StructName: typ.Name.Name}

		for _, field := range structType.Fields.List {
			fieldType := fmt.Sprintf("%s", GetFieldType(field.Type))

			for _, fieldName := range field.Names {
				structInfo.Fields = append(structInfo.Fields, StructField{
					FieldName: fieldName.Name,
					FieldType: fieldType,
				})
			}
		}

		hash := structInfo.Hash()
		typeName, ok := mp[hash]
		if !ok {
			mp[hash] = structInfo.StructName
		} else {
			if debug {
				fmt.Println("\tAlias found", structInfo.StructName, "EQUALS", typeName)
			}
			hasChanged = true
			typ.Assign = token.Pos(token.ASSIGN)
			typ.Type = &ast.Ident{Name: typeName}
		}

		return false
	})

	if hasChanged {
		buf := new(bytes.Buffer)

		if err := printer.Fprint(buf, fset, node); err != nil {
			return mp, err
		}

		formatedBytes, err := format.Source(buf.Bytes())
		if err != nil {
			return mp, err
		}

		_, fileName := path.Split(filename)

		if fullRewrite {
			// fullRewrite is true when targetFolder is different to src folder
			f, err := os.Create(path.Join(targetDir, fileName))
			if err != nil {
				return mp, err
			}
			_, err = f.Write(formatedBytes)
		} else {
			err = os.WriteFile(filename, formatedBytes, 0644)
		}
	}
	return mp, err
}

// CreateTypeAlias walks through a directory and extracts structs from all Go files
// it matches Equal struct and creates Type Alias
func CreateTypeAlias(srcDir string, targetDir string, debug bool) error {
	var fullRewrite bool
	if srcDir != targetDir {
		fullRewrite = true
	}

	mp := map[string]string{}
	modelsFilePath := filepath.Join(srcDir, "models.go")
	// Process models.go first
	mp, err := changeStructAndPrint(modelsFilePath, targetDir, mp, debug, fullRewrite)
	if err != nil {
		fmt.Println(err)
	}

	err = filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == modelsFilePath {
			return nil
		}
		if !info.IsDir() && filepath.Ext(path) == ".go" {
			mp, err = changeStructAndPrint(path, targetDir, mp, debug, fullRewrite)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func GetFieldType(field ast.Expr) string {
	switch t := field.(type) {
	case *ast.InterfaceType:
		return "interface{}"
	case *ast.ArrayType:
		return fmt.Sprintf("[]%s", GetFieldType(t.Elt))
	default:
		return fmt.Sprintf("%s", field)
	}
}
