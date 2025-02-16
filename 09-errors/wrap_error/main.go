package main

import (
    "errors"
    "fmt"
    "os"
)

func fileChecker(name string) error {
    f, err := os.Open(name)
    if err != nil {
        return fmt.Errorf("in fileChecker: %w", err)
    }

    f.Close()
    return nil
}

func main() {
    err := fileChecker("not_here.txt")
    if err != nil {
        fmt.Println(err) // in fileChecker: open not_here.txt: The system cannot find the file specified.
        if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
            fmt.Println(wrappedErr) // open not_here.txt: The system cannot find the file specified.
        }
    }
}
