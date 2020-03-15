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
	var wrongGuesses int

	reader := bufio.NewReader(os.Stdin)

	for word == "" {
		refreshScreen(99)
		fmt.Print("PLAYER 1\nEnter the secret word: ")
		word, _ = reader.ReadString('\n')
		word = strings.TrimSuffix(word, "\r\n")
		word = strings.ToLower(word)
		if !checkGuessBlank(word) {
			word = ""
		}
		if !checkGuessAlpha(word) {
			word = ""
		}
	}

	for i := range word {
		if string(word[i]) == " " {
			clue += "  "
		} else {
			clue += "_"
		}
	}

	refreshScreen(wrongGuesses)
	for !winner && !loser && curGuess != "exit" {
		player2Question(prevGuesses, clue, curGuess, "What's your guess? ", wrongGuesses)
		curGuess = readCurGuess(curGuess, reader)

		if curGuess != "exit" {
			for !checkGuessLength(curGuess) || !checkGuessPrevious(curGuess, prevGuesses) || !checkGuessAlpha(curGuess) {
				switch {
				case curGuess == "exit":
				case !checkGuessBlank(curGuess):
					player2Question(prevGuesses, clue, curGuess, "You've entered a blank guess! Try again: ", wrongGuesses)
					curGuess = readCurGuess(curGuess, reader)
				case !checkGuessLength(curGuess):
					player2Question(prevGuesses, clue, curGuess, "You've entered too many letters! Try again: ", wrongGuesses)
					curGuess = readCurGuess(curGuess, reader)
				case !checkGuessPrevious(curGuess, prevGuesses):
					player2Question(prevGuesses, clue, curGuess, "You've already guessed this letter! Try again: ", wrongGuesses)
					curGuess = readCurGuess(curGuess, reader)
				case !checkGuessAlpha(curGuess):
					player2Question(prevGuesses, clue, curGuess, "That's not a letter! Try again: ", wrongGuesses)
					curGuess = readCurGuess(curGuess, reader)
				}

			}
			prevGuesses = append(prevGuesses, curGuess)
			if letterCheck(curGuess, word) {
				clue = updateClue(curGuess, clue, word)
				if clue == word {
					winner = true
				}
			} else {
				wrongGuesses++
			}
		}
		refreshScreen(wrongGuesses)
		if curGuess == "exit" {
		} else if winner {
			fmt.Printf("PLAYER 2\nHere are your guesses so far: %q \n \n %s \n \n", prevGuesses, buildPrintClue(clue))
			fmt.Printf("Number of incorrect guesses so far: %d\n", wrongGuesses)
			fmt.Print(`
*******************
*WE HAVE A WINNER!*
*******************`)
		} else {
			if wrongGuesses == 9 {
				loser = true
			}
		}
	}
}

func player2Question(prevGuesses []string, clue, curGuess, errMessage string, wrongGuesses int) {
	refreshScreen(wrongGuesses)
	fmt.Printf("PLAYER 2\nHere are your guesses so far: %q \n \n %s \n \n", prevGuesses, buildPrintClue(clue))
	fmt.Printf("Number of incorrect guesses so far: %d\n", wrongGuesses)
	fmt.Printf(errMessage)
}

func readCurGuess(curGuess string, reader *bufio.Reader) string {
	curGuess, _ = reader.ReadString('\n')
	curGuess = strings.TrimSuffix(curGuess, "\r\n")
	curGuess = strings.ToLower(curGuess)
	return curGuess
}

func buildPrintClue(clue string) string {
	var printClue string
	for i := range clue {
		printClue += string(clue[i])
		printClue += " "
	}
	return printClue
}

func letterCheck(letter, word string) bool {
	for i := range word {
		if string(word[i]) == letter {
			return true
		}
	}
	return false
}

func updateClue(letter, clue, word string) string {
	var newClue string
	for i := range clue {
		if string(word[i]) == letter {
			newClue += letter
		} else if string(clue[i]) == "_" {
			newClue += "_"
		} else {
			newClue += string(word[i])
		}
	}
	return newClue
}

func checkGuessBlank(s string) bool {
	if len(s) == 0 {
		return false
	}
	return true
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

func checkGuessAlpha(s string) bool {

	if s < "a" || s > "z" {
		return false
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
