package number_of_changing_keys

func countKeyChanges(s string) int {
	prevLetter, changesCount := -1, -1
	bytes := []byte(s)
	for _, letter := range bytes {
		symbolInd, symbolSec := int(letter), 0
		if symbolInd >= 65 && symbolInd < 93 {
			symbolSec = symbolInd + 32
		} else if symbolInd >= 97 && symbolInd < 125 {
			symbolSec = symbolInd - 32
		}

		if !(prevLetter == symbolInd || prevLetter == symbolSec) {
			changesCount += 1
		}

		prevLetter = symbolInd
	}
	return changesCount
}
