package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)
type CmdFlags struct{
	Add string
	Del int
	Edit string 
	Toggle int 
	List bool 
}
func NewCmdFlags()*CmdFlags{
	cf:=CmdFlags{}
	flag.StringVar(&cf.Add,"add","","Add new todo by specifying the title")
	flag.StringVar(&cf.Edit,"edit","","Edit todo")
	flag.IntVar(&cf.Toggle,"toggle",-1,"toggle to change the completion of tasks")
	flag.BoolVar(&cf.List,"list",false,"list the todos")
	flag.Parse()
	return &cf
}
func (cf *CmdFlags) Execute(todos *Todos)  {
	switch{
	case cf.List:
		todos.print()
    case cf.Add!="":
		todos.add(cf.Add)	
    case cf.Edit!="":
		parts:=strings.SplitN(cf.Edit,":",2)
		if len(parts) !=2{
			fmt.Println("Error,Invalid format")
			os.Exit(1)
		}
		index,err:=strconv.Atoi(parts[0])
		if err!=nil{
			fmt.Println("invalid index for edit")
			os.Exit(1)
		}
		todos.edit(index,cf.Edit)			
    case cf.Toggle!=-1:
		todos.toggle(cf.Toggle)		
    case cf.Del!=-1:
		todos.delete(cf.Del)
    default:
		fmt.Println("invalid command")				
	}
	
	
}