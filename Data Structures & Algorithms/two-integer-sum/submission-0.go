func twoSum(nums []int, target int) []int {
    var neededNums map[int]int = make(map[int]int)

    for i, val := range nums {
        if neededNums[val] > 0 {
            return []int{neededNums[val] -1, i}
        }

        neededNums[target - val] = i + 1
    }

    fmt.Println(neededNums)
    return []int{}
}
