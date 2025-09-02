package strand

func ToRNA(dna string) string {
    toRNA := map[rune]rune{'G':'C', 'C':'G', 'T':'A', 'A':'U'}
    rna := ""
    for _, c := range dna {
        rna += string(toRNA[c])
    }
	return rna
}
