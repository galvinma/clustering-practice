package hierdenc

import (
    "testing"

    "gonum.org/v1/gonum/mat"
)

func TestHappyPath(t *testing.T) {
  zeros := mat.NewVecDense(5, []float64{0,0,0,0,0})
  ones := mat.NewVecDense(5, []float64{1,1,1,1,1})
  val, _ := HammingDistance(zeros, ones)
  // expect 5
  if val != 5 {
    t.Errorf("HammingDistance returned %v, wanted 5",
      val)
  }
}

func TestVectorMismatch(t *testing.T) {
  zeros := mat.NewVecDense(5, []float64{0,0,0,0,0})
  ones := mat.NewVecDense(4, []float64{1,1,1,1})
  err, _ := HammingDistance(zeros, ones)
  if err != -1 {
    t.Errorf("HammingDistance returned %v, wanted -1",
      err)
  }
}
