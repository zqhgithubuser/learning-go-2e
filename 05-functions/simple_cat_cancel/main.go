package main

import (
    "io"
    "log"
    "os"
)

func getFile(name string) (*os.File, func(), error) {
    file, err := os.Open(name)
    if err != nil {
        return nil, nil, err
    }
    return file, func() {
        file.Close()
    }, nil
}

func main() {
    f, closer, err := getFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer closer()

    data := make([]byte, 2048)
    for {
        count, err := f.Read(data)
        if err != nil {
            if err != io.EOF {
                log.Fatal(err)
            }
            break
        }
        os.Stdout.Write(data[:count])
    }
}
