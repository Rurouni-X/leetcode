func findWordsContaining(words []string, x byte) (res []int) {
    
    for i, word := range words {

        for j := 0; j < len(word); j++ {

            if word[j] == x {
                res = append(res, i)
                break
            }
        }
    }
    return
}