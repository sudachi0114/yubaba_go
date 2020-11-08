package main

import (
	"fmt"
)

func main()  {
	fmt.Println("契約書だよ。そこに名前を書きな。")

	var name string
	_, _ = fmt.Scan(&name)
	fmt.Printf("フン。%sというのかい。贅沢な名だねェ。\n", name)
}