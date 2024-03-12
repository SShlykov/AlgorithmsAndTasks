package container_with_most_water

func maxArea(height []int) (maxArea int) {
	left, right := 0, len(height)-1
	for left < right {
		area := (right - left) * min(height[left], height[right])
		if area > maxArea {
			maxArea = area
		}
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}
