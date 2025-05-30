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
        if errors.Is(err, os.ErrNotExist) {
            fmt.Println(errors.Unwrap(err))
            // open not_here.txt: The system cannot find the file specified.
        }
    }
}
