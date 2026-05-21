func maxArea(nums []int) int {
maxArea:=0
n:=len(nums)
l, r:= 0 ,n-1

for l < r{
	// calculate the area 

	area:= min(nums[l],nums[r])*(r-l)

	if area > maxArea{
		maxArea = area
	}

	// move the pillar which is of smaller height 

	if nums[l] < nums[r]{
		l++
	}else{
		r--
	}
}

return maxArea
}

func min (a, b int)int{
	if a < b {
		return a
	}

	return b
}
