import "strconv"

func mirrorDistance(n int) int {
    
    str := strconv.Itoa(n)
    arr := []rune(str)
    l := 0
    r := len(arr) - 1

    for l < r {
        arr[l], arr[r] = arr[r], arr[l]
        l++
        r--
    }

    num, _ := strconv.Atoi(string(arr))

    result := n - num

    if result < 0 {
        return -result
    }
    return result
}