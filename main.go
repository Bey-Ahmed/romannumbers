package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	number := Atoi(os.Args[1])

	if !IsNumeric(os.Args[1]) || number < 0 || number >= 4000 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	digits := []rune{}
	position := 1

	for number > 0 {
		remain := number % 10
		rang := remain * position
		if rang < 4 {
			for i := 0; i < remain; i++ {
				digits = append([]rune{'I'}, digits...)
			}
		} else if rang == 4 {
			digits = append([]rune{'V'}, digits...)
			digits = append([]rune{'I'}, digits...)
		} else if rang < 9 {
			for i := 0; i < remain-5; i++ {
				digits = append([]rune{'I'}, digits...)
			}
			digits = append([]rune{'V'}, digits...)
		} else if rang == 9 {
			digits = append([]rune{'X'}, digits...)
			digits = append([]rune{'I'}, digits...)
		} else if rang < 40 {
			for i := 0; i < remain; i++ {
				digits = append([]rune{'X'}, digits...)
			}
		} else if rang == 40 {
			digits = append([]rune{'L'}, digits...)
			digits = append([]rune{'X'}, digits...)
		} else if rang < 90 {
			for i := 0; i < remain; i++ {
				digits = append([]rune{'X'}, digits...)
			}
			digits = append([]rune{'L'}, digits...)
		} else if rang == 90 {
			digits = append([]rune{'C'}, digits...)
			digits = append([]rune{'X'}, digits...)
		} else if rang < 400 {
			for i := 0; i < remain; i++ {
				digits = append([]rune{'C'}, digits...)
			}
		} else if rang == 400 {
			digits = append([]rune{'D'}, digits...)
			digits = append([]rune{'C'}, digits...)
		} else if rang < 900 {
			for i := 0; i < remain; i++ {
				digits = append([]rune{'C'}, digits...)
			}
			digits = append([]rune{'D'}, digits...)
		} else if rang == 900 {
			digits = append([]rune{'M'}, digits...)
			digits = append([]rune{'C'}, digits...)
		} else if rang < 4000 {
			for i := 0; i < remain; i++ {
				digits = append([]rune{'M'}, digits...)
			}
		}
		number = number / 10
		position *= 10
	}

	for _, digit := range digits {
		z01.PrintRune(digit)
	}
	z01.PrintRune('\n')
}

func IsNumeric(s string) bool {
	if len(s) <= 0 {
		return false
	}
	store := []rune(s)
	size := len(s)
	start := 0
	if store[0] == '-' {
		start = 1
	}
	for i := start; i < size; i++ {
		if store[i] < '0' || store[i] > '9' {
			return false
		}
	}
	return true
}

func Atoi(s string) int {
	digits := []byte(s)
	size := len(digits)

	if size == 0 {
		return 0
	}

	if size == 1 {
		if digits[0] >= '0' && digits[0] <= '9' {
			return int(digits[0] % '0')
		}
		return 0
	}

	intValue := 0
	position := 1
	start := 0
	if digits[0] == '+' || digits[0] == '-' {
		start = 1
	}

	for i := start; i < size; i++ {
		if digits[i] < '0' || digits[i] > '9' {
			return 0
		}

		if digits[0] == '-' {
			intValue = intValue - int(digits[size-1-i+start]%'0')*position
		} else {
			intValue = intValue + int(digits[size-1-i+start]%'0')*position
		}

		position = position * 10
	}
	return intValue
}
