package ac_automaton

import "github.com/emirpasic/gods/v2/queues/linkedlistqueue"

type TrieNode struct {
	end     string
	matched bool
	fail    *TrieNode
	nexts   map[rune]*TrieNode
}

var root = &TrieNode{
	nexts: make(map[rune]*TrieNode),
}

func NewAcAutomaton(words ...string) {
	for _, word := range words {
		insert(word)
	}
	build()
}

func insert(word string) {
	cur := root
	for _, ch := range word {
		if _, ok := cur.nexts[ch]; !ok {
			cur.nexts[ch] = &TrieNode{
				nexts: make(map[rune]*TrieNode),
			}
		}
		cur = cur.nexts[ch]
	}
	cur.end = word
}

func build() {
	queue := linkedlistqueue.New[*TrieNode]()
	queue.Enqueue(root)
	for queue.Size() > 0 {
		cur, _ := queue.Dequeue()
		for ch, next := range cur.nexts {
			cfail := cur.fail
			for cfail != nil && cfail.nexts[ch] == nil {
				cfail = cfail.fail
			}
			if next.fail = root; cfail != nil {
				next.fail = cfail.nexts[ch]
			}
			queue.Enqueue(next)
		}
	}
}

func ContainWords(content string) (matches []string) {
	cur := root
	for _, ch := range content {
		for cur.nexts[ch] == nil && cur != root {
			cur = cur.fail
		}
		if cur = cur.nexts[ch]; cur == nil {
			cur = root
		}
		follow := cur
		for follow != root && !follow.matched {
			if follow.end != "" {
				matches = append(matches, follow.end)
				follow.matched = true
			}
			follow = follow.fail
		}
	}

	return
}
