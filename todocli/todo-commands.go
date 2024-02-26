package main

import (
	"flag"
	"fmt"
	"log"
	"todos/api"
)

var (
	todoFlags = flag.NewFlagSet("todo", flag.ExitOnError)
	todoID    = todoFlags.String("todo-id", "", "ID for todo when specifying a Todo")
	text      = todoFlags.String("text", "", "text used for todo create and update")
)

func todoCommands(args []string) {
	if len(args) < 1 {
		log.Println("An todo subcommand is required")
		return
	}

	token, err := readToken()
	if err != nil {
		log.Fatal("Error reading token", err)
	}

	todoFlags.Parse(args[1:])
	switch client := api.NewClient(*host, token); args[0] {
	case "mine":
		todos, err := client.ListTodos()
		if err != nil {
			log.Fatal("Error fetch your todos", err)
		}
		log.Println("Your Tasks")
		for _, t := range todos {
			printTodo(&t)
		}

	case "create":
		if *text == "" {
			log.Println("No text provided with -text flag")
			return
		}
		todo, err := client.CreateTodo(*text)
		if err != nil {
			log.Fatal("Failed to create todo:", err)
		}
		printTodo(todo)

	case "get":
		if *todoID == "" {
			log.Println("No id provided with -todo-id flag")
			return
		}
		todo, err := client.FetchTodo(*todoID)
		if err != nil {
			log.Fatal("Failed to fetch todo:", err)
		}
		printTodo(todo)

	case "complete":
		if *todoID == "" {
			log.Println("No id provided with -todo-id flag")
			return
		}
		t, err := client.FetchTodo(*todoID)
		if err != nil {
			log.Fatal("Failed to fetch todo:", err)
		}
		t, err = client.UpdateTodo(*todoID, t.Text, true)
		if err != nil {
			log.Fatal("Failed to update todo:", err)
		}
		printTodo(t)

	case "remove":
		if *todoID == "" {
			log.Println("No id provided with -todo-id flag")
			return
		}
		todo, err := client.DeleteTodo(*todoID)
		if err != nil {
			log.Fatal("Failed to delete todo:", err)
		}
		printTodo(todo)

	default:
		log.Println("Unknown todo command:", args[0])
	}
}

func printTodo(t *api.Todo) {
	var status string
	if t.Complete {
		status = "Done"
	} else {
		status = "Open"
	}
	fmt.Printf("\t- (%s) [%s] %s\n", t.ID, status, t.Text)
}
