package libs

import (
	"testing"
)

func TestParse(t *testing.T) {
	pos := []string{"名詞", "動詞"}
	parse := NewParse(pos, true)
	result := parse.Parse("寿司が食べたい。")

	if result != "寿司 食べる" {
		t.Errorf("Incorrect")
	}
}
