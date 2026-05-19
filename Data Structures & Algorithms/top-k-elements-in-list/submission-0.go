func topKFrequent(nums []int, k int) []int {
// optimal solution is using the bucket sort
    freq:=make(map[int]int)

    for i:=0;i<len(nums);i++{
        freq[nums[i]]++
    }

    bucket:=make([][]int,len(nums)+1)

    for num, f:=range freq{
        bucket[f]=append(bucket[f],num)
    }

    res:=[]int{}

    for i:=len(bucket)-1;i>=0;i--{
        for _ , v :=range bucket[i]{
            res=append(res , v)

            if len(res)==k {
                return res
            }
        }
    }

    return res
}
