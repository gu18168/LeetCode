package solution

import "testing"

func TestAsteroidCollision(t *testing.T) {
	cases := []struct {
		input  []int
		expect []int
	}{
		{
			input:  []int{5, 10, -5},
			expect: []int{5, 10},
		},
		{
			input:  []int{5, -5},
			expect: []int{},
		},
		{
			input:  []int{10, 2, -5},
			expect: []int{10},
		},
		{
			input:  []int{-2, -2, 1, -2},
			expect: []int{-2, -2, -2},
		},
	}

	for _, c := range cases {
		got := asteroidCollision(c.input)

		if len(got) != len(c.expect) {
			t.Fatalf("unexpect got, input=%v, expect=%v, got=%v", c.input, c.expect, got)
		}

		for i := range c.expect {
			if got[i] != c.expect[i] {
				t.Fatalf("unexpect got, input=%v, expect=%v, got=%v", c.input, c.expect, got)
			}
		}
	}
}
