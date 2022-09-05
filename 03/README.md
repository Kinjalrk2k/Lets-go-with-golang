# 03

### Golang installation and hello world

Video: [YouTube](https://youtu.be/62qGe9yhiJI)

---

- [GoLang](https://go.dev/)

## Hello World

- After install Go and VS Code extensions, we'll need to write our Hello World program
- Init (Something like `npm init`)

```sh
go mod init hello
```

- Go has a single entry point like in C/C++. `main` is a reserved keyword
- Automatic imports are done by VSCode
- main.go

```go
package main

import "fmt"

func main()  {
	fmt.Println("Hello, World!")
}
```

- To run the file:

```sh
go run main.go
```
