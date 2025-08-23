package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avg int) int {
    if avg <= 0 {
        return 2 * len(layers)
    }
    return avg * len(layers)
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64){
    sauceLayers := 0
    noodleLayers := 0
    for _, item := range layers {
        if item == "sauce" { sauceLayers += 1}
        if item == "noodles" { noodleLayers += 1}
    }
    return 50 * noodleLayers, 0.2 * float64(sauceLayers)
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList, myList []string) {
    lenMyList := len(myList)
    lenFriendList := len(friendList)
    if lenMyList > 0 && lenFriendList > 0 && myList[lenMyList - 1] == "?" {
        myList[lenMyList - 1] = friendList[lenFriendList - 1]
    }
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) []float64 {
	newAmounts := append([]float64{}, amounts...)
    for i, amount := range newAmounts {
        newAmounts[i] = amount * float64(portions) / 2
    } 
    return newAmounts
}


// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
