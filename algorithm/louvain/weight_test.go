package louvain

import (
    "testing"
    "path/filepath"

    "github.com/galvinma/agora/algorithm/common"
)

func getTestWeights(path string) [][]int {
    simplepath, _ := filepath.Abs(path)
    data := common.InitData(simplepath)
    weights := GetWeights(data)
    return weights
}

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

    weights := getTestWeights("./../../datasets/simple.data")
    for k,v := range(weights) {
        for m,n := range(v) {
            if n != val[k][m] {
                t.Errorf("GetWeights returned %v for weights[%v][%v]. Correct weight vector is %v.",
                n, k, m, val[k][m])
            }
        }
    }
}

func TestSumCommunityWeights(t *testing.T) {
    weights := getTestWeights("./../../datasets/simple.data")
    // Define communities
    var communities = map[int]int{
        // Object ID : Cluster ID
        0:0,
        1:0,
        2:0,
        3:0,
        4:5,
        5:5,
        6:5,
        7:7,
        8:8,
        9:9,
      }



    com := 0
    sum := 21.0
    comsum := SumCommunityWeights(communities, weights, com)
    if comsum != sum {
        t.Errorf("SumCommunityWeights returned %v. Correct sum for community %v is %v.",
        comsum, com, sum)

    }

    com = 5
    sum = 6.0
    comsum = SumCommunityWeights(communities, weights, com)
    if comsum != sum {
        t.Errorf("SumCommunityWeights returned %v. Correct sum for community %v is %v.",
        comsum, com, sum)

    }

    com = 9
    sum = 0.0
    comsum = SumCommunityWeights(communities, weights, com)
    if comsum != sum {
        t.Errorf("SumCommunityWeights returned %v. Correct sum for community %v is %v.",
        comsum, com, sum)

    }

}
