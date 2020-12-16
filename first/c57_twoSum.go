package first

// 暴力
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums);i++ {
		for j := i + 1; j < len(nums);j++ {
			if nums[i]+nums[j] == target {
				return []int{nums[i], nums[j]}
			}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	n := len(nums)
	j := n - 1
	for i := 0; i < j; {
		if nums[i] + nums[j] == target {
			return []int{nums[i], nums[j]}
		} else if nums[i] + nums[j] > target{
			j--
		} else {
			i++
		}
	}
	return nil
}