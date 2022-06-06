# Garbage Collector

## Go 1.3 Release Notes

### Stack

Go 1.3 has changed the implementation of goroutine stacks away from the old, "segmented" model to a contiguous model. When a goroutine needs more stack than is available, its stack is transferred to a larger single block of memory. The overhead of this transfer operation amortizes well and eliminates the old "hot spot" problem when a calculation repeatedly steps across a segment boundary. Details including performance numbers are in this [design document](https://go.dev/s/contigstacks).

### Changes to the garbage collector

For a while now, the garbage collector has been *precise* when examining values in the heap; the Go 1.3 release adds equivalent precision to values on the stack. This means that a non-pointer Go value such as an integer will never be mistaken for a pointer and prevent unused memory from being reclaimed.

Starting with Go 1.3, the runtime assumes that values with pointer type contain pointers and other values do not. This assumption is fundamental to the precise behavior of both stack expansion and garbage collection. Programs that use [package unsafe](https://go.dev/pkg/unsafe/) to store integers in pointer-typed values are illegal and will crash if the runtime detects the behavior. Programs that use [package unsafe](https://go.dev/pkg/unsafe/) to store pointers in integer-typed values are also illegal but more difficult to diagnose during execution. Because the pointers are hidden from the runtime, a stack expansion or garbage collection may reclaim the memory they point at, creating [dangling pointers](https://en.wikipedia.org/wiki/Dangling_pointer).

*Updating*: Code that uses `unsafe.Pointer` to convert an integer-typed value held in memory into a pointer is illegal and must be rewritten. Such code can be identified by `go vet`.

### Map iteration

Iterations over small maps no longer happen in a consistent order. Go 1 defines that “[The iteration order over maps is not specified and is not guaranteed to be the same from one iteration to the next.](https://go.dev/ref/spec#For_statements)” To keep code from depending on map iteration order, Go 1.0 started each map iteration at a random index in the map. A new map implementation introduced in Go 1.1 neglected to randomize iteration for maps with eight or fewer entries, although the iteration order can still vary from system to system. This has allowed people to write Go 1.1 and Go 1.2 programs that depend on small map iteration order and therefore only work reliably on certain systems. Go 1.3 reintroduces random iteration for small maps in order to flush out these bugs.

*Updating*: If code assumes a fixed iteration order for small maps, it will break and must be rewritten not to make that assumption. Because only small maps are affected, the problem arises most often in tests.

## Go 1.5 Release Notes

The garbage collector has been re-engineered for 1.5 as part of the development outlined in the [design document](https://go.dev/s/go14gc). Expected latencies are much lower than with the collector in prior releases, through a combination of advanced algorithms, better [scheduling](https://go.dev/s/go15gcpacing) of the collector, and running more of the collection in parallel with the user program. The "stop the world" phase of the collector will almost always be under 10 milliseconds and usually much less.

For systems that benefit from low latency, such as user-responsive web sites, the drop in expected latency with the new collector may be important.

Details of the new collector were presented in a [talk](https://go.dev/talks/2015/go-gc.pdf) at GopherCon 2015.

## Go 1.9 Release Notes

[Go 1.9 Release Notes](https://go.dev/doc/go1.9)

Library functions that used to trigger stop-the-world garbage collection now trigger concurrent garbage collection. Specifically, [`runtime.GC`](https://go.dev/pkg/runtime/#GC), [`debug.SetGCPercent`](https://go.dev/pkg/runtime/debug/#SetGCPercent), and [`debug.FreeOSMemory`](https://go.dev/pkg/runtime/debug/#FreeOSMemory), now trigger concurrent garbage collection, blocking only the calling goroutine until the garbage collection is done.

The [`debug.SetGCPercent`](https://go.dev/pkg/runtime/debug/#SetGCPercent) function only triggers a garbage collection if one is immediately necessary because of the new GOGC value. This makes it possible to adjust GOGC on-the-fly.

Large object allocation performance is significantly improved in applications using large (>50GB) heaps containing many large objects.

The [`runtime.ReadMemStats`](https://go.dev/pkg/runtime/#ReadMemStats) function now takes less than 100µs even for very large heaps.