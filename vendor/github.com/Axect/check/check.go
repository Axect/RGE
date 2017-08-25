package check

// Contains check contain
func Contains(x string, A []string) bool {
	for _, elem := range A {
		if x == elem {
			return true
		}
	}
	return false
}
