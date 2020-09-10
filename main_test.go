package main

import "testing"

func TestFunnyFunction(t *testing.T) {
	got := funnyFunction()
	if got != "Smile! :)" {
		t.Errorf("funnyFunction() = %s; expected %q", got, "Smile! :)")
	}
}
