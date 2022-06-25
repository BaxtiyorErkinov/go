package main

import (
    "fmt"
    "errors"
)


func main() {
    // mess, allowed := enterTheClub(6)
    // fmt.Println(mess, allowed)
    fmt.Println(prediction(10))

}

func enterTheClub(age int) (string, error) {
    if age >= 18 && age < 45 {
        return "Coming", nil
    } else if age >= 45 && age < 65 {
        return "Are you sure?", nil
    } else if age >= 65 {
        return "You are very old", errors.New("You are very old!!!")
    }
    return "Closed", errors.New("You are very young")
}

func prediction(dayOfWeek int) (string, error) {
    switch dayOfWeek {
        case 1: 
            return "Monday", nil
        case 2:
            return "Tuesday", nil
        case 3: 
            return "Wednesday", nil
        case 4: 
            return "Thursday", nil
        case 5:
            return "Friday", nil
        case 6: 
            return "Saturday", nil
        case 7:
            return "Sunday", nil
        default:
            return "Hapshanba", errors.New("Invalid day of the week")
    }
}