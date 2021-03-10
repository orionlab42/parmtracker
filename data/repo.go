package data

import "fmt"

var currentId int

var ToDos Todos

// Give us some seed data
func init() {
	RepoCreateTodo(Todo{Name: "Write presentation"})
	RepoCreateTodo(Todo{Name: "Host meetup"})
}

func RepoFindTodo(id int) Todo {
	for _, t := range ToDos {
		if t.Id == id {
			return t
		}
	}
	// return empty To do if not found
	return Todo{}
}

func RepoCreateTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	ToDos = append(ToDos, t)
	return t
}

func RepoDestroyTodo(id int) error {
	for i, t := range ToDos {
		if t.Id == id {
			ToDos = append(ToDos[:i], ToDos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("could not find Todo with id of %d to delete", id)
}
