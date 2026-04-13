func lengthOfLongestSubstring(s string) int {
    charIndex := make(map[byte]int)
    maxLen := 0
    left := 0

    for right := 0; right < len(s); right++ {
        ch := s[right]
        if idx, ok := charIndex[ch]; ok && idx >= left {
            left = idx + 1
        }
        charIndex[ch] = right
        if right - left + 1 > maxLen {
            maxLen = right - left + 1
        }
    }
    return maxLen
}