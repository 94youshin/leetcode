package main

import "fmt"

/**
 *给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个整数，并返回他们的数组下标。
 * 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
 */
func main() {
	nums := []int{3,2,4}
	target := 6

	fmt.Println(twoSum(nums, target))
}
// 解题思路: 暴力求解
// 拿到nums中的n1，在nums剩下的元素中找(target-n1)
func twoSum(nums []int, target int) []int{
	for i:=0;i<len(nums); i++ {
		n := nums[i]
		for j:=i+1;j<len(nums);j++ {
			if target-n == nums[j] {
				return []int{i,j}
			}
		}
	}
	return nil
}