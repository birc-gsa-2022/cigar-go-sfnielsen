// You can create modules at this level and they will be
// interpreted as under module birc.au.dk, so to import
// package `gsa` you need `import "birc.au.dk/gsa"`

package gsa

import "strings"

// Extract the edit operations from a pairwise alignment.
//
//  Args:
//      p: The first row in the pairwise alignment.
//      q: The second row in the pairwise alignment.
//
//  Returns:
//      The two strings without gaps and the list of edit operations
//      as a string.
func GetEdits(p, q string) (gapFreeP, gapFreeQ, edits string) {
	var sb_p, sb_q, sb_edits strings.Builder

	for i := 0; i < len(p); i++ {

		if p[i] == []byte("-")[0] {
			sb_edits.WriteString("I")
			sb_q.WriteByte(q[i])
			continue
		}
		if q[i] == []byte("-")[0] {
			sb_edits.WriteString("D")
			sb_p.WriteByte(p[i])
			continue
		}
		sb_edits.WriteString("M")
		sb_p.WriteByte(p[i])
		sb_q.WriteByte(q[i])

	}
	edits = sb_edits.String()
	gapFreeP = sb_p.String()
	gapFreeQ = sb_q.String()

	return gapFreeP, gapFreeQ, edits
}

//  Get the distance between p and the string that starts at x[i:]
//  using the edits.
//
//  Args:
//      p: The read string we have mapped against x
//      x: The longer string we have mapped against
//      i: The location where we have an approximative match
//      edits: The list of edits to apply, given as a string
//
//  Returns:
//      The distance from p to x[i:?] described by edits
func EditDist(p, x string, i int, edits string) int {
	pRow, xRow := LocalAlign(p, x, i, edits)
	result := 0
	for i := 0; i < len(pRow); i++ {
		if pRow[i] != xRow[i] {
			result++
		}
	}

	return result
}
