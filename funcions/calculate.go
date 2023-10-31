package funcions

import (
	"fmt"
	"strconv"
	"strings"
)

type Calculator struct {
	NumOne   int
	NumTwo   int
	Operator string
	Type     string
}

type RomanNumeralSystem struct {
	Value  int
	Symbol string
}

func GetArabicFromRomans(romanNum string) (int, error) {
	sliceRoman := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	arabicNum, err := sliceRoman[romanNum]
	if !err {
		return 0, fmt.Errorf("неверная римская цифра")
	}
	return arabicNum, nil
}

func GetRomanFromArabic(arabicNum int) (string, error) {
	var romanNumerals = []RomanNumeralSystem{
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	if arabicNum <= 0 || arabicNum > 100 {
		return "", fmt.Errorf("неверное римское число")
	}

	var romanNum string
	for _, numeral := range romanNumerals {
		for arabicNum >= numeral.Value {
			romanNum += numeral.Symbol
			arabicNum -= numeral.Value
		}
	}

	return romanNum, nil

}

func ParseExpression(expres string) (*Calculator, error) {
	parse := strings.Split(expres, " ")
	if len(parse) != 3 {
		return nil, fmt.Errorf("неверное построено выражение")
	}
	numOne, types, err := parseNumber(parse[0])
	if err != nil {
		return nil, err
	}
	numTwo, typesTwo, err := parseNumber(parse[2])
	if err != nil {
		return nil, err
	}
	if types != typesTwo {
		return nil, fmt.Errorf("два разных типа данных")
	}
	calc := &Calculator{
		NumOne:   numOne,
		NumTwo:   numTwo,
		Operator: parse[1],
		Type:     types,
	}

	return calc, nil
}

func parseNumber(str string) (int, string, error) {
	num, err := strconv.Atoi(str)
	if err == nil {
		if num < 1 || num > 10 {
			return 0, "arabic", fmt.Errorf("число должно быть от 1 до 10")
		}
		return num, "arabic", nil
	}
	romanNum, err := GetArabicFromRomans(str)
	if err != nil {
		return 0, "arabic", err
	}
	return romanNum, "roman", nil
}

func (calc *Calculator) WeConsider() (interface{}, error) {
	switch calc.Operator {
	case "+":
		answer := calc.NumOne + calc.NumTwo
		if calc.Type == "roman" {
			return GetRomanFromArabic(answer)
		}
		return answer, nil
	case "-":
		answer := calc.NumOne - calc.NumTwo
		if calc.Type == "roman" {
			return GetRomanFromArabic(answer)
		}
		return answer, nil

	case "*":
		answer := calc.NumOne * calc.NumTwo
		if calc.Type == "roman" {
			return GetRomanFromArabic(answer)
		}
		return answer, nil

	case "/":
		if calc.NumTwo == 0 {
			return 0, fmt.Errorf("деление на ноль запрещено")
		}
		answer := calc.NumOne / calc.NumTwo
		if calc.Type == "roman" {
			return GetRomanFromArabic(answer)
		}
		return answer, nil

	default:
		return 0, fmt.Errorf("недопустимая операция")
	}
}
