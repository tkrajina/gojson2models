package conv

import (
	"bytes"
	"errors"
	"fmt"
	"html"
	"os"
	"strings"
)

func init() {
	_ = fmt.Sprintf
	_ = errors.New
	_ = os.Stderr
	_ = html.EscapeString
}

// Generated code, do not edit!!!!

func TE__converter_script(args ConverterParams) (string, error) {
	__template__ := "converter_script.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	/* package main */
	result.WriteString(`package main
`)
	/*  */
	result.WriteString(`
`)
	/* import ( */
	result.WriteString(`import (
`)
	/* "{{=s args.ModelsPackage }}" */
	result.WriteString(fmt.Sprintf(`    "%s"
`, args.ModelsPackage))
	/* "github.com/tkrajina/gojson2models/jsonconv" */
	result.WriteString(`	"github.com/tkrajina/gojson2models/jsonconv"
`)
	/* ) */
	result.WriteString(`)
`)
	/*  */
	result.WriteString(`
`)
	/* func panicIfErr(err error) { */
	result.WriteString(`func panicIfErr(err error) {
`)
	/* if err != nil { */
	result.WriteString(`	if err != nil {
`)
	/* panic(err.Error()) */
	result.WriteString(`		panic(err.Error())
`)
	/* } */
	result.WriteString(`	}
`)
	/* } */
	result.WriteString(`}
`)
	/*  */
	result.WriteString(`
`)
	/* func main() { */
	result.WriteString(`func main() {
`)
	/* converter := jsonconv.New() */
	result.WriteString(`	converter := jsonconv.New()
`)
	/* !for _, str := range args.Structs */
	for _, str := range args.Structs {
		/* converter.Add({{=s str }}{}) */
		result.WriteString(fmt.Sprintf(`	converter.Add(%s{})
`, str))
		/* !end */
	}
	/* panicIfErr(converter.Parse()) */
	result.WriteString(`	panicIfErr(converter.Parse())
`)
	/* panicIfErr(converter.ConvertTo{{=s strings.Title(args.TargetLanguage) }}("{{=s args.TargetFile }}")) */
	result.WriteString(fmt.Sprintf(`	panicIfErr(converter.ConvertTo%s("%s"))
`, strings.Title(args.TargetLanguage), args.TargetFile))
	/* } */
	result.WriteString(`}
`)
	/*  */
	result.WriteString(``)

	return result.String(), nil
}

func T__converter_script(args ConverterParams) string {
	html, err := TE__converter_script(args)
	if err != nil {
		os.Stderr.WriteString("Error running template converter_script.tmpl:" + err.Error())
	}
	return html
}
