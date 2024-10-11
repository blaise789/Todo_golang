package main

import "fmt"

func main() {
	todos := Todos{}
	todos.add("taking breakfast")
	todos.add("going to sleep")
   fmt.Println(todos)
}
