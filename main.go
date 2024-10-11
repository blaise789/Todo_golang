package main

// import "fmt"

func main() {
	todos := Todos{}
	cmdFl:=NewCmdFlags()
	storage:=&Storage[Todos]{"todos.json"}
	storage.Load(&todos)
	cmdFl.Execute(&todos)
	storage.Save(todos)

}
