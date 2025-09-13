package grains

import (
    "fmt"
    "math"
    )

func pow(i, n uint64) uint64 {
    return uint64(math.Pow(float64(i), float64(n)))
}
func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
        return 0, fmt.Errorf("Square number has to be between 1 and 64 : %v", number)
    }
    return pow(2, uint64(number-1)), nil
}

func Total() uint64 {
    var total uint64 = 0
	for i:=1; i < 65; i++ {
        s, _ := Square(i)
        total += s
    }
    return total;
}
