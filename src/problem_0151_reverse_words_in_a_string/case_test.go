package solution

import "testing"

func TestReverseWords(t *testing.T) {
	cases := []struct {
		input  string
		expect string
	}{
		{
			input:  "the sky is blue",
			expect: "blue is sky the",
		},
		{
			input:  "  hello world  ",
			expect: "world hello",
		},
		{
			input:  "a good    example",
			expect: "example good a",
		},
	}

	for _, c := range cases {
		got := reverseWords(c.input)

		if got != c.expect {
			t.Fatalf("expect: %s, but got: %s\n", c.expect, got)
		}
	}
}
