package main

import (
	"DataStruct/trie"
	"fmt"
)

func main() {
	root := trie.NewTire()

	root.Insert("Hello")
	root.Insert("apple")

	fmt.Println(root.Search("app"))
	fmt.Println(root.WithStart("app"))

	root.StartAllWord("app")
}


