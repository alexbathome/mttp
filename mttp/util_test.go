package mttp

import "testing"

func TestInBetweenFunction(t *testing.T) {
	tt := map[string]struct {
		have          int
		haveStart     int
		haveEnd       int
		haveInclusion bool

		want bool
	}{
		"201 is between 200 and 300 (inclusive)": {
			have:          201,
			haveStart:     200,
			haveEnd:       300,
			haveInclusion: true,
			want:          true,
		},
		"154 is not between 200 and 300 (inclusive)": {
			have:          154,
			haveStart:     200,
			haveEnd:       300,
			haveInclusion: true,
			want:          false,
		},
		"200 is between 200 and 300 (inclusive)": {
			have:          200,
			haveStart:     200,
			haveEnd:       300,
			haveInclusion: true,
			want:          true,
		},
		"200 is not between 200 and 300 (exclusive)": {
			have:          200,
			haveStart:     200,
			haveEnd:       300,
			haveInclusion: false,
			want:          false,
		},
	}

	for _, test := range tt {
		result := isBetween(test.have, test.haveStart, test.haveEnd, test.haveInclusion)
		if result != test.want {
			t.Fail()
		}
	}
}
