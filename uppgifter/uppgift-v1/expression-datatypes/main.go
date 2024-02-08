package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println("2021, ", reflect.TypeOf("2021"))
	fmt.Println(2021, reflect.TypeOf(2021))
	fmt.Println("2021-12-24, ", reflect.TypeOf("2021-12-24"))
	fmt.Println(2021-12-24, reflect.TypeOf(2021-12-24)) // Date isn't formated so it substract due to contain "-" => 2021-12-24=1985
}
