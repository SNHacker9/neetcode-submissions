func minWindow(s string, t string) string {
    if len(t) > len(s){
		return ""
	}

	countT:= make(map[byte]int)

	for i:=0 ; i < len(t);i++ {
		countT[t[i]]++
	}

	need := len(countT)
	have:=0

	window:= make(map[byte]int)

	resLen:= len(s) +1

	left := 0
	leftStart := 0

	for right := 0 ; right < len(s) ; right ++{
		char := s[right]
		window[char]++

		// check if char is what we needed then update the have
		if countT[char]> 0 && window[char]==countT[char]{
			have++
		}

		// when window is the valid one 

		for have == need{
			// update the answer
			if right-left +1 < resLen{
				resLen = right-left +1
				leftStart= left
			}

			// shrink window 
			leftChar := s[left]
			window[leftChar]--

			if countT[leftChar] > 0 && window[leftChar]< countT[leftChar]{
				have--
			}

			left++
		}
	}

	if resLen==len(s)+1{
		return ""
	}

	return s[leftStart:leftStart+resLen]
}
