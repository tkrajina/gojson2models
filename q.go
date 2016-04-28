package main

import (
	"github.com/tkrajina/typescriptify-golang-structs/example"
	"github.com/tkrajina/typescriptify-golang-structs/jsonconv"
)

func main() {
	converter := jsonconv.NewEntityParser()
	converter.Add(example.Person{})
	err := converter.Parse()
	if err != nil {
		panic(err.Error())
	}
	err = converter.ConvertToJava("test.java")
	if err != nil {
		panic(err.Error())
	}
	err = converter.ConvertToTypescript("test.ts")
	if err != nil {
		panic(err.Error())
	}
}
