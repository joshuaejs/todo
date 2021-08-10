package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"joshuaejs/todo"
)

// default filename
var todoFileName = ".todo.json"

// getTask picks where to get the description for a new task from: arguments or STDIN
func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}
	if len(s.Text()) == 0 {
		return "", fmt.Errorf("task cannot be blank")
	}

	return s.Text(), nil
}

func main() {
	// add some defaults to the help message
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			"%s application, as demonstrated in Powerful Command-Line Applications in Go\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2021\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage information:\n")
		flag.PrintDefaults()
	}
	// parse command line flags
	add := flag.Bool("add", false, "Add task to the ToDo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item number to be completed")
    del := flag.Int("del", 0, "Item number to be deleted")

	flag.Parse()

	// check if ENV VAR for a custom filename is defined
	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}

	// define an items List
	l := &todo.List{}

	// use the Get method to read ToDo items from file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// decide what to do based upon the number of arguments provided
	switch {
	case *list:
		// list current ToDo items
		fmt.Print(l)
	case *complete > 0:
		// Complete the given item
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *add:
		// Add the task, any arguments, except flags, are used as the new task
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		l.Add(t)

		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
    case *del > 0:
        // Delete the given task
        if err := l.Delete(*del); err != nil {
            fmt.Fprintln(os.Stderr, err)
            os.Exit(1)
        }
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	// use a default case to check/concatenate/add a new item
	default:
		// Invalid flag provided
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}
