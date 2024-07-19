package main

type TrieNode struct {
	children    map[rune]*TrieNode
	isEndOfWord bool
}
