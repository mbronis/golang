[course link](https://www.coursera.org/learn/golang-getting-started/home/)

# Week 1
## Varia

Features of Go
- Runs fast
- Garbage collection
- Simple object (it is oo language)
- Concurrency build into the language


**Machine language**: cpu instructions represented in binary  
**Assembly language**: cpu instructions with mnemonics (no variables, just instructions and memory)  
**High-level language**: has variables, in general easier to use, needs to be translated  
**Translation**:
- Compiled: 
    - occurs only once, thus saving time, 
    - compiled before running the code, 
    - its faster (in general)
- Interpreted:
    - requires interpreter, 
    - translated while code is executed
    - manage memory automatically
    - has garbage collector

    Java compiles to bytecode, interpreted by jvm.
    Go, though being compiled, has a garbage collector

## Object-Oriented Programming
- Organize code through encapsulation
- Groups data and functions which are related (classes)
- user can define types which are specific to an application

Go does not uses term class  
Go uses structs, simplified implementation of class:
- No inheritance
- No constructors
- No generics

## Concurrency
### Parallelism
- increasing number of cores on processor
- Multiple tasks can be performed simultaneously
- However there are problems with parallelism:
    - When do tasks start/stop
    - What if task needs data from another task
    - Memory conflicts

### Concurrency
- Management of multiple tasks at the same time
- Concurrent programming enables parallelism:
    - Management of task execution
    - Communication between tasks
    - Synchronization between tasks

**Go** includes concurrency primitives:
- Goroutines represents concurrent task (a thread)
- Channels are used to communicate between tasks
- Select enables task synchronization

## Workspaces/Packages

Typically one workspace is used for multiple projects  
Go code is usually organized in 3 subdirectories:
- src: contains source code files
- pkg: contains packages you link in (libraries)
- bin: contains executable files

Workspace directory is defined by the **GOPATH** environment variable.  
Usually it is defined during installation

**Packages**:
- group of related source files
- Packages can be imported by other packages
- First line of file names the package
- There must be one package called main  
Building the main package generates an executable program  
Main package need a main function, where code execution begins

When importing a package Go searches **GOROOT** and **GOPATH**

## Go Tool
Go tool commands:
- `go build`: compiles the program  
Arguments can be a list of packages or a list of .go files  
Without arguments its compiles go files in current directory  
Creates an executable file for the main package, same name as the first .go file (.exe in windows)
- `go doc`: prints documentation for a package
- `go fmt`: formats source code files
- `go get`: downloads and installs packages
- `go list`: lists all installed packages
- `go run`: compiles .go files and runs the executable
- `go test`: runs tests

## Variables
Variable:
- Data stored in memory
- Must have name and type
- Must have declaration, eg: var x int
	
Type declaration  
Can define an alias (alternate name for a type), eg:
```go
type Celsius float64
type Pid int

var temperature Celsius
```

Initializing Variables  
```go
var x int = 100
var x = 100 // you don't need to specify the type
```

Initialize after declaration  
```go
var x int
x = 100
```

Uninitialized variables have a zero values (for its type)
```go
var x int // x = 0
var x string // x = “”
```

Short variable declaration (can be only done in functions)
```go
x := 100 // it interferes type by right hand value
```

# Week 2 (Basic data types)
## Pointers
Pointer: an address to data in memory
- **&** operator returns the address of a variable/function
- **\*** operator returns data at an address (dereferencing)
```go
var x int = 1
var y int
var ip *int	// ip is pointer to int
ip = &x		// ip now points to address of x
y = *ip		// y is now 1
```
```go
new() // function is a way to create variables
      // function creates a variable and returns a pointer to the variable
      // the variable is initialized to default
```
```go
ptr := new(int)
*ptr = 3		// value 3 is put at the address pointed by ptr
```
## Variable scope
Variable scopes are resolved using blocks.  
A block is a sequence of declarations and statements within matching pair og brackets {}

**Implicit blocks** (blocks that are defined without brackets):
- **Universe block**: all Go source
- **Package block**: all source code in a package
- **File block**: all source code in a file
- If, for, switch statements
- clauses in switch and select 

**Lexical scoping**  
b1 >= b2 if b2 is defined inside b1 block  
Variable is accessible from block b1 if variable is declared in block bx and bx>=b1

## Deallocating memory
**Stack vs. Heap**

**Stack**  
- area of memory dedicated to functions calls
- Stores local variables of functions
- Those are deallocated automatically when the function call completes
**Heap**
- Persistent region of memory
- Deallocation needs to be explicit 

## Garbage collection
- Unique feature of Go is that it is a compiled language and have garbage collector
- Compiler determines stack vs heap
- Garbage collection is automatic in background, at the cost of slight performance drop

## Comments, Printing, Integers, Floats, Strings
```go
// single line comment
/*
	Block comment
*/
```

**String concatenation in Go**: + operator
```go
“Hey ” + “Joe”
```

**Conversion** characters in Print statement  
`%s` is string conversion, so %s will be substituted by some string variable

```go
x := “Joe”
fmt.Printf(“Hi %s”, x)
```
**Type conversion**
```go
var x int32 = 1
var y int16 = 2
x = y			// error! Need to be the same type
x = int32(y) 	// ok! 
```
**Floats**
- float32 ~ 6 digits of precision
- float64 ~ 15 digits of precision

### Strings
- ASCII: 8 bits per character -> 2**8 (256) possible characters
- Unicode: 32 bits per character -> 2**32 (4.1*10**9) possible characters
- UTF-8:
    - subset of unicode,
    - 8 bits by default but can go up to 32 bits,
    - first 8 bits match ASCII
    - default In Go

**String**:
- sequence of arbitrary bytes, formatted in UTF-8
- read only

**String literals**:
- notated by double quotes
- Each byte is a rune (UTF-8 code point)

`unicode` package (functions on individual runes)
- ```IsDigit(r rune)```
- ```IsSpace(r rune)```
- ```IsLetter(r rune)```
- ```IsPunct(r rune)```
- ```IsLower(r rune)```
- ```ToUpper(r rune)```
- ```ToLower(r rune)```

`strings` package (string functions)
- ```Compare(a, b) // 0 if a==b, -1 if a < b, 1 if a>b```
- ```Contains(s, substring)```
- ```HasPrefix(s, prefix)```
- ```Index(s, substring) // index of the first instance of substring in s```
- ```Replace(s, old, new, n) // returns a copy of string with first n replaced```
- ```ToLower```
- ```ToUpper```
- ```TrimSpace```

`strconv` package
- `Atoi(s) // converts string to int`
- `Itoa(s) // converts int (base 10) to string`
- `FormatFloat(f, fmt, prec, bitSize) // converts floating point number to a string`
- `ParseFloat(s, bitSize) // converts string to a floating point number`

### Constants

Expression whose value is known at compile time  
Type is inferred from right hand side  
```go
const x = 1.3
const (
	y = 4
	z = “Hi”
)
```

**iota**
- Generates a set of related but distinct constants
- Can represent a property which has several distinct possible values  
Eg days of the week, colors
- Actual value of constant is not important
```go
type Grades int
const (
	A Grades = iota
	B
	C
	D
	F
)
// each constant is assigned to unique integer
// it starts with 1, but you should not depend on this
```

### Control flow
**for**
```go
for <init>; <cond>; <update> {
	<stmts>
}


// basic loop
for i:=0; i<10; i++ {
	fmt.Printf(“Hi”)
}

// alternative
i = 0
for i<10 {
	fmt.Printf(“Hi”)
	i++
}

// infinite loop
for {
	fmt.Printf(“Hi”)
}
```

**switch/case**
```go
switch x {
case 1:
	fmt.Printf(“case 1”)
case 2:
	fmt.Printf(“case 2”)
default:
	fmt.Printf(“no case”)
}

// tagless switch
switch {
case x > 1:
	fmt.Printf(“case 1”)
case x < -1:
	fmt.Printf(“case 2”)
default:
	fmt.Printf(“no case”)
}
```

**break/continue**
- `break`		exits containing loop
- `continue` 	goes to next iteration of loop


**fmt.Scan**  
- reads user input
- takes a pointer as an argument
- typed in data is written to pointer
- returns number of scanned items (tokens separated by space)
```go
var appleNum int

fmt.Printf(“Number of apples?”)
num, err := fmt.Scan(&appleNum)
fmt.Printf(appleNum)
```
# Week 3 (composite data types)
## Arrays
- fixed-length series of elements of the same type
- elements initialized to zero value (of the chosen type)
    ```go
    var x [5]int
    x[0] = 2
    ```

- Array literal - a way to initialize array elements
    ```go
    var x [5]int = {1, 2, 3, 4, 5}
    ```

- … infers size of an array from length of literal
    ```go
    x := [...]int{1, 2, 3, 4}
    ```

**Iterating an array**
```go
x := [3]int{1, 2, 3}
for i, v := range x {
	fmt.Printf(“index: %d, value: %d”, i, v)
}
```

## Slices
- window on an underlying array
- variable size, up to the size of whole array
- each slice has 3 properties
    - pointer: indicates the start of the slice
    - length: number of elements in the slice (len())
     - capacity: maximum number of elements (cap())  
    from start of slice to end of array
    ```go
    arr := [...]string{“a”, “b”, “c”, “d”, “e”, “f”, “g”}

    s1 := arr[1:3]
    s2 := arr[2:5]
    ```
- writing to a slice changes underlying array
- slice literals - creates underlying array and a slice that covers whole array
    ```go
    sli := []int{1, 2, 3} //no length or ... in [] -> it’s a slice literal
    ```
## Variable slices

- `make()` third method of creating slices
- 2-argument call - specify capacity (length = capacity)
    ```go
    sli = make([]int, 10)
    ```
- 3-argument call - specify length and capacity
    ```go
    sli = make([]int, 10, 15)
    ```
- `append()`  
    - increase size of slice (up to capacity of underlying array, or creates
    - a new array with increased size
    ```go
    sli = make([]int, 0, 3) // length of slice is 0
    sli = append(sli, 100)  // increases length to 1 and puts value 100
                            // in the first element
    ```

## Hash tables (maps)
- contains key/value pairs
- key has to be unique
- hash function is used to compute the slot for a key  

    Advantages:
    - faster lookup than lists (constant time vs linear time)
    - arbitrary keys (not ints like slices or arrays)

	Disadvantages
    - may have collisions (decreases speed of access, are rare)

- `map` is go implementation of hash table
    ```go
    // base initialization of a map
    var idMap map[string]int
    isMap = make(map[string]int)

    // map literal - inits a map with literal
    idMap := map[string]int {
        “joe”: 12345
    }

    delete(idMap, “Joe”) 	// removes key/value pair from map
    id, p := idMap[“Joe”] // p - bool indicating if key found in a map
    len(idMap)			// number of values in a map

    // iterate a map
    for k, v in range idMap {
        fmt.Println(k, v)
    }
    ```
## Structs
- groups together other objects of arbitrary type
    ```go
    type Person struct {
        name string  //field
        addr string
        phone string
    }
    var p1 Person

    p1.name = “Joe”
    x = p1.addr
    ```
- initialization
```go
// initializes a struct
// all fields will be initialized to zero value
p1 := new(Person)

// struct literal
p1 := Person{
	name: “joe”,
	addr: “a st.”,
	phone: “123”
}
```

# Week 4 (Protocols and Formats)
## RCF’s
- RCF - request for comments  
Definitions of Internet protocols and formats, eg:
    - HTML - hypertext markup language
    - URI - uniform resource identifier (eg URL)
    - HTTP - hypertext transfer protocol
## JSON
- all unicode
- human readable
- fairly compact representation (as small as you can get and still be readable)
- types can be combined recursively
- JSON Marshalling: JSON representation of an object
```go
p1 := Person(
	name: “joe”,
	addr: “a st.”,
	phone: “123”
)

barr, err := json.Marshal(p1) 	// returns JSON representation as []byte
						// and array of runes (byte array)

// unmarshall
var p2 Person
err := json.Unmarshal(barr, &p2)	// converts []byte to a Go object
						// accepts byte array and a pointer
						// to a Go object
```
## File Access: `ioutil`
- linear access (not random)
- basic operations:
    - open - get handle for access
    - read - read bytes into []byte
    - write - write []byte into file
    - close - release handle
    - seek - move read/write head
    - `io.ioutil` package
        ```go
        data, err := ioutil.ReadFile(“text.txt”) 	// data is []byte with
                                                    // contents of entire file
        ```
- Explicit open/close are not needed.
- Large files cause a problem (as it reads whole file)
```go
data := “Hello there”
err := ioutil.WriteFile(“outfile.txt”, data, 0777)
// 0777 is a (unix) permission to the newly created file
```

## File Access: `os`
- more precise package for files
- functions
    - `os.Open() // returns a file descriptor`
    - `os.Close() // closes a file`
    - `os.Read() // reads into []byte, controls the amount read`
    - `os.Write	// writes a []byte into a file`

- reading
    ```go
    f, err := os.Open(“dt.txt”)
    barr := make([]byte, 10)	// defines how much bytes to read
                        // running it again will read next 10 bytes
    nb, err := f.Read(barr)	// returns number of bytes read
    f.Close()
    ```
- writing
    ```go
    f, err := os.Create(“outfile.txt”)

    barr := []byte{1, 2, 3}
    nb, err := f.Write(barr)		// writes any Unicode sequence

    nb, err := f.WriteString(“Hi”)	// writes a string
    ```