package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
    lToN := make(map[string]int)
    for k, v := range in {
        for _, c := range v {
            lToN[strings.ToLower(c)] = k
        }
    }
    
	return lToN
}
