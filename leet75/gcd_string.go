package leet75

// GcdOfStrings finds the largest string X such that X divides both str1 and str2
// For example: str1 = "ABCABC", str2 = "ABC" -> "ABC"
func GcdOfStrings(str1 string, str2 string) string {
    // Check if strings can be divided by a common substring
    // If str1 + str2 != str2 + str1, no common divisor exists
    if str1+str2 != str2+str1 {
        return ""
    }

    // Calculate GCD of string lengths
    // This will give us the length of the largest common divisor
    gcdLength := gcd(len(str1), len(str2))
    
    // Return the substring of length GCD
    // This substring is guaranteed to be the largest common divisor
    return str1[:gcdLength]
}

// gcd calculates the Greatest Common Divisor of two numbers using Euclidean algorithm
// For example: gcd(6,3) = 3, gcd(12,8) = 4
func gcd(a, b int) int {
    // Use Euclidean algorithm:
    // Keep dividing larger number by smaller and take remainder
    // Until remainder becomes 0
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

// Example:
// str1 = "ABCABC", str2 = "ABC"
// str1 + str2 = "ABCABCABC"
// str2 + str1 = "ABCABCABC"
// gcd(6,3) = 3
// Result: "ABC"