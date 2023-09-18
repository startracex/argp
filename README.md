# argp

argp is a simple command line arguments parser for go.

## Import

```sh
go get -u github.com/startracex/argp
```

## Useage

```go
var ap = From([]string{"-f", "--fo", "--", "a", "=", "1", "b=2", "abc", "-xyz"})
ap.Bool("-f", "--f") 	//Query "-f" and "--fo" (true), (remove -f from Args).
ap.String("a")       	// Query "a" (1 true), (remove a , = , 1 from Args).
ap.String("b")       	// Query "b" (2 true), (remove b=2 from Args).
ap.After("--")       	// After "--" ([abc -xyz] 1).
ap.Short("-").Args   	// Separated by - "-" ([--fo -- abc -x -y -z]), (remove -xyz from Args).
```
