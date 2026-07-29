func reverseWords(s string) string {
    arr := []rune(s)

    i := 0

    for i < len(arr) {
        if arr[i] == ' ' {
            i++
            continue
        }
        start := i
        for i < len(arr) && arr[i] != ' ' {
            i++
        }
        reverse(arr, start, i-1)
    }
    return string(arr)
}

func reverse(arr []rune, start, end int) {
    for start < end {
        arr[start], arr[end] = arr[end], arr[start]
        start++
        end--
    }
}