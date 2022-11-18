[course link](https://www.coursera.org/learn/golang-concurrency/home/)

# Week 1 - Why concurrency
## Parallel execution
- build into the language
- parallel when executing at the same time
- in general one core executes one instruction at a time
- Dynamic power  
`P = a * CR(V**2)`
    - a - percent of time switching
    - C - capacitance (related to size)
    - F - clock frequency
    - V - voltage swing (from low to high)
- Dennard Scaling  
voltage should scale with transistor size

## Concurrent vs Parallel
- concurrent execution is not necessarily parallel
- concurrent start and end times overlap
- concurrent processes can be executed on one core
- programmer determines which tasks can be executed in parallel
- what will be executed in parallel is decided by os and go routine scheduler (hardware mapping)
- concurrency hides latency

# Week 2 - Concurrency Basics
## Processes
- process: an instance of running program
- it has:
    - its memory: address space
    - code
    - stack/heap
    - shared libraries
    - unique register: stores values
    - program counter (register that points to a instruction to be executed)
    - stack/heap register
- processes are switched quickly ~20ms
- schedule task (os): decides what process is executed when
- os gives fair access to cpu, mem, etc for all processes
- context switching: control flow change from one process to another
- a state (context) needs to be saved and swapped (from register to mem)

## Processes vs threads
- context switching can be slow
- thread is like process but with limited context part of which is shared with other threads

## Goroutines
- goroutine is a thread in Go
- many goroutines can be executed within a single os thread
- goroutine scheduling is covered by go runtime scheduler

## Interleavings
- order of execution within a task is known
- order of execution between concurrent tasks is unknown
- every time a concurrent program is run tasks interleaving can be different

## Race Conditions
- program should be deterministic, ie: outcome of the program should not depend on interleavings
- if it does it is called race condition
- race condition occur due to communication between the tasks

## Communication between tasks
- threads are largely independent but not completely independent
- eg.: web server, one thread per client, each thread is writing to the pace’s counter value

# Week 3 - Threads in Go
## Goroutines
- one goroutine is created automatically to execute main()
- other coroutines are created using the go keyword
    ```go
    a = 1
    go foo() 	// foo is sent for execution to a new goroutine
    a = 2		// main coroutine continues, without waiting for foo
    ```
- when main goroutines is completed, all other goroutines are forced to exit even if they didn't finish

## Synchronization
- goroutines are largely independent and doesn’t know about each other
- two possible interleavings of two goroutines
    ```
    Task1		Task2				Task1		Task2
    x = 1   						x = 1
                print(x)			x++
    x++								             print(x)
    ```
- we can synchronize them using some kind of global event
    ```
    Task1		Task2
    x = 1	
    x++
    GLOBAL EV	if GLOBAL EV: print(x)
    ```
- we reduce freedom of scheduler this way which reduces efficiency (but it is necessary)

## Wait groups
- `sync` package contains functions to synchronize between goroutines
- `sync.WaitGroup` forces a goroutine to wait for other goroutines  
Each WaitGroup instance has internal counter:
    - it is increased and decreased manually by programmer (Add/Done methods)
    - Wait method can be used to block further execution until counter is zero
    - idea is to increment the counter everytime a goroutine is spawned, and decrease it whenever it finishes its job
    - add wait to ensure all prior goroutines completed
    - Add and Wait should be used in the main thread
    - Done should be used by a new goroutine

## Communication (channels)
- channels are used to communicate between goroutines
- channels are typed
    ```go
    func Product(v1, v1 int, c chan int) {
        c <- v1 * v2
    }

    func main(){
        c := make(chan int)

        go Product(2, 3, c)
        go Product(4, 5, c)

        a := <- c
        b := <- c

        fmt.Println(a*b)
    }
    ```

## blocking in channels
- **unbuffered** channel cannot hold data in transit (default)
- sending blocks until data is received
- receiving blocks until data is sent
- channel can be used only for synchronization
    ```go
    <- c // without an assignment, similar as Wait
    ```
- **buffered** channels has capacity for holding a number of objects
    ```go
    c := make(chan int, 3)
    // 3 sends can be made without receivers
    // task can process (sent) up to 3 objects to channel wo blocking
    ```
- sending blocks only if a buffer is full
- receiving blocks only if a buffer is empty
- buffer is storing data in transit

# Week 4 - Synchronized Communications
## Blocking on channels
```go
ch := make(chan int)
for i := range ch {

}
// continues to read from channel until close happens
```

### Select (reading from multiple channels)
- scenario 1:  
a goroutine need results from two other goroutines (eg to calculate product).  
In such case we can read both channels sequentially
    ```go
    a := <- ch1
    b := <- ch2
    fmt.Println(a*b)
    ```
- scenario 2:  
We can read from either channel (whichever comes first)
    ```go
    select {
        case a = <- c1:
            fmt.Println(a)
        case b = <- c2:
            fmt.Println(b)
    }
    ```
- select can block also on sending data
    ```go
    select {
        case a = <- c1:
            fmt.Println(“Recieved a”)
        case c2 <- b:
            fmt.Println(“Sent b”)
    }
    ```
    Both cases are blocking, whichever channel becomes available first gets executed, either:
    - a value is sent to c1, or
    - c2 becomes available for receiving the data

### Abort channel
```go
for {
	select {
		case a <- c:
			fmt.Println(a)
		case <- abort:     // if whatever gets sent to abort
			return         // channel it will beak the infinite
	}					   // loop
}
```
### Default case
```go
select {
    case a = <- c1:
        fmt.Println(a)
    case b = <- c2:
        fmt.Println(b)
    default:
        fmt.Println(“nop”)
}
```
Adding default makes select non blocking. If both channel are empty than default will be executed

## Mutual Exclusion
- concurrency-safe  
function can be invoked concurrently without interfering with outher goroutine
- variable sharing between goroutines violates safety, eg
    ```go
    var i int = 0
    var wg sync.WaitGroup

    func inc(){
        i++
        wg.Done()
    }

    func main(){
        wg.Add(2)
        go inc()
        go inc()
        wg.Wait()
        fmt.Println(i)
    }
    ```
- interleaving happens on machine code level  
each increment consists of 3 machine code instructions:
    - read
    - increment
    - write

    thus in the example above interleaving can cause both goroutines can
    read from i before incrementing

## Mutex
- access to shared variables cannot be interleaved
- writting to shared variables should be mutually exclusive
- `sync.Mutex`
    - uses binary semaphires
    - flag up: shared variable is in use
    - flag down: shared variable is available
    - `Lock()`  
    method puts the flag up, if another gouroutine wants to call Lock() on the
    same variable it gets blocked and has to wait until flag is put down
    - `Unlock()`
    method that puts the flag down.
    ```go
    var i int = 0
    var mut sync.Mutex
    func int() {
        mut.Lock()		// all code region in between is 
        i++			// blocked
        mut.Unlock()
    }
    ```

## Once Synchronization
- initialization happen once
- must happen before everything else
- it can be realised by initializing in main, but it is not always possible
- another option is `sync.Once`
- it has one method: once.Do(f)
    - function is executed only one time, even if is called in multiple goroutines
    - `once.Do()` blocks until the first returns
    ```go
    var wg sync.WaitGroup

    func main(){
        wg.Add(2)
        go dostuff()
        go dostuff()
        wg.Wait()
    }

    var on sync.Once
    func setup(){
        fmt.Println(“Init”)
    }

    func dostuff(){
        on.Do(setup)
        fmt.Println(“hello”)
        wg.Done()
    }

    >> Init
    >> hello
    >> hello
    ```

## Deadlock
- comes from synchronization dependencies (execution dependency)
- a goroutine cannot continue if other gouroutine is done
- deadlock -> circular dependency
    ```go
    func dostuff(c1, c2 chan int) {
        <- c1
        c2 <- 1
        wg.Done()
    }
    // read from the first channel: wait for write onto c1
    // write to second channel: wait for read from c2

    func main(){
        c1 := make(chan int)
        c2 := make(chan int)

        wg.Add(2)
        go dostuff(c1, c2)
        go dostuff(c2, c1)
    wg.Wait()
    }
    ```

- golang interpreter automatically detects a deadlock if all goroutines are blocked -> raises an error  
it however cannot detect if only subset is blocked