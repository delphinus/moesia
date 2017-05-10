package mail

import (
	"strings"
	"testing"

	"github.com/delphinus/moesia/config"
)

func TestFailedToSend(t *testing.T) {
	m := New(&config.Config{})
	expected := "failed to DialAndSend"
	if err := m.Send(""); err != nil && strings.Index(err.Error(), "failed to DialAndSend") == -1 {
		t.Errorf("error has no expected words. expected: %v, result: %v", expected, err)
	}
}
