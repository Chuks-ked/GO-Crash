package main

import "fmt"


func main()  {
	fmt.Println("##### Hello Chuks, welcome to our Todolist App #####")

	var ShortGo = "Watch GO crash course."
	var TradingCourse = "Watch Trading crash course."
	var Leisure = "Rest and get back to work."
	var taskItems = []string {ShortGo, TradingCourse, Leisure}

	printTasks(taskItems)
	fmt.Println()
	taskItems = addTasks(taskItems, "Go for a run")

	fmt.Println("updated list")
	printTasks(taskItems)
}


func printTasks(taskItems []string) {
	fmt.Println("List of my Todos")
	for index, task := range taskItems {
		// fmt.Println(index+1,".", task)
		fmt.Printf("%d: %s\n", index+1, task)
	}
}


func addTasks(taskItems []string, newTask string) []string {
	var updatedTaskItems = append(taskItems, newTask)
	return updatedTaskItems
}
