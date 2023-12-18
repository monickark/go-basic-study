package lib

import "testing"

func TestPow(t *testing.T) {
	v := Pow(2, 3)
	if v != 8 {
		t.Error("TestFailed")
	}
}
