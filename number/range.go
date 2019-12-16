package number

func Chunk(n, p int, r func(l, r int)) {
	for idx := 0; idx < n; idx += p {
		right := idx + p
		if right > n {
			right = n
		}

		r(idx, right)
	}
}
