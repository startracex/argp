package argp

import (
	"os"
	"strings"
)

type Argp struct {
	Args   []string
	Origin []string
}

// Eq equal sign
var Eq = "="

// At attach sign
var At = "--"

// From set the Args and Origin from Origin and return the Argp
func From(Origin []string) *Argp {
	return &Argp{
		Args:   Origin,
		Origin: Origin,
	}
}

// New set the Args and Origin from os.Args and return the Argp
func New() *Argp {
	init := From(os.Args[1:])
	init.Short("-")
	return init
}

// Short separate short parameters into additional parameter lists based on short
func (A *Argp) Short(short string) *Argp {
	shortLen := len(short)
	for i, arg := range A.Args {
		if len(arg) > shortLen && arg[0:shortLen] == short && !strings.Contains(arg[shortLen:], short) {
			A.Remove(i, 1)
			for _, c := range arg[shortLen:] {
				A.Args = append(A.Args, short+string(c))
			}
		}
	}
	return A
}

// Bool query finds in the Args and returns its existence
func (A *Argp) Bool(finds ...string) bool {
	for _, find := range finds {
		i := A.IndexOf(find)
		if i > -1 {
			A.Remove(i, 1)
			return true

		}
	}
	return false
}

// BoolVar query finds in the Args and sets the exist result to value
func (A *Argp) BoolVar(value *bool, finds ...string) {
	*value = A.Bool(finds...)
}

// String query finds in the Args and returns the result and true or "" and false
func (A *Argp) String(finds ...string) (value string, exist bool) {
	for _, find := range finds {
		for i, arg := range A.Args {
			if arg == find && i+1 < len(A.Args) {
				value = A.Args[i+1]
				if value == Eq && i+2 < len(A.Args) {
					value = A.Args[i+2]
					A.Remove(i, 3)
				} else {
					A.Remove(i, 2)
				}
				return value, true
			}
			if strings.HasPrefix(arg, find+Eq) && len(arg) > len(find) {
				value = arg[len(find)+1:]
				A.Remove(i, 1)
				return value, true
			}
		}
	}
	return "", false
}

// StringVar query finds in the Args and set the parsed result to value
func (A *Argp) StringVar(value *string, finds ...string) {
	s, exist := A.String(finds...)
	if exist {
		*value = s
		return
	}
}

// Start query the Args which start with find and return the result and true or "" and false
func (A *Argp) Start(find string) (string, bool) {
	for i, arg := range A.Args {
		indexOf := strings.Index(arg, find)
		if indexOf != -1 && len(arg) > len(find) {
			A.Remove(i, 1)
			return arg[indexOf:], true
		}
	}
	return "", false
}

// End query the Args which start with find and return the result and true or "" and false
func (A *Argp) End(find string) (string, bool) {
	for i, arg := range A.Args {
		indexOf := strings.Index(arg, find)
		if indexOf != -1 && len(arg) > len(find) {
			A.Remove(i, 1)
			return arg[:indexOf], true
		}
	}
	return "", false

}

// Attach return the Args after the first At
func (A *Argp) Attach() []string {
	after, at := A.After(At)
	if at != -1 {
		A.Args = A.Args[:at]
	}
	return after
}

// Before return the Args before the first attr and index of attr or -1, do not delete the Args
func (A *Argp) Before(attr string) ([]string, int) {
	for i, arg := range A.Args {
		if arg == attr && i > 0 {
			value := A.Args[:i]
			return value, i
		}
	}
	return []string{}, -1
}

// After return the Args after the first attr and index of attr or -1, do not delete the Args
func (A *Argp) After(attr string) ([]string, int) {
	for i, arg := range A.Args {
		if arg == attr && i+1 < len(A.Args) {
			value := A.Args[i+1:]
			return value, i
		}
	}
	return []string{}, -1
}

// Remove length elements from the index of Args
func (A *Argp) Remove(index, length int) {
	A.Args = append(A.Args[:index], A.Args[index+length:]...)
}

// IndexOf return index of find or -1
func (A *Argp) IndexOf(find string) int {
	return A.IndexOfFunc(find, func(arg string, find string) bool {
		return arg == find
	})
}

// IndexOfFunc return index of fn returns true or -1
func (A *Argp) IndexOfFunc(find string, fn func(string, string) bool) int {
	for i, arg := range A.Args {
		if fn(arg, find) {
			return i
		}
	}
	return -1
}
