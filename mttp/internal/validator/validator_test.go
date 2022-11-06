package validator

import (
	"strings"
	"testing"
)

func TestValidateServerName(t *testing.T) {
	tt := map[string]struct {
		have             string
		wantErrorStrings []string
	}{
		"server name should not pass validation": {
			have: "my$$rver%name",
			wantErrorStrings: []string{
				"$ found in string at [2 3]\n",
				"found in string at [8]\n",
			},
		},
	}

	for _, test := range tt {
		r := ValidateServerName(test.have)
		for _, s := range test.wantErrorStrings {
			if !strings.Contains(r.Error(), s) {
				t.Fail()
			}
		}
	}
}
