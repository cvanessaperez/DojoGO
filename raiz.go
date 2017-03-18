package main

import (
  "fmt"
  "math"
)

func Sqrt(x float64) float64 {
  z := float64(x/2)

  //for i := 0; i < 10; i++{
    //z = z - (((z*z) - x)/(2*z))
  //}
  //return z

  counter := 0

  for delta := z; math.Abs(delta) > 1e-6; {
    zi := z
    z = z - (((z*z) - x)/(2*z))
    delta = z- zi
    counter++
  }
  fmt.Println(counter)
  return z

}

func main() {
  fmt.Println(Sqrt(10))
  fmt.Println(math.Sqrt(10))
}
