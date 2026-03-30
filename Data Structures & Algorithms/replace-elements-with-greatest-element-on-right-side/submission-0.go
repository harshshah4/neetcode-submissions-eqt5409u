func replaceElements(arr []int) []int {
	n := len(arr)
	max := arr[n - 1]
	for i:= n-1; i >= 0 ; i-- {
		temp := max
		if max < arr[i] {
			max = arr[i]
		}
		arr[i] = temp
	}
	arr[n-1] = -1
	return arr
}
