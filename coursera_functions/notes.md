[course link](https://www.coursera.org/learn/golang-functions-methods/home/)

# Week 1 - Functions
why use functions
- reusability -> makes code cleaner
- abstraction -> hides details

## Call by value/reference

**call by value**
- passed arguments are copied to parameters
- encapsulates data -> prevents bugs
- needs copy time
    ```go
    func foo(y int) {y = y + 1}
    func main() {
        x := 2
        foo(x)
        fmt.Print(x) // returns 2
    }
    ```
**call by reference**:
- pass a pointer as argument
- less copy time
- variables might change after function call
    ```go
    func foo(y *int) {*y = *y + 1}
    func main() {
        x := 2
        foo(&x)
        fmt.Print(x) // returns 3, the argument value has changed
    }
    ```
## Passing arrays and slices
- if you pass an array by value, the whole array needs to be copied
- better to pass by reference
    ```go
    func foo(x *[3]int){
        (*x)[0] = (*x)[0] + 1
    }
    func main(){
        a := [3]int{1, 2, 3}
        foo(&a)
        fmt.Print(a)
    }
    ```

- even better is to use slices  
Remember a slice is a structure with three attributes: pointer, length and capacity  
So passing by value copies just pointer and two other attributes
    ```go
    func foo(s []int){
        s[0] = s[0] + 1
    }
    func main(){
        a := []int{1, 2, 3}
        foo(a)
        fmt.Println(a)
    }
    ```
## Well written functions
- understandability/discoverability
- good function (and parameters) names
- cohesion: only one operation

# Week 2 - Function Types
## First-Class Values
- functions can be treated like other types
    - can be passed as function arguments
    - can be returned from a functions
    - can be created dynamically
    - can be stored in data structures


- functions as variables
    ```go
    var funcVar func(int) int 	// declare funcVar with a function
                        // signature

    func fooFn(x int) int {	// define a functions with signature
        return x + 1		// matching funcVar
    }

    func main() {
        funcVar = fooFn		// funcVar points to fooFn
        fmt.Print(funcVar(1))
    }
    ```
- functions as arguments
    ```go
    func applyIt(f func(int) int, val int) int {
        return f(val)
    }
    ```
- anonymous functions
    ```go
    func main(){
        v := applyIt(func(x int) int {return x + 1}, 2)
    }
    // anonymous function declaration
    ```
- returning functions
    ```go
    func MakeExpotent(b float64) func(float64)float64 {
        fn := func(v float64) float64 {
            return math.Pow(b, v)
        }
        return fn
    }
    ```  

    When returning/passing a function it comes with its environment, ie.: all variables that are available in its scope

## Variadic and Deferred
### variadic
- variadic: variable number of arguments
- ellipsis (…) is used to indicate any number of arguments
    ```go
    func getMax(vals ...int) int {
        maxV := -1
        for _, v := range vals {
            if v > maxV: {
                maxV = v
            }
        }
        return maxV
    }
    ```
- slices can be passed as variadic argument  
in such case a eclipsis suffix is needed
    ```go
    s := []int{1, 4, 2, 0}
    fmt.Println(getMax(s...))
    ```
### deferred
- deferred functions  
    - they don't get executed when they are called
    - they are executed when surrounding function completes  
    - usually used for cleanup activities
    ```go
    func main() {
        deferred fmt.Println(“Bye!”)
        fmt.Println(“Hello”)
    }
    ```
    !!! Arguments of deferred call are evaluated immediately
    ```go
    func main() {
        i := 1
        deferred fmt.Println(i + 1) // arg is evaluated to 2
        i++
        fmt.Println(“Hello”)
        // now deferred will be executed but with i value
        // equal to 2 (when it was called)
    }
    ```
# Week 3 - Object orientation in Go
## Receiver type
- encapsulation hides internals of class from class user
- no class keyword
- associating methods with data: receiver type
- those functions needs to be defined in the same module (file)
    ```go
    type MyInt int

        func (mi MyInt) Double () int {
            return int(mi*2)
        }
        // define receiver type for Double func

        func main(){
            v := MyInt(3)
            fmt.Println(v.Double())
        }
        // v is “silently” passed to Double
        // it is passed by value, so a copy of v is created
    ```

## Pointer Receiver
- we want functions to be able to change associated object
- solution is to define receiver as a pointer to object
    ```go
    type Point struct{
        x float64
        y float64
    }
    func (p *Point) OffsetX(v float){
        p.x = p.x + v
    }
    // observe that explicit dereferencing is not needed
    // its handled by compiler
    ```

## Controlling access
- var and functions whose names start with capital letters are “Public”
- they can be accessed when imported in another module
    ```go
    package data
    var x int = 1
    func PrintX() {fmt.Println(x)}

    package main
    import “data”

    func main(){
        data.PrintX()
    }
    ```
- similarly with struct attributes

# Week 4 - Interfaces
## Polymorphism
- ability for an object to have different forms depending on the context
    - identical at high level of abstraction
    - different at low level of abstraction
- in other languages polymorphism is handled with inheritance and overriding
- but not inheritance in Golang
## Interfaces
- how polymorphism is handled in Golang
- `interface`: set of method signatures
- used to express conceptual similarity between types
- a type satisfies and interface if it implements all methods specified in interface
    ```go
    type Shape2D interface {
        Area() float64
        Perimeter() float64
    }
    
    type Triangle {...}
    funct (t *Triangle) Area() float64 {...}
    funct (t *Triangle) Perimeter() float64 {...}
    ```
- Triangle satisfies the Shape2D interface (no need to state it explicitly)
- interface specifies only behavior, not data

## Pointer Interfaces
```go
type Speaker interface {Speak()}
type Dog struct {name string}
func (d *Dog) Speak() {
	fmt.Printf(“%s: woof!\n”, d)
}

func main(){
	var s Speaker
	s = &Dog{“Braian”}
	s.Speak()			//Braian: woof!

	// alternatively
	var d *Dog
	d = &Dog{“Braian”}
	s = d
	s.Speak()
}
```
## Interface Dynamic type and value
- when assigning value to an interface it sets two interface attributes:
        - dynamic type, and 
        - dynamic value
- type is equal to type of concrete interface implementation
- value is an concrete implementation instance
- Interface can have only dynamic type, still it knows how to call methods, eg:
    ```go
    type Speaker interface {Speak()}
    type Dog struct {name string}
    func (d *Dog) Speak() {
        if d == nill {
            fmt.Println(“<generic noise>”)
        } else {
        f   mt.Printf(“%s: woof!\n”, d)
        }
    }

    func main() {
        var s Speaker
        var d *Dog
        s = d
        s.Speak()
    }
    ```

- Nil dynamic type  
in such cases can't call interface methods

- Empty interface  
can be used for arguments of any type
    ```go
    func PrintMe(val interface{}) {	// empty interface
        fmt.Println(val)
    }
    ```
- Type assertions
    ```go
    type Shape2D interface {
        Area() float64
        Perimeter() float64
    }

    func DrawShape2D(s Shape2D) {
        rec, ok := s.(Rectangle)	// try extract concrete type
        if ok {
            DrawRec(rect)
        }
        tri, ok := s.(Triangle)	// try extract concrete type
        if ok {
            DrawTri(tri)
        }
    }

    // alternatively
    func DrawShape2D(s Shape2D){
        switch := sh := s.(type){
            case Rectangle:
                DrawRect(sh)
            case Triangle:
                DrawTri(sh)
        }
    }
    ```

## Error handling (error interface)
```go
func error interface{
	Error() string
}
```