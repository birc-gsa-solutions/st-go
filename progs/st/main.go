package main

import (
	"fmt"
	"os"

	"birc.au.dk/gsa"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: genome-file reads-file\n")
		os.Exit(1)
	}
	genome_file := os.Args[1]
	reads_file := os.Args[2]

	genome_st := map[string]gsa.SuffixTree{}
	for name, seq := range gsa.LoadFasta(genome_file) {
		genome_st[name] = *gsa.McCreight(seq)
	}

	gsa.ScanFastq(reads_file, func(rec *gsa.FastqRecord) {
		for chr_name, st := range genome_st {
			st.Search(rec.Read, func(i int) {
				cigar := fmt.Sprintf("%d%s", len(rec.Read), "M")
				gsa.PrintSam(rec.Name, chr_name, i, cigar, rec.Read)
			})
		}
	})
}
