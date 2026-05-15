func hasDuplicate(nums []int) bool {
    freq:=make(map[int]struct{})

    for i:=0 ; i< len(nums);i++{
        if _ , exist :=freq[nums[i]];exist{
            return true 
        }

        freq[nums[i]]= struct{}{}
    }

    return false
}
