package main

import (
    "fmt"
)

func main() {

    // text := 77
    // fmt.Println(&text) // 0xc000098000

    // address := &text // 0xc000098000
    // fmt.Println(*address) // 77
    // fmt.Println(*&text) // 77
    // fmt.Printf("address это %T\n", &text) // *int
    // var el *int //   здесь он ожидает pointer (ссылку). Например: &text
    // fmt.Println(el) // nil



    // mess := "Hello World" // 0x1234
    // fmt.Println(mess) // 0x1234
    // changeMess(&mess)
    // fmt.Println(mess) // 0x1234

    val1 := "val1"
    val2 := "val2"
    replaceOfValues(&val1, &val2)
    fmt.Println("Val1: ", val1)
    fmt.Println("Val2: ",  val2)
}



func changeMess(mess *string) { // 0x1234
    *mess += "Say Hello"
}

func printMess(mess string) {
    fmt.Println(mess) // 0x12345
}

func replaceOfValues(val1 *string, val2 *string) {
    newVal1 := *val1
    newVal2 := *val2
    *val1 = newVal2
    *val2 = newVal1
}