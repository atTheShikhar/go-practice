package main

func Sum(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func SumAll(numsToSum ...[]int) []int {
	var resultSlice []int

	for _, nums := range numsToSum {
		resultSlice = append(resultSlice, Sum(nums))
	}

	return resultSlice
}

func SumAllTails(numsToSum ...[]int) []int {
	var res []int
	for _, nums := range numsToSum {
		if len(nums) > 0 {
			tail := nums[1:]
			res = append(res, Sum(tail))
		} else {
			res = append(res, 0)
		}
	}
	return res
}
