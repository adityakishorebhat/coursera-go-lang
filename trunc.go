package main

import "fmt"

func main() {
    var floatNum float64
    fmt.Printf("Enter Float Number: ")
    num, _ := fmt.Scan(&floatNum)
    if num > 0 {
	var intNum int = int(floatNum)
	fmt.Println(intNum)
    }
}
