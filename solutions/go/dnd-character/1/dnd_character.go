package dndcharacter

import (
    "math/rand"
    "sort"
    "time"
    "math"
    )

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor((float64(score) - 10) / 2))
}

func throwDice() int {
	rand.Seed(time.Now().UnixNano()) // Seed with current time
    return rand.Intn(6) + 1 // 1..6  
}
func sum(ints []int) int {
    sum := 0
    for _, v := range ints {
        sum += v
    }
    return sum
}
// Ability uses randomness to generate the score for an ability
func Ability() int {
    throws := []int{throwDice(), throwDice(), throwDice(), throwDice()}
    sort.Ints(throws)
    return sum(throws[1:])
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
    constitution := Ability()
    hitpoints := 10 + Modifier(constitution)
	return Character{
		Strength : Ability(),
		Dexterity : Ability(),
		Constitution : constitution,
		Intelligence : Ability(),
		Wisdom  : Ability(),
		Charisma : Ability(),
		Hitpoints : hitpoints,
	}
}
