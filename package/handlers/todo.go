package handlers

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/MdSadiqMd/GoTasker/package/types"
	"github.com/aquasecurity/table"
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

func (todos *Todos) Print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	for index, t := range *todos {
		completed := "❌"
		completedAt := ""
		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)

			}

		}
		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}

	table.Render()
}
