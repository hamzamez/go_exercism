package collatzconjecture

import "fmt"

func CollatzConjecture(n int) (int, error) {
	if n < 1 {
    	return 0, fmt.Errorf("The input number muber be positive")
    }
	steps := 0
    for n != 1{
        if n % 2 == 0 {
        	n = n / 2
        } else {
        	n = 3 * n + 1
        }
        steps += 1 // increment the number of steps

    }
    
    return steps, nil
    
}
