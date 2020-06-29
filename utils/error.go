package utils

import "errors"

func ErrEq(e1, e2 error) bool {
	if e1 == nil || e2 == nil {
		if e1 == e2 {
			return true
		}

		return false
	}

	if e1 == e2 {
		return true
	}

	if errors.Is(e1, e2) || errors.Is(e2, e1) {
		return true
	}

	if e1.Error() == e2.Error() {
		return true
	}

	return false
}
