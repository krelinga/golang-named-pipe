package main

import (
    "fmt"
    "os"
    "time"
)

const pipePath = "/pipedir/pipe"

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
    fmt.Println("opening ", pipePath)
    const openFlags = os.O_WRONLY|os.O_SYNC|os.O_APPEND
    outFile, err := os.OpenFile(pipePath, openFlags , 0600)
    checkError(err)
    defer outFile.Close()
    sleepDuration, err := time.ParseDuration("10s")
    checkError(err)
    for _, entry := range args {
        outFile.Write([]byte(entry))
        outFile.Write([]byte("\n"))
        fmt.Println("wrote ", entry)
        time.Sleep(sleepDuration)
    }
}
