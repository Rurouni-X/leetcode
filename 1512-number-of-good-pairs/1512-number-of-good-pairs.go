func numIdenticalPairs(nums []int) (res int) {
    count := make(map[int]int)

    for _, num := range nums {
        count[num]++
    }

    for _, c := range count {
        res += c * (c - 1) / 2
    }
    return
}