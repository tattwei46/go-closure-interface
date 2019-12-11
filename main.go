package main

import (
	"fmt"
	"reflect"
)

type Apple struct {
	Name string
}

type Orange struct {
	Name   string
	UserID int
}

func getOrange() Orange {
	return Orange{
		Name:   "Orange",
		UserID: 1,
	}
}

func getApple() Apple {
	return Apple{Name: "Apple"}
}

func getCache(getObject func() interface{}) interface{} {
	return getObject()
}

func main() {

	var cached interface{}
	cached = getCache(func() interface{} {
		d := getOrange()
		return reflect.ValueOf(d).Interface()
	})
	oData := cached.(Orange)
	fmt.Println(oData.Name)

	cached = getCache(func() interface{} {
		d := getApple()
		return reflect.ValueOf(d).Interface()
	})
	aData := cached.(Apple)
	fmt.Println(aData.Name)
}
