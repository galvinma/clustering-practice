package louvain

import (
    "testing"
    "path/filepath"

    "github.com/galvinma/agora/algorithm/common"
)

func TestGetWeights(t *testing.T) {
    var val = [10][10]int{
          {0,3,3,5,3,2,5,3,3,2},
          {3,0,2,4,4,3,6,4,4,3},
          {3,2,0,4,2,1,4,2,2,1},
          {5,4,4,0,4,3,6,2,4,3},
          {3,4,2,4,0,1,2,2,2,1},
          {2,3,1,3,1,0,3,1,1,0},
          {5,6,4,6,2,3,0,4,4,3},
          {3,4,2,2,2,1,4,0,2,1},
          {3,4,2,4,2,1,4,2,0,1},
          {2,3,1,3,1,0,3,1,1,0},
    }
    simplepath, _ := filepath.Abs("./../../datasets/simple.data")
    data := common.InitData(simplepath)
    weights := GetWeights(data)

    for k,v := range(weights) {
        for m,n := range(v) {
            if n != val[k][m] {
                t.Errorf("GetWeights returned %v. Correct weight vector is %v.",
                n, val[k][m])
            }
        }
    }
}
