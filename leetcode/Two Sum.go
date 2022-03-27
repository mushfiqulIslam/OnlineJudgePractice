func twoSum(nums []int, target int) []int {
    var s []int
    for i, num := range nums {
        for j:= i+1; j<len(nums); j++ {
            if num + nums[j] == target {
                s = append(s, i)
                s = append(s, j)
                break
            }
        }
    }
    return s
    
}
