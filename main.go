package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) >= 2 {
		textFromOutside := os.Args[1]
		fileName := ""
		if len(os.Args) == 3 {
			if os.Args[2] == "shadow" || os.Args[2] == "standard" || os.Args[2] == "thinkertoy" {
				fileName = os.Args[2] // if something else standard 
			}
		}
		 fileLines := ReadStandardTxt(fileName)
		// fmt.Println(textFromOutside, fileLines)
		 asciiTemplates := return2dASCIIArray(fileLines)
		// fmt.Println(asciiTemplates)
		// printMultipleCharacter(textFromOutside, asciiTemplates)
		// fmt.Println(int('\n'))
		// test1 := "\nHello\nWorld\n\nMom\n\n\nAnd Dad\n"
		// fmt.Println(returnstring2EndlineArray(textFromOutside))	
		 printAllStringASCII(textFromOutside, asciiTemplates)

	}
}

func ReadStandardTxt(nameOfFile string) []string {
	if nameOfFile == "" {
		nameOfFile = "standard"
	}
	readFile, err := os.Open("artstyles/"+nameOfFile+".txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()
	return fileLines
}

func return2dASCIIArray(fileLines []string) [][]string {
	var asciiTemplates [][]string
	counter := 0
	var tempAsciArray []string
	for _, line := range fileLines {
		counter++
		// fmt.Println(index, line)
		if counter != 1 {
			tempAsciArray = append(tempAsciArray, line)
		}
		if counter == 9 {
			// fmt.Println("add to list") // but dont include the first line because it is empty line
			asciiTemplates = append(asciiTemplates, tempAsciArray)
			counter = 0
			tempAsciArray = nil
		} else {
		}
	}
	return asciiTemplates
}

// problem '\n' logic when we have 2 '\n' or 1 '\n' is different
func printMultipleCharacter(s string, asciiTemplates [][]string) {
	// for ex 'hello'
	// we gonna write all letters index 0 after $ after \n after index 1  after $ after \n
	tempIntArrLetter := returnAsciiCodeInt(s)
	for i := 0; i < 8; i++ {
		for _, v := range tempIntArrLetter {
			fmt.Print(asciiTemplates[v][i])
		}
		fmt.Println("$")
	}
}

func returnAsciiCodeInt(s string) []int {
	var tempIntArrLetter []int
	for _, v := range s {
		tempIntArrLetter = append(tempIntArrLetter, (int(v) - 32))
	}
	return tempIntArrLetter
}

func printAllStringASCII(text string, asciiTemplates [][]string) {
	/*
		if ends w \n it gonna print println $
		if you can see text after \n chec;
		before \n
		if yes  println $
		if no println
	*/

	/*
	   func to uses printMultipleCharacters print whole stringfrom outside
	*/
	// Split the input string into an array of strings
	// split the line into words if there is a "\r\n" symbol
	substrings := returnstring2EndlineArray(text)
	lenOfsubstrings := len(substrings)
	for index, v := range substrings {
		if v == "\\n" {
			// If it is last one
			if index == lenOfsubstrings-1 {
				fmt.Println("$")
			} else if index == 0 {
				fmt.Println("$") // no idea CHECK IT POTENTIAL ERROR
			} else {
				if substrings[index-1] == "\\n" {
					fmt.Println("$")

				} else {
					// "Hello\nWorld"
				}
			}
		} else {
			printMultipleCharacter(v, asciiTemplates)
		}
	}
	

}

func returnstring2EndlineArray(text string) []string {
	substrings := make([]string, 0)
	escapedN := "\\n"
	newline := "\n"

	for {
		idx := strings.Index(text, escapedN)
		if idx == -1 {
			substrings = append(substrings, text)
			break
		}

		substrings = append(substrings, text[:idx])

		if idx+len(escapedN) < len(text) && text[idx+len(escapedN)] == 'n' {
			substrings = append(substrings, newline)
			text = text[idx+len(escapedN)+1:]
		} else {
			substrings = append(substrings, escapedN)
			text = text[idx+len(escapedN):]
		}
	}
	// fmt.Printf("%#v\n", substrings)
	var mysubstring2 []string
	for _, mysub := range substrings {
		if mysub != "" {
			mysubstring2 = append(mysubstring2, mysub)
		}
	}
	// fmt.Printf("%#v\n", mysubstring2)
	return mysubstring2
}
