func reversePrefix(s string, k int) string {
    arr := []rune(s)
    l := 0
    r := k - 1

    for l <= r {
        arr[l], arr[r] = arr[r], arr[l]
        l++
        r--
    }
    return string(arr)
}