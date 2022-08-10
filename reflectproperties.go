package main

import (
	"fmt"
	"reflect"
)

type person struct {
	name  string
	grade int
}

func (p *person) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(p)

	if reflectValue.Kind() == reflect.Ptr{
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("Nama = ", reflectType.Field(i).Name)
		fmt.Println("Tipe = ", reflectType.Field(i).Type)
        fmt.Println("nilai     :", reflectValue.Field(i).Interface())
	}
}

func main(){
	var p1 = &person{name:"riqki",grade:10}
	p1.getPropertyInfo()
}