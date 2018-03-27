package hierdenc

import (
    "log"


)

/* Keep a list of all objects in space S and their associated density ranking. Thinking this should be a goroutine that feeds into the general hierdenc function. For each object in S, find the total number of objects that are dissimilar. Objects with a lower score have a high density, implying duplicate data. Might be better to consider a threshold or cap for this function. As the number of catagories grows, it will be less likely for the index to have a perfect match. Perhaps keep a tally of the number of objects within r of an object with a hamming distance threshold. Larger quantity of objects implies a denser space. */

func HierdencIndex(objects map[int][]int, radius int) []int {


    for key, _ := range objects {
        for _, val := range objects {



        }

    }



}
