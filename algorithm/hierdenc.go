package hierdenc

import (
    "bufio"
    "fmt"
    "os"
    "log"

    // "gonum.org/v1/gonum/mat"
)

// Function accepts a .txt that represents a catagorical dataset (S)
func CreateMatrix() {
    // Come back and handle error...
    file, err := os.Open("/home/galvin/Workspace/go/src/github.com/galvinma/agora/seed.text")
    if err != nil {
        log.Fatal(err)
    }
    reader := bufio.NewScanner(file)
    for reader.Scan()  {
        fmt.Println(reader.Text())
    }
}
