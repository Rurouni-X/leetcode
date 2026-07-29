func concatWithReverse(nums []int) (res []int) {

    res = append(res, nums...)
    l := 0
    r := len(nums) - 1

    for l < r {
        nums[l], nums[r] = nums[r], nums[l]
        l++
        r--
    }
    res = append(res, nums...)
    return
}