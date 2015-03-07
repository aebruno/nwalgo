package nwalgo

import (
    "fmt"
)

var UP    = 1
var LEFT  = 2
var NW    = 3
var NONE  = 4

func Align(a ,b string, match, mismatch, gap int) (string, string, int) {

    alen := len(a)+1
    blen := len(b)+1

    f := make([][]int, alen)
    pointer := make([][]int, alen)
    for i := range f {
        f[i] = make([]int, blen)
        pointer[i] = make([]int, blen)
    }

    for i:= 1; i < alen; i++ {
        f[i][0] = gap*i
        pointer[i][0] = UP
    }
    for j:= 1; j < blen; j++ {
        f[0][j] = gap*j
        pointer[0][j] = LEFT
    }

    pointer[0][0] = NONE
    for i:= 1; i < alen; i++ {
        for j:= 1; j < blen; j++ {
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

            p := NW;
            if max == hgap  {
                p = UP
            } else if max == vgap {
                p = LEFT
            }

            pointer[i][j] = p
            f[i][j] = max
        }
    }

    i := alen-1
    j := blen-1
    aln1 := ""
    aln2 := ""
    score := f[i][j]

    for p := pointer[i][j]; p != NONE; p = pointer[i][j] {
        if p == NW {
            aln1 = fmt.Sprintf("%c%s", a[i-1], aln1)
            aln2 = fmt.Sprintf("%c%s", b[j-1], aln2)
            i--
            j--
        } else if p == UP  {
            aln1 = fmt.Sprintf("%c%s", a[i-1], aln1)
            aln2 = fmt.Sprintf("%s%s", "-", aln2)
            i--
        } else if p == LEFT  {
            aln1 = fmt.Sprintf("%s%s", "-", aln1)
            aln2 = fmt.Sprintf("%c%s", b[j-1], aln2)
            j--
        }
    }

    return aln1,aln2,score
}
