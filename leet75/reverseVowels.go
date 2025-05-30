package leet75

func ReverseVowels(s string) string {
    // Convert string to byte slice for mutation
    bytes := []byte(s)
    start, end := 0, len(bytes)-1

    // Two-pointer approach to reverse vowels
    for start < end {
		toContinue := false
        if !isVowel(bytes[start]) {
            start++
            toContinue = true
        }
        if !isVowel(bytes[end]) {
            end--
            toContinue = true
        }
		if toContinue {
			continue
		}
        // Swap vowels
        bytes[start], bytes[end] = bytes[end], bytes[start]
        start++
        end--
    }
    
    return string(bytes)
}

func isVowel(c byte) bool {
    // Check both lowercase and uppercase vowels
    vowels := "aeiouAEIOU"
    for i := 0; i < len(vowels); i++ {
        if c == vowels[i] {
            return true
        }
    }
    return false
}