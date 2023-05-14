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

# 24. Working with files in golang

- Can read/erite text files
- `panic()` shuts down the program and handles the error

```go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content := "This needs to go in a file - Kinjal Raykarmakar"

	file, err := os.Create("./mygofile.txt")
	checkNillErr(err)
	defer file.Close()

	length, err := io.WriteString(file, content)
	checkNillErr(err)

	fmt.Println("Length is:", length)

	readFile("./mygofile.txt")
}

func readFile(filename string) {
	dataByte, err := os.ReadFile(filename) // data is not a string, but in bytes
	checkNillErr(err)

	fmt.Println("Text data inside the file is\n", string(dataByte))
}

func checkNillErr(err error) {
	if err != nil {
		panic(err)
	}
}
```

> `io/ioutils` is deprecated, hence using the `os` package for reading a file

# 25. Handling web request in golang

- [net/http](https://pkg.go.dev/net/http)
- We receive a [`Response`](https://pkg.go.dev/net/http#Response) type back when making a web request
  - The Close is important!
  - > // ReadResponse nor Response.Write ever closes a connection.
  - We need to manually close it
- The type of response is `*http.Response` (a pointer)
- `response.Body.Close()` is used to close the response
- Making a `GET` request and reading the response

```go
response, err := http.Get(url)
checkNillErr(err)
defer response.Body.Close()

fmt.Printf("Response is of type %T\n", response)

dataBytes, err := io.ReadAll(response.Body)

checkNillErr(err)
content := string(dataBytes)
fmt.Println("The content returned from the server is:\n", content)
```

# 26. Handling URL in golang

```go
result, _ := url.Parse(myurl)

fmt.Println("Scheme:", result.Scheme)
fmt.Println("Host:", result.Host)
fmt.Println("Path:", result.Path)
fmt.Println("Port:", result.Port())
fmt.Println("RawQuery:", result.RawQuery)
```

- Query params are a type of `url.Values` (Key-value pairs)
- Create a URL

```go
// need to pass reference
partsOfUrl := &url.URL{
	Scheme:   "https",
	Host:     "lco.dev",
	Path:     "/tutcss",
	RawQuery: "user=kinjal",
}

anotherUrl := partsOfUrl.String()
fmt.Println(anotherUrl)
```

# 28. How to make GET request in golang

```go
func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get((myurl))
	checkNillErr(err)
	defer response.Body.Close()

	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Content Length:", response.ContentLength)

	content, _ := (io.ReadAll(response.Body))
	fmt.Println(string(content))
}
```

- Alternate way to handle byte to string conversation. This gives us more control, as we still have the raw data in the `responseString`

```go
var responseString strings.Builder
byteCount, _ := responseString.Write(content)
fmt.Println("ByteCount is:", byteCount)
fmt.Println(responseString.String())
```

# 29. How to make POST request with JSON data in golang

```go
func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	// fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename":"Let's go with golang",
			"price": 0,
			"platform": "lco.in"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	checkNillErr(err)
	defer response.Body.Close()

	content, _ := (io.ReadAll(response.Body))

	fmt.Println(string(content))
}
```

# 30. How to send form data in golang

```go
func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	// form data
	data := url.Values{}
	data.Add("firstname", "Kinjal")
	data.Add("lastname", "Raykarmakar")
	data.Add("email", "kinjalrk2k@gmail.com")

	response, err := http.PostForm(myurl, data)
	checkNillErr(err)
	defer response.Body.Close()

	content, _ := (io.ReadAll(response.Body))

	fmt.Println(string(content))
}
```

# 31. How to create JSON data in golang

- Encoding of JSON: Converting key-value pairs into valid JSON string
- Struct defination

```go
type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"` // dash means it'll remove it while marshalling
	Tags     []string `json:"tags,omitempty"`
}
```

- The `json:....` act as an alias which renames the field name while marshalling
- `omitempty` skips the field alltogether when `nill` is encountered within it

- Encoding the JSON

```go
func EncodeJSON() {
	lcoCourses := []course{
		{Name: "ReactJS Bootcamp", Price: 299, Platform: "learncodeonline.in", Password: "abc123", Tags: []string{"web-dev", "js"}},
		{Name: "MERN Bootcamp", Price: 199, Platform: "learncodeonline.in", Password: "bcd123", Tags: []string{"full-stack", "js"}},
		{Name: "Angular Bootcamp", Price: 299, Platform: "learncodeonline.in", Password: "kin123", Tags: nil},
	}

	// package this data as JSON data
	// finalJSON, err := json.Marshal(lcoCourses)
	finalJSON, err := json.MarshalIndent(lcoCourses, "", "\t")
	checkNillErr(err)

	fmt.Printf("%s\n", finalJSON)
}
```

# 32. How to consume JSON data in golang

- Decoding the JSON

```go
func DecodeJson() {
	jsonDataFromWeb := []byte(`
		[
			{
							"coursename": "ReactJS Bootcamp",
							"price": 299,
							"website": "learncodeonline.in",
							"tags": [
											"web-dev",
											"js"
							]
			},
			{
							"coursename": "MERN Bootcamp",
							"price": 199,
							"website": "learncodeonline.in",
							"tags": [
											"full-stack",
											"js"
							]
			},
			{
							"coursename": "Angular Bootcamp",
							"price": 299,
							"website": "learncodeonline.in"
			}
		]
	`)

	var lcoCourses []course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON was valid!")
		json.Unmarshal(jsonDataFromWeb, &lcoCourses) // passing reference, 'cause i want to save it here
		fmt.Printf("%#v\n", lcoCourses)
	} else {
		fmt.Println("JSON was invalid")
	}
}
```

- When you do not want to create the struct

```go
// sometimes i don't waana create the struct
var myOnlineCourses []map[string]interface{} // interface - 'cause we do not know the type of it's value
json.Unmarshal(jsonDataFromWeb, &myOnlineCourses)
```

# 33. A long video on MOD in golang

- `go mod` is a tooling
- Tools are used for a lot of stuffs, like build, run etc
- We can still run Go code without modules, but it is not a good way
- Correct way to handle mods

```sh
go mod init github.com/kinjalrk2k/mymodules
```

- _`go.mod`_ file

```mod
module github.com/kinjalrk2k/mymodules

