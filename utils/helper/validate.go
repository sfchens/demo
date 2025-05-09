package helper

import "strconv"

type Number interface {
	~int | ~int32 | ~int64 | ~float64 | ~float32 | ~string | ~uint
}

func IsValidNumber[T Number](value T) bool {
	switch v := any(value).(type) {
	case int:
		return v > 0
	case int32:
		return v > 0
	case int64:
		return v > 0
	case float64:
		return v > 0
	case float32:
		return v > 0
	case string:
		if num, err := strconv.ParseFloat(v, 64); err == nil {
			return num > 0
		}
	case uint:
		return v > 0
	}
	return false
}
