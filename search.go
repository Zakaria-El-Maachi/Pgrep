package main

func preprocess(pattern string) []int64 {
	lps := make([]int64, len(pattern))
	l := int64(0)

	for i := 1; i < len(pattern); {
		if pattern[i] == pattern[l] {
			l++
			lps[i] = l
			i++
		} else {
			if l > 0 {
				l = lps[l-1]
			} else {
				i++
			}
		}
	}

	return lps
}

func kmp(pattern string, subject []byte, lps []int64) []int64 {
	if len(pattern) > len(subject) {
		return nil
	}
	indices := make([]int64, 0)

	for i, j := int64(0), int64(0); int64(len(subject))-i >= int64(len(pattern))-j; {
		if subject[i] == pattern[j] {
			j++
			i++
		}

		if j == int64(len(pattern)) {
			indices = append(indices, i-j)
			j = lps[j-1]
		} else if i < int64(len(subject)) && subject[i] != pattern[j] {
			if j > 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}

	return indices
}
