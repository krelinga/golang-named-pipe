package main

import (
    "fmt"
    "os"
    "time"
)

func checkError(err error) {
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func main() {
    args := os.Args[1:]
    if len(args) != 1 {
        fmt.Println("output file path must be specified.")
        os.Exit(1)
    }
    outFilePath := args[0]
    outFile, err := os.OpenFile(outFilePath, os.O_WRONLY|os.O_SYNC, 0600)
    checkError(err)
    outFile.Write([]byte("this is a test\n"))
    sleepDuration, err := time.ParseDuration("20s")
    checkError(err)
    time.Sleep(sleepDuration)
}
