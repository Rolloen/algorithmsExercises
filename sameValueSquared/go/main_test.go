package main

// Given 2 arrays, if the values of the 2nd array contains all the squared values in the 1st array
// => should return true
// The frequency of all values in the 1st array must be matched in the 2nd array
// Inputs : 2 arrays of integer (normalArray, squaredArray)
// Outputs : a boolean (result)

// EXAMPLES
// [1,2,3] & [4,1,9] => true : the order doesn't matter
// [1,2] & [4,2] => false
// [3, 4, 2] & [16, 4] => false  : because the frequency of 1st array's values is not matched in the 2nd array
// [4, 2] & [4, 16, 9, 25] => true : even if the 2nd array has more values, it still has all the squared values from the 1st array
// [3,3,3] & [9] => false : not same frequency
