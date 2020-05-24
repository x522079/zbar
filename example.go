package main

import (
	"fmt"
	"github.com/x522079/zbar/zbar"
)

func main() {
	zbar_ := &zbar.Zbar{}
	sources := make([]string, 0)
	sources = append(sources, "/home/xxx/xxx/046b4484de790acae5ec98cc520e83e8.jpg")
	//sources = append(sources, "/home/xxx/xxx/test.jpg")
	zbar_.Source = sources
	b, err := zbar_.Decode()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(b.Symbols())
}
