package main

import (
	"errors"
	"fmt"
	"time"
)
type Todo struct{
	Title string
    Completed bool
	CreatedAt time.Time
	CompletedAt *time.Time
	
}
type Todos []Todo
func (todos *Todos) add(title string){ 
	newTodo:=Todo{
		Title: title ,
		Completed: false,
		CreatedAt: time.Now(),
		CompletedAt: nil ,
	}
	*todos=append(*todos, newTodo)
}
func (todos *Todos) validateIndex(index int) error{
	fmt.Print(*todos)
	if  index<0 || index>len(*todos){
     err:=errors.New("invalid index")
	 fmt.Println(err)
	 return err
	}
  return nil
}
func (todos *Todos) delete(index int) error{
	t:=*todos
	err:=t.validateIndex(index)
	if err !=nil{
		return err
	}
	*todos=append(t[:index], t[index+1:]...)
    return nil
}
func (todos *Todos) toggle(index int) {
	//
	t:=*todos
	err:=t.validateIndex(index)
	if err!=nil{
		fmt.Println(err)
		return 
	}
	t[index].Completed=!t[index].Completed
	if !t[index].Completed{
      t[index].CompletedAt=nil
      
	}
	*t[index].CompletedAt=time.Now()
	*todos=t
	
	

}