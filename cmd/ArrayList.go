func RotateArray(nums []int, k int) {
    n := len(nums)
    if n == 0 {
        return
    }
    k %= n

    reverse(nums, 0, n-1) 
    reverse(nums, 0, k-1) 
    reverse(nums, k, n-1) 
}

func reverse(nums []int, l, r int) {
    for l < r {
        nums[l], nums[r] = nums[r], nums[l]
        l++
        r--
    }
}