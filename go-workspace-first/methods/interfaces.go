package main

import (
  "fmt"
  "math"
)

/*
定义一个接口
*/
type Abser interface {
  Abs() float64
}

type Myfloat float64

func (f Myfloat) Abs() float64 {
  if f < 0 {
    return float64(-f)
  }
  return float64(f)
}

