func flipAndInvertImage(image [][]int) [][]int {
    
   for _, val := range image {
    reverse(val)
   }
   return image
}

func reverse(nums []int) []int {
    l := 0
    r := len(nums) - 1

    for l < r {
        nums[l], nums[r] = nums[r], nums[l]

        if nums[l] == 1 {
            nums[l] = 0
        } else if nums[l] == 0 {
            nums[l] = 1
        }

        if nums[r] == 1 {
            nums[r] = 0
        } else if nums[r] == 0 {
            nums[r] = 1
        }

        l++
        r--
    }
    if l == r {
        if nums[l] == 1 {
            nums[l] = 0
        } else {
            nums[l] = 1
        }
    }
    return nums
}