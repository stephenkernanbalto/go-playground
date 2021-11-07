# The Python -> Go Cheatsheet

## Table of Contents

* [Lists](#lists)
* [Dictionaries](#dictionaries)
* [Loops](#loops)
* [Functions](#functions)
* [Important Differences](#differences)
    * [Pointers](#pointers)

## Strings

> type reference: `string`

Most string manipulation functions rely on the `fmt` module.

In order to use any of these methods, you'll need to import the library.

### print()

<a name="lists"></a>

## Lists -> Array, Slice

> type reference: `[length]type_for_content` for an array, `[]type_for_content` for a slice

### Basics

In Go, lists are either `arrays` (fixed-length) or `slices` (dynamic-length) and must contain items of the same type. To
distinguish between the two at declaration, you can omit a length between square brackets (`[]int`) to create a `slice`,
or provide a length (`[5]int`) to create a fixed-length `array`.

#### Creation

They are instantiated as follows:

1. Traditional Way

```gotemplate
{{/* creates an array consisting of 5 integers. by default they're all zeroes */}}
var list [5]int

{{/* creates a dynamic slice that is empty */}}
var list []int
```

2. Shorthand

```gotemplate
{{/* creates an array consisting of 5 integers. by default they're all zeroes */}}
list := [5]int{1,2,3,4,5}

{{/* creates a dynamic slice that is empty */}}
list := []int{1,2}
```

#### Accessing Values

You can access a value using the same syntax as Python:

```gotemplate
dictionary[key]
```

### Append

```python
list.append(value)
```

```gotemplate
append(list, value)
```

> Only works for `slice`. Due to the fixed-length nature of an `array`, an `array` cannot receive any additional items.

<a name="dictionaries"></a>

## Dictionaries -> Map

> type reference: `map[key_type]value_type`

### Basics

In Go, the hash table structure of a dictionary is referred to as a `map`.

#### Creation

To create a `map`, you use the built-in `make` function, like so:

1. Traditional

```gotemplate
var dictionary = make(map[string][int])
```

2. Shorthand

```gotemplate
dictionary := make(map[string][int])
```

#### Accessing Values

You can access a value using the same syntax as Python:

```gotemplate
dictionary[key]
```

<a name="loops"></a>

## Loops

### For Loops

#### Basics

This syntax is relatively similar to Python. You instantiate a variable, provide a condition at which the loop will end,
and increment the counter, like this:

```gotemplate
for i:= 0; i < 5; i++ {
  ... logic
}
```

### For-In Loops

#### Basics

This syntax is relatively similar to Python. You can unpack values from a slice and loop as follows:

```gotemplate
{{/* ARRAY */}}
arr := []string{"a", "b", "c"}

{{/* destructure the array args to variables of inferred types */}}
for index, value := range arr {
  ... logic
}

{{/* MAP */}}
dict := make(map[string]string)
dict["hello"] = "world"

{{/* destructure the map args to variables of inferred types */}}
for key, value := range m {
  ... logic
}
```

<a name="functions"></a>

## Functions

### Basics

Working with functions in Go is similar to working with functions in Python, although the `func` keyword and the static
typing provide slight differences.

The basic structure of a function is:

```gotemplate
func function_name(param_name param_type) return_value_type {
  logic
}
```

> **Note:** If returning more than one value, the return value types should go in a comma-separated list (e.g., `(string, int)`)

For a simple example, here's a function that simply prints and returns a string param if provided, or returns an error
if not:

```gotemplate
func PrintString(message string) (string, error) {
  if message == "" {
    return "", errors.New("No message provided")
  }
  
  fmt.Println(message)
  return message, nil
}
```

## Models -> Struct

### Basics

Similar to a database model in Flask/FastAPI & SQLAlchemy, a `struct` allows you to make a predictable preset of types
with a set of fields.

#### Creation

To create the structure itself:

```gotemplate
type model_structure struct {
  field_name field_type
  second_field second_type
}
```

To create a new object from that structure:

```gotemplate
model := model_structure{name: "Sly Cooper", favorite_book: "Thievius Raccoonus"}
```

#### Accessing Values

Similarly to Python, values in a model are accessed in dot notation:

```gotemplate
a := model.name
fmt.Println(a)
{{/* Prints 'Sly Cooper' */}}
```

<a name="differences"></a>

## Important Differences

<a name="pointers"></a>

## Pointers

> *!! PLEASE READ CAREFULLY, AS THIS IS AN ENTIRELY NEW CONCEPT* !!

In the following function, it looks as if the variable `i` should be incremented by 1, right?

```gotemplate
    func main() {
        i := 7
        inc(i)
        fmt.Println(i)
    }

    func inc(x int) {
        x++
    }
```

Not quite! In Go, variables are pointers to a place in memory where a value is stored. When this variable is passed
around to other functions, those functions receive *a copy* of the value rather than the actual value itself. So, in
this case, the `inc` function receives the value `7` but it has no access to change the actual value of `i`. So, it
increments `7` to `8` and then that value just disappears entirely. To access the actual value, we need a `pointer`.

This can be passed through with the following changes:

```gotemplate
    func main() {
        i := 7
        {{/* the ampersand tells Go to send the memory reference */}}
        inc(&i)
        fmt.Println(i)
    }

    {{/* the * tells the function it's receiving a pointer to an int */}}
    func inc(x *int) {
        {{/* the second * tells the function to retrieve the value at the pointer */}}
        *x++
    }
```