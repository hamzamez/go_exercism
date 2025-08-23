package hamming

import "fmt"

func Distance(a, b string) (int, error) {
    if len(a) != len(b) {
        return 0, fmt.Errorf("The sequences must be of equat length") 
    }
    as := []rune(a)
    bs := []rune(b)
    distance := 0
    for i, c := range as {
        if c != bs[i] {
            distance += 1
        }
    }
	return distance, nil
}
