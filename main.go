package main

import (
    "fmt"
)

func main() {
    var mess string = "Hello World" // string
    var num int = 1 // integer - целое число
    var num2 float32 = 1.1 // float - тип с плавающей запятой 
    var b bool = true // boolean - true | false
    var a byte = 68
    var r rune = 'a'


    fmt.Println(mess)
    fmt.Println(num)
    fmt.Println(num2)
    fmt.Println(b)
    fmt.Printf("%c\n", a)    
    fmt.Println(r)
}