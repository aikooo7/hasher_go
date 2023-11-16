package main

import (
	"crypto/sha512"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var input, filename string

	fmt.Printf("What file you want to hash: ")
	fmt.Scanln(&filename)

	file, err := os.ReadFile(filename)

	if err != nil {
		log.Fatalf("Error reading the file: %s", err)
	}

	sha_512 := sha512.New()
	sha_512.Write([]byte(file))

	fmt.Printf("Your hash: %x\n", sha_512.Sum(nil))

	fmt.Printf("Want to store the hash in a file? ")
	fmt.Scanln(&input)

	switch strings.ToLower(input) {
	case "yes":
		content := fmt.Sprintf("Filename: %s | Hash: %x\n", filename, sha_512.Sum(nil))

		f, err := os.OpenFile("hashes.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

		if err != nil {
			log.Fatalf("Error opening the file: %s", err)
		}

		_, err = f.WriteString(content)

		if err != nil {
			log.Fatalf("Error writing to the file: %s", err)
		} else {
			fmt.Printf("Done...\n")
		}

	case "no":
		fmt.Printf("Not writing to the file...\n")
	}

	fmt.Printf("Do you want to hash other file? ")
	fmt.Scanln(&input)

	switch strings.ToLower(input) {
	case "yes":
		fmt.Print("\033[H\033[2J")
		main()
	case "no":
		fmt.Printf("See you next time...\n")
	}
}
