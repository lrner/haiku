package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/mikkergimenez/haiku/lib/booktools"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Word struct {
	Text      string
	Syllables int
}

func WordArray() ([][]string, []map[byte][]string) {
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
	var wordsArray [][]string
	var wordsMDArray []map[byte][]string
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
		// fmt.Printf("%+v", newWord)
		words = append(words, newWord)
		if len(wordsMDArray) < syllables {
			for i := len(wordsMDArray) - 1; i <= syllables+2; i++ {
				fmt.Print("-", i, "-")
				wordsArray = append(wordsArray, []string{})
				wordsMDArray = append(wordsMDArray, make(map[byte][]string))
			}
		}
		wordsMDArray[syllables][byte(newWord.Text[0])] = append(wordsMDArray[syllables][byte(newWord.Text[0])], newWord.Text)
		wordsArray[syllables] = append(wordsArray[syllables], newWord.Text)

	}

	return wordsArray, wordsMDArray

}

type Handlers struct {
	Words         [][]string
	WordsByLetter []map[byte][]string
}

func Haiku(building int, room int, shelf int, row int, book int) {

}

func GetHaiku(words [][]string, startMod int, modifier int) []string {

	var haiku []string
	var firstFive string
	var secondSeven string
	var thirdFive string

	for i := 1; i <= 5; i++ {
		firstFive = fmt.Sprintf("%s %s", firstFive, words[1][0])
	}

	haiku = append(haiku, firstFive)

	for j := 1; j <= 7; j++ {
		secondSeven = fmt.Sprintf("%s %s", secondSeven, words[1][0])
	}

	haiku = append(haiku, secondSeven)

	for k := 1; k <= 5; k++ {
		if k == 4 && startMod > len(words[1]) {
			mod := startMod / len(words[1])
			thirdFive = fmt.Sprintf("%s %s", thirdFive, words[1][0+mod])
		} else if k == 5 {
			mod := startMod % len(words[1])
			thirdFive = fmt.Sprintf("%s %s", thirdFive, words[1][0+modifier+mod])
		} else {
			thirdFive = fmt.Sprintf("%s %s", thirdFive, words[1][0])
		}
	}

	haiku = append(haiku, thirdFive)

	return haiku
}

func StringArrayContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func NoDupsHaiku(words [][]string, startMod int, modifier int) []string {

	var haiku []string
	var firstFive []string
	var secondSeven []string
	var thirdFive []string
	noDupsMod := 0

	for i := 1; i <= 5; i++ {
		for len(firstFive) < i {
			if StringArrayContains(firstFive, words[1][0+noDupsMod]) {
				noDupsMod++
			} else {
				firstFive = append(firstFive, words[1][0+noDupsMod])
			}
		}
	}

	haiku = append(haiku, strings.Join(firstFive, " "))

	for k := 1; k <= 7; k++ {
		for len(secondSeven) < k {
			if StringArrayContains(firstFive, words[1][0+noDupsMod]) || StringArrayContains(secondSeven, words[1][0+noDupsMod]) {
				noDupsMod++
			} else {
				secondSeven = append(secondSeven, words[1][0+noDupsMod])
			}
		}
	}

	haiku = append(haiku, strings.Join(secondSeven, " "))

	for m := 1; m <= 5; m++ {
		for len(thirdFive) < m {
			if StringArrayContains(firstFive, words[1][0+noDupsMod]) || StringArrayContains(secondSeven, words[1][0+noDupsMod]) || StringArrayContains(thirdFive, words[1][0+noDupsMod]) {
				noDupsMod++
			} else {
				if m == 5 {
					thirdFive = append(thirdFive, words[1][0+noDupsMod+modifier])
				} else {
					thirdFive = append(thirdFive, words[1][0+noDupsMod])
				}
			}
		}
	}

	haiku = append(haiku, strings.Join(thirdFive, " "))

	return haiku
}

func (h *Handlers) pageHandler(w http.ResponseWriter, req *http.Request) {
	page := strings.TrimPrefix(req.URL.Path, "/page/")

	p, _ := strconv.Atoi(page)
	startMod := p * 10

	fmt.Printf("%d\n", startMod)

	var haiku []string
	var tenHaiku [][]string

	for i := 0; i <= 10; i++ {
		haiku = GetHaiku(h.Words, startMod, i)
		tenHaiku = append(tenHaiku, haiku)
	}

	str, _ := json.Marshal(tenHaiku)
	w.Write([]byte(str))
}

type ReqBody struct {
	Building int
	Room     int
	Row      int
	Series   int
	Shelf    int
	Volume   int
	Book     int
	Page     int
}

func (h *Handlers) bookHandler(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}
	rb := ReqBody{}
	json.Unmarshal(body, &rb)

	rows := booktools.GetBy(rb.Building, rb.Room, rb.Row, rb.Shelf, rb.Series, rb.Book, rb.Page)

	fmt.Printf("%+v", rows)

	var firstString string
	var secondString string
	var thirdString string

	var words []string
	var tenHaiku [][]string

	for x := 0; x <= 10; x++ {
		firstString = ""
		secondString = ""
		thirdString = ""
		for i, s := range rows.FirstSyllables {
			words = h.WordsByLetter[s][byte(rows.FirstLetters[i][0])]
			firstString = firstString + words[0] + " "
		}

		for i, s := range rows.SecondSyllables {
			words = h.WordsByLetter[s][byte(rows.SecondLetters[i][0])]
			secondString = secondString + words[0] + " "
		}

		for i, s := range rows.ThirdSyllables {
			words = h.WordsByLetter[s][byte(rows.ThirdLetters[i][0])]
			thirdString = thirdString + words[x] + " "
		}

		tenHaiku = append(tenHaiku, []string{
			firstString, secondString, thirdString,
		})
	}

	str, _ := json.Marshal(tenHaiku)
	w.Write([]byte(str))
}

func (h *Handlers) noDupsHandler(w http.ResponseWriter, req *http.Request) {
	page := strings.TrimPrefix(req.URL.Path, "/page/")

	fmt.Println("No Dups ", page)

	p, _ := strconv.Atoi(page)
	startMod := p * 10

	var haiku []string
	var tenHaiku [][]string

	for i := 0; i <= 10; i++ {
		haiku = NoDupsHaiku(h.Words, startMod, i)
		tenHaiku = append(tenHaiku, haiku)
	}

	str, _ := json.Marshal(tenHaiku)
	w.Write([]byte(str))
}

func main() {
	var handlers Handlers

	handlers.Words, handlers.WordsByLetter = WordArray()
	for i := 0; i <= len(handlers.Words)-1; i++ {
		fmt.Printf("%d: %d\n", i, len(handlers.Words[i]))
	}

	fmt.Printf("%+v", handlers.Words[1])
	http.HandleFunc("/page/1/no_dups", handlers.noDupsHandler)
	http.HandleFunc("/page", handlers.pageHandler)

	http.HandleFunc("/book", handlers.bookHandler)

	fmt.Println("Listening on Localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
