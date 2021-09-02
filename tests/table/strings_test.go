package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (part: %s) - index: expected (%d) <> found (%d)."

func TestIndex(t *testing.T) {
	tests := []struct {
		text     string
		part     string
		expected int
	}{
		{"Go is good", "Go", 0},
		{"", "", 0},
		{"Hi", "hi", -1},
		{"Ramon", "o", 3},
	}

	for _, test := range tests {
		t.Logf("Mass: %v", tests)
		current := strings.Index(test.text, test.part)

		if current != test.expected {
			t.Errorf(msgIndex, test.text, test.part, test.expected, current)
		}
	}
}
