package main

import "reflect"

func detection(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	default:
		if t := reflect.TypeOf(v); t != nil && t.Kind() == reflect.Chan {
			return "chan"
		}
		return "unknown"
	}
}

func main() {
	var (
		i       = 42
		s       = "hello"
		b       = true
		ch      chan int
		interfc interface{}
	)
	println(detection(i))
	println(detection(s))
	println(detection(b))
	println(detection(ch))
	println(detection(interfc))
}
