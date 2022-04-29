package template

import "testing"

func TestTemplate(t *testing.T) {
	content := EmailTemplate()
	if len(content) != 2268 {
		t.Fatalf("unexpected template size 2241 != %d", len(content))
	}
}
