package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	failStdoutMsg := "jakis tekst w stdout"
	failStderrMsg := "Error updating the Git index:"

	successMsg := "Files pulled successfully."

	// Seed the random number generator to ensure different results each run
	rand.Seed(time.Now().UnixNano())

	// Randomly decide whether to print an error or success message
	if rand.Intn(2) == 0 {
		fmt.Fprintf(os.Stdout, "%s\n", failStdoutMsg)
		fmt.Fprintf(os.Stderr, "%s\n", failStderrMsg)

	} else {
		fmt.Println(successMsg)
	}
}
