func recoverOrder(order []int, friends []int) []int {
    isFriend := make(map[int]struct{}, len(friends))
    res := make([]int, 0, len(friends))

    for _, id := range friends {
        isFriend[id] = struct{}{}
    }

    for _, id := range order {
        if _, ok := isFriend[id]; ok {
            res = append(res, id)
        }
    }
    return res
}