package main

import (
	"fmt"
	"math/rand"
	"time"
)

const maxValue int = 9999
const minValue uint = 1000

func main() {
	const figureDeletedFromSearch uint = 10

	var gameOver bool = false
	var userInput uint
	var value uint = generateValue()

	for !gameOver {
		fmt.Println("Введите число из четырех цифр:")
		fmt.Scan(&userInput)

		if userInput > uint(maxValue) {
			fmt.Println("Число слишком большое!")
		} else if userInput < minValue {
			fmt.Println("Число слишком маленькое!")
		} else {
			if userInput == value {
				fmt.Println("Победа! Вы угадали число.")
				gameOver = true
			} else {
				var userFigures [4]uint = getFigures(userInput)
				var computerFigures [4]uint = getFigures(value)

				var bulls uint = 0
				var cows uint = 0

				for i := 0; i < len(userFigures); i++ {
					if userFigures[i] == computerFigures[i] {
						bulls++
						userFigures[i] = figureDeletedFromSearch
						computerFigures[i] = figureDeletedFromSearch
					}
				}

				for i := 0; i < len(userFigures); i++ {
					for j := 0; j < len(userFigures); j++ {
						if (userFigures[j] != figureDeletedFromSearch) && (computerFigures[i] != figureDeletedFromSearch) {
							if userFigures[j] == computerFigures[i] {
								cows++
								userFigures[j] = figureDeletedFromSearch
								computerFigures[i] = figureDeletedFromSearch
							}
						}

					}
				}

				fmt.Println("Не угадали. Быки:", bulls, "Коровы:", cows)
			}
		}
	}

	fmt.Scan(&userInput)
}

func getFigures(value uint) [4]uint {
	var result [4]uint
	result[0] = value / 1000
	result[1] = (value % 1000) / 100
	result[2] = (value % 100) / 10
	result[3] = value % 10
	return result
}

func generateValue() uint {
	rand.Seed(time.Now().Unix())
	var value uint = uint(rand.Intn(maxValue))
	if value < minValue {
		value += minValue
	}
	return value
}
