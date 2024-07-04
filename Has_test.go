package extensions

import "testing"

func TestHas(t *testing.T) {
	if !Has("happy.txt", ".txt") {
		t.Fatal("happy.txt should have .txt")
	}
	if !Has("happy.jpg", ".jpg") {
		t.Fatal("happy.jpg should have .jpg")
	}
	if !Has("happy", "") {
		t.Fatal("happy has no extension and '' should return true")
	}
	if Has("happy", ".txt") {
		t.Fatal("happy has no extension and should have returned false")
	}
}
