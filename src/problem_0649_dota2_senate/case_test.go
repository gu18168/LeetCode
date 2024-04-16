package solution

import "testing"

func TestPredictPartyVictory(t *testing.T) {
	cases := []struct {
		input  string
		expect string
	}{
		{
			input:  "DRRDRDRDRDDRDRDR",
			expect: "Radiant",
		},
		{
			input:  "DR",
			expect: "Dire",
		},
		{
			input:  "RD",
			expect: "Radiant",
		},
		{
			input:  "RDD",
			expect: "Dire",
		},
		{
			input:  "DDRRR",
			expect: "Dire",
		},
		{
			input:  "DRDRRDDRDDR",
			expect: "Dire",
		},
	}

	for _, c := range cases {
		got := predictPartyVictory(c.input)
		if got != c.expect {
			t.Fatalf("input=%v, expect=%v, got=%v", c.input, c.expect, got)
		}
	}
}
