# iterext

`iterext` is a Go package that provides iterator extension methods, leveraging Go's generic `iter.Seq` type (available since Go 1.22). It aims to offer convenient and common functional programming patterns for working with iterators.

## Description

This package extends the capabilities of Go's standard library iterators by providing additional higher-order functions like `Map` and `Fold`. These functions allow for more expressive and concise data manipulation when working with sequences.

## Installation

To use `iterext` in your Go project, you can install it using `go get`:

````shell
go get github.com/aria3ppp/iterext@v0.1.0
````

## Usage

Here's how you can use the functions provided by `iterext`:

### `Fold`

The `Fold` function (also known as reduce or accumulate) applies a function against an accumulator and each element in the iterator (from left to right) to reduce it to a single value.

```go
slc := []int{1, 2, 3, 4, 5}
mp := make(map[string]int, len(slc))

// Fold the integer slice into a map[string]int
// where keys are string representations of the numbers
// and values are the numbers themselves.
mp = iterext.Fold(slices.Values(slc), mp, func(m map[string]int, v int) map[string]int {
	m[strconv.Itoa(v)] = v
	return m
})

fmt.Println(mp)
// Output: map[1:1 2:2 3:3 4:4 5:5]
```

### `Map`

The `Map` function transforms an iterator of one type into an iterator of another type by applying a given function to each element.

```go
slc := []int{1, 2, 3, 4, 5}

// Map the integer slice to a new iterator of strings.
stringsIter := iterext.Map(slices.Values(slc), func(v int) string {
	return strings.Repeat("1 ", v)[:2*v-1]
})

// Collect the results from the iterator into a slice.
stringsSlice := slices.Collect(stringsIter)

for i, s := range stringsSlice {
	fmt.Printf("Index %d: String Value %q, Integer Value %d\n", i, s, slc[i])
}
/*
Output:
Index 0: String Value "1", Integer Value 1
Index 1: String Value "1 1", Integer Value 2
Index 2: String Value "1 1 1", Integer Value 3
Index 3: String Value "1 1 1 1", Integer Value 4
Index 4: String Value "1 1 1 1 1", Integer Value 5
*/
```
