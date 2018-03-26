package hierdenc

import (
    "bufio"
    "encoding/csv"
    "io"
    "os"
    "log"
    "path/filepath"
    "strconv"

    // "gonum.org/v1/gonum/mat"
)

// Grabs catagorical data from file and creates a map
func InitSoybeanLarge() map[int][]int {
    // Initialize map
    var soybeans map[int][]int
    soybeans = make(map[int][]int)
    // Objects ID in the map
    count := 0
    // Open dataset
    soybeanpath, _ := filepath.Abs("./datasets/soybean-large.data")
    soybeanfile, err := os.Open(soybeanpath)
    if err != nil {
        log.Fatal(err)
    }
    // reader := bufio.NewScanner(soybeanfile)
    reader := csv.NewReader(bufio.NewReader(soybeanfile))
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
        soybeans[count] = holder
        count++
    }
    return soybeans
}
