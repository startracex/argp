package argp

import (
	"testing"
)

func Test_Common(t *testing.T) {
	// ap.New() // from os.Args
	ap := From([]string{"-f", "--fo", "--", "a", "=", "1", "b=2", "abc", "-xyz"})
	t.Log(ap.Bool("-f", "--f")) //Query "-f" and "--fo" (true), (remove -f from Args).
	t.Log(ap.String("a"))       // Query "a" (1 true), (remove a , = , 1 from Args).
	t.Log(ap.String("b"))       // Query "b" (2 true), (remove b=2 from Args).
	t.Log(ap.After("--"))       // After "--" ([abc -xyz] 1).
	t.Log(ap.Short("-").Args)   // Separated by - "-" ([--fo -- abc -x -y -z]), (remove -xyz from Args).
}
