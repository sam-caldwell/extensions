package extensions

import "testing"

func TestGet(t *testing.T) {
	t.Run("empty input", func(t *testing.T) {
		if Get("") != "" {
			t.Fatal("empty should return empty")
		}
	})
	t.Run("no extension", func(t *testing.T) {
		if Get("myFileWithoutExtension") != "" {
			t.Fatal("file without extension should return empty")
		}
	})
	t.Run("happy path 3-char", func(t *testing.T) {
		if Get("happy.txt") != ".txt" {
			t.Fatal("expect happy.txt to return .txt")
		}
	})
	t.Run("happy path 4-char", func(t *testing.T) {
		if Get("happy.text") != ".text" {
			t.Fatal("expect happy.text to return .text")
		}
	})
	t.Run("txt.happy", func(t *testing.T) {
		if Get("txt.happy") != ".happy" {
			t.Fatal("expect txt.happy to return .happy")
		}
	})
}
