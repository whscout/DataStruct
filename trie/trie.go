package trie

import (
	"container/list"
	"fmt"
)

type Trie struct {
	val int
	isWord bool
	children [256]*Trie
}

func NewTire() *Trie {
	return &Trie{}
}

func (t *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}
	node := t
	for _, v := range word {
		if node.children[v] == nil {
			tmp := &Trie{}
			tmp.val = int(v)
			node.children[v] = tmp
		}
		node = node.children[v]
	}
	node.isWord = true
}

func (t *Trie) Search(word string) bool {
	if len(word) == 0 {
		return false
	}
	node := t
	for _, v := range word {
		if node.children[v] == nil {
			return false
		}
		node = node.children[v]
	}
	return node.isWord
}

func (t *Trie) WithStart(word string) bool {
	if len(word) == 0 {
		return false
	}
	node := t
	for _, v := range word {
		if node.children[v] == nil {
			return false
		}
		node = node.children[v]
	}
	return true
}

func (t *Trie)StartAllWord(word string) {
	if len(word) == 0 {
		return
	}
	node := t
	for _, v := range word {
		if node.children[v] == nil {
			return
		}
		node = node.children[v]
	}
	// 此时node是前缀的最后一个字母

	//cur := node
	queue := list.New()
	//fmt.Println(node)
	for {
		for i := 0; i < 256; i++ {
			if node.children[i] != nil {
				queue.PushBack(node.children[i])
			}
		}
		tmp := queue.Front()
		if !(tmp.Value.(*Trie)).isWord {
			word += string((tmp.Value.(*Trie)).val)
			node = node.children[(tmp.Value.(*Trie)).val]
		} else {
			fmt.Println(word)
		}
		if queue.Len() == 0 {
			return
		}
	}
}
