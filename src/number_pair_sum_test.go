package main

import "testing"

func TestPairSum(t *testing.T) {
	cases := []struct {
		numbers []int
		total   int
		want    [][2]int
	}{
		{numbers: []int{1, 2, 3, 4, 2, 5}, total: 4, want: [][2]int{{1, 3}, {2, 2}}},
		{numbers: []int{1, 2, 3, 4, 2, 5}, total: 5, want: [][2]int{{2, 3}, {1, 4}}},
		{numbers: []int{1, 2, 3, 4, 2, 5}, total: 0, want: [][2]int{}},
		{numbers: []int{-8, 4, -2, 1, 1, 1, -3}, total: 2, want: [][2]int{{4, -2}, {1, 1}}},
		{numbers: []int{1, 1, 1, 1}, total: 2, want: [][2]int{{1, 1}, {1, 1}}},
	}

	for _, c := range cases {
		got := PairSum(c.numbers, c.total)

		if len(got) != len(c.want) {
			t.Errorf("PairSum(%v) = %v, want %v", c.numbers, got, c.want)
			// Test failed for this case, move onto next case
			continue
		}

		for i := range c.want {
			if len(got) == 0 || c.want[i][0] != got[i][0] || c.want[i][1] != got[i][1] {
				t.Errorf("PairSum(%v) = %v, want %v", c.numbers, got, c.want)
			}

			// Test failed for this case, we don't need to check the response any more.
			break
		}
	}
}
