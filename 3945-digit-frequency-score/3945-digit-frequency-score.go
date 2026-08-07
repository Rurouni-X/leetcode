import "strconv"

func digitFrequencyScore(n int) (res int) {
    
    str := strconv.Itoa(n)

    for _, v := range str {
        num, _ := strconv.Atoi(string(v))
        res += num
    }
    return
}