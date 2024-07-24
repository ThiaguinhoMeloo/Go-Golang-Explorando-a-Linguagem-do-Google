package functions

func SimpleRecursion(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * SimpleRecursion(n-1)
	}
}
