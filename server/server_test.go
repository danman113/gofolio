package server

import (
	"fmt"
	. "testing"
	"time"
)

var (
	_ = time.Sleep
	_ = fmt.Sprintf
)

func TestSetupTemplates(t *T) {
	tests := []struct {
		folder string
		err    string
	}{
		{"template", "open template: no such file or directory"},
		{"templates", ""},
		{"../test/templates/unmatchedbrace", "template: error.html:3: unexpected \"}\" in command"},
		{"../test/templates/missingvalue", "template: error.html:3: missing value for command"},
	}

	for _, test := range tests {
		t.Run(" - "+test.folder, func(t *T) {
			err := setupTemplates(test.folder)
			if !(err == nil && test.err == "") &&
				!(err != nil && err.Error() == test.err) {
				t.Errorf("Expected error \"%s\", got error \"%s\"\n",
					test.err,
					err,
				)
			}
		})
	}
}

func BenchmarkSetupTemplates(b *B) {
	for i := 0; i < b.N; i++ {
		setupTemplates("templates")
	}
}
