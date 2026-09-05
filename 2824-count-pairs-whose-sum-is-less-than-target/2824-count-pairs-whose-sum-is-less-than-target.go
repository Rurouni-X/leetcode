func countPairs(nums []int, target int) (res int) {
    
    for i := 0; i < len(nums); i++ {

        for j := i + 1; j < len(nums); j++ {
            if nums[i] + nums[j] < target {
                res++
            }
        }
    }
    return
}