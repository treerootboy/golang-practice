package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

func test1() {
    fmt.Println(os.Args[0])
}

func test2() {
    for idx, arg := range os.Args[0:] {
        s := "key:" + strconv.Itoa(idx) + " and value:" + arg
        fmt.Println(s)
    }   
}

func test3() {
    s := strings.Join(os.Args, " ")
    fmt.Println(s)
}

func main() {
    test1()
    test2()
    test3()
}