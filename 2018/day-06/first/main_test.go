package main

import "testing"

func TestFindBottomRight(t *testing.T) {
	coordinates := []xy{
		xy{2, 6},
		xy{8, 3},
		xy{3, 4},
		xy{1, 1},
		xy{5, 5},
		xy{7, 9},
	}

	bottomRight := xy{8, 9}

	br := findBottomRight(coordinates)
	if bottomRight != br {
		t.Errorf("expected bottomRight %+v, got %+v", bottomRight, br)
	}
}

func TestDistance(t *testing.T) {
	coordinates := [][]xy{
		[]xy{xy{1, 1}, xy{3, 4}},
		[]xy{xy{1, 4}, xy{2, 4}},
		[]xy{xy{9, 7}, xy{3, 6}},
	}

	expected := []int{5, 1, 7}

	for i, c := range coordinates {
		if distance(c[0], c[1]) != expected[i] {
			t.Errorf(
				"expected distance of %+v to %+v = %d, got %d",
				c[0],
				c[1],
				expected[i],
				distance(c[0], c[1]),
			)
		}
	}
}
