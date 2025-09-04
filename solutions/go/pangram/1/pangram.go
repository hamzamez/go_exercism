package pangram

import "strings"

func IsPangram(input string) bool {
    letters := make(map[rune]int)
    for _, c := range strings.ToLower(input) {
        if c >= 'a' && c <= 'z' {
            letters[c] += 1
        }
    }
    return len(letters) == 26
	
}
