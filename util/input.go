package util

import (
	"io"
	"os"
	"strings"
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

func MustInputLines(path string) []string {
	return strings.Split(MustInputString(path), "\n")
}
