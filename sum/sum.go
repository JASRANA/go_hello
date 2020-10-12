package sum

func Sum(numbers []int) (sum int) {
	sum = 0
	// _ 占位符
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	//lengthOfNumbers := len(numbersToSum)
	//sums = make([]int, lengthOfNumbers) // 制造特定长度的数组
	//
	//for i, numbers := range numbersToSum {
	//	sums[i] = Sum(numbers)
	//}
	//
	//return
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return
}