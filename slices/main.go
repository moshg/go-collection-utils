package slices

// Include returns true if the slice contains the target element.
func Include[T comparable](slice []T, target T) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}
