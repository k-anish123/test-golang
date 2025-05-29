package leet75
func CanPlaceFlowers(flowerbed []int, n int) bool {

	// Get the length of the flowerbed array
	len := len(flowerbed)

	// Iterate through each position in the flowerbed
	for i := 0; i < len; i++ {
		// Check if the left neighbor is empty or if it's the first position
		left := i == 0 || flowerbed[i-1] == 0

		// Check if the right neighbor is empty or if it's the last position
		right := i == len-1 || flowerbed[i+1] == 0

		// If both neighbors are empty and the current position is empty
		if left && right && flowerbed[i] == 0 {
			// Place a flower at the current position
			flowerbed[i] = 1
			// Decrease the count of flowers to place
			n--
		}
	}

	// Return true if all flowers have been placed, otherwise false
	return n <= 0

    
}