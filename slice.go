package main

import "fmt"
import "strconv"
import "os"
import "sort"

func main() {
  sliceLength := 3
  intSlice := make([]int, sliceLength)
  fmt.Printf("Enter an Integer : ")
  var inputString string
  num, _ := fmt.Scan(&inputString)
  if (num == 0) {
    os.Exit(3)
  }
  slicePointer := 0
  for (inputString != "X") {
    if (isNumeric(inputString)) {

      i, _ := strconv.Atoi(inputString)
      if (slicePointer < sliceLength) {
        fmt.Printf("Slice Pointer : %v\n", slicePointer)
        fmt.Printf("Slice Length : %v\n", sliceLength)
        intSlice[slicePointer] = i
        slicePointer++
      } else {
        intSlice = append(intSlice, i)
      }
      sliceToSort := make([]int, len(intSlice))
      copy(sliceToSort, intSlice)
      sort.Ints(sliceToSort)
      fmt.Printf("Original Slice : %v\n", intSlice)
      fmt.Printf("Sorted Slice: %v\n", sliceToSort)
    }

    fmt.Printf("Enter an Integer : ")
    num, _ := fmt.Scan(&inputString)
    if (num == 0) {
      os.Exit(3)
    }
  }
}

func isNumeric(s string) bool {
  _, err := strconv.ParseInt(s, 0, 64)
  return err == nil
}
