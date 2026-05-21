func threeSum(nums []int) [][]int {
	// sorting is required cause we need to make the skip logic work properly 
	sort.Ints(nums)
   n := len(nums)

   res:=[][]int{}

   for i:=0 ; i< n-2 ; i++{
	 // skip if duplucate 

	 if i > 0 && nums[i]==nums[i-1]{
		continue
	 }

	 left:=i+1
	 right := n-1

	 for left < right {
		sum:=nums[i]+nums[left]+nums[right]
		if sum==0{
			res=append(res , []int{nums[i], nums[left],nums[right]})

			for left < right && nums[left]==nums[left+1]{
				left++
			}

			for left < right && nums[right]==nums[right-1]{
				right--
			}

			left++
			right--
		}else if sum < 0 {
			left ++
		}else{
			right--
		}
	}
   }

   return res
}
