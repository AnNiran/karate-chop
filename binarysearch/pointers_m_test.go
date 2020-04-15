package binarysearch

import(
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMp(t *testing.T) {
	testInput := []struct {
		target int
		data   []int
		res    int
	}{
		{3,    []int{}, -1},
		{1,    []int{1}, 0},
		{3,    []int{165}, -1},
		{1,    []int{1, 324, 946, 70, 25, 566, 241, 3, 78, 5}, 0},
		{234,  []int{1, 3, 87, 980, 98, 6758, 234, 4532, 544, 1233, 32, 5}, 6},
		{3224, []int{1, 987, 987, 546, 564, 3224, 9879, 4234, 242344, 57567, 47355, 23556, 5657}, 5},
		{0,    []int{1, 3, 796, 34, 45, 353, 535, 3525, 87, 5}, -1},
		{2,    []int{1, 564, 87, 34, 354, 54534, 56546, 8778, 5, 65664, 2554, 12, 3, 23, 4, 343445, 5}, -1},
		{4,    []int{1, 76, 45, 6767, 6745, 7475, 8979, 6747, 36342, 24, 44, 6567, 747, 577, 654, 5}, -1},
		{677,  []int{1, 356, 3, 435, 456, 677, 89998, 5432, 453, 3455, 3434, 556, 78, 5}, 9},
		{1,    []int{1, 356, 665, 6446, 678, 8979, 3454, 546, 6577, 77, 7456, 7457, 34, 54, 23, 12, 5, 7}, 0},
		{546,  []int{1, 3, 9, 54, 35, 52, 43, 546, 897, 656, 645, 667, 75, 5, 7}, 10},
		{5,    []int{1, 3, 5, 7}, 2},
	}

	for _, test := range testInput {
		sData = test.data
		sData = []int{}
		keys = []int {}
		//po = 0

		target = test.target

		res := mp()

		assert.Equal(t, test.res, res)
	}
}

