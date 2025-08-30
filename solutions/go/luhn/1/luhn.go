package luhn

import(
    "regexp"

)

func Valid(id string) bool {
    if !isDigits(id) {
        return false
    }

    var s []int = cleanSpaces(id)

    if len(s) <= 1 {
        return false
    }
    
    s = doubleSeconds(s)
	sum := sumDigits(s)

    return sum % 10 == 0

}

func isDigits(id string) bool {
    re := regexp.MustCompile(`^[0-9 ]+$`)
	return re.MatchString(id)
}

func cleanSpaces(id string) []int {
    s := make([]int, 0, len([]rune(id)))
    for _, c := range id {
        if c != ' ' {
            s = append(s , int(c - '0'))
        }
    }
    return s
}

func doubleSeconds(s []int) []int {
    for i := len(s) - 2; i >= 0; i-=2 {
        s[i] *= 2
        if s[i] > 9 {
            s[i] -= 9
        }
    }   
    return s
}

func sumDigits(id []int) int {
    sum := 0
    for _, d := range id {
        sum += d
    }
    return sum
}


