package argp

import (
	"testing"
)

var ap = From([]string{"-f", "--fo", "--", "a", "=", "1", "b=2", "abc", "-xyz"})

func Test_Common(t *testing.T) {
	// ap.New() // from os.Args
	t.Log(ap.Bool("-f", "--f")) // true // Query "-f" and "--fo" , (remove -f from Args).
	t.Log(ap.String("a"))       // "1" true // Query "a", (remove "a", =, "1" from Args).
	t.Log(ap.String("b"))       // "2" true // Query "b", (remove "b=2" from Args).
	t.Log(ap.After("--"))       // ["abc", "-xyz"] "1" // After "--".
	t.Log(ap.Short("-").Args)   // ["--fo", "--", "abc", "-x", "-y", "-z"] Separated by - "-", (remove -xyz from Args).
}
