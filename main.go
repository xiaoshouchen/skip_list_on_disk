package main

import (
	"fmt"

	"github.com/xiaoshouchen/skip-list/skip_list"
)

func main() {

	sl := skip_list.NewSkipList(10)
	sl.Insert("1", "1")
	sl.Insert("2", "2")
	sl.Insert("3", "3")
	sl.Insert("1", "-1")
	sl.Remove("2")
	fmt.Println(sl.Get("1"))
	fmt.Println(sl.Size())
}
