// You can create modules at this level and they will be
// interpreted as under module birc.au.dk, so to import
// package `gsa` you need `import "birc.au.dk/gsa"`

package gsa

import "strings"

// Align two sequences from a sequence of edits.
//
//  Args:
//      p: The first sequence to align.
//      q: The second sequence to align
//      edits: The list of edits to apply, given as a string
//
//  Returns:
//      The two rows in the pairwise alignment
func Align(p, q, edits string) (pRow, qRow string) {

	var sb_p, sb_q strings.Builder
	var idx_p, idx_q int

	mid := []rune("MID")
	// Align p and q based on edits
	for _, v := range edits {
		if v == mid[0] {
			sb_p.WriteString(string(p[idx_p]))
			sb_q.WriteString(string(q[idx_q]))
			idx_p++
			idx_q++
			continue
		}
		if v == mid[1] {
			sb_q.WriteByte(q[idx_q])
			sb_p.WriteString("-")
			idx_q++
			continue
		}
		if v == mid[2] {
			sb_p.WriteByte(p[idx_p])
			sb_q.WriteString("-")
			idx_p++

		}
	}

	pRow = sb_p.String()
	qRow = sb_q.String()

	return pRow, qRow
}

// Align two sequences from a sequence of edits.
//
//  Args:
//      p: The first sequence to align
//      x: The second sequence to align; we only align locally
//      i: Start position of the alignment in x
//      edits: The list of edits to apply, given as a string
//
//  Returns:
//      The two rows in the pairwise alignment
func LocalAlign(p, x string, i int, edits string) (pRow, xRow string) {
	pRow, xRow = "", ""
	// Align p and q based on edits
	return pRow, xRow
}
