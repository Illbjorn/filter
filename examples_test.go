package filter

import (
	"fmt"
	"iter"
	"testing"
)

var (
	strings = []string{"abcdefz", "ab", "cz"}
)

func seq[T any](inputs ...T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := range len(inputs) {
			if !yield(inputs[i]) {
				return
			}
		}
	}
}

var _ Filter[string] = filterStrStartsWithA

func filterStrStartsWithA(v string) bool {
	if v == "" {
		return false
	}
	return v[0] == 'A' || v[0] == 'a'
}

func filterStrEndsWithZ(v string) bool {
	if v == "" {
		return false
	}
	return v[len(v)-1] == 'Z' || v[len(v)-1] == 'z'
}

func TestExamples(t *testing.T) {
	// iter.Seq[string] => []string, only strings which start with [Aa]
	results := AndSeqs(seq(strings...), filterStrStartsWithA)
	must(len(results) == 2)
	must(results[0] == "abcdefz")
	must(results[1] == "ab")

	// iter.Seq[string] => []string, only strings which start with [Aa] AND end
	// with [Zz]
	results = AndSeqs(seq(strings...), filterStrStartsWithA, filterStrEndsWithZ)
	must(len(results) == 1)
	must(results[0] == "abcdefz")

	// iter.Seq[string] => []string, only strings which start with [Aa] OR end
	// with [Zz]
	results = OrSeqs(seq(strings...), filterStrStartsWithA, filterStrEndsWithZ)
	must(len(results) == 3)
	must(results[0] == "abcdefz")
	must(results[1] == "ab")
	must(results[2] == "cz")

	// []string => []string, only strings which start with [Aa] AND end
	// with [Zz]
	results = Ands(strings, filterStrStartsWithA)
	must(len(results) == 2)
	must(results[0] == "abcdefz")
	must(results[1] == "ab")
}

func must(cond bool, msgAndArgs ...any) {
	if !cond {
		if len(msgAndArgs) == 0 {
			panic(0)
		}
		msg := fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
		panic(msg)
	}
}
