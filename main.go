package main

import "github.com/asqat/govue/static"

func main() {
	err := static.WebFS()
	if err != nil {
		panic(err)
	}
}
