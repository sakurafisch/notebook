# The Go Memory Model

[The Go Memory Model](https://go.dev/ref/mem)

[golang 中的 sync.Mutex 和 sync.RWMutex](https://www.jianshu.com/p/679041bdaa39)

## Introduction

The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine.

## Advice

Programs that modify data being simultaneously accessed by multiple goroutines must serialize such access.

To serialize access, protect the data in the following ways:

1. channel operations
2. synchronization primitives in the [`sync`](https://go.dev/pkg/sync/) and [`sync/atomic`](https://go.dev/pkg/sync/atomic/) packages

## Happens Before

### Definition

If event *e1* happens before event *e2*, then we say that *e2* happens after *e1*.Also, if *e1* does not happen before *e2* and does not happen after *e2*, then we say that *e1* and *e2* happen concurrently.

### Weak Conditions Pair

A read *r* of a variable `v` is *allowed* to observe a write *w* to `v` if both of the following hold:

1. *r* does not happen before *w*.
2. There is no other write *w'* to `v` that happens after *w* but before *r*.

### Strong Conditions Pair(recommend)

To guarantee that a read *r* of a variable `v` observes a particular write *w* to `v`, ensure that *w* is the only write *r* is allowed to observe. That is, *r* is *guaranteed* to observe *w* if both of the following hold:

1. *w* happens before *r*.
2. Any other write to the shared variable `v` either happens before *w* or after *r*.

### Within Single Goroutine

Within a single goroutine, there is no concurrency, so the two definitions are equivalent: a read *r* observes the value written by the most recent write *w* to `v`. When multiple goroutines access a shared variable `v`, they must use synchronization events to establish happens-before conditions that ensure reads observe the desired writes.

### Variable Initialization as Write

The initialization of variable `v` with the zero value for `v`'s type behaves as a write in the memory model.

### Multiple Machine-word-sized Operations

Reads and writes of values larger than a single machine word behave as multiple machine-word-sized operations in an unspecified order.

## Synchronization

Program initialization runs in a single goroutine, but that goroutine may create other goroutines, which run concurrently.

If a package `p` imports package `q`, the **completion** of `q`'s `init` functions happens **before** the **start** of any of `p`'s.

The **start** of the function `main.main` happens **after** all `init` functions have **finished**.

### Channel communication

Channel communication is the main method of synchronization between goroutines. Each send on a particular channel is matched to a corresponding receive from that channel, usually in a different goroutine.

A send on a channel happens before the corresponding receive from that channel completes.

This program:

```go
var c = make(chan int, 10)
var a string

func f() {
	a = "hello, world"
	c <- 0
}

func main() {
	go f()
	<-c
	print(a)
}
```

is guaranteed to print `"hello, world"`. The write to `a` happens before the send on `c`, which happens before the corresponding receive on `c` completes, which happens before the `print`.

The closing of a channel happens before a receive that returns a zero value because the channel is closed.

In the previous example, replacing `c <- 0` with `close(c)` yields a program with the same guaranteed behavior.

A receive from an unbuffered channel happens before the send on that channel completes.

This program (as above, but with the send and receive statements swapped and using an unbuffered channel):

```go
var c = make(chan int)
var a string

func f() {
	a = "hello, world"
	<-c
}

func main() {
	go f()
	c <- 0
	print(a)
}
```

is also guaranteed to print `"hello, world"`. The write to `a` happens before the receive on `c`, which happens before the corresponding send on `c` completes, which happens before the `print`.

If the channel were buffered (e.g., `c = make(chan int, 1)`) then the program would not be guaranteed to print `"hello, world"`. (It might print the empty string, crash, or do something else.)

The *k*th receive on a channel with capacity *C* happens before the *k*+*C*th send from that channel completes.

This rule generalizes the previous rule to buffered channels. It allows a counting semaphore to be modeled by a buffered channel: the number of items in the channel corresponds to the number of active uses, the capacity of the channel corresponds to the maximum number of simultaneous uses, sending an item acquires the semaphore, and receiving an item releases the semaphore. This is a common idiom for limiting concurrency.

This program starts a goroutine for every entry in the work list, but the goroutines coordinate using the `limit` channel to ensure that at most three are running work functions at a time.

```go
var limit = make(chan int, 3)

func main() {
	for _, w := range work {
		go func(w func()) {
			limit <- 1
			w()
			<-limit
		}(w)
	}
	select{}
}
```

### Locks

The `sync` package implements two lock data types, `sync.Mutex` and `sync.RWMutex`.

For any `sync.Mutex` or `sync.RWMutex` variable `l` and *n* < *m*, call *n* of `l.Unlock()` happens before call *m* of `l.Lock()` returns.

For any call to `l.RLock` on a `sync.RWMutex` variable `l`, there is an *n* such that the `l.RLock` happens (returns) after call *n* to `l.Unlock` and the matching `l.RUnlock` happens before call *n*+1 to `l.Lock`.

#### Mutex（互斥锁）

- Mutex 为互斥锁，Lock() 加锁，Unlock() 解锁
- 在一个 goroutine 获得 Mutex 后，其他 goroutine 只能等到这个 goroutine 释放该 Mutex
- 使用 Lock() 加锁后，不能再继续对其加锁，直到利用 Unlock() 解锁后才能再加锁
- 在 Lock() 之前使用 Unlock() 会导致 panic 异常
- 已经锁定的 Mutex 并不与特定的 goroutine 相关联，这样可以利用一个 goroutine 对其加锁，再利用其他 goroutine 对其解锁
- 在同一个 goroutine 中的 Mutex 解锁之前再次进行加锁，会导致死锁
- 适用于读写不确定，并且只有一个读或者写的场景

#### RWMutex（读写锁）

- RWMutex 是单写多读锁，该锁可以加多个读锁或者一个写锁
- 读锁占用的情况下会阻止写，不会阻止读，多个 goroutine 可以同时获取读锁
- 写锁会阻止其他 goroutine（无论读和写）进来，整个锁由该 goroutine 独占
- 适用于读多写少的场景

##### Lock() 和 Unlock()

- Lock() 加写锁，Unlock() 解写锁
- 如果在加写锁之前已经有其他的读锁和写锁，则 Lock() 会阻塞直到该锁可用，为确保该锁可用，已经阻塞的 Lock() 调用会从获得的锁中排除新的读取器，即写锁权限高于读锁，有写锁时优先进行写锁定
- 在 Lock() 之前使用 Unlock() 会导致 panic 异常

##### RLock() 和 RUnlock()

- RLock() 加读锁，RUnlock() 解读锁
- RLock() 加读锁时，如果存在写锁，则无法加读锁；当只有读锁或者没有锁时，可以加读锁，读锁可以加载多个
- RUnlock() 解读锁，RUnlock() 撤销单词 RLock() 调用，对于其他同时存在的读锁则没有效果
- 在没有读锁的情况下调用 RUnlock() 会导致 panic 错误
- RUnlock() 的个数不得多余 RLock()，否则会导致 panic 错误

### Once

The `sync` package provides a safe mechanism for **initialization** in the presence of multiple goroutines through the use of the `Once` type. Multiple threads can execute `once.Do(f)` for a particular `f`, but only one will run `f()`, and the other calls block until `f()` has returned.

A single call of `f()` from `once.Do(f)` happens (returns) before any call of `once.Do(f)` returns.

In this program:

```go
var a string
var once sync.Once

func setup() {
	a = "hello, world"
}

func doprint() {
	once.Do(setup)
	print(a)
}

func twoprint() {
	go doprint()
	go doprint()
}
```

calling `twoprint` will call `setup` exactly once. The `setup` function will complete before either call of `print`. The result will be that `"hello, world"` will be printed twice.

## Incorrect synchronization

Note that a read *r* may observe the value written by a write *w* that happens concurrently with *r*. Even if this occurs, it does not imply that reads happening after *r* will observe writes that happened before *w*.

In this program:

```go
var a, b int

func f() {
	a = 1
	b = 2
}

func g() {
	print(b)
	print(a)
}

func main() {
	go f()
	g()
}
```

it can happen that `g` prints `2` and then `0`.

This fact invalidates a few common idioms.

Double-checked locking is an attempt to avoid the overhead of synchronization. For example, the `twoprint` program might be incorrectly written as:

```go
var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
}

func doprint() {
	if !done {
		once.Do(setup)
	}
	print(a)
}

func twoprint() {
	go doprint()
	go doprint()
}
```

but there is no guarantee that, in `doprint`, observing the write to `done` implies observing the write to `a`. This version can (incorrectly) print an empty string instead of `"hello, world"`.

Another incorrect idiom is busy waiting for a value, as in:

```go
var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
}

func main() {
	go setup()
	for !done {
	}
	print(a)
}
```

As before, there is no guarantee that, in `main`, observing the write to `done` implies observing the write to `a`, so this program could print an empty string too. Worse, there is no guarantee that the write to `done` will ever be observed by `main`, since there are no synchronization events between the two threads. The loop in `main` is not guaranteed to finish.

There are subtler variants on this theme, such as this program.

```go
type T struct {
	msg string
}

var g *T

func setup() {
	t := new(T)
	t.msg = "hello, world"
	g = t
}

func main() {
	go setup()
	for g == nil {
	}
	print(g.msg)
}
```

Even if `main` observes `g != nil` and exits its loop, there is no guarantee that it will observe the initialized value for `g.msg`.

In all these examples, the solution is the same: use explicit synchronization.