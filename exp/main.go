package main

import (
	"html/template"
	"os"
)

type Test struct {
	Name          string
	Description   string
	IntValue      int64
	FloatValue    float64
	MultipleValue []string
	IsLoading     bool
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}
	data := Test{
		Name: "John Smith",
		Description: "lorem	ipsum is dfdf dfdf dfdfdfdfdfd",
		IntValue:      2345,
		FloatValue:    2334.232,
		MultipleValue: []string{"one", "two", "three"},
		IsLoading:     false,
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
