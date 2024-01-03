package main

import (
    "fmt"
    "os"
)

func main() {
    args := os.Args[1:]
    if len(args) != 1 {
        fmt.Println("output file path must be specified.")
        os.Exit(1)
    }
    fmt.Println(os.Args[1:])
}
