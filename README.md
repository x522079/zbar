# zbar
qrcode lib based on ZBar

# usage
```go
zbar_ := &zbar.Zbar{}
	sources := make([]string, 0)
	sources = append(sources, "/home/xxx/xxx/046b4484de790acae5ec98cc520e83e8.jpg")
	sources = append(sources, "/home/xxx/xxx/test.jpg")
	sources = append(sources, "/home/xxx/xxx/d129887b69c1f492b7adfcf9e482a602.jpg")
	zbar_.Source = sources
	b, err := zbar_.Decode()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(b.Symbols())
```