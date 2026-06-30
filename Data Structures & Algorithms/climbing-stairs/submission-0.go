func climbStairs(n int) int {
      if n <=1 {
		return 1
	  }

	  prev1 , prev2 := 1,1 

	  for i:=2;i<=n ;i++{
		 curr := prev1+prev2
		 prev2=prev1
		 prev1=curr
	  }

	  return prev1
    
}
