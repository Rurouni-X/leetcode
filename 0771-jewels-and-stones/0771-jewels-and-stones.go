func numJewelsInStones(jewels string, stones string) (res int) {
    count := make(map[rune]struct{})

    for _, char := range jewels {
        count[char] = struct{}{}
    }

    for _, char := range stones {
        if _, ok := count[char]; ok {
            res++
        }
    }
    return
}