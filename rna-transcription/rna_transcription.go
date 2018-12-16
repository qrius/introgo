package strand

import "strings"

//ToRNA returns RNA complement of DNA input
func ToRNA(dna string) string {

	if dna == "" {
		return dna
	}
	rna := map[string]string{
		"G": "C",
		"C": "G",
		"T": "A",
		"A": "U",
	}
	var s strings.Builder

	for _, c := range dna {
		s.WriteString(rna[string(c)])
	}
	return s.String()
}
