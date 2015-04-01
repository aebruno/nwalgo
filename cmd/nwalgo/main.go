// Copyright 2015 Andrew E. Bruno. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
    "fmt"
	"flag"
	"log"
    "github.com/aebruno/nwalgo"
)

var seq1 = flag.String("seq1", "", "first sequence")
var seq2 = flag.String("seq2", "", "second sequence")
var match = flag.Int("match", 1, "match score")
var mismatch = flag.Int("mismatch", -1, "mismatch score")
var gap = flag.Int("gap", -1, "gap penalty")

func main() {
    flag.Parse()
    if *seq1 == "" || *seq2 == "" {
        log.Fatal("Please provide 2 sequences to align. See nwalgo --help")
    }

    aln1, aln2, score := nwalgo.Align(*seq1, *seq2, *match, *mismatch, *gap)
    fmt.Printf("%s\n%s\nScore: %d\n", aln1, aln2, score)
}
