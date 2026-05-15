func twoSum(nums []int, target int) []int {
    m:=make(map[int]int)

    for i:=0;i<len(nums);i++{
        if v, exist:=m[target-nums[i]];exist{
            return []int{v,i}
        }

        m[nums[i]]=i
    }

    return []int{}
}
