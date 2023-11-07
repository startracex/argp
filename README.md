# argp

argp is a simple command line arguments parser for go.

## Import

```sh
go get -u github.com/startracex/argp
```

## Usage

```go
//ap := argp.New()
ap := argp.From([]string{"-f", "--fo", "--", "a", "=", "1", "b=2", "abc", "-xyz"})
ap.Bool("-f", "--f")  // true // Query "-f" and "--fo" , (remove -f from Args).
ap.String("a")        // "1" true // Query "a", (remove "a", =, "1" from Args).
ap.String("b")        // "2" true // Query "b", (remove "b=2" from Args).
ap.After("--")        // ["abc", "-xyz"] "1" // After "--".
ap.Short("-").Args    // ["--fo", "--", "abc", "-x", "-y", "-z"] Separated by - "-", (remove -xyz from Args).
```
