Array, Slices and Maps
===

## Array

An array is a numbered sequence of elements of a single type with a **fixed**
length. You can not delete or append element from an array.

```
ArrayType   = "[" ArrayLength "]" ElementType
```


```go
var arr [5]int
arr[4] = 100
fmt.Println(arr) //[0 0 0 0 100]

var arr2 := [5]int{1, 2, 3, 4, 5}
```

## Slice

Like arrays, slices are indexable and have a length. The length of a slice s
can be discovered by the built-in function `len`; unlike with arrays it may
change during execution.

The value of an uninitialized slice is `nil`.

Slices are always associated with some array, and although they can never be
longer than the array, they can be smaller.

```go
SliceType = "[" "]" ElementType
make([]T, length, capacity) // slice length, capacity of the associated array
```

Another way to create slices is to use the [low : high] expression.

```go
make([]int, 50, 100)
new([100]int)[0:50]
```
There are valid expressions like `[:5], [5:], [:]`

Append to slice and copy slice, for slice type S with value type T:

```go
append(s S, x ...T) S  // T is the element type of S
copy(dst, src []T) int
copy(dst []byte, src string) int
```

For exmaple:

```go
slice := []int{1,2,3,4}
append(slice, 5, 6) // [1,2,3,4,5,6]
var slice2 = make([]int, 3)
copy(slice2, slice) // [1,2,3]
```

## Maps

A map is an unordered collection of key-value pairs.

```go
MapType     = "map" "[" KeyType "]" ElementType
```

A new, empty map value is made using the built-in function [make](https://golang.org/ref/spec#Making_slices_maps_and_channels),
which takes the map type and an optional capacity hint as arguments:

```go
var x := make(map[string]int)
x["bob"] = 10
x["jack"] = 11
fmt.Println(len(x)) //2
fmt.Println(x["bob"]) //10
delete(x, "bob")
```

Map returns *zero value* for the *ElementType* if it does exist.

```go
v, ok := x["bill"] // 0, false
```

To check if an element is in a map:

```go
if (v, ok) := x["bill"]; ok {
  //do something
}
```

Iterate through a map:

```go
for k, v := range x {
  fmt.Println(k, v)
}

// if you only want the keys
for k := range x {
  fmt.Println(k)
}
```
