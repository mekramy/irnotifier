package irnotifier

import "fmt"

func alter[T comparable](v *T, alt T) T {
	var empty T
	if v == nil || *v == empty {
		return alt
	} else {
		return *v
	}
}

func responseError(status int) error {
	return fmt.Errorf("failed with status %d", status)
}

func isSuccess(status int) bool {
	return status >= 200 && status < 300
}
