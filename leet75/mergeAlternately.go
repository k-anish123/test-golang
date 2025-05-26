package leet75

func MergeAlternately(word1, word2 string) string {
    // Get lengths of both input strings
    n1, n2 := len(word1), len(word2)
    // Create a byte slice with enough space for both strings
    merged := make([]byte, n1+n2)
    // Initialize indices:
    // i: tracks position in word1
    // j: tracks position in word2
    // k: tracks position in merged result
    i, j, k := 0, 0, 0

    // Phase 1: Merge characters alternately while both strings have characters
    for i < n1 && j < n2 {
        merged[k] = word1[i]     // Add character from word1
        merged[k+1] = word2[j]   // Add character from word2
        i++                      // Move to next character in word1
        j++                      // Move to next character in word2
        k += 2                   // Move merged index by 2 positions
    }

    // Phase 2: Handle remaining characters from word1 (if any)
    for i < n1 {
        merged[k] = word1[i]
        i++
        k++
    }

    // Phase 3: Handle remaining characters from word2 (if any)
    for j < n2 {
        merged[k] = word2[j]
        j++
        k++
    }

    // Convert byte slice back to string and return
    return string(merged)
}
