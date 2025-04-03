package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface{
	Save() error
}

type outPut interface {
	//Save() error
	saver
	Display()
}

// type value interface{
// 	switch value {
// 	case condition:
		
// 	}
// }

func main() {


	PrintfSameThing(1)
	PrintfSameThing(1.5)
	PrintfSameThing("hello string")
	
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")
	
	userTodo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	PrintfSameThing(userTodo)

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	//err = userTodo.Save()
	err = saveData(userNote)
	err = saveData(userTodo)
	// userNote.Display()
	// userTodo.Display()

	// if err != nil {
	// 	fmt.Println("Saving the todo failed.")
	// 	return
	// }

	// fmt.Println("Saving the todo succeeded!")

	//err = userNote.Save()

	
}

func PrintfSameThing(value interface{}){
	switch value.(type){
	case int:
		fmt.Println("value is integer", value)
	case float64:
		fmt.Println("value is float64", value)
	case string:
		fmt.Println("value is string", value)
	default:
		fmt.Println("Unknow value")
	}
}
func saveData(data outPut) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the  failed.")
		return err
	}
	data.Display()
	fmt.Println("Saving the  succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}