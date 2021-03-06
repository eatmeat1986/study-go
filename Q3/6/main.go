package main

import (
	"io"
	"os"
	"strings"
)

var (
	computer    = strings.NewReader("COMPUTER")
	system      = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

func main() {
	buildASCII()
}

func buildASCII() {
	var stream io.Reader
	// ここにioパッケージを使ってコードを書く
	stream = io.MultiReader(
		io.NewSectionReader(programming, 5, 1),
		io.LimitReader(system, 1),
		io.LimitReader(computer, 1),
		io.NewSectionReader(programming, 8, 1),
		io.NewSectionReader(programming, 8, 1),
	)
	io.Copy(os.Stdout, stream)
}
