/* Regina Richardson ik6089 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Your file will include one interface IPlayer
It has just one function protoype, guess(), which takes no parameters and returns an integer
*/
type IPlayer interface {
	guess()
}

/*
Game which includes
An IPlayer for the player making the guesses
Integers for the number of guesses made so far, the (random) number the user is trying to match, and the most recent guess
*/
type Game struct {
	randomNum    int
	guessesCount int
	IPlayer
}

/* Human which is defined as an empty struct */
type Human struct{}

/* Asks for the useer for the next number to guess */
func (h *Human) guess() (num int) {
	// asks the user for the next number to guess
	fmt.Println("Enter your next guess: ")
	var g int // next guess
	fmt.Scanf("%d", &g)
	return g
}

/*
Autoguess for an object that generates guesses. It includes
Intgers for the min and mac possible values
A pointer to the Game
*/
type Autoguess struct {
	minValue int
	maxValue int
	ptr      *Game
}

/* Will return an appropriate guess based on choosing the middle value of the possible remaining values*/
func (a *Autoguess) guess() (num int) {
	// will return an appropriate guess based on choosing the middle value of the possible
	var appropriateGuess = ((a.maxValue + a.minValue) / 2)
	fmt.Printf("The computer has chosen %d \n", appropriateGuess)
	return appropriateGuess
}

/*  This plays the game by calling the player's guess() to get the next guess and outputting the appropriate response*/
func (g *Game) play() {
	// This plays the game by calling the player's guess() to get the next guess and
	//outputting the appropriate response
	const one = 1
	const three = 3
	const win = "You win."
	const tooHigh = "Too high"
	const tooLow = "Too low"
	const lose = "You ran out of guesses. Game Over."
	fmt.Println("(y/n -- if n guesses will be generated for you) :")
	var ynchoice string // This variable holds the users y/n choice
	fmt.Scanf("%s", &ynchoice)
	switch ynchoice {
	case "y":
		var h Human
		for g.guessesCount < three { // While guesses is less than 3
			var usersGuess int = h.guess()
			if usersGuess > g.randomNum { // if the users number is larger than the generated number
				fmt.Println(tooHigh)
				g.guessesCount++
			} else if usersGuess < g.randomNum { // if the users number is less than the generated number
				fmt.Println(tooLow)
				g.guessesCount++
			} else {
				fmt.Println(win)
				break
			}
		}
		if g.guessesCount == three { // number of guesses equals 3
			fmt.Println(lose)
		}
		break

	case "n":
		a := Autoguess{
			minValue: 1,
			maxValue: 10,
			ptr:      g,
		}
		for a.ptr.guessesCount < three { // While guesses is less than 3
			var usersGuess2 int = a.guess()
			if usersGuess2 > a.ptr.randomNum { // if the users number is larger than the generated number
				fmt.Println(tooHigh)
				a.ptr.guessesCount++
				a.maxValue = (usersGuess2 - one)
			} else if usersGuess2 < a.ptr.randomNum { // if the users number is less than the generated number
				fmt.Println(tooLow)
				a.ptr.guessesCount++
				a.minValue = (usersGuess2 + one)
			} else {
				fmt.Println(win)
				break
			}
		}
		if a.ptr.guessesCount == three {
			fmt.Println(lose)
		}
		break
	}

}

func main() {
	fmt.Println("You have 3 guesses to guess a number from 1 to 10")
	fmt.Println("Do you want to make the guesses?")

	const randHigh = 9 // This is the high end used for the random number function
	const randLow = 1  // This is the low end used for the random number function
	// Creating random number for game
	rand.Seed(time.Now().UnixNano())       // Makes the random number different each time
	num := (rand.Intn(randHigh) + randLow) // Creates the random number and assigns it to num
	user := Game{
		randomNum:    num,
		guessesCount: 0,
	}
	ptrGame := &user
	ptrGame.play()

}
