package main

import (
	"fmt"
	"strings"
	"math/rand"
)

func main()  {
	fmt.Println("契約書だよ。そこに名前を書きな。")

	var name string = "すだち"
	// _, _ = fmt.Scan(&name)
	fmt.Printf("フン。%sというのかい。贅沢な名だねェ。\n", name)

	nameArray := strings.Split(name, "")
	fmt.Println(nameArray, len(nameArray))
	var r int = rand.Intn(len(nameArray))
	fmt.Println(r)
}