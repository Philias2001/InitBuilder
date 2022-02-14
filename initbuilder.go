package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var name string

func init() {
	fmt.Println("Welcome to the Initials builder")
	fmt.Println()
	fmt.Println("There are two types:")
	fmt.Println("1: EU standard")
	fmt.Println("2: International")
	fmt.Println()
}

func main() {
	fmt.Print("Type [1] EU or [2] International: ")
	reader := bufio.NewReader(os.Stdin)

	// Readstring will block until the delimiter is entered
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("ERROR! Please try again. ", err)
		return
	}


	// Remove the delimter from the string
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	switch input {
	case "1":
		name = eu()
	case "2":
		name = international()
	case "close":
		fmt.Println("---------------------------")
		fmt.Println("closing the program")
		return
	default:
		fmt.Println("---------------------------")
		fmt.Println("Please enter [1] or [2]")
		fmt.Println("---------------------------")

		main()
		return
	}

	if name == "" {
		return
	}

	fmt.Println()
	fmt.Println("Your initial name is:",name)
	fmt.Println()

	main()
}

func eu()string {
	var initName string
	fmt.Print("Enter your first name(s): ")
	readerFirstName := bufio.NewReader(os.Stdin)

	// Readstring will block until the delimiter is entered
	inputFirstName, err := readerFirstName.ReadString('\n')

	fmt.Print("Enter your preposistion(s): ")
	readerPreposistion := bufio.NewReader(os.Stdin)

	// Readstring will block until the delimiter is entered
	inputPreposistion, err := readerPreposistion.ReadString('\n')

	fmt.Print("Enter your last name: ")
	readerLastName := bufio.NewReader(os.Stdin)

	// Readstring will block until the delimiter is entered
	inputLastName, err := readerLastName.ReadString('\n')

	if err != nil {
		fmt.Println("ERROR! Please try again. ", err)
		return ""
	}


	// Remove the delimter from the string
	inputFirstName = strings.TrimSuffix(inputFirstName, "\n")
	inputFirstName = strings.TrimSuffix(inputFirstName, "\r")

	inputPreposistion = strings.TrimSuffix(inputPreposistion, "\n")
	inputPreposistion = strings.TrimSuffix(inputPreposistion, "\r")

	inputLastName = strings.TrimSuffix(inputLastName, "\n")
	inputLastName = strings.TrimSuffix(inputLastName, "\r")

	firstNameArray := strings.Split(inputFirstName, " ")

	for _, c := range firstNameArray {
			c = strings.ToUpper(c[0:1])
			initName += c + "."
	}

	count := len(initName) - 1
	initName = initName[0:count]

	if inputPreposistion == "" {
		initName += " " + inputLastName
	} else {
		initName += " " + inputPreposistion + " " + inputLastName
	}

	return initName
}

func international()string {
	var initName string
	fmt.Print("Enter full name: ")
	reader := bufio.NewReader(os.Stdin)

	// Readstring will block until the delimiter is entered
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("ERROR! Please try again. ", err)
		return ""
	}

	// Remove the delimter from the string
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	inputArray := strings.Split(input, " ")

	for _, c := range inputArray {
		matched, err := regexp.Match(`^(?i)(van|der|de|den)`, []byte(c))

		if err != nil {
			fmt.Println("Something went wrong. ", err)
			return ""
		}

		if !matched {
			c = strings.ToUpper(c[0:1])
			initName += c + "."
		}
	}

	count := len(initName) - 1
	initName = initName[0:count]

	return initName
}