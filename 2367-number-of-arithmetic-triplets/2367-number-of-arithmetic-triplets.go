func arithmeticTriplets(nums []int, diff int) (res int) {
    
    mp := make(map[int]bool)

    for _, x := range nums {

        if mp[x-diff] && mp[x-2*diff] {
            res++
        }
        mp[x] = true
    }
    return
}