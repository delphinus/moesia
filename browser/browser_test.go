package browser

import (
	"github.com/sclevine/agouti"
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

func TestGetTexts(t *testing.T) {
	isTest = true
	getTextTexts = []string{"", "hoge", "fuga", ""}
	b, _ := New()
	defer b.End()
	var texts []string
	var err error
	if texts, err = b.getTexts(&agouti.MultiSelection{}); err != nil {
		t.Errorf("getTexts() failed: %v", err)
	}
	if len(texts) != 2 || texts[0] != getTextTexts[1] || texts[1] != getTextTexts[2] {
		t.Errorf("getTexts() returns invalid data: %v", texts)
	}
}
