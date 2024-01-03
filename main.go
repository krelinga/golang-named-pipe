package main

import (
    "fmt"
    "os"
    "time"
)

const pipePath = "./pipe"

func checkError(err error) {
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func main() {
    args := os.Args[1:]
    if len(args) < 1 {
        fmt.Println("Some entries to output must be specified.")
        os.Exit(1)
    }
    outFile, err := os.OpenFile(pipePath, os.O_WRONLY|os.O_SYNC, 0600)
    checkError(err)
    sleepDuration, err := time.ParseDuration("10s")
    checkError(err)
    for _, entry := range args {
        outFile.Write([]byte(entry))
        outFile.Write([]byte("\n"))
        fmt.Println("wrote ", entry)
        time.Sleep(sleepDuration)
    }
}
