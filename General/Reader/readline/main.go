package main

/*
NewReader returns a new Reader whose buffer has the default size (bufio's default buffer size is 4K).
By putting os.Stdin as the variable in the NewReader function, we indicate the Reader to read from the terminal.
We will show how we can use the ReadString and the ReadLine methods to read data from the terminal.

ReadLine does a buffered read up to a line terminator. It handles either \n or \r\n, and returns
just the line without the new line character.

*/
import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
Readline() call is a byte array representation of the input string,
and we can use string(line) to convert the byte array into the original string.
As you can see, ReadLine is a low-level line-reading primitive. Most callers
should use ReadBytes('\n') or ReadString('\n') instead or use a Scanner.
*/
func main() {
	reader := bufio.NewReader(os.Stdin) // SETTING UP THE READER TERMINAL TO TERMINAL INPUT
	fmt.Println("Enter the Name :")     //ENTER THE NAME: RAMS
	line, _, err := reader.ReadLine()   // READLINE func (b *reader) ReadLine()(line []byte,isprefix bool,err error)
	//If the line was too long for the buffer then isPrefix is set and the beginning of the line is returned.
	if err != nil {
		log.Fatalf("Error reading the input", err)
	}
	fmt.Println(line)         //PRINTS BYTE OUTPUT
	fmt.Println(string(line)) //PRINTS THE NAME ENTERED IN TERMINAL
}
