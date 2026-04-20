# Slices

- You can create a new slice from an existing slice by getting a range of elements. Once again using square-bracket notation, but specifying both a starting (inclusive) and ending (exclusive) index. If you don't specify a starting index, it defaults to 0. If you don't specify an ending index, it defaults to the length of the slice.

```go
    newSlice := withData[2:4]
    // => []int{2,3}
    newSlice := withData[:2]
    // => []int{0,1}
    newSlice := withData[2:]
    // => []int{2,3,4,5}
    newSlice := withData[:]
    // => []int{0,1,2,3,4,5}
```
