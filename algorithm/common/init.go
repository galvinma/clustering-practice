package common

import (
    "bufio"
    "encoding/csv"
    "io"
    "os"
    "log"
    "path/filepath"
    "strconv"
)

// Grabs catagorical data from file and creates a map
func initData(path string) map[int][]int {
    // Initialize map
    var data map[int][]int
    data = make(map[int][]int)
    // Objects ID in the map
    count := 0
    // Open dataset
    file, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }
    // reader := bufio.NewScanner(soybeanfile)
    reader := csv.NewReader(bufio.NewReader(file))
    // Add each line to the map as a seperate vector
    for {
        // Read CSV line
        line, error := reader.Read()
        if error == io.EOF {
            break
        } else if error != nil {
            log.Fatal(error)
        }
        // Create a slice to hold values.
        var holder []int
        for i := range line {
            // Convert string to int
            val, err := strconv.Atoi(line[i])
            // This check gets rid of unwanted catagorical data
            if err == nil {
                holder = append(holder, val)
            }
        }
        // Apend list to the map
        data[count] = holder
        count++
    }
    return data
}

func InitSoybeanLarge() map[int][]int {
    // Initialize map
    soybeanpath, _ := filepath.Abs("./datasets/soybean-large.data")
    soybeans := initData(soybeanpath)
    return soybeans
}

func InitSimple() map[int][]int {
    simplepath, _ := filepath.Abs("./datasets/simple.data")
    simple := initData(simplepath)
    return simple
}
