package hierdenc

import (
    "bufio"
    "encoding/csv"
    "io"
    "fmt"
    "os"
    "log"
    "path/filepath"
    "strconv"

    "gonum.org/v1/gonum/mat"
)

// Grabs catagorical data from file and creates a map
func InitSoybeanLarge() {
    // Initialize map
    var soybeans map[int][]*mat.VecDense
    soybeans = make(map[int][]*mat.VecDense)
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
        // Use slice as an intermediary to filter out values that cannot be
        // converted to float64. NewVecDense requires the length of th vector
        // at the time the variable is created.
        var holder []float64
        for i := range line {
            // Convert string to float64
            val, err := strconv.ParseFloat(line[i], 64)
            // This check gets rid of unwanted catagorical data
            if err == nil {
                holder = append(holder, val)
            }
        }
        // Create a vector from holder slice.
        vector := mat.NewVecDense(len(holder), nil)
        for i := 0; i < len(holder); i++ {
            vector.SetVec(i, holder[i])
        }
        // Apend vector to the map
        soybeans[count] = append(soybeans[count], vector)
        count++
    }

    // Return a map of this data
    for key, value := range soybeans {
    fmt.Println("Key:", key, "Value:", value)
    } 

}
