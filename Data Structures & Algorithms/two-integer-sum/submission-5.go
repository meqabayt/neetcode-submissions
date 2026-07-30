func twoSum(nums []int, target int) []int {
    var neededNums map[int]int = make(map[int]int)

    for leftPointer:=0; leftPointer<=len(nums)/2; leftPointer++ {
        rightPointer := len(nums) - leftPointer - 1

        leftNum := nums[leftPointer]
        rightNum := nums[rightPointer]

        if neededNums[target - leftNum] > 0 {
            return resPrep(neededNums[target-leftNum]-1, leftPointer)
        }
        neededNums[leftNum] = leftPointer + 1
        
        if neededNums[target - rightNum] > 0 {
            return resPrep(neededNums[target-rightNum]-1, rightPointer)
        }
        neededNums[rightNum] = rightPointer + 1
    }

    return []int{}
}

func resPrep(num1, num2 int) []int {
    fmt.Println(num1, num2)
    if num1 > num2 {
        return []int{num2, num1}
    } else {
        return []int{num1, num2}
    }
}
