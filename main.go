package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	var notable string
	var notablecount int
	var mapping = make(map[string]int)
	var notableamount []string
	var decoder = make(map[string]int)
	// Ask user to enter a desired notable amount to search and store every notable in a different value
	fmt.Println("Enter a notable amount to search:")
	fmt.Println("Minimum 4 is recommended, as search is not showing anything lower than 3.")
	fmt.Scanln(&notablecount)
	if notablecount > 8 {
		fmt.Println("Please enter a number between 3 and 8.")
		os.Exit(1)
	}
	for i := 0; i < notablecount; i++ {
		// Take an user oinput for notable
		fmt.Printf("Enter notable %d> ", i)
		// Scan notable with buffer
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		notable = scanner.Text()
		// Append notable to array
		notableamount = append(notableamount, notable)
	}
	notablefile, err := os.Open("notables.txt")
	fmt.Println("Enter desired lethal pride seed:")
	timenow := time.Now()
	decodefile, err := os.Open("decode.txt")
	if err != nil {
		panic(err)
	}
	lethal := bufio.NewScanner(os.Stdin)
	lethal.Scan()
	lethalpridemod := lethal.Text()
	var decodeCheck string
	scannerdecode := bufio.NewScanner(decodefile)
	for scannerdecode.Scan() {
		decode := scannerdecode.Text()
		reg, err := regexp.Compile("[^a-zA-Z]+")
		if err != nil {
			log.Fatal(err)
		}
		decodeCheck = reg.ReplaceAllString(decode, "")
		// Check if decodeCheck is equal to lethalpridemod
		if decodeCheck == lethalpridemod {
			fmt.Println("Decode found!")
			fmt.Println("Decode:", decode)
			regnum2, _ := regexp.Compile("[^0-9 ]+")
			numbers := regnum2.ReplaceAllString(decode, "")
			numbers_integer, _ := strconv.Atoi(numbers)
			decoder[decodeCheck] = numbers_integer
			break
		}
	}
	// Scan the lines of notablefile
	scanner := bufio.NewScanner(notablefile)
	for scanner.Scan() {
		for line := range notableamount {
			notable = scanner.Text()                  // Scan the words of notables
			reg, err := regexp.Compile("[^a-zA-Z ]+") // Trim non-alphanumeric characters
			if err != nil {
				log.Fatal(err)
			}
			notableCheck := reg.ReplaceAllString(notable, "")
			if notableCheck == notableamount[line] { // If trimmed string is equal to user input
				reg, err := regexp.Compile("[^0-9]+")
				if err != nil {
					log.Fatal(err)
				}
				processedString := reg.ReplaceAllString(notable, "")
				processedStringInteger, err := strconv.Atoi(processedString)
				// Search the processedString in the Lethal pride seeds.csv file
				seedfile, err := os.Open("Lethal pride seeds.csv")
				if err != nil {
					panic(err)
				}
				csv_reader := csv.NewReader(seedfile)
				data, err := csv_reader.ReadAll()
				if err != nil {
					panic(err)
				}
				for _, row := range data {
					// Match row[processedString] to the decode.txt file
					replacednode := row[processedStringInteger]
					replaced_reg_int, _ := strconv.Atoi(replacednode)
					fmt.Println(decoder[decodeCheck], replaced_reg_int)
					if decoder[decodeCheck] == replaced_reg_int {
						if mapping[row[1]] == 0 {
							mapping[row[1]] = 1
						}
						mapping[row[1]] = mapping[row[1]] + 1
						continue
					}
				}
			}
		}
	}
	seedfile, err := os.Open("Lethal pride seeds.csv")
	if err != nil {
		panic(err)
	}
	csv_reader := csv.NewReader(seedfile)
	data, err := csv_reader.ReadAll()
	fmt.Println("Only showing the seeds that higher than 3 occurences, if you dont see anything, then it doesnt exist.")
	// Create the seeds.csv file
	file, err := os.Create("seeds.csv")
	for _, row := range data {
		for i := range mapping {
			if mapping[i] >= 3 {
				fmt.Println("Seed>", row[1])
				// Append row[1] to the seeds.csv file
				file.WriteString(row[1] + "\n")
			}
		}
	}
	fmt.Println("Time taken:", time.Since(timenow))
}
