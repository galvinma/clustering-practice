package common

import (
    "testing"

)

func TestHappyPath(t *testing.T) {
  zeros := []int{0,0,0,0,0}
  ones := []int{1,1,1,1,1}
  val, _ := HammingDistance(zeros, ones)
  // expect 5
  if val != 5 {
    t.Errorf("HammingDistance returned %v, wanted 5",
      val)
  }
}

func TestVectorMismatch(t *testing.T) {
  zeros := []int{0,0,0,0}
  ones := []int{1,1,1,1,1}
  err, _ := HammingDistance(zeros, ones)
  if err != -1 {
    t.Errorf("HammingDistance returned %v, wanted -1",
      err)
  }
}
