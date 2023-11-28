package main

func Sum(nums []int) int {
	res := 0
	for i := range nums {
		res += nums[i]
	}
	return res
}

func SumAll(numbersToSum ...[]int) []int {
	res := []int{}
	for _, nums := range numbersToSum {
		res = append(res, Sum(nums))
	}
	return res
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, nums := range numbersToSum {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {
			tail := nums[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
