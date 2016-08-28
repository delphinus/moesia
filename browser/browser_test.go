package browser

import (
	"testing"
)

func TestBrowser(t *testing.T) {
	isTest = true
	b, err := New()
	if err != nil {
		t.Errorf("New() failed: %v", err)
	}
	if err = b.Start(); err != nil {
		t.Errorf("Start() failed: %v", err)
	}
	if err = b.End(); err != nil {
		t.Errorf("End() failed: %v", err)
	}
}
