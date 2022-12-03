package util

import (
	"io"
	"os"
)

func MustInputBytes(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	contents, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	return contents
}

func MustInputString(path string) string {
	return string(MustInputBytes(path))
}
