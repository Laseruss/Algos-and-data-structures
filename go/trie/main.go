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

func collectWords(trieNode *Trie, word string, words []string) []string {
	currNode := trieNode
	for key, value := range currNode.children {
		if key == '*' {
			words = append(words, word)
		} else {
			word += string(key)
			words = collectWords(value, word, words)
			word = word[:len(word)-1]
		}
	}
	return words
}

func (t *Trie) FindCompletions(prefix string) []string {
	currNode := t.Search(prefix)

	if currNode == nil {
		return []string{}
	}

	words := collectWords(currNode, prefix, []string{})

	return words
}

func main() {
	trie := Trie{map[rune]*Trie{}}
	words := []string{"ace", "act", "bad", "bake", "bat", "batter", "cab", "cat", "catnap", "catnip"}
	for _, word := range words {
		trie.Insert(word)
	}

	fmt.Println(trie.FindCompletions("cat"))
}
