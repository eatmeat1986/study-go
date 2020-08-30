package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestBuildASCII(t *testing.T) {
	result := captureStdout(buildASCII)

	if result != "ASCII" {
		t.Errorf("ASCIIではなかったです。%v", result)
	}
}

// ヘルパー関数
func captureStdout(f func()) string {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	stdout := os.Stdout
	os.Stdout = w

	f()

	os.Stdout = stdout
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}
