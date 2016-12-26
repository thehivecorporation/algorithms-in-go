//go:generate ./generate.sh
package main

import (
	"os"
	"text/template"

	"log"

	algorithmsTmpl "github.com/thehivecorporation/algorithms-in-go/template"
)

var header string = `package common

`

func main() {
	toGenerate := []algorithmsTmpl.GenerateImplAndTest{
		{
			Impl: swapTmpl,
			Test: swapTest,
		},
		{
			Impl: isLessTmpl,
			Test: isLessTest,
		},
	}

	generateWithTemplate(toGenerate, "../common")
}

func generateWithTemplate(toGenerate []algorithmsTmpl.GenerateImplAndTest, filename string) {
	codeFile, err := os.Create(filename + ".go")
	if err != nil {
		log.Fatal(err)
	}
	defer codeFile.Close()

	codeFile.WriteString(algorithmsTmpl.Warning)
	codeFile.WriteString(header)

	testFile, err := os.Create(filename + "_test.go")
	if err != nil {
		log.Fatal(err)
	}
	defer testFile.Close()

	testFile.WriteString(algorithmsTmpl.Warning)
	testFile.WriteString(header)

	for _, tmp := range toGenerate {
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
