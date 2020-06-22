package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ghodss/yaml"

	"github.com/open-policy-agent/frameworks/constraint/pkg/apis/templates/v1beta1"
	"github.com/spf13/cobra"
	r "github.com/xUnholy/konstraint/pkg/rego"
)

var (
	policyFilePath = "policy"
	libFilePath    = "libs"
	outputType     = "json"
)

func TemplateCli() *cobra.Command {
	var templateCmd = &cobra.Command{
		Use:   "template",
		Short: "",
		Run:   templateCmd,
	}
	templateCmd.Flags().StringVarP(&policyFilePath, "policy", "p", policyFilePath, "Path to the Rego policy files directory. For the test command, specifying a specific .rego file is allowed. (default \"policy\")")
	templateCmd.Flags().StringVarP(&libFilePath, "libs", "l", libFilePath, "Path to the Rego library files directory. For the test command, specifying a specific .rego file is allowed.")
	templateCmd.Flags().StringVarP(&outputType, "output", "o", outputType, "Output format. One of: json|yaml (default \"json\")")
	return templateCmd
}

func templateCmd(cmd *cobra.Command, args []string) {
	// Determine if Path or File for Policies
	policyFiles := []string{policyFilePath}
	if IsDirectory(policyFilePath) {
		policyFiles = GetRegoFiles(policyFilePath)
	}
	// Determine if Path or File for Libraries
	libraryFiles := []string{libFilePath}
	if IsDirectory(libFilePath) {
		libraryFiles = GetRegoFiles(libFilePath)
	}
	// Create map[string]string of library name:rego
	libsMap := make(map[string]string)
	for i := range libraryFiles {
		rego := GetFileContent(libraryFiles[i])
		name := ParseRegoLibrary(rego)
		libsMap[name] = rego
	}
	// Iterate through Policies Rego
	for i := range policyFiles {
		rego := GetFileContent(policyFiles[i])
		name, libs := ParseRegoPolicy(rego)
		if strings.Contains(name, "lib.") {
			continue
		}
		fLibs := []string{}
		for l := range libs {
			if val, ok := libsMap[libs[l]]; ok {
				fLibs = append(fLibs, val)
			}
		}
		WriteFile(policyFiles[i], t.ConstraintTemplate(name, fLibs, rego))
	}
}

func ParseRegoPolicy(rego string) (string, []string) {
	regoAst := r.Parse(rego)
	libs := r.GetImports(regoAst)
	name := r.GetPackageName(regoAst)
	return name, libs
}

func WriteFile(path string, out *v1beta1.ConstraintTemplate) {
	var err error
	var file []byte
	output := strings.ToLower(outputType)
	if output == "yaml" {
		file, err = yaml.Marshal(&out)
		if err != nil {
			log.Fatal(err)
		}
	} else if output == "json" {
		file, err = json.MarshalIndent(&out, "", "	")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatalln("Error: Output format is not supported. Must be one of json|yaml")
	}
	err = ioutil.WriteFile(fmt.Sprintf("%v.%v", path, output), file, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func GetFileContent(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("unable to open path: %v", path)
	}
	rego, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", file)
	}
	return string(rego)
}

func ParseRegoLibrary(rego string) string {
	regoAst := r.Parse(rego)
	if regoAst.Imports != nil {
		fmt.Println(regoAst.Imports)
		log.Fatal("Error: Libraries importing other packages is not supported")
	}
	return r.GetPackageName(regoAst)
}

func GetRegoFiles(path string) []string {
	regoFileExt := ".rego"
	regoTestFileSuffix := "_test"
	files := []string{}
	err := filepath.Walk(path,
		func(path string, f os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if filepath.Ext(path) == regoFileExt {
				if !strings.Contains(f.Name(), regoTestFileSuffix) {
					files = append(files, path)
				} else {
					fmt.Println("Ignoring test file:", f.Name())
				}
			}
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func IsDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}
	return fileInfo.IsDir()
}
