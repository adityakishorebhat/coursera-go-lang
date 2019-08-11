package main

import "fmt"
import "strings"

// func main() {
//   fmt.Printf("Enter a string : ")
//   var inputString string
//   inputString = strings.ToLower(inputString)
//   num, _ := fmt.Scan(&inputString)
//   if num > 0 {
//     if (strings.Contains(inputString, "a")) {
//       var found bool = true
//       for i, c := range inputString {
//         if i == 0 && c != "i" {
//           found = false
//           fmt.Println("Not found!")
//           break
//         } else if i == len(inputString - 1) && c != "n" {
//           found = false
//           fmt.Println("Not found!")
//         }
//       }
//       if found == true {
//         fmt.Println("Found!")
//       }
//     }
//   }
// }

func main() {
    fmt.Printf("Enter a string : ")
    var inputString string
    num, _ := fmt.Scan(&inputString)
    if (num > 0) {
      inputString = strings.ToLower(inputString)
      fmt.Println(inputString)
      if len(inputString) < 3 {
        fmt.Println("Not found!")
      } else {
        if strings.Contains(inputString, "a") {
          var found bool = true
          for i, c := range inputString {
            if i == 0 && string(c) != "i" {
              found = false
              fmt.Println("Not found!")
              break
            }
            if i == len(inputString) - 1 && string(c) != "n" {
              found = false
              fmt.Println("Not found!")
            }
          }
          if found == true {
            fmt.Println("Found!")
          }
        } else {
          fmt.Println("Not found!")
        }
      }
    }
}
