package mathtools

func PowUint64(base, exp uint64) uint64 {
	var val, i uint64 = 1, 0
	for ; i < exp; i++ {
		val *= base
	}
	return val
}
