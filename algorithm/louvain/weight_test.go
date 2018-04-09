package louvain

import (
    "testing"
    "path/filepath"

    "github.com/galvinma/agora/algorithm/common"
)

func getData(path string) map[int][]int {
    filepath, _ := filepath.Abs(path)
    data := common.InitData(filepath)
    return data
}

func getTestWeights(path string) [][]int {
    simplepath, _ := filepath.Abs(path)
    data := common.InitData(simplepath)
    weights := GetWeights(data)
    return weights
}

func forcedSimpleCommunities() map[int]int {
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
      return communities
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
    communities := forcedSimpleCommunities()

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

func TestSumIncidentCommunityWeights(t *testing.T) {
      data := getData("./../../datasets/simple.data")
      weights := getTestWeights("./../../datasets/simple.data")
      communities := forcedSimpleCommunities()

      com := 0
      sum := 76.0
      comsum := SumIncidentCommunityWeights(data, communities, weights, com)
      if comsum != sum {
          t.Errorf("SumIncidentCommunityWeights returned %v. Correct sum for community %v is %v.",
          comsum, com, sum)

      }

      com = 5
      sum = 61.0
      comsum = SumIncidentCommunityWeights(data, communities, weights, com)
      if comsum != sum {
          t.Errorf("SumIncidentCommunityWeights returned %v. Correct sum for community %v is %v.",
          comsum, com, sum)

      }

      com = 9
      sum = 15.0
      comsum = SumIncidentCommunityWeights(data, communities, weights, com)
      if comsum != sum {
          t.Errorf("SumIncidentCommunityWeights returned %v. Correct sum for community %v is %v.",
          comsum, com, sum)

      }
}

func TestGetIncidentWeights(t *testing.T) {
    data := getData("./../../datasets/simple.data")
    weights := getTestWeights("./../../datasets/simple.data")

    object := 1
    expected := 33.0
    sum := GetIncidentWeights(data, weights, object)
    if sum != expected {
        t.Errorf("GetIncidentWeights returned %v. Correct sum for object %v is %v.",
        sum, object, expected)
    }

    object = 5
    expected = 15.0
    sum = GetIncidentWeights(data, weights, object)
    if sum != expected {
        t.Errorf("GetIncidentWeights returned %v. Correct sum for object %v is %v.",
        sum, object, expected)
    }

}

func TestGetIncidentCommunityWeights(t *testing.T) {
    weights := getTestWeights("./../../datasets/simple.data")
    communities := forcedSimpleCommunities()

    object := 1
    community := 0
    expected := 9.0
    sum := GetIncidentCommunityWeights(communities, weights, community, object)
    if sum != expected {
        t.Errorf("GetIncidentCommunityWeights returned %v. Correct sum for object:community %v:%v is %v.",
        sum, object, community, expected)
    }

    object = 1
    community = 5
    expected = 13.0
    sum = GetIncidentCommunityWeights(communities, weights, community, object)
    if sum != expected {
        t.Errorf("GetIncidentCommunityWeights returned %v. Correct sum for object:community %v:%v is %v.",
        sum, object, community, expected)
    }

    object = 5
    community = 9
    expected = 0.0
    sum = GetIncidentCommunityWeights(communities, weights, community, object)
    if sum != expected {
        t.Errorf("GetIncidentCommunityWeights returned %v. Correct sum for object:community %v:%v is %v.",
        sum, object, community, expected)
    }

    object = 8
    community = 8
    expected = 0.0
    sum = GetIncidentCommunityWeights(communities, weights, community, object)
    if sum != expected {
        t.Errorf("GetIncidentCommunityWeights returned %v. Correct sum for object:community %v:%v is %v.",
        sum, object, community, expected)
    }

    object = 4
    community = 0
    expected = 13.0
    sum = GetIncidentCommunityWeights(communities, weights, community, object)
    if sum != expected {
        t.Errorf("GetIncidentCommunityWeights returned %v. Correct sum for object:community %v:%v is %v.",
        sum, object, community, expected)
    }

    object = 4
    community = 7
    expected = 2.0
    sum = GetIncidentCommunityWeights(communities, weights, community, object)
    if sum != expected {
        t.Errorf("GetIncidentCommunityWeights returned %v. Correct sum for object:community %v:%v is %v.",
        sum, object, community, expected)
    }

}

func TestSumNetworkLinks(t *testing.T) {
    weights := getTestWeights("./../../datasets/simple.data")
    sum := SumNetworkLinks(weights)
    expected := 125.0
    if sum != expected {
        t.Errorf("Network sum returned %v instead of expected %v.",
        sum, expected)
    }
}
