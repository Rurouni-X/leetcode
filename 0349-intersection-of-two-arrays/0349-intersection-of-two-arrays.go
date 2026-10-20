func intersection(nums1 []int, nums2 []int) []int {
    res := make([]int, 0)
    count := make(map[int]struct{})

    for _, num := range nums1 {
        count[num] = struct{}{}
    }

    for _, num := range nums2 {
        if _, ok := count[num]; ok {
            res = append(res, num)
            delete(count, num)
        }
    }
    return res
}