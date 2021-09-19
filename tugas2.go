package main

import "fmt"

var angka int16 = 11

func main(){
  if angka%2==0 {
    fmt.Printf("%d merupakan bilangan genap",angka)
  }else {
    fmt.Printf("%d merupakan bilangan ganjil",angka)
  }
}
