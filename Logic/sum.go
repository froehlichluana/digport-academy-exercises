package logic 

import "fmt"

func Sum() {
	var entryNumber int
	var sum int

	fmt.Println("Enter a number: ")
	fmt.Scanf("%d", &entryNumber)

	if entryNumber > 0 {
		sum = entryNumber + 10
		fmt.Printf("The result of the sum is %d", sum)
	} else if entryNumber == 0 {
		sum = entryNumber + 2
		fmt.Printf("The result of the sum is %d", sum)
	} else {
		sum = entryNumber + 23
		fmt.Printf("The result of the sum is %d", sum)
	}

}