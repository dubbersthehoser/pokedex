package cli

import (
	"testing"
)

func TestCleanInpu(t *testing.T) {
	cases := []struct {
		input string
		expect []string
	}{
		{
			input: "  hello  world  ",
			expect: []string{"hello", "world"},
		},
		{
			input: "booo\nba",
			expect: []string{"booo", "ba"},
		},
		{
			input: "  \n leet\t\nno   \n",
			expect: []string{"leet", "no"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expect) {
			t.Errorf("len(actual)(%d) != len(expect)(%d)", len(actual), len(c.expect))
			continue
		}

		for i := range actual {
			if actual[i] != c.expect[i] {
				t.Errorf("actual: '%s' != expect: '%s'", actual[i], c.expect[i])
			}
			
		}
	}

}
