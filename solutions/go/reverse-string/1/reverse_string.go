package reverse

func Reverse(input string) string {
    reversed := make([]rune,0, len(input))
    for _, c := range input {
    	reversed = append([]rune{c}, reversed...)  
    }
	return string(reversed)
}
