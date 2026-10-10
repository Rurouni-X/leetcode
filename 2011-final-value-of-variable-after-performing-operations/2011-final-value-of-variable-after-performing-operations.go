import "strings"

func finalValueAfterOperations(operations []string) (res int) {
    for _, val := range operations {
        exist := strings.Contains(val, "+")

        if exist {
            res++
        } else {
            res--
        }
    }
    return
}