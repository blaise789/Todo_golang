package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}
type Todos []Todo

func (todos *Todos) add(title string) {
	newTodo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	*todos = append(*todos, newTodo)
	fmt.Println("new todo added successfully")

}
func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index > len(*todos) {
		err := errors.New("invalid index")
		return err
	}
	return nil
}
func (todos *Todos) delete(index int) error {
	t := *todos
	err := t.validateIndex(index)
	if err != nil {
		fmt.Println(err)

		return err
	}
	*todos = append(t[:index], t[index+1:]...)
	return nil
}
func (todos *Todos) toggle(index int) {
	//
	t := *todos
	err := t.validateIndex(index)
	if err != nil {
		fmt.Println(err)
		return
	}
	t[index].Completed = !t[index].Completed
	if !t[index].Completed {
		t[index].CompletedAt = nil

	}
	completedTime:=time.Now()
	t[index].CompletedAt = &completedTime
	*todos = t

}
func (todos *Todos) edit(index int,value string) error{
	t:=*todos
	err:=t.validateIndex(index)
	if err!=nil{
        fmt.Println(err)
        return err
    }
	t[index].Title=value
	*todos=t
	return nil
}

func (todos *Todos) print(){
	table:=table.New(os.Stdout)
	table.AddHeaders("#","title","completed","createdAt","completedAt")
	for i,todo:=range *todos{
		completed:="❌"
		completedAt:=""
		if todo.Completed{
			completed="✅"
			if todo.CompletedAt !=nil{
				completedAt=todo.CompletedAt.Format(time.RFC1123)
			}

		}
		table.AddRow(strconv.Itoa(i),todo.Title,completed,todo.CreatedAt.Format(time.RFC1123),completedAt)
	}
	table.Render()
}