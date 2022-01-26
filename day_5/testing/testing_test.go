package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

/*
Unit testing is an important part of writing principled Go programs.
The testing package provides
the tools we need to write unit tests and the go test command runs tests.

Testing code typically lives in the same package as the code it tests.

Writing tests can be repetitive, so it’s idiomatic to use
table-driven style, where test inputs and expected outputs are listed
in a table and a single loop walks over them and performs the test logic.
*/
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	//t.Run enables running “subtests”,
	//one for each table entry.
	//These are shown separately when executing go test -v.

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := MinInteger(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}

	// also you can skip execution of some tests using
	//t.Skip()
}

func TestIntMinTableDrivenTestify(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	//t.Run enables running “subtests”,
	//one for each table entry.
	//These are shown separately when executing go test -v.

	for _, tc := range tests {
		testname := fmt.Sprintf("%d,%d", tc.a, tc.b)
		t.Run(testname, func(t *testing.T) {
			//t.Skip()
			got := MinInteger(tc.a, tc.b)
			require.EqualValues(t, tc.want, got)
		})
	}
}

// go test ./...
// go test ./... -v
// we can also generate output file that can be used by SonarQube for example
// go test -coverprofile=coverage.out.tmp ./...
// and we can check code cov of functions
// go tool cover -func=coverage.out
// go tool cover -html=coverage.out


// +build integration
//go test ./... -test.v -tags integration