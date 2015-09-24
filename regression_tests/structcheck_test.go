package regression_tests

import "testing"

func TestStructcheck(t *testing.T) {
	source := `package test

type test struct {
	unused int
}
`
	expected := Issues{
		{Linter: "structcheck", Severity: "warning", Path: "test.go", Line: 4, Col: 2, Message: "unused struct field .test.unused"},
	}
	ExpectIssues(t, "structcheck", source, expected)
}