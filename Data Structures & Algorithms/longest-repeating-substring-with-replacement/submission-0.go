func characterReplacement(s string, k int) int {
  

  maxLength := 0
  left := 0
  hash:= make([]int,26)
  maxFreq:=0

  for right := 0 ; right< len(s); right++{
	hash[s[right]-'A']++
    maxFreq= max(maxFreq,hash[s[right]-'A'] )

	windowLen:= right -left +1

	if windowLen - maxFreq > k {
		hash[s[left]-'A']--
		left ++
	}

    maxLength= max(maxLength , right -left +1)
	
  }

  return maxLength
}

func max (a, b int)int{
	if  a> b {
		return a
	}

	return b
}