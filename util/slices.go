package util

func MaxValue[T Ordered](slice []T) T {
	max := slice[0]

	for _, item := range slice {
		if item > max {
			max = item
		}
	}

	return max
}

func Sum[T Addable](slice []T) (sum T) {
	for _, value := range slice {
		sum += value
	}

	return
}
