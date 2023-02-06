package main

import (
	"fmt"

	"github.com/xiaoshouchen/skip-list/sorted_string_table"
)

func main() {

	sst := sorted_string_table.NewSST()
	sst.Insert("hello", "world")
	sst.Insert("你好", "世界")
	get, err := sst.Get("你好")
	if err != nil {
		return
	}
	fmt.Println(get)
	get1, err := sst.Get("hello")
	if err != nil {
		return
	}
	fmt.Println(get1)
}
