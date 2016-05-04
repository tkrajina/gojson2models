package main

import (
	"github.com/tkrajina/gojson2models/example"
	"github.com/tkrajina/gojson2models/jsonconv"
	"reflect"
)

func panicIfErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	converter := jsonconv.New()
	converter.Add(example.Person{})
	converter.AddType(reflect.TypeOf(example.Person{}.Bu))
	panicIfErr(converter.Parse())
	panicIfErr(converter.ConvertToJava("example/example.java"))
}
