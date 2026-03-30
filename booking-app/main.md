Simulate a long processing task

Generate and send ticket

By Default sequential code execution
processing 1 task or 1 code after another

concurrency to make out program more efficient

go keyword

go starts a new goroutine

A goroutine is a lightweight thread managed by the Go runtime


Waitgroup

Waits for the launched goroutine to finish

Package "sync" provides basic synhronization functionality

Add: Sets the number of goroutines to wait for (increases the counter by the provided number)

Wait: Blocks until the WaitGroup counter is 0

Done: Decrements the waitGroup counter by 1, So this is called by the goroutine to indicate that it's finished

### Comparison to other languages

Writing concurrent code: Writing concurrent code is more complex, More overhead

Threads vs Goroutines


Goroutine
-> Go is using, what's called a "Green thread"
-> Abstraction of an actual thread
-> Managed by the go runtime, we are only interacting with these high level goroutines
-> Cheaper & lightweight
-> Built-in functionality fro goroutines to talk with one another

OS Thread
-> Managed by Kernel
-> Are hardware dependent
-> Cost of these threads are higher
-> Higher start up time
-> No easy communication between threads
-> Reaciving and Sending Data

