package main

import (
    "fmt"
)

var message string


func main() {
    fmt.Println(message)

    res := increment()
    fmt.Println(res())
    fmt.Println(res())
    fmt.Println(res())
    fmt.Println(res())
}

func increment() func() int {
    count := 0
    return func() int {
        count++
        return count
    } 
} 

func init() {
    message = "hello world"
}