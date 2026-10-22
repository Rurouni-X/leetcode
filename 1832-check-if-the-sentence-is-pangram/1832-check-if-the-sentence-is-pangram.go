func checkIfPangram(sentence string) bool {
    count := make(map[rune]struct{})

    for _, char := range sentence {
        count[char] = struct{}{}
    }

    return len(count) >= 26
}