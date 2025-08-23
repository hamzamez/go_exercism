package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	re  := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
    return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re :=  regexp.MustCompile(`<(~|\*|=|-)*>`)
    return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`".*(?i)password(?-i).*"`)
    count := 0
    for _, s := range lines {
        matches := re.FindAllString(s, -1)
        if matches != nil {
            count += len(matches)
        }
    }
    return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line([0-9])*`)
    return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
    re := regexp.MustCompile(`User\s+(\w*)\s`)
	for i, l := range lines {
        username := re.FindStringSubmatch(l)
        if username != nil {
            lines[i] = "[USR] "+username[1]+" "+l
        }
    }
    return lines
}
