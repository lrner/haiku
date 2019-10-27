package syllables

import (
	"fmt"
	"log"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func calcSquaredSum(syllableSlice [][]int) int {
	var result int
	for _, intSlice := range syllableSlice {
		result = result + int(math.Pow(26, float64(len(intSlice))))
	}

	return result
}

func eachAddsTo(t *testing.T, syllableSlice [][]int, expectedSum int) {
	var thisSum int
	for _, intSlice := range syllableSlice {
		thisSum = 0
		for _, j := range intSlice {
			thisSum += j
		}
		assert.Equal(t, thisSum, expectedSum, fmt.Sprintf("Row %+v doesnt add to %d", intSlice, expectedSum))
	}
}

func TestSyllables(t *testing.T) {
	fiveSyllablesSquaredSum := (calcSquaredSum(FiveSyllableSlice()))
	sevenSyllablesSquaredSum := (calcSquaredSum(SevenSyllableSlice()))

	eachAddsTo(t, FiveSyllableSlice(), 5)
	eachAddsTo(t, SevenSyllableSlice(), 7)

	placesArray := PlacesArray()

	log.Printf("%d, %d", fiveSyllablesSquaredSum, placesArray[4])
	log.Printf("%d, %d", sevenSyllablesSquaredSum, placesArray[6])
	assert.Equal(t, fiveSyllablesSquaredSum, placesArray[4], "For five syllables, the sums should be the same.")
	assert.Equal(t, sevenSyllablesSquaredSum, placesArray[6], "For seven syllables, the sums should be the same.")

}
