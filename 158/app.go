package main

import (  
    "fmt"
    "io/ioutil"
	"crypto/sha256"
    "encoding/base64"
)

func main() {  
    data, err := ioutil.ReadFile("test.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
	fmt.Println("Contents of file:", string(data))
	genHash([]byte(string(data)))
}

   

func genHash(bv []byte) {
    hasher := sha256.New()
    hasher.Write(bv)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	fmt.Println("sha value : ",sha)
}