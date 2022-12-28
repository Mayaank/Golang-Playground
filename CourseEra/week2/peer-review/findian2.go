package main

import (
  "fmt"
  "strings"
)

func main() {

  var userInput string
  fmt.Scan(&userInput)

  userInput = strings.ToLower(userInput)

  lengthMask := len(userInput) < 3

  if lengthMask {
    fmt.Println("Not found!")

  } else {

    firstCharMask := userInput[0] == 'i'
    lastCharMask := userInput[len(userInput) - 1] == 'n'

    // Checks if 'a' is in the "middle" of the user input.
    middleCharMask := strings.ContainsRune(userInput[1:len(userInput) - 1], 'a')

    matchMask := firstCharMask && middleCharMask && lastCharMask

    if matchMask {
      fmt.Println("Found!")

    } else {
      fmt.Println("Not found!")
    }

  }

}
