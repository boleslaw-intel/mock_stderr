package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// TODO:
//	1. dodaj opcje --always-fail,  ktora zastapi random success/fail
//  2. dodaj parametr --cmd, ktory zostanie przekazany do cmd.Run() w przypadku 'success case'

func main() {

	failStdoutMsg := []string{"some non-error text in stdout", "STDouT"}
	failStderrMsg := []string{"Error updating the Git index:", "Some error but not critical", "Error 123", "without err or keyword"}

	successMsg := "Files pulled successfully."

	// Seed the random number generator to ensure different results each run
	rand.Seed(time.Now().UnixNano())

	// Randomly decide whether to print an error or success message
	if rand.Intn(2) == 0 {
		fmt.Fprintf(os.Stdout, "%s\n", failStdoutMsg[rand.Intn(len(failStdoutMsg))])
		fmt.Fprintf(os.Stderr, "%s\n", failStderrMsg[rand.Intn(len(failStderrMsg))])

	} else {
		fmt.Println(successMsg)
		//wywolaj oryuginalny cmd call jako argument
	}
}
