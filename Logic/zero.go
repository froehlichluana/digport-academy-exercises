package logic 

import "fmt"

func Zero(){
	var number int 

	fmt.Println("Enter a number: ")
	fmt.Scan(number)

	if (number > 0){
		fmt.Println("Greater than 0.")
	if (number == 0) {
		fmt.Println("Zero.")
	}else{
		fmt.Println("Lessa than zero.")
	}
	}
}