package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/tkrajina/gojson2models/conv"
)

func main() {
	var packagePath, language, target string
	flag.StringVar(&packagePath, "package", "", "Path of the package with models")
	flag.StringVar(&language, "language", "", "Target language")
	flag.StringVar(&target, "target", "", "Target models file")
	flag.Parse()

	structs := []string{}
	for _, structOrGoFile := range flag.Args() {
		if strings.HasSuffix(structOrGoFile, ".go") {
			fmt.Println("Parsing:", structOrGoFile)
			fileStructs, err := GetGolangFileStructs(structOrGoFile)
			if err != nil {
				panic(fmt.Sprintf("Error loading/parsing golang file %s: %s", structOrGoFile, err.Error()))
			}
			for _, s := range fileStructs {
				structs = append(structs, s)
			}
		} else {
			structs = append(structs, structOrGoFile)
		}
	}

	if len(packagePath) == 0 {
		fmt.Fprintln(os.Stderr, "No package given")
		flag.Usage()
		os.Exit(1)
	}
	if len(target) == 0 {
		fmt.Fprintln(os.Stderr, "No target file")
		flag.Usage()
		os.Exit(1)
	}
	if len(language) == 0 {
		fmt.Fprintln(os.Stderr, "No target language")
		flag.Usage()
		os.Exit(1)
	}

	packageParts := strings.Split(packagePath, string(os.PathSeparator))
	pckg := packageParts[len(packageParts)-1]

	filename, err := ioutil.TempDir(os.TempDir(), "")
	panicIfErr(err)

	filename = fmt.Sprintf("%s/gojson2models_%d.go", filename, time.Now().Nanosecond())

	f, err := os.Create(filename)
	panicIfErr(err)
	defer f.Close()

	structsArr := make([]string, 0)
	for _, str := range structs {
		str = strings.TrimSpace(str)
		if len(str) > 0 {
			structsArr = append(structsArr, pckg+"."+str)
		}
	}

	params := conv.ConverterParams{
		Structs:        structsArr,
		ModelsPackage:  packagePath,
		TargetLanguage: language,
		TargetFile:     target,
	}
	code := conv.T__converter_script(params)
	f.Write([]byte(code))

	cmd := exec.Command("go", "run", filename)
	fmt.Println(strings.Join(cmd.Args, " "))
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(output))
		panicIfErr(err)
	}
	fmt.Println(string(output))
}

func GetGolangFileStructs(filename string) ([]string, error) {
	fset := token.NewFileSet() // positions are relative to fset

	f, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		return nil, err
	}

	v := &StructVisitor{}
	ast.Walk(v, f)

	return v.structs, nil
}

type StructVisitor struct {
	structNameCandidate string
	structs             []string
}

func (v *StructVisitor) Visit(node ast.Node) ast.Visitor {
	if node != nil {
		switch t := node.(type) {
		case *ast.Ident:
			v.structNameCandidate = t.Name
		case *ast.StructType:
			if len(v.structNameCandidate) > 0 {
				v.structs = append(v.structs, v.structNameCandidate)
				v.structNameCandidate = ""
			}
		default:
			fmt.Println("t=", t)
			v.structNameCandidate = ""
		}
	}
	return v
}

func panicIfErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
