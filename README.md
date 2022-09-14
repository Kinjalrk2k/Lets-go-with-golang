# Let's go with golang

> Hitesh Choudhary

Course: [YouTube Playlist](https://www.youtube.com/playlist?list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa)

# 02. Before you start with golang

- Go enthusiasts are called "Gophers"
- Go is termed as the Next-Gen programming language
  - A lot of people go ahead and say that if C was built in present day, it would be have been built like Go
- Go utilizes the true power of Cloud infrastructure

## Features

- Golang is compiled
- Runtime is built into the final executable. So it can run directly without a VM
- Executables are different for different OS
- Golang can be used to build from system apps to web apps
- Go code is already in production
- Object oriented - Both yes and no
  - We don't have classes in Golang (we have structs)
  - There's no overloading
- Lexer does a lot of work

# 03. Golang installation and hello world

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

# 04. GOPATH and reading go docs

- Help:

```sh
go help
```

- `go run` compiles and run the program

## GOPATH

- The path where Go is installed
- Default: subdirectory `go` in thw home directly
- Get current GOPATH

```sh
go env GOPATH
```

# 05. Lexer in golang and Types

## Semicolons

- Semicolons should be put up everywhere. But at some places the Lexer comes up to put the semicolon itself
- [Docs](https://go.dev/ref/spec#Lexical_elements)
- Lexer checks the grammar of the language

## Types

- Types are case sensitive
- Varibale type should be known in advance
- Everything is a Type
- Case determines the access level:
  - In `fmt.Println("Hello, World!")`, the `P` in `Println` is capital which denotes that this function was exported publically

### List of Types

- String
- Bool
- Integer
  - uint8
  - uint64
  - int8
  - int64
  - uintptr
- Floating
  - float32
  - float64
- Complex
- **Advanced Types**
  - Array
  - Slices
  - Maps
  - Structs
  - Pointers
  - _Function_
  - _Channels_
  - _...almost everything_

# 06. Variables, types and constants

## Variables

- To create a variable

```go
var username string = "Kinjal"
fmt.Println(username)
fmt.Printf("Variable is of type: %T \n", username)
```

- `var` is used to create a variable
- Type of the varibale is explicitly mentioned
- `%T` is a placeholder which prints the type of the variable
- Integers [Docs](https://go.dev/ref/spec#Numeric_types)
- Floating
  - For `float32` we get 5 values after the decimal
  - For `float64` we get more precise data
- Default value
  - For integers it's 0
  - For strings it's ""
- Implicit type

  - Sometimes we do not need to specify the type of the variable, the lexer decides it

  ```go
  var website = "kinjal.com" // implicit as string
  website = 3 // error
  ```

- No var style

```go
// no var style
numberOfUsers := 100
fmt.Println(numberOfUsers)
fmt.Printf("Variable is of type: %T \n", numberOfUsers)
```

> `:=` is the Walrun operator

- We can use Walrus operator inside a method. But outside we cannot

- Constants

```go
const LoginToken string = "dssad"
```

- When a variable name starts with a capital letter it is a public variable and can be exported

# 07. Comma ok syntax and packages in golang

- [bufio](https://pkg.go.dev/bufio)
- We do not have try-catch in Go
- Here is where comma ok syntax comes into play
- Reading user inputs

```go
reader := bufio.NewReader(os.Stdin)
fmt.Println("Enter the rating for our Pizza:")

// comma ok syntax
// err ok syntax
input, _ := reader.ReadString('\n') // \n is an ender which says till where should input be read
fmt.Println("Thanks for rating, ", input)
fmt.Printf("Thanks for this rating is %T", input)
```

- `input, err := reader.ReadString('\n')` can be used, but we are not using the error, hence we're using `_`

# 08. Conversions in golang

- When using `reader.ReadString('\n')`, there is a trailing character `\n` which also gets added to the input
- Errors are predictable!
- Error handling

```go
numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Added 1 to your rating", numRating+1)
}
```

# 09. Handling time in golang

```go
presentTime := time.Now()
```

- Formating date

```go
fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
```

- Go uses the above timestamp to format the date

> This is very strange

- New date

```go
createdDate := time.Date(2000, time.September, 22, 23, 23, 0, 0, time.UTC)
fmt.Println(createdDate.Format("01-02-2006 Monday"))
```

# 10. Build for Windows, Linux and Mac

- Go gives us the tools to create executables

```sh
go build
```

- For other OSs

```sh
GOOS="windows" go build
GOOS="linux" go build
```

# 11. Memory management in golang

- Memory allocation and deallocation happens automatically
- Methods for memory management
  - `new()`
    - Allocate memory but no INIT
    - We'll get the memory address
    - Zeroed Storage (data cannot be put initially)
  - `make()`
    - Allocate memory and INIT
    - We'll get the memory address
    - Non-zeroed storage
- GC or Garbage collection happens automatically
  - However there are certain requirements for the GC to work

> The GOGC variable sets the initial garbage collection target percentage. A collection is triggered when the ratio of freshly allocated data to live data remaining after the previous collection reaches this percentage. The default is GOGC=100. Setting GOGC=off disables the garbage collector entirely. runtime/debug.SetGCPercent allows changing this percentage at run time.
> [runtime](https://pkg.go.dev/runtime)

# 12. Pointers in golang

- Whenever we declare a variable, an array etc, it is just a reference to it's memory
- Sometimes while passing these variables to a function, these variables do not get passed on directly. A copy of these variables gets passed on which creats irregularities in the program. So, instead of passing the variables, if we pass the memory address of the variables, it garuntees that the value we pass is the actual value from the memory
- Defaut value of a pointer id `<nil>`

```go
var ptr *int // pointer that'll hold integer value
fmt.Println("Value of pointer is", ptr)
```

- `&` is usef for referencing
- `*` is used to denote a pointer

```go
myNumber := 23
var ptr = &myNumber
fmt.Println("Value of actual pointer is", ptr)  // 0x1400011c00
fmt.Println("Value of actual pointer is", *ptr) // 23

*ptr = *ptr + 2
fmt.Println("New value is: ", myNumber) // 25
```

# 13. Array in golang

- Array is very less used in golang
- Need to explictly mention the amount of space we need

```go
var fruitList [4]string

fruitList[0] = "apple"
fruitList[1] = "tomato"
fruitList[3] = "peach"

fmt.Println("Fruit List is", fruitList)
fmt.Println("Fruit List lenght is", len(fruitList))

var vegList = [3]string{"potato", "beans", "mushroom"}
fmt.Println("Veggy list is", vegList)
```

# 14. Slices in golang

- Mush more powerful and mostly used
- Under the hood Arrays
- Useful when working with databases
- When using thr `var` syntax, we need to initialize slices too

```go
var fruitList = []string{"Apple", "Tomato", "Peach"}

// fmt.Println("Type of fruitList is %T", fruitList)

fruitList = append(fruitList, "Mango", "Banana")
fmt.Println(fruitList)

// used for slicing
fruitList = (fruitList[1:3]) // [1:] or [:3] also works
fmt.Println(fruitList)

// define a slice with make
highScores := make([]int, 4)
highScores[0] = 234
highScores[1] = 945
highScores[2] = 465
highScores[3] = 867

// highScores[4] = 777 // crashed
highScores = append(highScores, 555, 666, 777) // doesn't crash - append reallocated memory
fmt.Println(highScores)

fmt.Println(sort.IntsAreSorted(highScores))
sort.Ints(highScores) // works in place
fmt.Println(highScores)
fmt.Println(sort.IntsAreSorted(highScores))
```

# 15. How to remove a value from slice based on index in golang

> Need to understand this in detail

```go
var courses = []string{"react", "javascript", "swift", "python", "go"}
fmt.Println(courses)

var index int = 2
courses = append(courses[:index], courses[index+1:]...)
fmt.Println(courses)
```

- `...` is used for varargs

# 16. Maps in golang

- Format in key-value pairs
- `make` can be used to create maps

```go
languages := make(map[string]string)
languages["js"] = "JavaScript"
languages["rb"] = "Ruby"
languages["go"] = "Go"
languages["py"] = "Python"

fmt.Println("List of all languages:", languages)
fmt.Println("JS is:", languages["js"])

// deleting a value
delete(languages, "rb")
fmt.Println("List of all languages:", languages)

// loop through maps
for key, value := range languages { // can be for _, value := (comma-ok syntax)
	fmt.Printf("For key %v, value is %v\n", key, value)
}
```

# 17. Structs in golang

- There is no inheritence in golang
- No `super()`

```go
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
```

- Capitals as we need to access them from outside

```go
kinjal := User{"Kinjal", "kinjalrk2k@gmail.com", true, 22}
fmt.Println(kinjal)
fmt.Printf("Details are: %+v\n", kinjal)

kinjalAgain := User{Name: "Kinjal", Email: "kinjalrk2k@gmail.com", Status: true, Age: 22}
fmt.Println(kinjalAgain)

fmt.Printf("Name is %v and email is %v", kinjal.Name, kinjal.Email)
```

# 18. If else in golang

- Control flow

```go
if loginCount < 10 {
	result = "Regular user"
} else if loginCount > 10 {
	result = "Watch Out!"
} else {
	result = "Exactly 10"
}
```

- `{` should be in the same line. Same thing with `} else {`
- Initializing on the go

```go
if num := 3; num < 10 {
	fmt.Println("num is less than 10")
} else {
	fmt.Println("num is not less than 10")
}
```

# 19. Switch case in golang and online playground

- Generate random number

```go
rand.Seed(time.Now().UnixNano())
diceNumber := rand.Intn(6) + 1
```

- Switch case
  - No need for `break` statement

> Need to figure out `fallthrough`

```go
switch diceNumber {
case 1:
	fmt.Println("Dice value is 1 and you can open")
case 2:
	fallthrough
case 3:
	fallthrough
case 4:
	fallthrough
case 5:
	fmt.Printf("You can move %v spots\n", diceNumber)
case 6:
	fmt.Println("You can move 6 spots and roll the dice again!")
default:
	fmt.Println("What was that?!?")
}
```

# 20. Loop break continue and goto in golang

- Only `for` loop is present
- Nothing like `++d` in Go

```go
for d := 0; d < len(days); d++ {
	fmt.Println(days[d])
}
```

- `range` automatically loops through array or slices and yields index

```go
for i := range days { // i is index here
	fmt.Println(days[i])
}
```

- For of loop

```go
for index, day := range days {
	fmt.Printf("Index is %v and Day is %v\n", index, day)
}
```

- While loop

```go
rougeValue := 1
for rougeValue < 10 {
  fmt.Println("Value is:", rougeValue)
	rougeValue++
}
```

- `continue`, `break` and `goto` is also present in Go
- Using `break`

```go
for rougeValue < 10 {

	if rougeValue == 5 {
		break
	}

	fmt.Println("Value is:", rougeValue)
	rougeValue++
}
```

- Using `goto`

```go
	rougeValue := 1
	for rougeValue < 10 {

		if rougeValue == 5 {
			goto lco
		}

		fmt.Println("Value is:", rougeValue)
		rougeValue++
	}

lco:
	fmt.Println("Jumping...")
```

# 21. Functions in golang

- `main()` is the entry point in a Go program

```go
func greeter() {
	fmt.Println("Namastee from GoLang")
}
```

- Functions cannot be written inside functions

> IIFEs and lambda functions are possible in GoLang

```go
func adder(valOne int, valTwo int) int /* this is return type */ {
	return valOne + valTwo
}
```

- Varargs

```go
func proAdder(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}

	return total
}
```

- Returning multiple values

```go
func proAdder(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}

	return total, "Hi from proAdder function"
}
```

# 22. Methods in golang

- Not much of a big difference though
- Mothods go into `structs` and we call them methods

```go
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active?", u.Status)
}

func (u User /* a copy is passed */) NewMail() {
	u.Email = "test@go.dev" // does not actually change the property
	fmt.Println("Email of this user is:", u.Email)
}
```

# 23. Defer in golan

- [Docs](https://go.dev/ref/spec#Defer_statements)

> A "defer" statement invokes a function whose execution is deferred to the moment the surrounding function returns, either because the surrounding function executed a return statement, reached the end of its function body, or because the corresponding goroutine is panicking.

- When we use `defer` for a statement, then that statement is run at the end of the function

> deferred functions are invoked immediately before the surrounding function returns, in the reverse order they were deferred.

- LIFO

```go
func main() {
	defer fmt.Println("Hello World")
	fmt.Println("Hello")
}

// OUTPUT :-
// Hello
// Hello World
```

```go
func main() {
	defer fmt.Println("Hello World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
}

// OUTPUT :-
// Hello
// Two
// One
// Hello World
```
