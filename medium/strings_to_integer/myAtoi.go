package strings_to_integer

import (
	"math"
	"strconv"
	"strings"
)

const (
	plusSign  = "+"
	minusSign = "-"
)

var allowedSymbols = map[string]struct{}{
	plusSign:  {},
	minusSign: {},
	"0":       {},
	"1":       {},
	"2":       {},
	"3":       {},
	"4":       {},
	"5":       {},
	"6":       {},
	"7":       {},
	"8":       {},
	"9":       {},
}

var edge = math.Pow(2, 31)

func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}

	splitted := strings.Split(strings.TrimSpace(str), "")
	if len(splitted) == 0 {
		return 0
	}

	if _, ok := allowedSymbols[splitted[0]]; !ok {
		return 0
	}

	if splitted[0] != plusSign && splitted[0] != minusSign {
		delete(allowedSymbols, plusSign)
		delete(allowedSymbols, minusSign)
	}

	out := ""
	counts := 0
	for _, symbol := range splitted {
		if _, ok := allowedSymbols[symbol]; !ok {
			break
		}

		if symbol == plusSign || symbol == minusSign {
			counts++
			if counts > 1 {
				break
			}
		}

		out += symbol
	}

	allowedSymbols[plusSign] = struct{}{}
	allowedSymbols[minusSign] = struct{}{}

	res, err := strconv.ParseFloat(out, 64)
	if err != nil {
		return 0
	}

	if math.Abs(res) >= edge {
		if res < 0 {
			return -int(edge)
		}

		return int(edge - 1)
	}

	return int(res)
}
