package main

import (
	"fmt"
	"strings"
	"math/rand"
	"time"
)

func main()  {
	fmt.Println("契約書だよ。そこに名前を書きな。")

	var name string
	_, _ = fmt.Scan(&name)
	fmt.Printf("フン。%sというのかい。贅沢な名だねェ。\n", name)

	nameArray := strings.Split(name, "")

	rand.Seed( time.Now().UnixNano() )
	newname := nameArray[ rand.Intn(len(nameArray)) ]
	fmt.Printf("今からお前の名前は%sだ。いいかい、%s。わかったら返事をするんだ、%s!!\n", newname, newname, newname)
}
