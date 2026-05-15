func isAnagram(s string, t string) bool {

    if len(s)!=len(t){
        return false
    }

     freq:=make(map[byte]int)

     for i:=0 ; i<len(s);i++{
        freq[s[i]]++
     }


     for i:=0; i< len(t);i++{
        freq[t[i]]--
     }

     for _, f:= range freq{
        if f > 0 {
            return false
        }
     }

     return true
}
