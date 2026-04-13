func nextGreaterElement(nums []int) []int {
    n := len(nums)
    result := make([]int, n)
    stack := []int{} 

    for i := 0; i < n; i++ {
        result[i] = -1 
    }

    for i := 0; i < n; i++ {
        for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
            idx := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            result[idx] = nums[i] 
        }
        stack = append(stack, i)
    }

    return result
}