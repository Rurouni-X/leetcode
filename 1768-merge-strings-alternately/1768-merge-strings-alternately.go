func mergeAlternately(word1 string, word2 string) (res string){
    p1 := 0
    p2 := 0

    for p1 < len(word1) && p2 < len(word2) {
        res += string(word1[p1]) + string(word2[p2])
        p1++
        p2++
    }
    res += word1[p1:]
    res += word2[p2:]
    return
}