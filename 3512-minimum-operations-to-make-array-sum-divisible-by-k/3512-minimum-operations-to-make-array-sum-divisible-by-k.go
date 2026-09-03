func sum(arr []int) (res int) {
    for _, v := range arr {
        res += v
    }
    return res
}

func minOperations(nums []int, k int) int {
    
    s := sum(nums)

    if s % k == 0 {
        return 0
    }
    return s % k
}