package util

import (
	"testing"
)

func TestMoesiaParseInLocation(t *testing.T) {
	_, err := MoesiaParseInLocation("hogehoge")
	if err == nil {
		t.Errorf("Expected error does not appear")
	}
	_, err = MoesiaParseInLocation("2016年8月29日")
	if err != nil {
		t.Errorf("Parsed valid string validly")
	}
}

func TestMoesiaFormat(t *testing.T) {
	tm, _ := MoesiaParseInLocation("2016年8月29日")
	expected := "8/29 (月)"
	if str := tm.MoesiaFormat(); str != expected {
		t.Errorf("MoesiaFormat() has error: result '%s' expected '%s'", str, expected)
	}
}
