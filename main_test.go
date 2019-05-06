package main

import "testing"

//Test 1
func TestIt(t *testing.T) {
	if false {
		t.Errorf("Test 1 failed; wanted true got false.")
	}
}