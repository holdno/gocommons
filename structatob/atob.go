package structatob

import (
	"errors"
	"reflect"
	"strings"
)

type User struct {
	Name string `a:"name"`
}

type Man struct {
	Car int `a:"name"`
}

var tagname = "json"

var (
	ErrNotPoint  = errors.New("target param is not a point type")
	ErrNotStruct = errors.New("target param is not a struct")
)

func SetTagName(name string) {
	tagname = name
}

// AtoB copy a param to b param by struct tag
func AtoB(a, b interface{}) error {
	at := reflect.TypeOf(a)
	bt := reflect.TypeOf(b)
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)

	if bt.Kind() != reflect.Ptr || at.Kind() != reflect.Ptr {
		return ErrNotPoint
	}

	if bt.Elem().Kind() != reflect.Struct || at.Elem().Kind() != reflect.Struct {
		return ErrNotStruct
	}

	// build value index
	vi := make(map[string]reflect.Value)
	nf := at.Elem().NumField()
	for i := 0; i < nf; i++ {
		// filed and value to map
		vi[getTagFieldName(at.Elem().Field(i).Tag.Get(tagname))] = av.Elem().Field(i)
	}

	nf = bt.Elem().NumField()
	for i := 0; i < nf; i++ {
		if value, exist := vi[getTagFieldName(bt.Elem().Field(i).Tag.Get(tagname))]; exist && bv.Elem().Field(i).Kind() == value.Kind() {
			bv.Elem().Field(i).Set(value)
		}
	}

	return nil
}

func getTagFieldName(tag string) string {
	return strings.Split(tag, ",")[0]
}
