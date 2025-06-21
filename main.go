package main

import "fmt"


func main()  {
	ShortGo := "Watch GO crash course."
	TradingCourse := "Watch Trading crash course."
	Leisure := "Rest and get back to work."

	taskItems := []string {ShortGo, TradingCourse, Leisure}

	fmt.Println("##### Hello Chuks, welcome to our Todolist App #####")

	for index, task := range taskItems {
		// fmt.Println(index+1,".", task)
		fmt.Printf("%d: %s\n", index+1, task)
	}
}
