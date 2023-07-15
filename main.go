package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	playKey := rand.Intn(int(time.Now().UnixNano()))

	balance := 100
	bet := 0
	numberDice := 2

	fmt.Println("Welcome to dice bet")
	fmt.Println("Your balance is: ", balance)

	fmt.Println("Your play key is:", playKey)

	for balance > 0 {
		fmt.Println("Select your value for bet")
		fmt.Scan(&bet)

		if balance < bet {
			fmt.Println("No found!")
			break
		}

		fmt.Println("Select a one value beetwen 2 & 12")
		fmt.Scan(&numberDice)

		if bet == 0 {
			fmt.Println("Bet not valid!")
			break
		}

		if bet > balance {
			fmt.Println("Your bet exceeds your balance")
			continue
		}

		diceOne := rand.Intn(6) + 1
		diceTwo := rand.Intn(6) + 1

		valuesOfDices := diceOne + diceTwo

		fmt.Println("The dices run", diceOne, diceTwo)
		fmt.Println("---------------------")

		if valuesOfDices == numberDice {
			balance += bet
			fmt.Println("You win! - Your balance updated", balance)
		} else {
			balance -= bet
			fmt.Println("You lose! - Your balance updated", balance)
		}

		if balance == 0 {
			fmt.Println("Your balance is 0, for continue, add more balance")
		}
	}

}
