package main

import (
	"fmt"
	"net/http"
)

var ShortGo = "Watch GO crash course."
var TradingCourse = "Watch Trading crash course."
var Leisure = "Rest and get back to work."
var taskItems = []string {ShortGo, TradingCourse, Leisure}



func main()  {
	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)
	http.ListenAndServe(":8080", nil)
}
func showTasks(writer http.ResponseWriter, request *http.Request){
	for _, task := range taskItems {
		fmt.Fprintln(writer, task)
	}
}

func helloUser(writer http.ResponseWriter, request *http.Request){
	var greeting = "Welcome to our todolist APP"
	fmt.Fprintln(writer, greeting)
}




// func printTasks(taskItems []string) {
// 	fmt.Println("List of my Todos")
// 	for index, task := range taskItems {
// 		// fmt.Println(index+1,".", task)
// 		fmt.Printf("%d: %s\n", index+1, task)
// 	}
// }


// func addTasks(taskItems []string, newTask string) []string {
// 	var updatedTaskItems = append(taskItems, newTask)
// 	return updatedTaskItems
// }
