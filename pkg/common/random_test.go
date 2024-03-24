package common

import "testing"

func TestRandomString(t *testing.T) {
	s := RandomString(32)
	if len(s) != 32 {
		t.Error("expected rnadom string of 32 characters atleast.")
	}
}
