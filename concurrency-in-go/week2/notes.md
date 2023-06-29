## Concurrency Basics

#### 1. Processes: 
- Process is an instance of a running program
- Things inside a process
    - Memory (code, stack, heap, libraries, ...)
    - Registers

- OS allows processes to execute concurrently
- Processes switch quickly: ~20ms -> it seems like they're running at the same time

#### 2. Scheduling:
- OS schedules processes for execution -> illusion of parallelism

#### 3. Threads & Go routines
- One process can have many threads, which share some context with each other (memory, register - which is unique to process)
- Go routine: like a thread in Go
    - Many Goroutines execute within a single OS thread
    - The Go Runtime Scheduler schedule goroutins inside an OS thread
    (like a little OS inside a single OS thread)

#### 4. Interleavings
- Concurrent code can be hard to debug
- Order of execution between concurrent task is unknown
- You should consider all possibilities of order of execution (interleaving)
- In the end, ordering is non-deterministic

#### 5. Race condition
- Race condition is a issued where the outcome of the program depends on interleavings
- Make sure the outcome doesn't depend on interleaving
- Raise the need of communicating between two tasks (two routines)
- Threads are largely independent but not completely independent