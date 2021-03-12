package main

import (
	"fmt"

	"github.com/comdori-web/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	err := dictionary.Update(baseWord, "Second")
	if err != nil {
		fmt.Println(err)
	}

	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)

	fmt.Println(dictionary)
	dictionary.Delete(baseWord)
	fmt.Println(dictionary)
}