go 1.19
```

- GoLang follows a SemVer version - Semantic Versioning
  - `major.minor.patch`
  - [Docs](https://semver.org/)
- When dependencies are used, they do not come in directly in the working directory
- Modules came into exsistence since 2019 (Previously, they had something like Workspaces)

  - [Blog](https://go.dev/blog/using-go-modules)
    > Go 1.11 and 1.12 include preliminary support for modules, Go’s new dependency management system that makes dependency version information explicit and easier to manage.

## Trying to install a dependency

- Installing `gorrila/mux` [Docs](https://github.com/gorilla/mux#install)
- We'll use a Go toolchain
- `go get` is used to pull in dependencies from the **web**

```sh
go get -u github.com/gorilla/mux
```

- A _`go.sum`_ file is created and the `go.mod` file is updated
- _`go.mod`_ file
  - `indirect` means we're not using this dependency in our code yet
  - Run `go mod tidy` to remove this once used

```mod
module github.com/kinjalrk2k/mymodules

go 1.19

require github.com/gorilla/mux v1.8.0 // indirect
```

- _`go.sum`_ file
  - Stores hashes of the dependencies for security reasons
  - `go mod verify` verifies the hashes

```sum
github.com/gorilla/mux v1.8.0 h1:i40aqfkR1h2SlN9hojwV5ZA91wcXFOvkdNIeFDP5koI=
github.com/gorilla/mux v1.8.0/go.mod h1:DVbg23sWSpFRCP0SfiEN6jmj59UnW/n46BH5rLB71So=
```

- The dependencies are installed, but we do not see them in our working directory
  - From `go env` we can see that `GOPATH="/Users/kinjal/go"`
  - The dependencies go in here: `/Users/kinjal/go/pkg/mod/cache/download`

## `go list`

- `go list` - lists all the dependencies
- `go list all` - dumps all the packages installed
- `go list -m all` - lists all the dependencies in our module
- `go list -m -versions github.com/gorilla/mux` - lists all the versions of this package
- `go mod why github.com/gorilla/mux` - shows which module is dependent in this module
- `go mod graph` - prints the dependency graph
- `go mod edit -go 1.16` - changes the go version in the _`go.mod`_ file
- `go mod edit -module 1.16` - changes the module version in the _`go.mod`_ file

## A simple server

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in GoLang")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r)) // logs if something fails
}

func greeter() {
	fmt.Println("Hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to gorilla/mux</h1>"))
}
```

## Vendors

```sh
go mod vendor
```

- Something like node_modules
- Pull up the dependency code directly in our working directly, so than we can edit them if we want to
- `vendor.modules.txt` lists down all the modules
- To run the go file using the dependencies in the vendor
  - It first looks into the Vendor folder to grab the dependency. If not present, then goes to the cached version or from the web!

```sh
go run -mod=vendor main.go
```

# 34. Building API in golang - Models

```go
// Model for courses - file
type Course struct {
	CourseId    string  `json:"course_id"`
	CourseName  string  `json:"course_name"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"` // type is a pointer
}

type Author struct {
	Fullname string `json:"full_name"`
	Website  string `json:"website"`
}
```

# 35. Sending a API json response for all courses in golang

```go
// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API</h1>"))
}

