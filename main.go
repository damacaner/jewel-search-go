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
	reg, err := regexp.Compile("[^a-zA-Z]+")
	regnum, err := regexp.Compile("[^0-9]+")
	var notable string
	var notablecount int
	type mapping map[string]int
	var mapping2d []mapping
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
			numbers := regnum.ReplaceAllString(decode, "")
			numbers_integer, _ := strconv.Atoi(numbers)
			decoder[decodeCheck] = numbers_integer
			break
		}
	}
	seedfile, err := os.Open("Lethal pride seeds.csv")
	if err != nil {
		panic(err)
	}
	csv_reader := csv.NewReader(seedfile)
	data, err := csv_reader.ReadAll()
	if err != nil {
		panic(err)
	}
	// Scan the lines of notablefile
	scanner := bufio.NewScanner(notablefile)
	for scanner.Scan() {
		notable = scanner.Text() // Scan the words of notables
		// Trim non-alphanumeric characters
		if err != nil {
			log.Fatal(err)
		}
		notableCheck := reg.ReplaceAllString(notable, "")
		for line := range notableamount {
			if notableCheck == notableamount[line] { // If trimmed string is equal to user input
				if err != nil {
					log.Fatal(err)
				}
				processedString := regnum.ReplaceAllString(notable, "")
				processedStringInteger, err := strconv.Atoi(processedString)
				if err != nil {
					log.Fatal(err)
				}
				// Search the processedString in the data
				// Search the processedString in the Lethal pride seeds.csv file
				for i := 0; i < 20000; i++ {
					// Append dummy value to mapping2d
					mapping2d = append(mapping2d, make(mapping))
				}
				for _, row := range data {
					// Match row[processedString] to the decode.txt file
					replacednode := row[processedStringInteger]
					replaced_reg_int, _ := strconv.Atoi(replacednode)
					row_int, err := strconv.Atoi(row[1])
					if err != nil {
						log.Fatal(err)
					}
					// Append mapping2d[row_int][string(line)] to the mapping2d array
					if decoder[decodeCheck] == replaced_reg_int {
						// Increment the value of the key in the mapping2d array
						mapping2d[row_int][string(line)] = mapping2d[row_int][string(line)] + 1
					}
				}
			}
		}
		// Create the seeds.csv file
	}
	file, err := os.Create("seeds.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for i := range mapping2d {
		for jewel := range mapping2d[i] {
			if mapping2d[i][jewel] >= 3 {
				fmt.Fprintf(file, "%d,%d\n", i)
				// 	Print to console
				fmt.Println(i, mapping2d[i][jewel])
			}
		}
	}
	fmt.Println("Only showing the seeds that higher than 3 occurences, if you dont see anything, then it doesnt exist.")
	fmt.Println("Time taken:", time.Since(timenow))
}
