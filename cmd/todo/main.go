package main

import (
    "fmt"
    "os"
    "strings"

    "joshuaejs/todo"
)

// hardcoding the filename, for now
const todoFileName = ".todo.json"

func main() {
    // define an items List
    l := &todo.List{}

    // use the Get method to read ToDo items from file
    if err := l.Get(todoFileName); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

    // decide what to do based upon the number of arguments provided
    switch {
    // for no extra arguments, print the list
    case len(os.Args) == 1:
        // list current ToDo items
        for _, item := range *l {
            fmt.Println(item.Task)
        }
    // use a default case to check/concatenate/add a new item
    default:
        // contatenate all arguments with a space
        item := strings.Join(os.Args[1:], " ")

        // Add the task
        l.Add(item)

        // Save the new List
        if err := l.Save(todoFileName); err != nil {
            fmt.Fprintln(os.Stderr, err)
            os.Exit(1)
        }
    }

}

