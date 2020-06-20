package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"strings"

	"github.com/open-policy-agent/frameworks/constraint/pkg/apis/templates/v1beta1"
	"github.com/open-policy-agent/opa/ast"
	"github.com/spf13/cobra"
	"github.com/xUnholy/konstraint/internal/template"
)

func TemplateCli() *cobra.Command {
	var templateCmd = &cobra.Command{
		Use:   "template",
		Short: "",
		Run:   templateCmd,
	}
	return templateCmd
}

func templateCmd(cmd *cobra.Command, args []string) {
	for a := range args {
		path := args[a]
		if Exists(path) {
			rego, err := GetFileContent(path)
			if err != nil {
				log.Fatalln(err)
			}
			name, libs, err := ParseRego(rego)
			if err != nil {
				log.Fatalln(err)
			}
			WriteFileToYaml(path, template.GenerateConstraintTemplate(name, libs, rego))
		}
	}
}

func WriteFileToJson(path string, out v1beta1.ConstraintTemplate) {
	file, _ := json.MarshalIndent(out, "", "	")
	_ = ioutil.WriteFile(fmt.Sprintf("%v.json", path), file, 0644)
}


func WriteFileToYaml(path string, out v1beta1.ConstraintTemplate) {
	file, _ := yaml.Marshal(out)
	_ = ioutil.WriteFile(fmt.Sprintf("%v.yaml", path), file, 0644)
}

func GetFileContent(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("unable to open path: %v", path)
	}

	rego, err := ioutil.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("unable to read file: %v", file)
	}
	return string(rego), nil
}

func ParseRego(rego string) (string, []string, error) {
	libs := []string{}
	regoAst := ast.MustParseModule(rego)
	if regoAst == nil {
		return "", []string{}, fmt.Errorf("unable to parse rego file")
	}

	for i := range regoAst.Imports {
		libs = append(libs, strings.Trim(fmt.Sprintf("%v",regoAst.Imports[i]), "import data."))
	}

	name := strings.Trim(fmt.Sprintf("%v", regoAst.Package.Path), "data.")
	return name, libs, nil
}

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
			if os.IsNotExist(err) {
					return false
			}
	}
	return true
}
