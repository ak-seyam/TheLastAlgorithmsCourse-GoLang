package ds

import (
	"errors"
	"strings"
	"unicode"
)

type EnglishWordsTrieNode struct {
	Children       [26]*EnglishWordsTrieNode
	Value          rune
	WordTerminator bool
}

type EnglishWordsRetrievalTree struct {
	root *EnglishWordsTrieNode
}

type TrieTraversalCB func(
	word string,
)

func NewEnglishWordsRetrievalTree() EnglishWordsRetrievalTree {
	return EnglishWordsRetrievalTree{
		root: &EnglishWordsTrieNode{
			Children: [26]*EnglishWordsTrieNode{},
			Value:    0,
		},
	}
}

func (t *EnglishWordsRetrievalTree) AddWord(str string) error {
	neutralizedStr := strings.ToLower(str)
	zero := byte('a')
	current := t.root
	for i, char := range neutralizedStr {
		respectiveValue := byte(char) - zero
		if respectiveValue > 25 {
			return errors.New("invalid character, not an english letter")
		}
		if current.Children[respectiveValue] == nil {
			current.Children[respectiveValue] = &EnglishWordsTrieNode{
				Value: char,
			}
		}
		if i == len(str)-1 {
			current.Children[respectiveValue].WordTerminator = true
		}
		current = current.Children[respectiveValue]
	}
	return nil
}

func (t EnglishWordsRetrievalTree) GetNextPossibleLetters(currentChar rune) ([]rune, error) {
	res := []rune{}
	zero := byte('a')
	neuCurrChar := byte(unicode.ToLower(currentChar))
	idx := neuCurrChar - zero
	if idx > 25 {
		return res, errors.New("not an english rune")
	}
	for _, child := range t.root.Children[idx].Children {
		if child != nil {
			res = append(res, child.Value)
		}
	}
	return res, nil
}

func (t EnglishWordsTrieNode) IsEnd() bool {
	for _, v := range t.Children {
		if v != nil {
			return false
		}
	}
	return true
}

func (t EnglishWordsRetrievalTree) GetWordsFrom(startingFrom rune) []string {
	root := t.root.Children[idx(startingFrom)]
	res := []string{}
	sb := &[]rune{}
	trieDFS(root, root, sb, func(word string) { res = append(res, word) })
	return res
}

func trieDFS(
	rootPtr *EnglishWordsTrieNode,
	n *EnglishWordsTrieNode,
	currentWord *[]rune,
	cb TrieTraversalCB,
) {
	if n == nil {
		return
	}
	if n.Value != 0 && !n.WordTerminator {
		val := append(*currentWord, n.Value)
		currentWord = &val
	}
	if n.WordTerminator {
		val := append(*currentWord, n.Value)
		cb(string(val))
		if !n.IsEnd() {
			currentWord = &val
		}
	}
	if n == rootPtr {
		currentWord = &[]rune{}
		val := append(*currentWord, n.Value)
		currentWord = &val
	}
	for _, v := range n.Children {
		if v != nil {
			trieDFS(rootPtr, v, currentWord, cb)
		}
	}
}

func (t *EnglishWordsRetrievalTree) DeleteWord(word string) {
	neutralizedWord := strings.ToLower(word)
	deleteFromTrie(t.root, t.root.Children[idx(rune(neutralizedWord[0]))], neutralizedWord, 0)
}

func deleteFromTrie(
	p *EnglishWordsTrieNode,
	n *EnglishWordsTrieNode,
	word string,
	currentIndex int,
) {
	if n == nil {
		return
	}
	char := word[currentIndex]
	if n.IsEnd() {
		index := idx(rune(char))
		p.Children[index] = nil
		return
	}
	if currentIndex == len(word)-1 {
		if n.WordTerminator && !n.IsEnd() {
			n.WordTerminator = false
		}
	}
	nxtIdx := currentIndex + 1
	if nxtIdx >= len(word) {
		return
	}
	newVal := n.Children[idx(rune(word[nxtIdx]))]
	deleteFromTrie(
		n,
		newVal,
		word,
		nxtIdx,
	)
}

func idx(r rune) rune {
	return r - 'a'
}
