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

func main() {
	var resultChanel = make(chan [2]int)

	go func() { //check
		for i := 1; i <= 10; {
			fmt.Println("Enter next fibo value:")
			msg := <-resultChanel
			currentFibo := fibo(i)
			if msg[1] == i {
				fmt.Printf("You late. True answer is: %d", currentFibo)
				fmt.Println()
			} else {
				if msg[0] != currentFibo {
					fmt.Printf("Your answer wrong. True answer is: %d", currentFibo)
					fmt.Println()
				} else {
					fmt.Println("Correct")
				}
			}
			i++
		}
		fmt.Println("Press enter for exit...")
	}()

	for nextOrder := 1; nextOrder <= 10; { //input
		ticker := time.NewTicker(time.Second * 10)
		go func() { //ticker
			for range ticker.C {
				var result = [2]int{0, nextOrder}
				resultChanel <- result
				nextOrder++
			}
		}()

		var nextInput int
		fmt.Scanf("%d\n", &nextInput)
		ticker.Stop()
		if nextOrder <= 10 {
			var result = [2]int{nextInput, 0}
			resultChanel <- result
			nextOrder++
		}
	}

	var input string
	fmt.Scanln(&input)
}
