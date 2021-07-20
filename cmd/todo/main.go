package main

import (
    "flag"
    "fmt"
    "os"

    "joshuaejs/todo"
)

// hardcoding the filename, for now
const todoFileName = ".todo.json"

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
    task := flag.String("task", "", "Task to be added to/included in the ToDo list")
    list := flag.Bool("list", false, "List all tasks")
    complete := flag.Int("complete", 0, "Item number to be completed")

    flag.Parse()

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
        for _, item := range *l {
            if !item.Done {
                fmt.Println(item.Task)
            }
        }
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
    case *task != "":
        // Add the task
        l.Add(*task)
        // Save the new list
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

