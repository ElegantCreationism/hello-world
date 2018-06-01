package main

import (
	"testing"
)

func TestSayHello(t *testing.T) {
	expectedGreeting := "hello, john"
	greeting := SayHello("john")


	if expectedGreeting != greeting {
		t.Fail()
	}

}
