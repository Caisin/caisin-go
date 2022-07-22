package math

func Min[T float64 | int64 | int | float32](x, y T) T {
	if x > y {
		return y
	}
	return x
}

func Max[T float64 | int64 | int | float32](x, y T) T {
	if x > y {
		return x
	}
	return y
}

func MulMin[T float64 | int64 | int | float32](nums ...T) (T, []T) {
	var min = nums[0]
	other := make([]T, 0)
	for i, num := range nums {
		if i == 0 {
			continue
		}
		if min > num {
			other = append(other, min)
			min = num
		} else {
			other = append(other, num)
		}
	}
	return min, other
}

func MulMax[T float64 | int64 | int | float32](nums ...T) (T, []T) {
	var max = nums[0]
	other := make([]T, 0)
	for i, num := range nums {
		if i == 0 {
			continue
		}
		if max < num {
			other = append(other, max)
			max = num
		} else {
			other = append(other, num)
		}
	}
	return max, other
}
