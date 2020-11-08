package main

import (
	"fmt"
	"strings"
	"math/rand"
	"time"
	"bufio"
	"os"
)

func main()  {
	var name string

	for len(strings.TrimSpace(name)) == 0 {
		fmt.Println("契約書だよ。そこに名前を書きな。")

		stdin := bufio.NewScanner(os.Stdin)
		stdin.Scan()
		name = stdin.Text()
	}

	fmt.Printf("フン。%sというのかい。贅沢な名だねェ。\n", name)

	nameArray := strings.Split(name, "")

	rand.Seed( time.Now().UnixNano() )
	newname := nameArray[ rand.Intn(len(nameArray)) ]
	fmt.Printf("今からお前の名前は%sだ。いいかい、%s。わかったら返事をするんだ、%s!!\n", newname, newname, newname)
}
