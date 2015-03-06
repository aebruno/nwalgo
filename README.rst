===============================================================================
nwalgo - Needleman-Wunsch Alignment in Go
===============================================================================

-------------------------------------------------------------------------------
About
-------------------------------------------------------------------------------

An implementation of the Needleman-Wunsch global alignment algorithm [1] in Go.
Computes the alignment score and optimal global alignment.

-------------------------------------------------------------------------------
Install
-------------------------------------------------------------------------------

Fetch from github::

    $ go get install github.com/aebruno/nwalgo
    $ go get install github.com/aebruno/nwalgo/nwalgo

-------------------------------------------------------------------------------
Usage
-------------------------------------------------------------------------------

Align 2 DNA sequences::

    $ nwalgo -seq1 GAAAAAAT -seq2 GAAT 
    GAAAAAAT
    GAA----T
    Score: 0

From code::

    package main

    import (
        "github.com/aebruno/nwalgo"
    )

    func main() {
        aln1, aln2, score := nwalgo.Align("GAAAAAAT", "GAAT", 1, -1, -1)
    }


-------------------------------------------------------------------------------
License
-------------------------------------------------------------------------------

nwalgo was written by Andrew E. Bruno and released under the GNU General
Public License ("GPL") Version 3.0.  See the LICENSE file.

-------------------------------------------------------------------------------
References
-------------------------------------------------------------------------------

[1] http://en.wikipedia.org/wiki/Needleman-Wunsch_algorithm
