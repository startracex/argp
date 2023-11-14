# argp

argp is a simple command line arguments parser for go.

## Import

```sh
go get -u github.com/startracex/argp
```

## Usage

```go
// from os.Args
// ap.New()

// from []string
var ap = From([]string{"-f", "--fo", "--", "a", "=", "1", "b=2", "abc", "-xyz"})

// true // Query "-f" and "--fo" , (remove -f from Args).
ap.Bool("-f", "--fo")

// "1" true // Query "a", (remove "a", =, "1" from Args).
ap.String("a")

// "2" true // Query "b", (remove "b=2" from Args).
ap.String("b")

// ["abc", "-xyz"] 1 // After "--".
ap.After("--")

// ["--fo", "--", "abc", "-x", "-y", "-z"] Separated by - "-", (remove -xyz from Args).
ap.Short("-").Args
```
