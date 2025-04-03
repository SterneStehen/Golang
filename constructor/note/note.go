package note

import (
	"fmt"
	"time"
	"errors"
	"strings"
	"encoding/json"
	"os"
)

type Note struct{
	Title string `json:"title"`
	Content string `json:"content"`
	CurTime time.Time `json:"crated_at"`
}

func  NewNote(titleN, ContentN string) (Note, error){
	if titleN == "" || ContentN == ""{
		return Note{}, errors.New("errors NewNote")
	}
	return Note{
	Title: titleN,
	Content: ContentN,
	CurTime: time.Now(),
	}, nil
	
}

func (note Note) Display(){
	fmt.Printf("Your node titled %v contain next Content: %v\ntime is: %v", note.Title, note.Content, note.CurTime)
}

func (note Note) SaveJson() error{
	fileName := strings.ReplaceAll(note.Title, "", "")
	fileName = strings.ToLower(fileName)
	fileName += ".json"
	json, err := json.Marshal(note)
	if err != nil{
		return errors.New("fail json")
	}

	os.WriteFile(fileName, json, 0644)
	return nil
}