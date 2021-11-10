# Concurrency In Go

## Table of Contents

* [Introduction](#introduction)
    * [Communicating Sequential Processes](#csp)
* [Thinking Concurrently In Go](#thinking-concurrently)
    * [Goroutines](#goroutines)
    * [Channels](#channels)
    * [Select Switches](#select)
    * [Contexts](#contexts)

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

#### Important Concepts

For an *excellent* explanation of channel behavior, check out [this class](https://youtu.be/fCkxKGd6CVQ?t=540) by Matt Holiday, starting at 9m00s.

##### Channels **block** unless ready to read or write.

A channel is ready to write if either of the following is true:

* it has buffer space, or
* at least one reader is ready to read

A channel is ready to read if any one of the following is true:

* it has unread data in the buffer, or
* at least one writer is ready to write, or
* it is closed

#### Channels are One-Directional (That's What Makes Them Beautiful)

A channel can **either** receive data *or* send data, but not both. To make this behavior more explicit and predictable,
we can actually limit the operations that can be performed on a channel by passing either a read-only or write-only
channel like this:

```gotemplate
{{/* declares a function that accepts a write only channel*/}}
func get(ch chan<- string) { ... }

{{/* declares a function that accepts a read only channel*/}}
func get(ch <-chan string) { ... }
```

#### Arrow Syntax

In the example above, we already see a glimpse of the purpose of arrow syntax. This syntax shows the flow of the data
through the channel. If the arrow points toward the channel, then we are `writing` to the channel. If the arrow points
away from the channel, we are `reading` data from the channel. The basic arrow syntax can be used in two ways.

1. Description - Like the example above, sometimes we need to describe the operations that *can* be performed with our
   channel.
    1. `<-chan` - If we are passing a channel in a function, this syntax means that we are allowed to read from the channel, but not to write to it.
    2. `chan<-` - Inversely, when the arrow points to the channel, it means data will flow *into* the channel. Therefore, we are `writing` to it.
2. Action - We also use arrow syntax to send to/read from a channel.
   1. `msg := <- chan` - In this case, `msg` is equal to *whatever is read from the channel*
   2. `chan <- msg` - In this example, `chan` receives `msg` in a write operation.
<a name="context"></a>



### Contexts

Contexts are helpful for maintaining synchronization of communication among channels. Contexts form a tree structure in
which child nodes of the tree communicate by allowing child nodes to monitor their parent nodes. This upward reference
enables us to cascade cancellations down the tree. This is especially helpful for things like timeout functions.

```
X = a timeout event for child_1_ctx

                            parent_process_ctx
                                    |
                            ------------------
                            |                |
                      child_1_ctx       child_2_ctx
                            X
                    -----------------
                    X               X
            grandchild_1_ctx    grandchild_2_ctx    
```

Contexts give two parameters to work with at each level of the tree:

1. A `channel` that closes when the context is cancelled
2. An error that's readable once the channel closes.

#### Creation

To create a context, it is most common to use the `Background` method like so:

```gotemplate
ctx := context.Background()
ctx = context.withValue(ctx, "name", "Sly Cooper")
```

#### Adding Properties to Context

You may notice that in the above example, we reassign the `ctx` value using another context method. It is common to add
in properties such as values or additional functionality to the channel before passing it down to the subtree. Some
common properties we may want to add include:

1. `context.withValue(parent_ctx, key, value)` - returns `context` - This lets us add an immutable value to our context,
   which can be helpful many reasons including logging purposes
2. `context.withTimeout(parent_ctx, time)` - returns `context`, `cancel` - Enables us to add a built-in timeout to our
   context at which we will close the subtree and any child contexts.