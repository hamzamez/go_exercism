package isogram

import "strings"

func IsIsogram(word string) bool {
    m := make(map[rune]int)
    for _, c := range strings.ToUpper(word) {
        m[c] += 1
        if m[c] > 1 && c != ' ' && c != '-' {
            return false
        }
    }
    return true

}
