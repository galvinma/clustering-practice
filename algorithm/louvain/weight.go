package louvain

import (
    "log"

    "github.com/galvinma/agora/algorithm/common"
)

func GetWeights(data map[int][]int) map[int][]int {
    log.Println("Getting vertex weights...")
    var weights map[int][]int
    weights = make(map[int][]int)

    for k,v := range(data) {
        var hd []int
        for _,m := range(data) {
            dif, err := common.HammingDistance(v,m)
            if err != nil {
                log.Fatal(err)
            }
            hd = append(hd, dif)
        }
        weights[k] = hd
    }
    return weights
}
