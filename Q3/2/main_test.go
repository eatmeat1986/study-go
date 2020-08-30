package main

import (
	"os"
	"testing"
)

func TestGenerateRandBinary1024(t *testing.T) {
	generateRandBinary1024()
	file, err := os.Open("new.txt")
	if err != nil {
		panic(err)
	}

	rb1024 := make([]byte, 0)
	l, err := file.Read(rb1024)
	if err != nil {
		panic(err)
	}

	if l != 1024 {
		t.Errorf("invalid length, %d", l)
	}
}
