func maximumCount(nums []int) int {
    
    var plus int
    var minus int

    for _, v := range nums {

        if v == 0 {
            continue
        }
        if v < 0 {
            minus++
        } else {
            plus++
        }
    }

    if plus > minus {
        return plus
    }
    return minus
}