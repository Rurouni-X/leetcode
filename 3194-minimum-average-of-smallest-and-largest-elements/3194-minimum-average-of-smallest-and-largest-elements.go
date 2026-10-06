import "slices"

func minimumAverage(nums []int) float64 {
    slices.Sort(nums)
    ln := len(nums)
    avg := make([]float64, 0, ln / 2)

    left := 0
    right := ln - 1

    for left < right {
        sum := (float64(nums[left]) + float64(nums[right])) / 2.0
        avg = append(avg, sum)
        left++
        right--
    }
    slices.Sort(avg)
    return avg[0]
}