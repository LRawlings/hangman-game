package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var word string
	var prevGuesses []string
	var clue string
	var winner bool
	var loser bool
	var curGuess string
	wrongGuesses := 0
	reader := bufio.NewReader(os.Stdin)

	refreshScreen(99)
	fmt.Print("Setter, enter the secret word: ")
	word, _ = reader.ReadString('\n')
	word = strings.TrimSuffix(word, "\r\n")

	for i := 0; i < len(word); i++ {
		clue += "_ "
	}

	refreshScreen(wrongGuesses)
	for !winner && !loser && curGuess != "exit" {
		fmt.Printf("Here are your guesses so far: %q \n \n %s \n \n", prevGuesses, clue)

		fmt.Print("Guesser, what is your first guess? ")
		curGuess, _ = reader.ReadString('\n')
		curGuess = strings.TrimSuffix(curGuess, "\r\n")

		if curGuess != "exit" {

			for !checkGuessLength(curGuess) || !checkGuessPrevious(curGuess, prevGuesses) {
				if !checkGuessLength(curGuess) {
					refreshScreen(wrongGuesses)
					fmt.Printf("Here are your guesses so far: %q \n \n %s \n \n", prevGuesses, clue)
					fmt.Printf("You enetered %q: you can only guess one letter! Try again: ", curGuess)
				}
				if !checkGuessPrevious(curGuess, prevGuesses) {
					refreshScreen(wrongGuesses)
					fmt.Printf("Here are your guesses so far: %q \n \n %s \n \n", prevGuesses, clue)
					fmt.Printf("You enetered %q: you've already guessed %q! Try again: ", curGuess, curGuess)
				}
				curGuess, _ = reader.ReadString('\n')
				curGuess = strings.TrimSuffix(curGuess, "\r\n")
			}
			prevGuesses = append(prevGuesses, curGuess)
			wrongGuesses++
		}
		refreshScreen(wrongGuesses)
		if wrongGuesses == 9 {
			loser = true
		}
	}
}

func checkGuessLength(s string) bool {
	if len(s) != 1 {
		return false
	}
	return true
}

func checkGuessPrevious(s string, p []string) bool {
	for i := range p {
		if p[i] == s {
			return false
		}
	}
	return true
}

func refreshScreen(n int) {
	// Clear screen
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	header := `
	*************************************
	*                                   *
	*     Luke's First Go Game!         *
	*                                   *
	*            Hangman                *
	*                                   *
	*************************************
	`
	var hangman string

	switch n {
	case 99:
		hangman = `
     ______
      |   \|
    O_|    |
    /|\    |
    / \    |
__________/_\_`
	case 0:
		hangman = `





__________/_\_`
	case 1:
		hangman = `

           |
           |
           |
           |
__________/_\_`
	case 2:
		hangman = `
     ______
          \|
           |
           |
           |
__________/_\_`
	case 3:
		hangman = `
     ______
     |    \|
     |     |
           |
           |
__________/_\_`
	case 4:
		hangman = `
     ______
     |    \|
   O_|     |
           |
           |
__________/_\_`
	case 5:
		hangman = `
     ______
     |    \|
   O_|     |
    |      |
           |
__________/_\_`
	case 6:
		hangman = `
     ______
     |    \|
   O_|     |
   /|      |
           |
__________/_\_`
	case 7:
		hangman = `
     ______
     |    \|
   O_|     |
   /|\     |
           |
__________/_\_`
	case 8:
		hangman = `
     ______
     |    \|
   O_|     |
   /|\     |
   /       |
__________/_\_`
	case 9:
		hangman = `
     ______
     |    \|
   O_|     |
   /|\     |
   / \     |
__________/_\_
***********
*GAME OVER*
***********`
	}

	fmt.Println(header, hangman)
}
