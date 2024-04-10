package solution

import "testing"

func TestCloseStrings(t *testing.T) {
	cases := []struct {
		word1  string
		word2  string
		expect bool
	}{
		{
			word1:  "abbzccca",
			word2:  "babzzczc",
			expect: true,
		},
	}

	for _, c := range cases {
		got := closeStrings(c.word1, c.word2)
		if got != c.expect {
			t.Fatalf("word1: %s, word2: %s, expect: %v", c.word1, c.word2, c.expect)
		}
	}
}
