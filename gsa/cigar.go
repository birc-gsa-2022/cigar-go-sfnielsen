// You can create modules at this level and they will be
// interpreted as under module birc.au.dk, so to import
// package `gsa` you need `import "birc.au.dk/gsa"`

package gsa

import (
	"strconv"
	"strings"
)

// Expand the compressed CIGAR encoding into the full list of edits.
//
//  Args:
//      cigar: A CIGAR string
//
//  Returns:
//      The edit operations the CIGAR string describes.
func CigarToEdits(cigar string) (edits string) {
	var sb strings.Builder
	i := 0
	for i < len(cigar) {
		numb, _ := strconv.Atoi(string(cigar[i]))
		op := string(cigar[i+1])
		i += 2

		sb.WriteString(strings.Repeat(op, numb))
	}
	edits = sb.String()
	return edits
}

// Encode a sequence of edits as a CIGAR.
//
//  Args:
//      edits: A sequence of edit operations
//
//  Returns:
//      The CIGAR encoding of edits.
func EditsToCigar(edits string) (cigar string) {
	cigar = ""
	return cigar
}
