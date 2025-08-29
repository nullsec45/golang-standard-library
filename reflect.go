package main

import (
	"fmt"
	"reflect"
)


type Sample struct {
	Name string
}

type Person struct {
	Name, Address, Email string
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type Name", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var valueField reflect.StructField = valueType.Field(i)
		fmt.Printf(valueField.Name, " with type ", valueField.Type)
	}
}
func main(){

	readField(Person{"Rama","Indonesia","anon@gmail.com"})
	readField(Sample{"Fajar"})
}