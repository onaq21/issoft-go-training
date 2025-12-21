package main

import "fmt"

func main() {
	var maxGrade, quantity, taskGrade, overallScore int
	fmt.Print("Enter the maximum score for the task: ")
	fmt.Scanln(&maxGrade)
	fmt.Print("Enter the number of tasks: ")
	fmt.Scanln(&quantity)

	for i := 1; i <= quantity; i++ {
		fmt.Printf("Enter your grade for the task №%d: ", i)
		fmt.Scanln(&taskGrade)
		if taskGrade < 0 || taskGrade > maxGrade {
			fmt.Println("Invalid grade, try again")
			i--
			continue
		}
		overallScore += taskGrade
	}

	var mark string = "F"
	switch result := overallScore * 100 / (maxGrade * quantity); {
		case 90 <= result && result <= 100:
			mark = "A"
		case 80 <= result && result <= 89:
			mark = "B"
		case 70 <= result && result <= 79:
			mark = "C"
		case 65 <= result && result <= 69:
			mark = "D"
	}
	fmt.Printf("Your mark is %s\n", mark)
}

