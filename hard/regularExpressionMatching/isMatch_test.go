package regularExpressionMatching

import "testing"

func TestFunc(t *testing.T) {
	for i, c := range getCases() {
		result := isMatch(c.InputString, c.InputPattern)
		if result != c.Expected {
			t.Fatalf("[%d] InputString: %v; InputPattern: %v; Result: %v; Expected: %v;",
				i, c.InputString, c.InputPattern, result, c.Expected)
		}
	}
}

type Case struct {
	InputString  string
	InputPattern string
	Expected     bool
}

func getCases() []Case {
	return []Case{
		{
			InputString:  "aa",
			InputPattern: "a",
			Expected:     false,
		},
		{
			InputString:  "aa",
			InputPattern: "a*",
			Expected:     true,
		},
		{
			InputString:  "aab",
			InputPattern: "c*a*b",
			Expected:     true,
		},
		{
			InputString:  "mississippi",
			InputPattern: "mis*is*p*.",
			Expected:     true,
		},
		{
			InputString:  "aa",
			InputPattern: "a",
			Expected:     false,
		},
	}
}
