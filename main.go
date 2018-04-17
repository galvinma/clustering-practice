package main

import (
    "log"

    "github.com/galvinma/agora/algorithm/common"
    "github.com/galvinma/agora/algorithm/hierdenc"
    // "github.com/galvinma/agora/algorithm/louvain"
)

func main() {
    log.Println("Initializing data...")
    data := common.InitSoybeanLarge()

    clusters := hierdenc.HIERDENC(data)
    log.Println("Assigned clusters to",len(clusters),"out of",len(data),"objects")
    for k,v := range(clusters) {
        log.Println("Object",k,"assigned to cluster",v)
    }


}
