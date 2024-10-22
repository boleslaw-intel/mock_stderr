package main

import (
	"fmt"
	"os"
)

func main() {
	// Define the text to be printed to stderr
	text1 := "jakis tekst w stdout"
	text2 := "Error updating the Git index:"

	// Print the text to stdout using fmt.Fprintf
	fmt.Fprintf(os.Stdout, "%s\n", text1)

	// Print the text to stderr using fmt.Fprintf
	fmt.Fprintf(os.Stderr, "%s\n", text2)
}
