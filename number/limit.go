package number

func Max(args ...int) (max int) {
	for idx, arg := range args {
		if idx == 0 || arg > max {
			max = arg
		}
	}

	return
}

func Min(args ...int) (min int) {
	for idx, arg := range args {
		if idx == 0 || arg < min {
			min = arg
		}
	}

	return
}

func Limit(limit int, max int, args ...int) int {
	if limit > 0 && limit < max {
		return limit
	}

	return Min(append(args, max)...)
}
