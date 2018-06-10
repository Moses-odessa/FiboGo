package main

import (
	"fmt"
	"time"
)

func fibo(i int) int {
	var result int
	if i < 1 {
		result = 0
	} else if i == 1 {
		result = 1
	} else {
		result = fibo(i-1) + fibo(i-2)
	}
	return result
}

func check(trueAnswerQty int, mistakeAnswerQty int, resultChanel chan resulData) {
	var mistakeCount = 0
	var trueCount = 0
	for i := 1; trueCount < trueAnswerQty && mistakeCount < mistakeAnswerQty; i++ {
		fmt.Printf("Enter %d fibo value (within 10 seconds):\n", i)
		msg := <-resultChanel
		currentFibo := fibo(i)
		if !msg.fromUser || (msg.value != currentFibo && msg.fromUser) {
			mistakeCount++
			trueCount = 0
			fmt.Printf("Number of mistake: %d (of %d). True answer is: {order: %d, value: %d}\n", mistakeCount, mistakeAnswerQty, i, currentFibo)
		} else {
			trueCount++
			fmt.Printf("Correct!!! Number of correct answers: %d (need %d)\n", trueCount, trueAnswerQty)

		}
	}
	if mistakeCount == mistakeAnswerQty {
		fmt.Printf("It was %d mistakes.\n", mistakeAnswerQty)
	} else {
		fmt.Printf("You have %d true answers.\n", trueCount)
	}
	fmt.Println("Press enter for exit...")
}

func userInterface(trueAnswerQty int, mistakeAnswerQty int, answerTime time.Duration, resultChanel chan resulData) {
	var mistakeCount = 0
	var trueCount = 0

	for nextOrder := 1; mistakeCount < mistakeAnswerQty && trueCount < trueAnswerQty; nextOrder++ {
		ticker := time.NewTicker(time.Second * answerTime)
		go func() { //ticker
			for range ticker.C {
				var result = resulData{0, false}
				resultChanel <- result
				nextOrder++
				mistakeCount++
				trueCount = 0
			}
		}()

		var nextInput int
		fmt.Scanf("%d\n", &nextInput)
		ticker.Stop()
		if nextInput == fibo(nextOrder) {
			trueCount++
		} else {
			mistakeCount++
			trueCount = 0
		}
		if trueCount <= trueAnswerQty && mistakeCount <= mistakeAnswerQty {
			var result = resulData{nextInput, true}
			resultChanel <- result
		}
	}
	if mistakeCount == mistakeAnswerQty || trueCount == trueAnswerQty {
		var input string
		fmt.Scanln(&input)
	}
}

type resulData struct {
	value    int
	fromUser bool
}

func main() {
	var resultChanel = make(chan resulData)
	const trueAnswerQty = 10
	const mistakeAnswerQty = 3
	const answerTime = 10

	go check(trueAnswerQty, mistakeAnswerQty, resultChanel)

	userInterface(trueAnswerQty, mistakeAnswerQty, answerTime, resultChanel)

}
