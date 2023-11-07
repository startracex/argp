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
func (X *Argp) Short(short string) *Argp {
	for i, v := range X.Args {
		if strings.HasPrefix(v, short) && len(v) > len(short)+1 && v[len(short)] != short[0] {
			X.Args = append(X.Args[:i], X.Args[i+1:]...)
			for _, c := range v[len(short):] {
				X.Args = append(X.Args, "-"+string(c))
			}
		}
	}
	return X
}

// Bool query finds in the Args and returns its existence
func (X *Argp) Bool(finds ...string) bool {
	for _, find := range finds {
		for i, v := range X.Args {
			if v == find {
				X.Args = append(X.Args[:i], X.Args[i+1:]...)
				return true
			}
		}
	}
	return false
}

// BoolVar query finds in the Args and sets the exist result to value
func (X *Argp) BoolVar(value *bool, finds ...string) {
	for _, find := range finds {
		*value = X.Bool(find)
		if *value {
			return
		}
	}
}

// String query finds in the Args and returns the result and true or "" and false
func (X *Argp) String(finds ...string) (value string, exist bool) {
	for _, find := range finds {
		for i, v := range X.Args {
			if v == find && i+1 < len(X.Args) {
				// 删除相同值和之后的值
				value = X.Args[i+1]
				if value == Eq && i+2 < len(X.Args) {
					value = X.Args[i+2]
					X.Args = append(X.Args[:i], X.Args[i+1:]...)
				}
				X.Args = append(X.Args[:i], X.Args[i+2:]...)
				return value, true
			} else if strings.HasPrefix(v, find+Eq) && len(v) > len(find) {
				value = v[len(find)+1:]
				X.Args = append(X.Args[:i], X.Args[i+1:]...)
				return value, true
			}
		}
	}
	return "", false
}

// StringVar query finds in the Args and set the parsed result to value
func (X *Argp) StringVar(value *string, finds ...string) {
	for _, find := range finds {
		v, e := X.String(find)
		if e {
			*value = v
			return
		}
	}
}

// Start query the Args which start with find and return the result and true or "" and false
func (X *Argp) Start(find string) (string, bool) {
	for i, v := range X.Args {
		if strings.HasPrefix(v, find) && len(v) > len(find) {
			X.Args = append(X.Args[:i], X.Args[i+1:]...)
			return v[strings.Index(v, find):], true
		}
	}
	return "", false
}

// Attach return the Args after the "--"
func (X *Argp) Attach() []string {
	after, at := X.After(At)
	if at != -1 {
		X.Args = X.Args[:at]
	}
	return after
}

// After return the Args after the "--", do not delete the Args
func (X *Argp) After(attr string) ([]string, int) {
	for i, v := range X.Args {
		if v == attr && i+1 < len(X.Args) {
			value := X.Args[i+1:]
			return value, i
		}
	}
	return []string{}, -1
}

// Before return the Args before the "--", do not delete the Args
func (X *Argp) Before(attr string) ([]string, int) {
	for i, v := range X.Args {
		if v == attr && i > 0 {
			value := X.Args[:i]
			return value, i
		}
	}
	return []string{}, -1
}
