package solution

import "testing"

var cases = []struct {
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

func TestReverseWords(t *testing.T) {
	for _, c := range cases {
		got := reverseWords(c.input)

		if got != c.expect {
			t.Fatalf("expect: %s, but got: %s\n", c.expect, got)
		}
	}
}

func TestReverseWordsInplace(t *testing.T) {
	for _, c := range cases {
		got := reverseWordsInplace(c.input)

		if got != c.expect {
			t.Fatalf("expect: %s, but got: %s\n", c.expect, got)
		}
	}
}
