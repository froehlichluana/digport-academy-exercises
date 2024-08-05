package logic 

import "fmt"

func ZeroCheck(){

	var number int

	fmt.Println("Enter a number: ")
	fmt.Scanf("%d", &number)

	if number > 0{
		fmt.Println("Greater than zero.")
	} else if number < 0  {
		fmt.Println("Less than zero.")
	} else{
		fmt.Println("Zero. ")
}

}