package util

import "testing"

func TestUnique(t *testing.T) {
	result := Unique([]string{"t", "k", "a", "d"})

	if result != true {
		t.Fatal("A unique slice should return true but it doesn't")
	}
}

func TestUniqueNoUnique(t *testing.T) {
	result := Unique([]string{"t", "k", "k", "d"})

	if result != false {
		t.Fatal("A non unique slice should return false but it doesn't")
	}
}
