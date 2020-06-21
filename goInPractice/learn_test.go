package main

import "testing"

// To be used with Listing 1.10 in learn.go
func Testname(t *testing.T) {
	name := getName()

	if name != "World" {
		t.Error("invalid response form getName")
	}
}
