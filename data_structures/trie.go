package data_structures

const (
	ALPHABET_LENGTH = 26
)

type trieNode struct {
	childs    [ALPHABET_LENGTH]*trieNode
	value     string
	isWordEnd bool
}

type trie struct {
	root *trieNode
}

func InitTrie() *trie {
	return &trie{root: &trieNode{}}
}

func (t *trie) Insert(word string) {
	current := t.root
	wordLength := len(word)

	for i := 0; i < wordLength; i++ {
		to := 1
		if i == wordLength-1 {
			to = 0
		}

		index := word[i] - 'a'
		if current.childs[index] == nil {
			current.childs[index] = &trieNode{}
		}

		current.childs[index].value = word[:i+to]
		current = current.childs[index]

	}

	current.isWordEnd = true

}

func (t *trie) findSimilarWords(tn *trieNode, words *[]string, count int) {
	var current *trieNode
	if tn == nil || count == 4 {
		return
	}

	for _, v := range tn.childs {
		if v != nil {
			*words = append(*words, v.value)
			current = v
			break
		}
	}

	t.findSimilarWords(current, words, count+1)
	return
}

func (t *trie) Find(word string) []string {
	words := make([]string, 0, 3)
	current := t.root

	for _, v := range word {
		if v < 'a' || v > 'z' {
			return nil
		}
		index := v - 'a'
		if current.childs[index] == nil {
			return nil
		}

		current = current.childs[index]
	}

	t.findSimilarWords(current, &words, 1)

	if len(words) == 0 {
		return nil
	}

	return words
}
