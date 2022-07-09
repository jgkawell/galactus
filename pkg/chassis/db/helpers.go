package db

// doesStringSliceContainString tells whether a string slice `a` contains a string of  `x``.
func doesStringSliceContainString(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}