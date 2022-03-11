package main

func main() {

}

func removeDuplicates(nums []int) int {
	// Check for edge cases.
	if nums == nil {
		return 0
	}
	// Use the two pointer technique to remove the duplicates in-place.
	// The first element shouldn't be touched; it's already in its correct place.
	var writePointer int = 1
	// Go through each element in the Array.
	for readPointer := 1; readPointer < len(nums); readPointer++ {
		// If the current element we're reading is *different* to the previous
		// element...
		if nums[readPointer] != nums[readPointer-1] {
			nums[writePointer] = nums[readPointer]
			// And we need to now increment writePointer, because the next element
			// should be written one space over.
			writePointer++
		}
	}
	// This turns out to be the correct length value.
	return writePointer
}
