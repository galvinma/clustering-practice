package hierdenc

import (
    "log"
    "sort"

    "github.com/galvinma/agora/algorithm/common"
)
// struct keeps Object ID & density
type Object struct {
    ID int
    Density int
}

/* Generate a list of all objects in space S and their associated density ranking for a given radius r. */
func HierdencIndex(objects map[int][]int, r int) ([]Object) {
    // sorted list of densest hypercubes for radius r
    var index []Object
    for key, _ := range objects {
        // Bucket for object's density count.
        var density int
        for _, val := range objects {
            dist, err := common.HammingDistance(objects[key], val)
            if err != nil {
                log.Fatal(err)
            } else {
                /* If the hamming distance between two objects is <= to the radius, increment the density. Two objects will have a hamming distance of zero if they have the same catagorical attributes. */
                if dist <= r {
                    density++
                }
            }
        }
        index = append(index, Object{key, density})
    }
    sort.Slice(index, func(i, j int) bool { return index[i].Density > index[j].Density })
    return index
}
