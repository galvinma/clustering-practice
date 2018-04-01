package common

import (
    "errors"
)

func HammingDistance(n []int, m []int) (int, error) {
  if len(n) != len(m) {
    return -1, errors.New("Lists to HammingDistance must be the same length.")
  }
  // Create bucket and sum the differential
  var differential int
  for i := 0; i < len(n); i++  {
      if n[i] != m[i] {
          differential++
      }
  }
  return differential, nil
}
