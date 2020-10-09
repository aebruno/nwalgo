// Copyright 2015 Andrew E. Bruno. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package nwalgo

import (
	"testing"
)

func TestAlign(t *testing.T) {
	seqs := [][]string{
		[]string{"CGAGAGA", "GAGAGA", "CGAGAGA", "-GAGAGA", "CGAGAGA", "-GAGAGA"},
		[]string{"CGAGAGA", "GACC", "CGAGAGA", "-GACC--", "CGAGAGA--", "-----GACC"},
	}

	for _, a := range seqs {
		aln1, aln2, _ := Align(a[0], a[1], 1, -1, -1, false)
		if aln1 != a[2] || aln2 != a[3] {
			t.Errorf("Align(%s, %s)\n***GOT***\n%s\n%s\n***WANT***\n%s\n%s", a[0], a[1], aln1, aln2, a[2], a[3])
		}
		aln1, aln2, _ = Align(a[0], a[1], 1, -1, -1, true)
		if aln1 != a[4] || aln2 != a[5] {
			t.Errorf("Align(%s, %s)\n***GOT***\n%s\n%s\n***WANT***\n%s\n%s", a[0], a[1], aln1, aln2, a[4], a[5])
		}
	}

}

func BenchmarkAlign(b *testing.B) {
	seq1 := "GGAATTAATCCAGGTAATGGACCCCAAGAT"
	seq2 := "GCCAGGATTCCCAGATATGGCCAAGGTTCC"

	for i := 0; i < b.N; i++ {
		Align(seq1, seq2, 1, -1, -1, false)
	}
}
