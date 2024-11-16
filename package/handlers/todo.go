package handlers

import (
	"errors"
	"fmt"
	"time"

	"github.com/MdSadiqMd/GoTasker/package/types"
)

type Todos []types.Todo

func (todos *Todos) Add(title string) {
	todo := types.Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}
	*todos = append(*todos, todo)
}

func (todos *Todos) ValidateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) Delete(index int) error {
	if err := todos.ValidateIndex(index); err != nil {
		return err
	}
	*todos = append((*todos)[:index], (*todos)[index+1:]...)
	return nil
}

func (todos *Todos) Toggle(index int) error {
	t := (*todos)
	if err := t.ValidateIndex(index); err != nil {
		return err
	}

	isCompleted := t[index].Completed
	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}
	t[index].Completed = !isCompleted

	return nil
}

func (todos *Todos) Edit(index int, title string) error {
	t := *todos
	if err := t.ValidateIndex(index); err != nil {
		return err
	}
	t[index].Title = title

	return nil
}
