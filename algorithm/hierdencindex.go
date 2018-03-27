package hierdenc

import (
    "log"
)

/* Keep a list of all objects in space S and their associated density ranking. Thinking this should be a goroutine that feeds into the general hierdenc function. For each object in S, find the total number of objects that are dissimilar. Objects with a lower score have a high density, implying duplicate data. Might be better to consider a threshold or cap for this function. As the number of catagories grows, it will be less likely for the index to have a perfect match. Perhaps keep a tally of the number of objects within r of an object with a hamming distance threshold. Larger quantity of objects implies a denser space. */

func HierdencIndex(objects map[int][]int, r int) (int, map[int]int) {
    var index map[int]int
    index = make(map[int]int)
    var densest int
    for key, _ := range objects {
        // Bucket for object's density count.
        var density int
        for _, val := range objects {
            dist, err := HammingDistance(objects[key], val)
            if err != nil {
                log.Fatal(err)
            } else {
                /* If the hamming distance between two objects is <= to the radius, increment the density. Two objects will have a hamming distance of zero if they have the same catagorical attributes. */
                if dist <= r {
                    density++
                }
            }
        }
        index[key] = density
        /* This will return the first value found with given density. Keys with equivilant densities will be "densest" in a later cycle.*/
        if density > index[densest] {
            densest = key
        }
    }
    return densest, index
}
