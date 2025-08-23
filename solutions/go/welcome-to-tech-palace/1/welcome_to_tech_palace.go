package techpalace
import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    border := strings.Repeat("*", numStarsPerLine)
	return border + "\n" + welcomeMsg + "\n" + border
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    old := []rune(oldMsg)
    startIndx := 0
    endIndx := len(old) - 1
	for _, c := range oldMsg {
        if c != '*' && c != '\n' && c != ' ' {
			break
        }
        startIndx += 1
    }

    for i := endIndx; i >= 0 ; i-- {
        c := old[i]
        if c != '*' && c != '\n' && c != ' '{
            break
        }
        endIndx -= 1
    }
    return string(old[startIndx:endIndx+1])
}
