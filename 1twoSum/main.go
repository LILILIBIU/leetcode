package main

import "fmt"

func main() {
	var (
		testNum [][]int
		testTag []int
	)

	nums1 := []int{2, 7, 11, 15}
	targetNum1 := 9
	testNum, testTag = append(testNum, nums1), append(testTag, targetNum1)

	nums2 := []int{1, 4, 7, 45}
	targetNum2 := 52
	testNum, testTag = append(testNum, nums2), append(testTag, targetNum2)

	for i, v := range testNum {
		sum := twoSum(v, testTag[i])
		fmt.Println(sum)
	}

}
func twoSum(nums []int, target int) []int {
	flag := map[int]int{}
	for i, num := range nums {
		res := target - num
		v, ok := flag[res]
		if ok {
			return []int{i, v}
		}
		flag[num] = i
	}
	return []int{}
}
