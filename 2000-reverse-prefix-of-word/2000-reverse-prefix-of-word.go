func reversePrefix(word string, ch byte) string {
    arr := []byte(word)
    l := 0
    var r int

    for i, v := range arr {
        if v == ch {
            r = i
            break
        }
    }

    for l <= r {
        arr[l], arr[r] = arr[r], arr[l]
        l++
        r--
    }
    if r == 0 && arr[0] != ch {
        return word
    }
    return string(arr)
}