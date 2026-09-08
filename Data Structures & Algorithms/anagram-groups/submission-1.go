func groupAnagrams(strs []string) [][]string {
    m := make(map[[26]int][]string)
    var res [][]string

    for _, str := range strs {
        var tmp [26]int

        for _, ch := range str {
            tmp[ch-'a']++
        }

        m[tmp] = append(m[tmp], str)
    }

    for _, strs := range m {
        res = append(res, strs)
    }

    return res
}