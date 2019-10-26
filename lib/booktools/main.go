package booktools

import (
	"math"

	"github.com/mikkergimenez/haiku/lib/syllables"
)

func GetLetters(letterSum float64, lenSlice int) []string {
	retLetterArray := []string{}

	intermediateLetterSum := letterSum

	letterArray := []string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
		"h",
		"i",
		"j",
		"k",
		"l",
		"m",
		"n",
		"o",
		"p",
		"q",
		"r",
		"s",
		"t",
		"u",
		"v",
		"w",
		"x",
		"y",
		"z",
	}

	var mod float64
	for i := 0; i <= lenSlice; i++ {
		if i == 0 {
			mod = 1
		} else {
			mod = 0
		}
		if intermediateLetterSum > 0 {
			mathMod := math.Mod(float64(intermediateLetterSum), 26)
			retLetterArray = append(retLetterArray, letterArray[int(mathMod-mod)])

			intermediateLetterSum = (intermediateLetterSum - mathMod) / 26
		}
	}

	for j := len(retLetterArray); j < lenSlice; j++ {
		retLetterArray = append(retLetterArray, "a")
	}

	s := retLetterArray
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	retLetterArray = s

	return retLetterArray
}

func FiveRow(identifier int) ([]int, []string) {
	fiveSliceList := syllables.FiveSyllableSlice()
	return TestNumberRow(fiveSliceList, float64(identifier))
}

func SevenRow(identifier int) ([]int, []string) {
	sevenSliceList := syllables.SevenSyllableSlice()
	return TestNumberRow(sevenSliceList, float64(identifier))
}

func TestNumberRow(sliceList [][]int, identifier float64) ([]int, []string) {
	var retSyllableArray []int
	var retLetterArray []string

	var intermediateSum float64
	intermediateSum = 0

	var sliceLength float64
	var letterSum float64
	for _, slice := range sliceList {
		sliceLength = float64(len(slice))
		if identifier > (math.Pow(26, sliceLength) + intermediateSum) {
			intermediateSum += math.Pow(26, sliceLength)
		} else {
			letterSum = identifier - intermediateSum

			retLetterArray := GetLetters(letterSum, len(slice))

			return slice, retLetterArray
		}
	}

	return retSyllableArray, retLetterArray

}

type Rows struct {
	FirstSyllables  []int
	FirstLetters    []string
	SecondSyllables []int
	SecondLetters   []string
	ThirdSyllables  []int
	ThirdLetters    []string
}

func GetBy(building int, room int, row int, shelf int, series int, book int, page int) Rows {
	var rows Rows

	rows.FirstSyllables, rows.FirstLetters = FiveRow(building)
	rows.SecondSyllables, rows.SecondLetters = SevenRow(room)
	rows.ThirdSyllables, rows.ThirdLetters = FiveRow(row)

	return rows
}