// list all courses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all courses")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(courses)
}
```

# 36. Get one course based on request id in golang

```go
// get only one course
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Geting one course")

	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through the courses, find matching id and return
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode(`{"msg": "No course found with the given ID"}`)
}
```

# 37. Add a course controller in golang

```go
// create a course
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode(`{"msg": "Please send some data"}`)
		return
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	// what about - {}
	if course.IsEmpty() {
		json.NewEncoder(w).Encode(`{"msg": "No data in JSON"}`)
		return
	}

	// generate a unique id, convert to string
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	// append this course to course slice
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}
```

# 38. Update a course controller in golang

```go
// update one course
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through courses, remove it and add with exsisting id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			var updatedCourse Course
			_ = json.NewDecoder(r.Body).Decode(&updatedCourse)

			updatedCourse.CourseId = params["id"]
			courses = append(courses, updatedCourse)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode(`{"msg": "No course found with the given ID"}`)
}
```

# 39. Delete a course controller in golang

```go
// delete a course
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through courses, remove it
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(`{"msg": "Delete successful!"}`)
			return
		}
	}

	json.NewEncoder(w).Encode(`{"msg": "No course found with the given ID"}`)
}

```

# 40. Handling routes and testing routes in golang

```go
func main() {
	fmt.Println("API")
	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{CourseId: "1", CourseName: "Golang", CoursePrice: 100, Author: &Author{Fullname: "Hitesh", Website: "https://lco.dev"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "Full stack Web", CoursePrice: 100, Author: &Author{Fullname: "Colt", Website: "https://udemy.com"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen
	log.Fatal(http.ListenAndServe(":4000", r))
}
```

# 41. MongoDB setup for API in golang

https://go.mongodb.org/

- There should be only one go file: `main.go` in the root directory

# 42. Defining models for netflix in golang

_model/netflix.model.go_

```go
package netflixModel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}
```

# 43. Making a connection to database in golang

- https://pkg.go.dev/context
- WHenever we're connecting to something outside our system, we'll need to work with contexts

```go
const connectionString = "mongodb+srv://..."
const dbName = "netflix"
const collectionName = "watchlist"

// important!
var collection *mongo.Collection

// connect with MongoDB

// specialized method - runs on first time the application starts
func init() {
	// client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to MongoDB!")

	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection Reference is ready!")
}
```

# 44. Insert data in mongodb in golang

- Install go tools for VSCode

```go
// insert a record
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted with ID", inserted.InsertedID)
}
```

# 45. Update a record in mongodb in golang

- https://stackoverflow.com/questions/64281675/bson-d-vs-bson-m-for-find-queries

```go
// update a record
func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id} // read more about bson.M vs bson.D
	update := bson.M{"$set": bson.M{"watched": true}}

	updated, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count", updated.ModifiedCount)
}
```

# 46. Delete one and delete many in mongodb in golang

```go
// delete all records
func deleteAllMovies() int64 {
	deleted, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted count", deleted.DeletedCount)
	return deleted.DeletedCount
}
```

# 47. Get all collection in mongodb in golang

- https://www.mongodb.com/blog/post/quick-start-golang--mongodb--how-to-read-documents

```go
// get all movies
func getAllMovies() []primitive.M {
	curr, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer curr.Close(context.Background())

	var movies []primitive.M
	for curr.Next(context.Background()) {
		var movie bson.M
		err := curr.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	return movies
}
```

# 48. Get all movies from DB in golang

```go
func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}
```

# 49. Mark movie as watched in golang

```go
func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)

	json.NewEncoder(w).Encode(movie)
}
```

```go
func MarkWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	updateOneMovie(params["id"])

	json.NewEncoder(w).Encode(`{"msg": "Marked as watched"}`)
}
```

# 50. Delete 1 and all movie in golang

```go
func DeleteOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])

	json.NewEncoder(w).Encode(`{"msg": "Deleted successfully"}`)
}
```

```go
func DeleteAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	deleteAllMovies()

	json.NewEncoder(w).Encode(`{"msg": "Deleted all successfully"}`)
}
```

# 51. Creating routes and testing API in golang

```go
package router

import (
	"github.com/gorilla/mux"
	"github.com/kinjalrk2k/mongoapi/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movie", controller.GetAll).Methods("GET")
	router.HandleFunc("/api/movie", controller.Create).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteOne).Methods("DELETE")
	router.HandleFunc("/api/movie", controller.DeleteAll).Methods("DELETE")

	return router
}
```

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kinjalrk2k/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB API")

	r := router.Router()

	fmt.Println("Server is getting started...")
	fmt.Println("Listening on PORT 4000")
	log.Fatal(http.ListenAndServe(":4000", r))
}
```
