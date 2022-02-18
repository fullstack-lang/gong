// package models aims to store uml diagrams into a go variable
// the aim is to be robust to refactoring tasks such as rename
package models

import (
	"fmt"
	"os"
)

// indent prints nbIndentation tabs to the file
func indent(file *os.File, nbIndentation int) {
	for i := 0; i < nbIndentation; i++ {
		fmt.Fprintf(file, "\t")
	}
}
