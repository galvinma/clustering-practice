package hierdenc

import (
    "bufio"
    "fmt"
    "os"
    "log"
    "path/filepath"

    // "gonum.org/v1/gonum/mat"
)

// Grabs catagorical data from file and creates a map
func InitSoybeanLarge() {
    soybean, _ := filepath.Abs("../datasets/soybean-large.data")
    file, err := os.Open(soybean)
    if err != nil {
        log.Fatal(err)
    }
    reader := bufio.NewScanner(file)
    for reader.Scan()  {
        fmt.Println(reader.Text())
    }

    // Return a Map of this data
}
