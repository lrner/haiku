package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

type Word struct {
	Text      string
	Syllables int
}

func main() {
	const SPACE = 32
	const SEPERATOR = 65533
	file, err := os.Open("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var words []Word

	scanner := bufio.NewScanner(file)
	var syllables int
	var wordString string
	var newWord Word
	// var wordsMDArray [][]Word
	for scanner.Scan() {
		syllables = 1
		wordString = ""
		for _, c := range scanner.Text() {
			value := strconv.Itoa(int(c))
			intValue := int(c)
			// fmt.Println(i, " => ", string(c), " => ", value)
			if intValue == SPACE {
				wordString = fmt.Sprintf("%s%s", wordString, string(c))
				syllables++
			} else if intValue == SEPERATOR {
				syllables++
			} else if int(c) > 65000 {
				fmt.Println(scanner.Text())
				fmt.Println(value)
			} else {
				wordString = fmt.Sprintf("%s%s", wordString, string(c))
			}
		}

		newWord = Word{
			Text:      wordString,
			Syllables: syllables,
		}
		if newWord.Syllables == 1 {
			fmt.Println(newWord)
			time.Sleep(time.Second)
		}
		// fmt.Printf("%+v", newWord)
		words = append(words, newWord)
		// wordsMDArray[syllables] = append(wordsMDArray[syllables], newWord)
	}

	// wrFile, err := os.Open("words.json")

	jsonData, _ := json.MarshalIndent(words, "", "")

	_ = ioutil.WriteFile("words.json", jsonData, 0644)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
