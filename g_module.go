package gmodule

func Add(nums ...int) (sum int) {
	for _, v := range nums {
		sum += v
	}
	return sum
}
