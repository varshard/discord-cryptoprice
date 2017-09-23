package main

import (
	"fmt"
	"os"
	"testing"
)

func TestEnvironmentAreSet(t *testing.T) {
	token := os.Getenv("SECRET")

	if len(token) == 0 {
		t.Errorf("TOKEN is not set")
	} else {
		fmt.Printf("TOKEN is %s", token)
	}
}
