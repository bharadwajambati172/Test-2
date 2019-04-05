package main


import (  
	"fmt"
	"strings"
	"sort"
)

func main(){
	input := []string{"war","eat","donut","pizza","raw","ate","apizz"}
	// x := sortWord(input[0])
	result := [][]string{}
	for i:=0; i<len(input);i++{
		fmt.Println(input[i])
		words := []string{}
		for j:=i; j<len(input); j++{
			word := sortWord(input[j])
			if strings.Compare(input[j],word) == 0 {
				words = append(words,input[j])
			}
		}
		result=append(result,words)
	}
	fmt.Println(result)

}
func sortWord(word string) string{
	words := strings.Split(word, "")
    sort.Strings(words)
    return strings.Join(words, "")
}