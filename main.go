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
	timenow := time.Now()
	var notable string
	var mapping = make(map[string]int)
	var decoder = make(map[int]string)
	// Take user input for notableinput variable
	fmt.Println("Enter a notable:")
	linescanner := bufio.NewScanner(os.Stdin)
	linescanner.Scan()
	line := linescanner.Text()
	notablefile, err := os.Open("notables.txt")
	// Scan one more inputs from user for notableinput2 variable
	fmt.Println("Enter another notable:")
	linescanner.Scan()
	line2 := linescanner.Text()
	// Scan one more inputs from user for notableinput3 variable
	fmt.Println("Enter another notable:")
	linescanner.Scan()
	line3 := linescanner.Text()
	if err != nil {
		panic(err)
	}
	fmt.Println("Enter desired lethal pride seed:")
	lethal := bufio.NewScanner(os.Stdin)
	lethal.Scan()
	lethalpridemod := lethal.Text()
	// Scan the lines of notablefile
	scanner := bufio.NewScanner(notablefile)
	for scanner.Scan() {
		notable = scanner.Text()                  // Scan the words of notables
		reg, err := regexp.Compile("[^a-zA-Z ]+") // Trim non-alphanumeric characters
		if err != nil {
			log.Fatal(err)
		}
		notableCheck := reg.ReplaceAllString(notable, "")
		if notableCheck == line { // If trimmed string is equal to user input
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
			// Open decode.txt file
			decodefile, err := os.Open("decode.txt")
			if err != nil {
				panic(err)
			}
			scanner := bufio.NewScanner(decodefile)
			for scanner.Scan() {
				decode := scanner.Text()
				reg, err := regexp.Compile("[^a-zA-Z]+")
				if err != nil {
					log.Fatal(err)
				}
				decodeCheck := reg.ReplaceAllString(decode, "")
				regnum2, err := regexp.Compile("[^0-9 ]+")
				numbers := regnum2.ReplaceAllString(decode, "")
				numbers_integer, _ := strconv.Atoi(numbers)
				decoder[numbers_integer] = decodeCheck
			}
			for _, row := range data {
				// Match row[processedString] to the decode.txt file
				replacednode := row[processedStringInteger]
				replaced_reg_int, _ := strconv.Atoi(replacednode)
				if decoder[replaced_reg_int] == lethalpridemod {
					if mapping[row[1]] == 0 {
						mapping[row[1]] = 1
					}
					mapping[row[1]]++
				}
			}
		}
		if notableCheck == line2 { // If trimmed string is equal to user input
			reg, err := regexp.Compile("[^0-9]+")
			if err != nil {
				log.Fatal(err)
			}
			processedString := reg.ReplaceAllString(line2, "")
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
			// Open decode.txt file
			decodefile, err := os.Open("decode.txt")
			if err != nil {
				panic(err)
			}
			scanner := bufio.NewScanner(decodefile)
			for scanner.Scan() {
				decode := scanner.Text()
				reg, err := regexp.Compile("[^a-zA-Z]+")
				if err != nil {
					log.Fatal(err)
				}
				decodeCheck := reg.ReplaceAllString(decode, "")
				regnum2, err := regexp.Compile("[^0-9 ]+")
				numbers := regnum2.ReplaceAllString(decode, "")
				numbers_integer, _ := strconv.Atoi(numbers)
				decoder[numbers_integer] = decodeCheck
			}
			for _, row := range data {
				// Match row[processedString] to the decode.txt file
				replacednode := row[processedStringInteger]
				replaced_reg_int, _ := strconv.Atoi(replacednode)
				if decoder[replaced_reg_int] == lethalpridemod {
					if mapping[row[1]] == 0 {
						mapping[row[1]] = 1
					}
					mapping[row[1]]++
				}
			}
		}
		if notableCheck == line3 { // If trimmed string is equal to user input
			reg, err := regexp.Compile("[^0-9]+")
			if err != nil {
				log.Fatal(err)
			}
			processedString := reg.ReplaceAllString(line3, "")
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
			// Open decode.txt file
			decodefile, err := os.Open("decode.txt")
			if err != nil {
				panic(err)
			}
			scanner := bufio.NewScanner(decodefile)
			for scanner.Scan() {
				decode := scanner.Text()
				reg, err := regexp.Compile("[^a-zA-Z]+")
				if err != nil {
					log.Fatal(err)
				}
				decodeCheck := reg.ReplaceAllString(decode, "")
				regnum2, err := regexp.Compile("[^0-9 ]+")
				numbers := regnum2.ReplaceAllString(decode, "")
				numbers_integer, _ := strconv.Atoi(numbers)
				decoder[numbers_integer] = decodeCheck
			}
			for _, row := range data {
				// Match row[processedString] to the decode.txt file
				replacednode := row[processedStringInteger]
				replaced_reg_int, _ := strconv.Atoi(replacednode)
				if decoder[replaced_reg_int] == lethalpridemod {
					if mapping[row[1]] == 0 {
						mapping[row[1]] = 1
					}
					mapping[row[1]]++
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
	for _, row := range data {
		for i := range mapping {
			if mapping[i] >= 3 {
				fmt.Println("Seed>", row[1])
			}
		}
	}
	fmt.Println("Time taken:", time.Since(timenow))
}
