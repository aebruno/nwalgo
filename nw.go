// Copyright 2015 Andrew E. Bruno. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package nwalgo

const (
	Up   = 1
	Left = 2
	NW   = 3
	None = 4
)

func Align(a, b string, match, mismatch, gap int) (alignA, alignB string, score int) {

	alen := len(a) + 1
	blen := len(b) + 1

	maxLen := alen
	if maxLen < blen {
		maxLen = blen
	}

	aBytes := make([]byte, 0, maxLen)
	bBytes := make([]byte, 0, maxLen)

	f := make([][]int, alen)
	pointer := make([][]int, alen)
	for i := range f {
		f[i] = make([]int, blen)
		pointer[i] = make([]int, blen)
	}

	for i := 1; i < alen; i++ {
		f[i][0] = gap * i
		pointer[i][0] = Up
	}
	for j := 1; j < blen; j++ {
		f[0][j] = gap * j
		pointer[0][j] = Left
	}

	pointer[0][0] = None
	for i := 1; i < alen; i++ {
		for j := 1; j < blen; j++ {
			match_mismatch := mismatch
			if a[i-1] == b[j-1] {
				match_mismatch = match
			}

			max := f[i-1][j-1] + match_mismatch
			hgap := f[i-1][j] + gap
			vgap := f[i][j-1] + gap
			if hgap > max {
				max = hgap
			}
			if vgap > max {
				max = vgap
			}

			p := NW
			if max == hgap {
				p = Up
			} else if max == vgap {
				p = Left
			}

			pointer[i][j] = p
			f[i][j] = max
		}
	}

	i := alen - 1
	j := blen - 1

	score = f[i][j]

	for p := pointer[i][j]; p != None; p = pointer[i][j] {
		if p == NW {
			aBytes = append(aBytes, a[i-1])
			bBytes = append(bBytes, b[j-1])
			i--
			j--
		} else if p == Up {
			aBytes = append(aBytes, a[i-1])
			bBytes = append(bBytes, '-')
			i--
		} else if p == Left {
			aBytes = append(aBytes, '-')
			bBytes = append(bBytes, b[j-1])
			j--
		}
	}

	reverse(aBytes)
	reverse(bBytes)

	return string(aBytes), string(bBytes), score
}

func reverse(a []byte) {
	for i := 0; i < len(a)/2; i++ {
		j := len(a) - 1 - i
		a[i], a[j] = a[j], a[i]
	}
}
