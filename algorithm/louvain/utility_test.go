package louvain

import (
    "testing"
    "path/filepath"

    "github.com/galvinma/agora/algorithm/common"
)

func checkList(id int, memberlist []int) bool {
    for _,v := range(memberlist) {
        if id == v {
            return true
        }
    }
    return false
}

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

    // Placement in community "0"
    zeromembers := []int{0,1,2,3}
    com := 0
    members := GetMembers(communities, com)
    for _,v := range(zeromembers) {
        if checkList(v, members) != true {
            t.Errorf("GetMembers failed to place object ID %v in community ID %v.",
            v, com)
        }
    }

    // Placement in community "5"
    fivemembers := []int{4,5,6}
    com = 5
    members = GetMembers(communities, com)
    for _,v := range(fivemembers) {
        if checkList(v, members) != true {
            t.Errorf("GetMembers failed to place object ID %v in community ID %v.",
            v, com)
        }
    }

}
