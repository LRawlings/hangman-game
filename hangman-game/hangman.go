package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	refreshScreen(0)
	var word string
	prevGuesses := []string{"s"}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Setter, enter the secret word: ")

	word, _ = reader.ReadString('\n')
	word = strings.TrimSuffix(word, "\r\n")

	refreshScreen(1)
	var clue string
	for i := 0; i < len(word); i++ {
		clue += "_ "
	}

	fmt.Printf("Here are your guesses so far: %q \n \n %s \n \n", prevGuesses, clue)

	fmt.Print("Guesser, what is your first guess? ")
	curGuess, _ := reader.ReadString('\n')
	curGuess = strings.TrimSuffix(curGuess, "\r\n")

	for !checkGuessLength(curGuess) || !checkGuessPrevious(curGuess, prevGuesses) {
		if !checkGuessLength(curGuess) {
			refreshScreen(1)
			fmt.Printf("Here are your guesses so far: %q \n \n %s \n \n", prevGuesses, clue)
			fmt.Printf("You enetered %q: you can only guess one letter! Try again: ", curGuess)
		}
		if !checkGuessPrevious(curGuess, prevGuesses) {
			refreshScreen(1)
			fmt.Printf("Here are your guesses so far: %q \n \n %s \n \n", clue)
			fmt.Printf("You enetered %q: you've already guessed %q! Try again: ", curGuess, curGuess)
		}
		curGuess, _ = reader.ReadString('\n')
		curGuess = strings.TrimSuffix(curGuess, "\r\n")
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
	case 0:
		hangman = `
     ______
      |   \|
    O_|    |
    /|\    |
    / \    |
__________/_\_`
	case 1:
		hangman = `





__________/_\_`
	case 2:
		hangman = `

           |
           |
           |
           |
__________/_\_`
	case 3:
		hangman = `
     ______
          \|
           |
           |
           |
__________/_\_`
	}

	fmt.Println(header, hangman)
}
