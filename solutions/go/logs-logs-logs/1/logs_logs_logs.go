package logs
import "unicode/utf8"
import "fmt"

// Application identifies the application emitting the given log.
func Application(log string) string {
    runes := []rune(log)
	for i, c := range runes {
        fmt.Println("i = ", i, " c = ", c)
        switch c {
            case 0x2757 :
            	return "recommendation"
            case 0x1F50D :
            	return "search"
            case 0x2600  :
            	return "weather"
        }
    }
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
    runes := []rune(log)
    for i, c := range runes {
        if c == oldRune {
            runes[i] = newRune
        }
    }
    
	return string(runes)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
