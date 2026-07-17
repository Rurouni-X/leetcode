func twoSum(nums []int, target int) []int {
    
    maps := make(map[int]int)

    for i, num := range nums {

        sum := target - num

        if index, ok := maps[sum]; ok {
            return []int{index, i}
        }
        maps[num] = i
    }
    return nil
}