package argp

import (
	"testing"
)

// from os.Args
// var ap = New()

// from []string
var ap = From([]string{"-f", "--fo", "--", "a", "=", "1", "b=2", "abc", "-xyz"})

func Test_Common(t *testing.T) {

	// true // Query "-f" and "--fo" , (remove -f from Args).
	t.Log(ap.Bool("-f", "--fo"))

	// "1" true // Query "a", (remove "a", =, "1" from Args).
	t.Log(ap.String("a"))

	// "2" true // Query "b", (remove "b=2" from Args).
	t.Log(ap.String("b"))

	// ["abc", "-xyz"] 1 // After "--".
	t.Log(ap.After("--"))

	// ["--fo", "--", "abc", "-x", "-y", "-z"] Separated by - "-", (remove -xyz from Args).
	t.Log(ap.Short("-").Args)

}
