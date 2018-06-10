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

func check(mumbersAmount int, resultChanel chan resulData) {
	for i := 1; i <= mumbersAmount; {
		fmt.Println("Enter next fibo value (within 10 seconds):")
		msg := <-resultChanel
		currentFibo := fibo(i)
		if !msg.fromUser {
			fmt.Printf("You late. True answer is: %d", currentFibo)
			fmt.Println()
		} else {
			if msg.value != currentFibo {
				fmt.Printf("Your answer wrong. True answer is: %d", currentFibo)
				fmt.Println()
			} else {
				fmt.Println("Correct")
			}
		}
		i++
	}
	fmt.Println("Press enter for exit...")
}

func userInterface(mumbersAmount int, answerTime time.Duration, resultChanel chan resulData) {
	var nextOrder = 1
	for nextOrder <= mumbersAmount { //input
		ticker := time.NewTicker(time.Second * answerTime)
		go func() { //ticker
			for range ticker.C {
				var result = resulData{0, false}
				resultChanel <- result
				nextOrder++
			}
		}()

		var nextInput int
		fmt.Scanf("%d\n", &nextInput)
		ticker.Stop()
		if nextOrder <= mumbersAmount {
			var result = resulData{nextInput, true}
			resultChanel <- result
		}
		nextOrder++
	}

	if nextOrder < (mumbersAmount + 2) {
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
	const mumbersAmount = 10
	const answerTime = 10

	go check(mumbersAmount, resultChanel)

	userInterface(mumbersAmount, answerTime, resultChanel)

}
