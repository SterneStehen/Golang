package main

import (
	"fmt"
	"errors"
	"os"
	"bufio"
	"strings"
	"constructor/note"
	"constructor/todo"
	
)

func main(){
	title, content, err := getInput("Enter title: ", "Enter content: ")
	
	if err != nil{
		fmt.Println(err)
		return
	}
	//fmt.Println("title is ", title, "content is ", content)
	
	var userNote note.Note
	userNote, err = note.NewNote(title, content)
	if(err != nil){
		fmt.Println(err)
		return
	}

	

	//fmt.Println(st.title, st.content, st.curTime)
	
	err = userNote.SaveJson()
	if(err != nil){
		fmt.Println(err)
	}
	var usTodo todo.Todo
	textTodo, err := getTodoImput("Enter text: ")
	usTodo, err = todo.NewTodo(textTodo)
	if(err != nil){
		fmt.Println(err)
	}
	err = usTodo.SaveJson()
	if(err != nil){
		fmt.Println(err)
	}
	userNote.Display()
	usTodo.Display()
	fmt.Println("\nJson succesfull writed")
}

func getInput(titleText, contText string)(string, string, error){
	var resTitle, resCont string
	fmt.Print(titleText)
	fmt.Scanln(&resTitle)
	if resTitle == ""{
		return "", "", errors.New("error input title")
	}
	fmt.Print(contText)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil{
		return "", "", errors.New("error input title")
	}
	text = strings.TrimSuffix(text, "\n")
	resCont = strings.TrimSuffix(text, "\r")
	//fmt.Scanln(&resCont)
	if resCont == ""{
		return "", "", errors.New("error input content")
	}
	return resTitle, resCont, nil
}

func getTodoImput(todoText string) (string, error){
	//var text string
	fmt.Print(todoText)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil{
		return "", errors.New("error input text")
	}
	text = strings.TrimSuffix(text, "\n")
	resCont := strings.TrimSuffix(text, "\r")
	//fmt.Scanln(&resCont)
	if resCont == ""{
		return "", errors.New("error input content")
	}
	return resCont, nil
}