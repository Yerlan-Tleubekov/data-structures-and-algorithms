package main

import (
	"awesomeProject2/data_structures"
	"fmt"
	"os"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	words := []string{"hello", "he", "me", "yerlan", "be", "find", "insert", "inser", "inse", "ins", "in", "i"}

	//finder := []string{"hello", "he", "me", "yerlan", "be", "find", "insert", "mal", "idiot", "lol"}

	trie := data_structures.InitTrie()

	wg.Add(len(words))
	for _, v := range words {

		go func(w string) {
			trie.Insert(w)
			defer wg.Done()
		}(v)
	}
	wg.Wait()

	for {
		word := ""
		count, err := fmt.Fscan(os.Stdin, &word)
		if word == "\\quit" {

			break
		}
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		if count != 1 {
			fmt.Println("enter the one word")
			break
		}
		wrds := trie.Find(word)
		fmt.Println(wrds)

	}

}
