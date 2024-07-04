package extensions

import "testing"

func TestHasOneOf(t *testing.T) {
	if !HasOneOf("happy.txt", ".txt", ".jpg") {
		t.Fatal("happy.txt should have .txt")
	}
	if !HasOneOf("happy.jpg", ".txt", ".jpg") {
		t.Fatal("happy.txt should have .jpg")
	}
	if HasOneOf("happy.bad", ".txt", ".jpg") {
		t.Fatal("happy.bad should return false on .txt and .jpg")
	}
}
