package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string "namestr"
	age  int    "ageint"
}

func ShowTag(i interface{}) {
	switch t := reflect.TypeOf(i); t.Kind() {
	case reflect.Ptr:
		tag := t.Elem().Field(0).Tag
		fmt.Println(tag)
	}
}

func main() {
	x := Person{"HAHA", 13}
	ShowTag(&x)
}
