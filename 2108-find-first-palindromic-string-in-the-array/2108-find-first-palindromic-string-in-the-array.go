func firstPalindrome(words []string) string {
    
    for _, word := range words {
        
        if palindrome(word) {
            return word
        }
    }
    return ""
}

func palindrome(word string) bool {

    l := 0
    r := len(word) - 1

    for l <= r {

        if word[l] != word[r] {
            return false
        }
        l++
        r--
    }
    return true
}