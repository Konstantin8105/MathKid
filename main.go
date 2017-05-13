package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var maxResult = 25
var minResult = 0
var amountTasks = 20

type sign bool

const (
	plus  sign = true
	minus      = false
)

type numbers struct {
	values [3]int
	signs  [2]sign
}

func generateValue() int {
	return rand.Intn(maxResult)
}

func generateSign() sign {
	return sign(rand.Intn(2) == 0)
}

func calculuteNumber(n numbers) (result int) {
	result = n.values[0]
	if n.signs[0] == plus {
		result += n.values[1]
	} else {
		result -= n.values[1]
	}
	if n.signs[1] == plus {
		result += n.values[2]
	} else {
		result -= n.values[2]
	}
	return result
}

func generateNumber() (n numbers) {
	for i := range n.values {
		n.values[i] = generateValue()
	}
	for i := range n.signs {
		n.signs[i] = generateSign()
	}
	if calculuteNumber(n) > maxResult || calculuteNumber(n) < minResult {
		return generateNumber()
	}
	if n.values[0] == n.values[1] && n.values[1] == n.values[2] {
		return generateNumber()
	}
	if n.signs[0] == plus && n.values[0]+n.values[1] > maxResult {
		return generateNumber()
	}
	if n.signs[0] == plus && n.signs[1] == plus && n.values[1]+n.values[2] > maxResult {
		return generateNumber()
	}
	if n.signs[1] == plus && n.values[0]+n.values[2] > maxResult {
		return generateNumber()
	}
	if n.signs[0] == minus && n.values[0] < n.values[1] {
		return generateNumber()
	}
	return n
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	var summaryResult []bool

	for i := 0; i < amountTasks; i++ {
		n := generateNumber()
	AGAIN:
		// out
		fmt.Printf("\n\n_______________\n\n")
		fmt.Println("Пример №", i+1, " из ", amountTasks)
		fmt.Printf("Дан следующий пример:\n")
		fmt.Printf("%v", n.values[0])
		if n.signs[0] == plus {
			fmt.Printf("+")
		} else {
			fmt.Printf("-")
		}
		fmt.Printf("%v", n.values[1])
		if n.signs[1] == plus {
			fmt.Printf("+")
		} else {
			fmt.Printf("-")
		}
		fmt.Printf("%v", n.values[2])
		fmt.Printf("\n")
		fmt.Println("Реши и запиши ответ:")
		// get result
		var result int
		var line string
		_, err := fmt.Scanf("%s", &line)
		if err != nil {
			fmt.Println("\nНе понятный ответ")
			//fmt.Println("Неверный ответ. Попробуй снова")
			//fmt.Println("t = ", t)
			//fmt.Println("err1 = ", err)
			goto AGAIN
		}
		r, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			fmt.Println("\nНе понятный ответ")
			//fmt.Println("Неверный ответ. Попробуй снова")
			//fmt.Println("err2 = ", err)
			goto AGAIN
		}
		result = int(r)
		// output fail or ok
		if result != calculuteNumber(n) {
			fmt.Println("Неверный ответ. Попробуй снова")
			summaryResult = append(summaryResult, false)
			goto AGAIN
		}
		fmt.Println("Правильно")
		summaryResult = append(summaryResult, true)
	}

	// results
	positive := 0
	for i := range summaryResult {
		if summaryResult[i] {
			positive++
		}
	}
	fmt.Println("\n\n_______________")
	fmt.Println("Количество правильных ответов : ", positive)
	fmt.Println("Количество неправильных ответов : ", len(summaryResult)-positive)

}
