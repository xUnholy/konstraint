package rego

import (
	"fmt"
	"log"
	"strings"

	"github.com/open-policy-agent/opa/ast"
)

func GetPackageName(rego *ast.Module) string {
	return strings.TrimPrefix(fmt.Sprintf("%v", rego.Package.Path), "data.")
}

func GetImports(rego *ast.Module) []string {
	imports := []string{}
	for i := range rego.Imports {
		imports = append(imports, strings.TrimPrefix(fmt.Sprintf("%v", rego.Imports[i]), "import data."))
	}
	return imports
}

func Parse(rego string) *ast.Module {
	regoAst := ast.MustParseModule(rego)
	if regoAst == nil {
		log.Fatal("unable to parse rego file")
	}
	return regoAst
}
