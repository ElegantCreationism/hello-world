package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSayHello(t *testing.T) {
	expectedGreeting := "hello, john"
	greeting := SayHello("john")

	assert.Equal(t, expectedGreeting, greeting)
}
