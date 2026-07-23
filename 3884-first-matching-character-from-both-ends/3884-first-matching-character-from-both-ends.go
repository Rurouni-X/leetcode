func firstMatchingIndex(s string) int {
    arr := []rune(s)
    l := 0
    r := len(arr) - 1

    for l <= r {

        if arr[l] == arr[r] {
            return l
        }
        l++
        r--
    }
    return -1
}