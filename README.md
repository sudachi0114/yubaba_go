# Golang で、湯婆婆を実装する

## Usage:

```sh
go run src/main/main.go
```

or

```sh
go build src/main/main.go
./main
```

## Example

* 正常系

```console
(*'-') < go run src/main/main.go 
契約書だよ。そこに名前を書きな。
山田太郎
フン。山田太郎というのかい。贅沢な名だねェ。
今からお前の名前は山だ。いいかい、山。わかったら返事をするんだ、山!!
```

* 湯婆婆クラッシュ

```console
(;x_x) < go run src/main/main.go 
契約書だよ。そこに名前を書きな。

フン。というのかい。贅沢な名だねェ。
panic: invalid argument to Intn

goroutine 1 [running]:
math/rand.(*Rand).Intn(0xc000064180, 0x0, 0x117a4e0)
        /Users/sudachi/.goenv/versions/1.15.2/src/math/rand/rand.go:169 +0x9d
math/rand.Intn(...)
        /Users/sudachi/.goenv/versions/1.15.2/src/math/rand/rand.go:337
main.main()
        /Users/sudachi/Documents/yubaba_go/src/main/main.go:23 +0x286
exit status 2
```

## Librarys:

* [fmt](https://golang.org/pkg/fmt/)
* [strings](https://golang.org/pkg/strings/#Split)
* [math/rand](https://golang.org/pkg/math/rand/)
* [time.UnixNano](https://golang.org/pkg/time/#Time.UnixNano)
* [bufio.NewScanner](https://golang.org/pkg/bufio/#NewScanner)
* [os](https://golang.org/pkg/os/)

## Links:
* [元ネタ様](https://qiita.com/RyotaNakaya/items/1c160932c21d69db5786)