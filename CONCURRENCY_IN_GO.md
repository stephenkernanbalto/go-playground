# Concurrency In Go

## Table of Contents

* [Introduction](#introduction)
    * [Communicating Sequential Processes](#csp)
* [Thinking Concurrently In Go](#thinking-concurrently)
    * [Goroutines](#goroutines)
    * [Channels](#channels)
    * [Select Switches](#select)

## Introduction

### Communicating Sequential Processes (CSP)

CSP, or Communicating Sequential Processes, is foundational for understanding Go's concurrency model. For our case, it
should be sufficient to understand the basic concepts. For a more detailed understanding of CSP, here
is [an excellent deep dive](https://www.youtube.com/watch?v=zJd7Dvg3XCk) on Youtube.

As broken down by Arne Claus in his excellent
presentation, [Concurrency Patterns In Go](https://www.youtube.com/watch?v=rDRa23k70CU), CSP can be summarized with the
following points:

1. ***Each process is built for sequential execution.***

   This means, simply, that our code should be built as if it's running in order. This should be a relatively
   straightforward concept. In fact, I would venture to say that the *majority* of all programs we've written have been
   written for sequential execution.


2. ***Data is *communicated* between processes via channels***

   In other words, our processes should be designed in such a way that they maintain a local state, which can be *
   communicated* (rather than shared) to other processes. There is no universally shared state, but rather a network of
   communicative processes working with relative independence. This is the backbone of CSP.

3. ***Scale by adding more of the same***

   Additional workers should share the load.

## Thinking Concurrently in Go

"But wait! Why would we build for sequential execution when our goal is concurrency?"

The power of CSP, and therefore Go, is that by making our sequential processes communicative, we allow these
individually-concurrent processes to *become* concurrent as a group.

The result is a much simpler development experience in which we can write asynchronous processes as if they were
synchronous. By spreading this work across multiple cores, the result is a concurrent (or even parallel) program.

So how we accomplish concurrency in Go?

### Goroutines

> A `goroutine` is a unit of independent execution.

Goroutines are *not* threads. Goroutines are independent functions which Go schedules and assigns to a variable number
of threads based on hardware.

#### Creation

It's simple to create a goroutine. All you have to do is add the `go` keyword in front of a function call, like this:

```gotemplate
go do_func()
```

### Channels

Channels are the one-way entities through which `goroutines` communicate.



#### Creation

A `channel` is actually a datatype in Go, and can be created using the `make` function, like so:

```gotemplate
channel := make(chan int)
```
