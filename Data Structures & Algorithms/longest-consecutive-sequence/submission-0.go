func longestConsecutive(nums []int) int {
    m:=make(map[int]struct{})
    longest:=0

	for _, v :=range nums{
		m[v]=struct{}{}
	}
	
	for _, num:=range nums{
		if _ , exist:= m[num-1];!exist{
			current:=num
			count:=1

			for {
				if _, exist:=m[current+1];exist{
					current++
					count++
				}else{
                    break
				}
			}

			if count > longest {
				longest=count
			}
		}


	}

	return longest
}
