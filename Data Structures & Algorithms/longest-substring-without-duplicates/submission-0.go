func lengthOfLongestSubstring(s string) int {
    longest := 0 
    m := make(map[byte]bool)
    l := 0

    for i := 0; i < len(s); i++ {
        // Shrink the window from the left until the duplicate of s[i] is removed
        for m[s[i]] {
            m[s[l]] = false
            l++
        }

        // Now that the duplicate is gone, add the current character to the window
        m[s[i]] = true

        // Calculate the window size at EVERY step to ensure we don't miss anything
        ws := i - l + 1
        if ws > longest {
            longest = ws
        }
    }

    return longest
}
