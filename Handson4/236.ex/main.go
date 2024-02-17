package main

import (
	"fmt"
	"os"
)

type errxface struct {
	errormsg string
	errorcd  error
}

func (errx errxface) Error() string {
	return fmt.Sprintln("Fatal error : ", errx.errormsg, errx.errorcd)
}

func foo() error {
	f, error1 := os.Open("error.txt")
	if error1 != nil {
		errora := fmt.Sprintf("Opening the file : %v", f)
		errorb := fmt.Errorf("with error : ", error1)
		return errxface{errora, errorb}
	} else {
		f.WriteString("Test file")
		return nil
	}
}

func main() {
	errors := foo()
	fmt.Println(errors)
}
