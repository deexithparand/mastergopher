package slices

// Write a function that takes an array of integers and returns a new slice containing only the even numbers from the array. Then, write another function that appends a given integer to this slice.
func SliceMethod(arr [127]int32, size int) []int32 {

	var slice []int32

	// allows array size = 127 bits , 2^8-1, each integer will be of 4 bytes, with a range until 32 bit
	for i := 0; i < size; i++ {

		if arr[i]%2 == 0 {
			slice = append(slice, arr[i])
		}

	}

	// slice is a part of array, unlimited sizing,... can't have fixed length
	return slice
}
