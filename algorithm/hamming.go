package hierdenc

import (
    "errors"

    "gonum.org/v1/gonum/mat"
)

func HammingDistance(n *mat.VecDense, m *mat.VecDense) (int, error) {
  if n.Len() != m.Len() {
    return -1, errors.New("Vectors to HammingDistance must be the same length.")
  }
  // Create bucket and sum the differential
  // Larger differential imp
  var differential int
  for i := 0; i < n.Len(); i++  {
      if n.AtVec(i) != m.AtVec(i)  {
          differential++
      }
  }
  return differential, nil
}
