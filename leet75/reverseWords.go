package leet75

import "strings"

func ReverseWords(s string) string {
    // Trim spaces and split into words
    words := strings.Fields(s)
    
    // Reverse words array
    start, end := 0, len(words)-1
    for start < end {
        // Swap words
        words[start], words[end] = words[end], words[start]
        start++
        end--
    }
    
    // Join words with single space
    return strings.Join(words, " ")
}