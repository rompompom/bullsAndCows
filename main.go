package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const numberOfDigits int = 4

func main() {
	var computerValues [numberOfDigits]uint = generateValue()
	var gameOver bool = false
	in := bufio.NewReader(os.Stdin)

	for !gameOver {
		var bulls uint = 0
		var cows uint = 0
		fmt.Println("Введите число из четырех цифр:")

		userInput, err := in.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка ввода: ", err)
		}

		if len(userInput) != 6 {
			fmt.Println("Ошибка! Значение должно содержать 4 цифры")
		} else {
			userValues, err := getFigures(userInput)
			if err != nil {
				fmt.Println("Введено некорректное значение. (Нужно ввести четыре неповторяющиеся цифры)")
			} else if !isValuesUnic(userValues) {
				fmt.Println("Числа не должны повторяться")
			} else {
				for i := 0; i < numberOfDigits; i++ {
					for j := 0; j < numberOfDigits; j++ {
						if computerValues[i] == userValues[j] {
							if i == j {
								bulls++
							} else {
								cows++
							}
						}
					}
				}
				if bulls == 4 {
					fmt.Println("Победа!")
					gameOver = true
				} else {
					fmt.Println("Не угадали. Быки:", bulls, "Коровы:", cows)
				}
			}
		}
	}

	var temp string
	fmt.Scan(&temp)
}

//Generates array with four unic value in range 0...9
func generateValue() [numberOfDigits]uint {
	var result [numberOfDigits]uint
	rand.Seed(time.Now().Unix())
	result[0] = uint(rand.Intn(10))

	for i := 1; i < numberOfDigits; i++ {
		var newValue uint = uint(rand.Intn(10))
		var checkValueIsUnic bool = true
		for checkValueIsUnic {
			checkValueIsUnic = false
			for j := 0; j < i; j++ {
				if newValue == result[j] {
					newValue++
					if newValue > 9 {
						newValue = 0
					}
					checkValueIsUnic = true
					break
				}
			}
		}
		result[i] = newValue
	}

	return result
}

//Converts symbols from str to digits array
//Returns error if str contains symbols that not digits
func getFigures(str string) ([numberOfDigits]uint, error) {
	var result [numberOfDigits]uint
	var err error = nil

	for i := 0; i < numberOfDigits; i++ {
		result[i], err = getFigure(str[i])
		if err != nil {
			break
		}
	}

	return result, err
}

//Converts ASCII code to digit
//If code isn't digit code returns error
func getFigure(symbol byte) (uint, error) {
	const asciiCode0 byte = 0x30
	var result uint = 0
	var err error = nil

	if (symbol >= asciiCode0) && (symbol <= 0x39) {
		result = uint(symbol - 0x30)
	} else {
		err = errors.New("symbol is incorrect")
	}
	return result, err
}

//Retturns true if all values in array is unic
//otherwise returns false
func isValuesUnic(values [numberOfDigits]uint) bool {
	var result bool = true
	for i := 0; i < numberOfDigits; i++ {
		var identicalValues uint = 0
		for j := 0; j < numberOfDigits; j++ {
			if values[i] == values[j] {
				identicalValues++
			}

			if identicalValues > 1 {
				result = false
				break
			}
		}
	}
	return result
}
