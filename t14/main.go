package main

import (
	"fmt"
	"reflect"
)

func main() {
	// создаем переменные типа interface{}
	var intValue interface{} = 5
	var stringValue interface{} = "hello!"
	var boolValue interface{} = true
	var chanValue interface{} = make(chan int)
	// печатаем их типы
	PrintType(intValue)
	PrintType(stringValue)
	PrintType(boolValue)
	PrintType(chanValue)
}

// функция для печатания типов переменных
func PrintType(value interface{}) {
	// используем функцию из пакета reflect
	t := reflect.TypeOf(value)
	fmt.Println(t)
}
