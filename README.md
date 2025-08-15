# goclamp

simple clamp for golang

provide `clamp` function for `ordered` types
```go
clamp(3, 10, 7)   // = 7
clamp(3, 10, 1)   // = 3
clamp(3, 10, 11)  // = 10
```

and `clampFunc` function with custom `Less` function
```go
clampFunc(
		4,
		7,
		2,
		func(a, b int) bool {
			return a%b == 0
		}) // = 2
```

