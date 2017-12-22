package slices

type Strings []string

/**/
func StringSlicesJoin(stringSlices ...[]string) []string {
	var (
		totalSize int = 0
	)
	for _, stringSlice := range stringSlices {
		totalSize += len(stringSlice)
	}

	var (
		result []string = make([]string, totalSize)
		offset int      = 0
	)
	for _, stringSlice := range stringSlices {
		for i, str := range stringSlice {
			result[offset+i] = str
		}
		offset += len(stringSlice)
	}

	return result
}

/**/
func StringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, valA := range a {
		if valA != b[i] {
			return false
		}
	}
	return true
}

/*
IMPORTANT: does not sort in-place, allocates 2x the size of the input slice
as intermediate storage, returns a sorted duplicate of the input slice
*/
func (this Strings) Mergesort() []string {
	strings := []string(this)
	return StringsMergesort(strings)
}

func StringsMergesort(strings []string) []string {
	interim := make([]string, len(strings))
	result := make([]string, len(strings))
	for i, str := range strings {
		result[i] = str
	}
	stringsMergesort(result, interim)
	return result
}

func stringsMergesort(result []string, interim []string) {
	l := len(interim)
	if l <= 1 {
		// already sorted, do nothing
	} else if l == 2 {
		// sort the two children ascending
		if result[1] < result[0] {
			x := result[0]
			result[0] = result[1]
			result[1] = x
		}
	} else {
		// split and sort children, then merge
		la := int(l / 2) // la is the smaller part always, as fraction is removed
		// section 'a' is 0:a, whilst section 'b' is a:l
		stringsMergesort(result[0:la], interim[0:la])
		stringsMergesort(result[la:l], interim[la:l])
		// merge
		for i, ia, ib := 0, 0, la; i < l; i++ {
			if ia == la {
				interim[i] = result[ib]
				ib++

			} else if ib == l {
				interim[i] = result[ia]
				ia++

			} else if result[ia] < result[ib] {
				interim[i] = result[ia]
				ia++

			} else {
				interim[i] = result[ib]
				ib++
			}
		}
		// copy back
		for i := 0; i < l; i++ {
			result[i] = interim[i]
		}
	}
}
