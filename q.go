package main

import (
	"fmt"
	"reflect"
)

func panicIfErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

type Address struct {
}

type Person struct {
	Bu struct {
		Aaa string `json:"aaa"`
	} `json:"bu"`
	Add Address `json:"addr"`
}

func main() {
	fields := deepFields(reflect.TypeOf(Person{}))
	for _, field := range fields {
		fmt.Println(field.Name, field.Type.Name())
	}
}

func deepFields(typeOf reflect.Type) []reflect.StructField {
	fields := make([]reflect.StructField, 0)

	if typeOf.Kind() == reflect.Ptr {
		typeOf = typeOf.Elem()
	}

	if typeOf.Kind() != reflect.Struct {
		return fields
	}

	for i := 0; i < typeOf.NumField(); i++ {
		f := typeOf.Field(i)

		kind := f.Type.Kind()
		if f.Anonymous && kind == reflect.Struct {
			//fmt.Println(v.Interface())
			fields = append(fields, deepFields(f.Type)...)
		} else {
			fields = append(fields, f)
		}
	}

	return fields
}
