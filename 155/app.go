package main

import (  
    "bufio"
    "strings"
    "strconv"
    "fmt"
    "log"
    "os"
)

func main() {  
    f, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    s := bufio.NewScanner(f)
    fs, err := os.Create("output.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    for s.Scan() {
        fmt.Println(s.Text())
        length := len(strings.Replace(s.Text()," ","",-1))
        t := strconv.Itoa(length)
        data := s.Text()+" : "+t
        fmt.Fprintln(fs, data)
        if err != nil {
            fmt.Println(err)
            fs.Close()
            return
        }
    }
    err = s.Err()
    if err != nil {
        log.Fatal(err)
    }
}