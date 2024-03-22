package main

import "testing"

type testTableStruct struct {
	name  string
	input [][]int
	want  bool
}

// Given 2 arrays, if the values of the 2nd array contains all the squared values in the 1st array
// => should return true
// The frequency of all values in the 1st array must be matched in the 2nd array
// Inputs : 2 arrays of integer (normalArray, squaredArray)
// Outputs : a boolean (result)

// EXAMPLES
// [] & [] => false : empty inputs
// [1,2,3] & [4,1,9] => true : the order doesn't matter
// [1,2] & [4,2] => false
// [3, 4, 2] & [16, 4] => false  : because the frequency of 1st array's values is not matched in the 2nd array
// [4, 2] & [4, 16, 9, 25] => true : even if the 2nd array has more values, it still has all the squared values from the 1st array
// [3,3,3] & [9] => false : not same frequency
// [1,1,1,1,4,3,4,9,9] & [16,16,1,1,1,1,9, 81,81] => true
func TestSameFunc(t *testing.T) {

	testTables := []testTableStruct{
		{
			name: "Given two empty arrays [] & [], should return false",
			input: [][]int{
				{}, {},
			},
			want: false,
		},
		{
			name: "Given two arrays [1,2,3] & [4,1,9], should return true",
			input: [][]int{
				{1, 2, 3}, {4, 1, 9},
			},
			want: true,
		},
		{
			name: "Given two arrays [1,2] & [4,2], should return false",
			input: [][]int{
				{1, 2}, {4, 2},
			},
			want: false,
		},
		{
			name: "Given two arrays [3, 4, 2] & [16, 4], should return false",
			input: [][]int{
				{3, 4, 2}, {16, 4},
			},
			want: false,
		},
		{
			name: "Given two arrays [4, 2] & [4, 16, 9, 25], should return true",
			input: [][]int{
				{4, 2}, {4, 16, 9, 25},
			},
			want: true,
		},
		{
			name: "Given two arrays [3,3,3] & [9], should return false",
			input: [][]int{
				{3, 3, 3}, {9},
			},
			want: false,
		},
		{
			name: "Given two arrays [1,1,1,1,4,3,4,9,9] & [16,16,1,1,1,1,9, 81,81], should return true",
			input: [][]int{
				{1, 1, 1, 1, 4, 3, 4, 9, 9}, {16, 16, 1, 1, 1, 1, 9, 81, 81},
			},
			want: true,
		},
	}

	for _, tt := range testTables {
		ttValue := tt
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			normalArray := ttValue.input[0]
			squaredArray := ttValue.input[1]
			// Act
			ans := same(normalArray, squaredArray)
			// Assert
			if ans != ttValue.want {
				t.Errorf("got %v, want %v", ans, ttValue.want)
			}
		})
	}
}
