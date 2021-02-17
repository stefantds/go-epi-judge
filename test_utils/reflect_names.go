package test_utils

import (
	"reflect"
	"runtime"
	"strings"
)

// GetFuncName returns the name of the function (only the last part, without the package part)
func GetFuncName(function interface{}) string {
	fValue := reflect.ValueOf(function)
	if fValue.Kind() != reflect.Func {
		panic("expected a function")
	}

	funcFullName := runtime.FuncForPC(fValue.Pointer()).Name()
	funcNameParts := strings.Split(funcFullName, ".")
	return funcNameParts[len(funcNameParts)-1]
}

// GetTypeName returns the name of a struct type (also accepts a pointer to struct)
func GetTypeName(val interface{}) string {
	t := reflect.TypeOf(val)
	if t.Kind() != reflect.Ptr && t.Kind() != reflect.Struct {
		panic("expected a struct or pointer to struct")
	}

	if t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	}
	return t.Name()
}
