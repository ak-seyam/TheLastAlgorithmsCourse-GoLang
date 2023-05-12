package ds

import "testing"

func TestTrieInsertion(t *testing.T) {
	trieRoot := NewEnglishWordsRetrievalTree()
	trieRoot.AddWord("abdullah")
	trieRoot.AddWord("andy")
	trieRoot.AddWord("ahmed")
	trieRoot.AddWord("cat")
	trieRoot.AddWord("hat")
	nextPos, err := trieRoot.GetNextPossibleLetters('a')
	if err != nil {
		t.Fatal(err.Error())
	}
	if len(nextPos) != 3 {
		t.Fatal("invalid next pos #items")
	}
	nextPostSet := ToSet(nextPos)
	expectedNext := []rune{'b', 'n', 'h'}
	for _, v := range expectedNext {
		if !nextPostSet.Contains(v) {
			t.Fatalf("posibile values set should contain %c but it is not found", v)
		}
	}
}

func TestIsWordEnding(tester *testing.T) {
	wordsTable := NewEnglishWordsRetrievalTree()
	wordsTable.AddWord("cat")
	root := wordsTable.root
	c := root.Children[idx('c')]
	a := c.Children[idx('a')]
	t := a.Children[idx('t')]
	if !t.WordTerminator {
		tester.Fatal("invalid word terminator")
	}
}

func TestGettingAllWordsStartingWith(t *testing.T) {
	wordsTree := NewEnglishWordsRetrievalTree()
	wordsTree.AddWord("cat")
	wordsTree.AddWord("car")
	wordsTree.AddWord("cattle")
	wordsTree.AddWord("abdullah")
	wordsTree.AddWord("card")
	allWords := ToSet(wordsTree.GetWordsFrom('c'))
	expectedWords := []string{"cat", "car", "cattle", "card"}
	for _, word := range expectedWords {
		if !allWords.Contains(word) {
			t.Fatalf("words should contain '%s' but it doesn't", word)
		}
	}
	allWords_a := ToSet(wordsTree.GetWordsFrom('a'))
	expectedWords_a := []string{"abdullah"}
	for _, word := range expectedWords_a {
		if !allWords_a.Contains(word) {
			t.Fatalf("words should contain '%s' but it doesn't", word)
		}
	}
}

func TestWordDelete(t *testing.T) {
	wordsTree := NewEnglishWordsRetrievalTree()
	wordsTree.AddWord("cat")
	wordsTree.AddWord("car")
	wordsTree.AddWord("cattle")
	wordsTree.AddWord("abdullah")
	wordsTree.DeleteWord("cat")
	wordsTree.DeleteWord("car")
	allWords := ToSet(wordsTree.GetWordsFrom('c'))
	for k := range allWords {
		if k == "cat" || k == "car" {
			t.Fatal("words 'cat' and 'car' shouldn't exist")
		}
	}
}
