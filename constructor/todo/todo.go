package todo

import (
	"fmt"
	"errors"
	"encoding/json"
	"os"
)

type Todo struct{
	Text string `json:"text"`
}

func  NewTodo(textN string) (Todo, error){
	if textN == "" {
		return Todo{}, errors.New("errors NewTodo")
	}
	return Todo{
	Text: textN,
	}, nil
}

func (todo Todo) Display(){
	fmt.Printf("\nYour text: %v", todo.Text)
}

func (todo Todo) SaveJson() error{
	fileName := "todo.json"
	json, err := json.Marshal(todo.Text)
	if err != nil{
		return errors.New("fail json")
	}

	os.WriteFile(fileName, json, 0644)
	return nil
}