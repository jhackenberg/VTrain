package session

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (session *Session) Run() bool {
	pair, _ := session.Vocabulary.RandomPair()
	fmt.Printf("%s: ", pair.A.Phrase)
	r := bufio.NewReader(os.Stdin)
	answer, err := r.ReadString('\n')
	if err != nil {
		return false
	}
	answer = strings.Trim(answer, " \n")
	if answer == pair.B.Phrase {
		fmt.Println("\x1b[32mCorrect\x1b[0m")
	} else {
		fmt.Printf("\x1b[31mIncorrect!\x1b[0m Correct: %s\n", pair.B.Phrase)
	}

	return true
}
