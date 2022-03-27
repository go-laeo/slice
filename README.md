# slice

Slice provides some funcs to you to manipulates slice variables. Go1.18+

# Requirements

Go 1.18 and later.

# Install

```shell
go get github.com/go-laeo/slice
```

# Usage

```go
# Contain
slice.Contain([]string{"hello", "world", "slice"}, "hello") // => true

# Chunk
slice.Chunk([]int{1,2,3,4,5}, 2) // => [][]int{{1,2}, {3,4}, {5}}

# Unique
slice.Unique([]string{"go", "go", "golang", "1.18", "1.18", "generics"}) // => []string{"go", "golang", "1.18", "generics"}
```

# License

Apache 2.0

Any PRs are welcome and appreciated.
