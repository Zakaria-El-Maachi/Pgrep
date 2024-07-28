package main

import (
	"testing"
)

func TestPreprocess(t *testing.T) {
	str := []string{"ABABCABAB", "KEKEKEKEKEKKKKKEEEK", "ABABDABACDABABCABAB"}
	lpss := [][]int64{
		{0, 0, 1, 2, 0, 1, 2, 3, 4},
		{0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 1, 1, 1, 2, 0, 0, 1},
		{0, 0, 1, 2, 0, 1, 2, 3, 0, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4},
	}

	for i, v := range str {
		l := preprocess(v)
		for j, k := range l {
			if k != lpss[i][j] {
				t.Errorf("Failed Preprocessing at test %d", i)
			}
		}
	}

}

func TestKMP(t *testing.T) {
	pattern := "ABABCABAB"
	indices := kmp(pattern, []byte("ABABDABACDABABCABAB"), preprocess(pattern))
	tr := []int64{10}

	if len(tr) != len(indices) {
		t.Errorf("Failed KMP")
	}

	for i, k := range indices {
		if k != tr[i] {
			t.Errorf("Failed KMP")
		}
	}

}
