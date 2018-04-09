package louvain

import (
    "testing"
    "path/filepath"

    "github.com/galvinma/agora/algorithm/common"
)

// Test initial communities
func TestInitialCommunities(t *testing.T) {
    var communities = map[int]int{
        // Object ID : Cluster ID
        0:0,
        1:1,
        2:2,
        3:3,
        4:4,
        5:5,
        6:6,
        7:7,
        8:8,
        9:9,
      }

      simplepath, _ := filepath.Abs("./../../datasets/simple.data")
      data := common.InitData(simplepath)
      network := InitialCommunities(data)
      for k,v := range(network) {
          if v != communities[k] {
              t.Errorf("InitialCommunities returned %v for object %v. Correct community is %v.",
              v, k, communities[k])
          }
      }
}

func TestGetMembers(t *testing.T) {
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

    members := GetMembers(communities, 0)
    for _,v := range(members) {
        if communities[v] != 0 {
            t.Errorf("GetMembers returned an incorrect member list. Got %v, wanted %v.",
            v, communities[v])
        }

    }
}
