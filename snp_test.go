package main

import "testing"

func TestDnaComplement(t *testing.T) {
	test := "AG"
	complement := dnaComplement(test)
	if complement != "TC" {
		t.Errorf("Expected TC, got %s\n", complement)
	}
}
