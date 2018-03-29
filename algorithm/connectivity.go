package hierdenc

/* HIERDENC cut-off level. */

func avgScore(scores []float64) float64{
    total := 0
    for _, v := range scores {
        total += v
    }
    avg := total/float64(len(scores))
    return avg

}

func inList(i int, c []int) {
    for _, cluster := range(c) {
        if i == cluster {
            return true
        }
    }
    return false
}

func CalculateSilhouetteScore(clusters map[int]int) float64 {
    // Create a list of known clusters
    var clustlist []int
    // Object ID, Cluster ID
    for _, v := range(clusters) {
        // If cluster ID is not in list, add it
        if inList(v, clustlist) = false {
            clustlist = append(clustlist, v)
        }
    }

    // Bucket for average score of each point in map
    var scores []float64
    // Calculate silhouette coefficient for each point
    for c := range clustlist {
        // Get all objects in given cluster
        // Object ID, Cluster ID
        for k,v := range clusters {
            if c == v {


            }

        }

    }




    // Average the list and return score
    average :=  avgScore(scores)
    return average
}
