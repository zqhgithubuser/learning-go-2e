package main

import (
    "encoding/json"
    "fmt"
    "log/slog"
    "os"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func ProcessPerson() error {
    toFile := Person{
        Name: "Fred",
        Age:  40,
    }

    // Write it out
    tmpFile, err := os.CreateTemp(os.TempDir(), "sample-")
    if err != nil {
        panic(err)
    }
    defer os.Remove(tmpFile.Name())
    err = json.NewEncoder(tmpFile).Encode(toFile)
    if err != nil {
        panic(err)
    }
    err = tmpFile.Close()
    if err != nil {
        panic(err)
    }

    // // Read it back
    tmpFile2, err := os.Open(tmpFile.Name())
    if err != nil {
        panic(err)
    }
    var fromFile Person
    err = json.NewDecoder(tmpFile2).Decode(&fromFile)
    if err != nil {
        panic(err)
    }
    err = tmpFile2.Close()
    if err != nil {
        panic(err)
    }
    fmt.Printf("%+v\n", fromFile) // {Name:Fred Age:40}
    return nil
}

func main() {
    err := ProcessPerson()
    if err != nil {
        slog.Error("error in processPerson", "msg", err)
    }
}
