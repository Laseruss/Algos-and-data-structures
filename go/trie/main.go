package main

import "fmt"

type Trie struct {
	children map[rune]*Trie
}

func (t *Trie) Search(prefix string) *Trie {
	node := t

	for _, r := range prefix {
		if node == nil {
			return nil
		}

		if child, ok := node.children[r]; ok {
			node = child
		} else {
			return nil
		}
	}
	return node
}

func (t *Trie) Insert(word string) {
	node := t

	for _, r := range word {
		if _, ok := node.children[r]; !ok {
			node.children[r] = &Trie{map[rune]*Trie{}}
		}

		node = node.children[r]
	}

	node.children = map[rune]*Trie{'*': nil}
}

func main() {
	trie := Trie{map[rune]*Trie{}}
	words := []string{"ace", "act", "bad", "bake", "bat", "batter", "cab", "cat", "catnap", "catnip"}
	for _, word := range words {
		trie.Insert(word)
	}

	fmt.Println(trie.Search("catn"))
}
