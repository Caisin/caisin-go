package opcv

// P 对象转指针
func P[T any](obj T) *T {
	return &obj
}

// O 取指针对象
func O[T any](obj *T) T {
	return *obj
}
