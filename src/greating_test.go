package main

import (
	"testing"
)

// test function
func TestReturnGeeks(t *testing.T) {
	actualString := SayHello("Hasanboy")
	expectedString := "Hello Hasanboy"
	if actualString != expectedString{
		t.Errorf("Expected String(%s) is not same as"+
		" actual string (%s)", expectedString,actualString)
	}
}

