package main

import (
    "fmt"
	"encoding/json"
	"io/ioutil"
	"os"
)

type Jason struct {
    Data string
}

func main() {
	data, err := ioutil.ReadFile("test.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
	fmt.Println("Contents of file:", string(data))

    a := &Jason{string(data)}

    out, err := json.Marshal(a)
    if err != nil {
        panic (err)
    }

	fmt.Println(string(out))
	
	fs, err := os.Create("output.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    _, err = fs.WriteString(string(out))
    if err != nil {
        fmt.Println(err)
        fs.Close()
        return
    }
}