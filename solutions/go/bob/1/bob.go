package bob

import(
    "strings"
    "unicode"
    )

func Hey(remark string) string {
    isQuestion := strings.HasSuffix(strings.TrimSpace(remark), "?")
    isYelling  := !stringHas(remark, func(c rune) bool{ return unicode.IsLetter(c) && !unicode.IsUpper(c) }) && 
                   stringHas(remark, unicode.IsLetter)
    isSilence  := !stringHas(remark, func(c rune) bool{ return !unicode.IsSpace(c)}) || remark == ""

    switch{
        case isQuestion && isYelling : return "Calm down, I know what I'm doing!"
        case isQuestion : return "Sure."
        case isSilence  : return "Fine. Be that way!"
        case isYelling  : return "Whoa, chill out!"

    }

	return "Whatever."
}


func stringHas(s string, f func(rune) bool) bool {
    for _, c := range s {
        if f(c) {
            return true
        }
    }
    return false
}
