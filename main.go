package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
)

func main() {
	reg, err := regexp.Compile("[^a-zA-Z]+")
	regnum, err := regexp.Compile("[^0-9]+")
	var root string
	var seedcnt int
	var notable string
	var notablefile *os.File
	var seedfile *os.File
	var nodesearched string
	var scanner *bufio.Scanner
	var decoderoption int
	var decodefile *os.File
	var data [][]string
	var notablecount int
	type mapping map[string]int
	var decodeCheck string
	var totalcnt int
	var appendamount int
	var mapping2d []mapping
	var notableamount []string
	var decoder = make(map[string]int)
	var elegantseedpath string
	var lethalseedpath string
	root, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

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
		fmt.Printf("Enter %d. notable without spaces and non-alphanumeric characters > ", i)
		// Scan notable with buffer
		scannernotables := bufio.NewScanner(os.Stdin)
		scannernotables.Scan()
		notable = scannernotables.Text()
		// Append notable to array
		notableamount = append(notableamount, notable)
	}
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		// Check if file name is notables.txt
		if filepath.Base(path) == "notables.txt" {
			// Open file
			notablefile, err = os.Open(path)
			if err != nil {
				log.Fatal(err)
			}
		}
		// Check if file name is decode.txt
		if filepath.Base(path) == "decode.txt" {
			// Open file
			decodefile, err = os.Open(path)
			if err != nil {
				log.Fatal(err)
			}

		}
		// Check if file name is Lethal Pride seeds.csv
		if filepath.Base(path) == "Lethal Pride seeds.csv" {
			// Store the path in lethalseedpath
			lethalseedpath = path
		}
		// Check if file name is Elegant Hubris seeds.csv
		if filepath.Base(path) == "Elegant Hubris seeds.csv" {
			// Store the path in elegantseedpath
			elegantseedpath = path
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Enter if you want to search lethal pride or elegant hubris nodes:")
	fmt.Println("1. Lethal Pride")
	fmt.Println("2. Elegant Hubris")
	fmt.Scanln(&decoderoption)
	timenow := time.Now()
	if err != nil {
		panic(err)
	}
	fmt.Println("Enter the desired node without spaces and non-alphanumeric characters:")
	fmt.Scanln(&nodesearched)
	if decoderoption == 1 {
		scannerdecode := bufio.NewScanner(decodefile)
		for scannerdecode.Scan() {
			decode := scannerdecode.Text()
			reg, err := regexp.Compile("[^a-zA-Z]+")
			if err != nil {
				log.Fatal(err)
			}
			decodeCheck = reg.ReplaceAllString(decode, "")
			// Check if decodeCheck is equal to lethalpridemod
			if decodeCheck == nodesearched {
				fmt.Println("Decode found!")
				fmt.Println("Decode:", decode)
				numbers := regnum.ReplaceAllString(decode, "")
				numbers_integer, _ := strconv.Atoi(numbers)
				decoder[decodeCheck] = numbers_integer
				break
			}
		}
	} else if decoderoption == 2 {
		scannerdecode := bufio.NewScanner(decodefile)
		for scannerdecode.Scan() {
			decode := scannerdecode.Text()
			reg, err := regexp.Compile("[^a-zA-Z]+")
			if err != nil {
				log.Fatal(err)
			}
			decodeCheck = reg.ReplaceAllString(decode, "")
			// Check if decodeCheck is equal to lethalpridemod
			if decodeCheck == nodesearched {
				fmt.Println("Decode found!")
				fmt.Println("Decode:", decode)
				numbers := regnum.ReplaceAllString(decode, "")
				numbers_integer, _ := strconv.Atoi(numbers)
				decoder[decodeCheck] = numbers_integer
				break
			}
		}
	}
	if decoderoption == 2 {
		seedfile, err = os.Open(elegantseedpath)
		if err != nil {
			panic(err)
		}
	} else if decoderoption == 1 {
		seedfile, err = os.Open(lethalseedpath)
		if err != nil {
			panic(err)
		}
	}
	csv_reader := csv.NewReader(seedfile)
	data, err = csv_reader.ReadAll()
	if err != nil {
		panic(err)
	}
	scanner = bufio.NewScanner(notablefile)
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
				if decoderoption == 1 {
					appendamount = 30000
				}
				if decoderoption == 2 {
					appendamount = 170000
				}
				for i := 0; i < appendamount; i++ {
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
	defer file.Close()
	totalcnt = 0
	seedcnt = 0
	for i := range mapping2d {
		for jewel := range mapping2d[i] {
			if mapping2d[i][jewel] > 0 {
				totalcnt++
			}
		}
		if totalcnt >= 3 {
			fmt.Fprintf(file, "Found seed: %d, Seed Number: %d\n", seedcnt, i)
			seedcnt++
		}
		totalcnt = 0
	}
	fmt.Println("Only showing the seeds that higher than 3 occurences in the seeds.txt, if you dont see anything, then it doesnt exist.")
	fmt.Println("Time taken:", time.Since(timenow))
}
