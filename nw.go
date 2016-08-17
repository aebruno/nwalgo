// Copyright 2015 Andrew E. Bruno. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package nwalgo

import (
	"fmt"
)

const (
	Up   = 1
	Left = 2
	NW   = 3
	None = 4
)

func Align(a, b string, match, mismatch, gap int) (string, string, int) {

	alen := len(a) + 1
	blen := len(b) + 1

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
	aln1 := ""
	aln2 := ""
	score := f[i][j]

	for p := pointer[i][j]; p != None; p = pointer[i][j] {
		if p == NW {
			aln1 = fmt.Sprintf("%c%s", a[i-1], aln1)
			aln2 = fmt.Sprintf("%c%s", b[j-1], aln2)
			i--
			j--
		} else if p == Up {
			aln1 = fmt.Sprintf("%c%s", a[i-1], aln1)
			aln2 = fmt.Sprintf("%s%s", "-", aln2)
			i--
		} else if p == Left {
			aln1 = fmt.Sprintf("%s%s", "-", aln1)
			aln2 = fmt.Sprintf("%c%s", b[j-1], aln2)
			j--
		}
	}

	return aln1, aln2, score
}
