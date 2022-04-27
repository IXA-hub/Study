package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int    "ユーザーID"
	Name string "名前"
	Age  uint   "年齢"
	Sex  string "性別"
}

func main() {
	u := User{Id: 1, Name: "Taro", Age: 32, Sex: "Men"}

	t := reflect.TypeOf(u)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Tag)
	}
}
