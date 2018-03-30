package hierdenc

func Cluster(clusters map[int]int, objects map[int][]int, index []Object, r int, id int) (map[int]int, []Object, int) {
  // Start cluster count @ 1.
  // Map lookup relies on nil value of 0 for non-existant key
  for _, object := range index {
      // If density is 1, increase r
      // Density of 1 implies no other objects within r
      if object.Density <= 1 {
          // Update the index
          index = HierdencIndex(objects, r)
          break
      }

      // If object not in map, check if it can be added to an existing cluster
      if clusters[object.ID] == 0 {
        // for obj #, cluster #
        for k,v := range clusters {
            hd, _ := HammingDistance(objects[object.ID], objects[k])
            if hd <= r {
                // Add object to found cluster
                clusters[object.ID] = v
                break
            }
        }
      }

      // If object cannot be added to an existing cluster, create a new cluster
      if clusters[object.ID] == 0 {
          clusters[object.ID] = id
          // for obj #, attributes
          for k,v := range objects {
              hd, _ := HammingDistance(objects[object.ID], v)
              if hd <= r {
                  // Add object to current cluster
                  clusters[k] = id
              }
          }
          // Increment the counter
          id++
      }
  }
  return clusters, index, id
}
