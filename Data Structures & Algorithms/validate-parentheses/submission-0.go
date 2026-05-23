func isValid(s string) bool {
    st:=New()

	m:=map[byte]byte{
		')':'(',
		'}':'{',
		']': '[',
	}

	for i:=0 ; i < len(s);i++{
		char:= s[i]

		if char =='(' || char =='{' || char=='['{
			st.Push(char)
		}

		if char ==')' || char =='}' || char==']' {
			if len(st.eles) == 0 {
				return false
			}

			ele:= st.Pop()
            if ele != m[char]{
				return false
			}

		}
	}

	return len(st.eles) == 0
}

type stack struct{
	eles []byte
}

func New()*stack{
	return &stack{
		eles :make([]byte,0),
	}
}

func (s *stack)Push(ele byte){
    if s==nil{
		return
	}

	s.eles=append(s.eles, ele)
}

func (s *stack)Pop()byte{
	n:=len(s.eles)
	ele:=s.eles[n-1]
	s.eles=s.eles[:n-1]

	return ele
}
