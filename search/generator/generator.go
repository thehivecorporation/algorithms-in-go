//go:generate ./generate.sh
package main

import (
	"log"
	"os"
	"text/template"

	algorithmsTmpl "github.com/thehivecorporation/algorithms-in-go/template"
)

var header string = `package search

`

func main() {
	toGenerate := []algorithmsTmpl.GenerateImplAndTest{
		{
			Impl:     binaryImpl,
			Test:     binaryTest,
			FileName: "binary",
		},
		{
			Impl:     sequentialImpl,
			Test:     sequentialTest,
			FileName: "sequential",
		},
	}

	generateWithTemplate(toGenerate)
}

func generateWithTemplate(toGenerate []algorithmsTmpl.GenerateImplAndTest) {
	for _, tmp := range toGenerate {
		codeFile, err := os.Create("../" + tmp.FileName + ".go")
		if err != nil {
			log.Fatal(err)
		}
		defer codeFile.Close()

		codeFile.WriteString(algorithmsTmpl.Warning)
		codeFile.WriteString(header)

		testFile, err := os.Create("../" + tmp.FileName + "_test.go")
		if err != nil {
			log.Fatal(err)
		}
		defer testFile.Close()

		testFile.WriteString(algorithmsTmpl.Warning)
		testFile.WriteString(header)
		tmpl, err := template.New("tmpl").Parse(tmp.Impl)
		if err != nil {
			log.Fatal(err)
		}

		for _, num := range algorithmsTmpl.NumericTypes {
			err = tmpl.Execute(codeFile, num)
			if err != nil {
				log.Fatal(err)
			}
		}

		tmpl, err = template.New("tmpl").Parse(tmp.Test)
		if err != nil {
			log.Fatal(err)
		}

		for _, num := range algorithmsTmpl.NumericTypes {
			err = tmpl.Execute(testFile, num)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
