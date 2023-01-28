package reflection

import (
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	// walk the interface, find all strings, run the func on all the found strings
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}
