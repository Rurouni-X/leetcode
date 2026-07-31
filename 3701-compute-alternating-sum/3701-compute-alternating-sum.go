func alternatingSum(nums []int) (res int) {
    
   for i := 0; i < len(nums); i++ {
    if i % 2 == 0 {
        res += nums[i]
    } else {
        res -= nums[i]
    }
   }
    return
}