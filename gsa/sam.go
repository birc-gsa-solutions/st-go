package gsa

import "fmt"

func PrintSam(sname, rname string, pos int, cigar, read string) {
	fmt.Printf("%s\t%s\t%d\t%s\t%s\n", sname, rname, pos+1, cigar, read)
}
