package main

import (
	"io"
)

func main() {

}

func copyN(dest io.Writer, src io.Reader, length int64) (int64, error) {
	limitSrc := io.LimitReader(src, length)
	n, err := io.Copy(dest, limitSrc)
	if err != nil {
		return n, err
	}

	return n, nil
}
