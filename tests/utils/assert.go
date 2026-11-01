package utils

import "testing"

func AssertNoPanic(t *testing.T, fn func()) {
	t.Helper()
	defer func() {
		if recovered := recover(); recovered != nil {
			t.Errorf("unexpected panic: %v", recovered)
		}
	}()
	fn()
}
