func getSneakyNumbers(nums []int) []int {
    
    mp := make(map[int]struct{})
    res := make([]int, 0, 2)
    for _, num := range nums {
        if _, ok := mp[num]; ok {
            res = append(res, num)
        } else {
            mp[num] = struct{}{}
        }
    } 
    return res
}