package nwalgo

import (
    "fmt"
)

var UP   = 1
var LEFT = 2
var NW   = 3

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
    }
    for j:= 1; j < blen; j++ {
        f[0][j] = gap*j
    }

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

    for i != 0 && j != 0 {
        base1 := a[i-1]
        base2 := b[j-1]
        p := pointer[i][j]
        if p == NW {
            i--
            j--
            aln1 = fmt.Sprintf("%c%s", base1, aln1)
            aln2 = fmt.Sprintf("%c%s", base2, aln2)
        } else if p == UP  {
            i--
            aln1 = fmt.Sprintf("%c%s", base1, aln1)
            aln2 = fmt.Sprintf("%s%s", "-", aln2)
        } else if p == LEFT  {
            j--
            aln1 = fmt.Sprintf("%s%s", "-", aln1)
            aln2 = fmt.Sprintf("%c%s", base2, aln2)
        }
    }

    return aln1,aln2,score
}
