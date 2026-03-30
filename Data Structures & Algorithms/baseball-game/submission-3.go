func calPoints(operations []string) int {
	nums := make([]int, len(operations))
	pointer := -1
	for _, op := range operations {
		switch op {
			case "+" :
				if pointer >= 1 {
					pointer++
					nums[pointer] = nums[pointer-1] + nums[pointer-2]
				}
			case "C" :
				if pointer >= 0 {
					pointer--
				}
			case "D" :
				if pointer >= 0 {
					pointer++
					nums[pointer] = nums[pointer-1] + nums[pointer-1]
				}
			default :
				pointer++
				if op[0] == '-' {
					nums[pointer] = convert(op[1:]) * -1
				} else {
					nums[pointer] = convert(op)
				}
		}
	}
	sum := 0
	for i:=0; i<=pointer; i++ {
		sum += nums[i]
	}
	return sum
}

func convert(ops string) int {
	val := 0
	for i:=0; i< len(ops); i++ {
		val = val*10 + int(ops[i] - '0')
	}
	return val
}
