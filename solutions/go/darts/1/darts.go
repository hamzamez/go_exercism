package darts

import "math"

func Score(x, y float64) int {
    dart_r := math.Sqrt(x*x + y*y)
    switch{
        case dart_r <= 1 : return 10
        case dart_r <= 5 : return 5
        case dart_r <= 10 : return 1
    } 
    
	return 0
}
